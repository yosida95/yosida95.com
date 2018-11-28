# -*- coding: utf-8 -*- #

AUTHOR = 'yosida95'
SITENAME = 'yosida95'
SITEURL = ''
RELATIVE_URLS = True

PATH = 'content'
PATH_METADATA = '(?P<date>\d{4}/\d{2}/\d{2})/(?P<slug>[^\.]*).rst'
USE_FOLDER_AS_CATEGORY = False
SUMMARY_MAX_LENGTH = 10
DEFAULT_PAGINATION = 5

THEME = './_themes/purity'
DIRECT_TEMPLATES = ['archives']

STATIC_PATHS = ['_static']
EXTRA_PATH_METADATA = {
    '_static/.htaccess': {'path': '.htaccess'},
    '_static/robots.txt': {'path': 'robots.txt'},
    '_static/favicon.ico': {'path': 'favicon.ico'}}

ARTICLE_PATHS = ['2009',
                 '2010', '2011', '2012', '2013', '2014', '2015',
                 '2016', '2017', '2018',
                 ]
ARTICLE_URL = '{date:%Y}/{date:%m}/{date:%d}/{slug}.html'
ARTICLE_SAVE_AS = ARTICLE_URL

INDEX_SAVE_AS = None

PAGE_PATHS = ['']
PAGE_EXCLUDES = ['_static']
PAGE_URL = '{slug}/'
PAGE_SAVE_AS = '{slug}/index.html'

ARCHIVES_SAVE_AS = 'archives.html'

TIMEZONE = 'Asia/Tokyo'

DEFAULT_LANG = 'ja'

# Feed generation is usually not desired when developing
FEED_ALL_ATOM = None
CATEGORY_FEED_ATOM = None
TRANSLATION_FEED_ATOM = None
AUTHOR_FEED_ATOM = None
AUTHOR_FEED_RSS = None

PLUGIN_PATHS = ['../plugins/pelican-plugins']
PLUGINS = ['pelican_alias']

JINJA_ENVIRONMENT = {'lstrip_blocks': False,
                     'trim_blocks': False}

PYGMENTS_RST_OPTIONS = {}
