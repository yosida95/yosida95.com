package store_mysql

import (
	"context"
	"database/sql"
	"io"
	"time"

	"cloud.google.com/go/storage"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yosida95/yosida95.com/photos"
)

type Config struct {
	DSN       string
	Bucket    string
	KeyPrefix string
}

func New(c Config) (photos.StoreFactory, error) {
	conn, err := sql.Open("mysql", c.DSN)
	if err != nil {
		return nil, err
	}
	conn.SetMaxOpenConns(4)
	conn.SetConnMaxLifetime(60 * time.Second)
	if err := conn.Ping(); err != nil {
		return nil, err
	}

	s := &factory{
		conn: conn,
		cfg:  c,
	}
	return s, nil
}

type factory struct {
	conn *sql.DB
	cfg  Config
}

func (s *factory) Close() error {
	return s.conn.Close()
}

func (s *factory) Begin(ctx context.Context) (photos.Store, error) {
	tx, err := s.conn.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	cli, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	c := &cursor{
		ctx: ctx,
		tx:  tx,
		cli: cli,
		cfg: s.cfg,
	}
	return c, nil
}

type cursor struct {
	ctx context.Context
	tx  *sql.Tx
	cli *storage.Client
	cfg Config
}

func (c *cursor) Rollback() error {
	c.cli.Close()
	return c.tx.Rollback()
}

func (c *cursor) Commit() error {
	c.cli.Close()
	return c.tx.Commit()
}

func (c *cursor) FetchPhoto() photos.PhotoFetcher {
	return &photoFetcher{
		tx: c.tx,
		stmt: sq.
			Select("id", "mime_type", "comment", "created_at").
			From("photos").
			OrderBy("created_at"),
	}
}

type photoFetcher struct {
	tx   *sql.Tx
	stmt sq.SelectBuilder
}

func (f *photoFetcher) Id(id photos.PhotoId) photos.PhotoFetcher {
	f.stmt = f.stmt.Where(sq.Eq{"id": id.String()})
	return f
}

func (f *photoFetcher) Limit(i int64) photos.PhotoFetcher {
	f.stmt = f.stmt.Limit(uint64(i))
	return f
}

func (f *photoFetcher) Offset(i int64) photos.PhotoFetcher {
	f.stmt = f.stmt.Offset(uint64(i))
	return f
}

func (f *photoFetcher) First() (*photos.Photo, error) {
	f.stmt = f.stmt.Limit(1)
	ret, err := f.All()
	if err != nil {
		return nil, err
	}
	if len(ret) == 0 {
		return nil, photos.ErrPhotoNotFound
	}
	return ret[0], nil
}

func (f *photoFetcher) All() ([]*photos.Photo, error) {
	rows, err := f.stmt.RunWith(f.tx).Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ret := []*photos.Photo{}
	for rows.Next() {
		var id string
		var photo photos.Photo
		if err := rows.Scan(&id, &photo.MIMEType, &photo.Comment, &photo.CreatedAt); err != nil {
			return nil, err
		}
		photo.Id = photos.PhotoId(id)
		ret = append(ret, &photo)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ret, nil
}

func (c *cursor) CountPhoto() (int64, error) {
	var count int64
	err := sq.
		Select("COUNT(id)").
		From("photos").
		RunWith(c.tx).
		QueryRow().
		Scan(&count)
	return count, err
}

func (f *cursor) BlobGet(key string) (io.ReadCloser, error) {
	b := f.cli.Bucket(f.cfg.Bucket)
	obj := b.Object(f.cfg.KeyPrefix + key)
	r, err := obj.NewReader(f.ctx)
	if err != nil {
		if err == storage.ErrObjectNotExist {
			return nil, photos.ErrPhotoNotFound // FIXME(yosida95)
		}
		return nil, err
	}
	return r, nil
}
