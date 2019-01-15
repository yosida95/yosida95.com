# -*- coding: utf-8 -*- #

AUTHOR = 'yosida95'
SITENAME = 'yosida95'
SITEURL = ''
TIMEZONE = 'Asia/Tokyo'
DEFAULT_LANG = 'ja'

PATH = 'content'
PATH_METADATA = '(?P<date>\d{4}/\d{2}/\d{2})/(?P<slug>[^\.]*).rst'
USE_FOLDER_AS_CATEGORY = False
RELATIVE_URLS = True

THEME = './_themes/purity'
SUMMARY_MAX_LENGTH = 10
DEFAULT_PAGINATION = 5
PAGINATION_PATTERNS = (
    (1, '{base_name}/', '{name}{extension}'),
    (2, '{base_name}/page{number}{extension}', '{base_name}/page{number}{extension}'),
)
JINJA_ENVIRONMENT = {
    'lstrip_blocks': False,
    'trim_blocks': False,
}
PYGMENTS_RST_OPTIONS = {}

PLUGIN_PATHS = ['../plugins/pelican-plugins']
PLUGINS = ['pelican_alias']

INDEX_SAVE_AS = 'archives/index.html'
INDEX_URL = 'archives/'

ARCHIVES_SAVE_AS = ''
YEAR_ARCHIVE_SAVE_AS = '{date:%Y}/index.html'
MONTH_ARCHIVE_SAVE_AS = '{date:%Y}/{date:%m}/index.html'
DAY_ARCHIVE_SAVE_AS = '{date:%Y}/{date:%m}/{date:%d}/index.html'

ARTICLE_PATHS = ['']
ARTICLE_URL = '{date:%Y}/{date:%m}/{date:%d}/{slug}.html'
ARTICLE_SAVE_AS = ARTICLE_URL

PAGE_PATHS = ['_pages']
PAGE_URL = '{slug}/'
PAGE_SAVE_AS = '{slug}/index.html'

AUTHORS_SAVE_AS = ''
AUTHOR_SAVE_AS = ''

CATEGORIES_SAVE_AS = ''
CATEGORY_SAVE_AS = 'categories/{slug}/index.html'
CATEGORY_URL = 'categories/{slug}/'

TAGS_SAVE_AS = ''
TAG_SAVE_AS = 'tags/{slug}/index.html'
TAG_URL = 'tags/{slug}/'

STATIC_PATHS = ['_static']
EXTRA_PATH_METADATA = {
    '_static/.htaccess': {'path': '.htaccess'},
    '_static/robots.txt': {'path': 'robots.txt'},
    '_static/favicon.ico': {'path': 'favicon.ico'},
}

# Feed generation is usually not desired when developing
FEED_ALL_ATOM = None
CATEGORY_FEED_ATOM = None
TRANSLATION_FEED_ATOM = None
AUTHOR_FEED_ATOM = None
AUTHOR_FEED_RSS = None
