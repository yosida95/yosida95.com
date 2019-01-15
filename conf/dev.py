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
DIRECT_TEMPLATES = ['archives']
SUMMARY_MAX_LENGTH = 10
DEFAULT_PAGINATION = 5
JINJA_ENVIRONMENT = {
    'lstrip_blocks': False,
    'trim_blocks': False,
}
PYGMENTS_RST_OPTIONS = {}

PLUGIN_PATHS = ['../plugins/pelican-plugins']
PLUGINS = ['pelican_alias']

INDEX_SAVE_AS = None

ARCHIVES_SAVE_AS = 'archives.html'

ARTICLE_PATHS = ['']
ARTICLE_URL = '{date:%Y}/{date:%m}/{date:%d}/{slug}.html'
ARTICLE_SAVE_AS = ARTICLE_URL

PAGE_PATHS = ['_pages']
PAGE_URL = '{slug}/'
PAGE_SAVE_AS = '{slug}/index.html'

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
