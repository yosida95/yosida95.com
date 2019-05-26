# -*- coding: utf-8 -*- #

AUTHOR = 'yosida95'
SITENAME = 'yosida95'
SITEURL = ''
RELATIVE_URLS = True

TIMEZONE = 'Asia/Tokyo'
DEFAULT_LANG = 'ja'

PATH = 'content'
STATIC_PATHS = ['_static']

PATH_METADATA = r'(?P<slug>(?P<date>\d{4}/\d{2}/\d{2})/[^\.]+)\.rst'
EXTRA_PATH_METADATA = {
    '_static/robots.txt': {'path': 'robots.txt'},
    '_static/favicon.ico': {'path': 'favicon.ico'},
}

DEFAULT_CATEGORY = 'Uncategorized'
USE_FOLDER_AS_CATEGORY = False

THEME = './_themes/purity'
SUMMARY_MAX_LENGTH = 10

DEFAULT_PAGINATION = 5
PAGINATION_PATTERNS = (
    (1, '{base_name}/', '{name}{extension}'),
    (2, '{base_name}/page{number}{extension}',
        '{base_name}/page{number}{extension}'),
)


INDEX_URL = 'archives/'
INDEX_SAVE_AS = INDEX_URL + 'index.html'

ARCHIVES_SAVE_AS = ''

YEAR_ARCHIVE_URL = '{date:%Y}/'
YEAR_ARCHIVE_SAVE_AS = YEAR_ARCHIVE_URL + 'index.html'

MONTH_ARCHIVE_URL = '{date:%Y}/{date:%m}/'
MONTH_ARCHIVE_SAVE_AS = MONTH_ARCHIVE_URL + 'index.html'

DAY_ARCHIVE_URL = '{date:%Y}/{date:%m}/{date:%d}/'
DAY_ARCHIVE_SAVE_AS = DAY_ARCHIVE_URL + 'index.html'

ARTICLE_PATHS = ['']
ARTICLE_URL = '{slug}.html'
ARTICLE_SAVE_AS = ARTICLE_URL

DRAFT_URL = ARTICLE_URL
DRAFT_SAVE_AS = ARTICLE_SAVE_AS

PAGE_PATHS = ['_pages']
PAGE_URL = '{slug}/'
PAGE_SAVE_AS = PAGE_URL + 'index.html'

DRAFT_PAGE_URL = PAGE_URL
DRAFT_PAGE_SAVE_AS = PAGE_SAVE_AS

AUTHORS_SAVE_AS = ''
AUTHOR_SAVE_AS = ''

CATEGORIES_URL = 'categories/'
CATEGORIES_SAVE_AS = CATEGORIES_URL + 'index.html'

CATEGORY_URL = CATEGORIES_URL + '{slug}/'
CATEGORY_SAVE_AS = CATEGORY_URL + 'index.html'

TAGS_URL = 'tags/'
TAGS_SAVE_AS = TAGS_URL + 'index.html'

TAG_URL = TAGS_URL + '{slug}/'
TAG_SAVE_AS = TAG_URL + 'index.html'


JINJA_ENVIRONMENT = {
    'lstrip_blocks': False,
    'trim_blocks': False,
}

PYGMENTS_RST_OPTIONS = {}

PLUGIN_PATHS = ['../plugins/pelican-plugins']
PLUGINS = ['pelican_alias']


# Feed generation is usually not desired when developing
FEED_ALL_ATOM = None
CATEGORY_FEED_ATOM = None
TRANSLATION_FEED_ATOM = None
AUTHOR_FEED_ATOM = None
AUTHOR_FEED_RSS = None
