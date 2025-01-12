package photos

import (
	"fmt"
	"strings"
	"time"
)

type Photo struct {
	Id        PhotoId
	CreatedAt time.Time
	MIMEType  string
	Comment   string
	Scope     Scope
}

func (p *Photo) Key() string {
	return fmt.Sprintf("%s.%s", p.Id, p.Ext())
}

func (p *Photo) KeyResized() string {
	return fmt.Sprintf("%s.resized.%s", p.Id, p.Ext())
}

func (p *Photo) KeyCropped() string {
	return fmt.Sprintf("%s.thumbnail.%s", p.Id, p.Ext())
}

func (p *Photo) Ext() string {
	if strings.HasPrefix(p.MIMEType, "image/") {
		return p.MIMEType[len("image/"):]
	}
	return "ext"
}

type PhotoId string

func (id PhotoId) String() string { return string(id) }

type Scope string

const (
	Photo_Public  Scope = "PUBLIC"
	Photo_Private Scope = "PRIVATE"
)

func (scope Scope) String() string { return string(scope) }
