# -*- coding: utf-8 -*- #

import os
import sys
sys.path.append(os.path.dirname(os.path.abspath(__file__)))

from dev import *  # noqa

SITEURL = 'https://yosida95.com'
RELATIVE_URLS = False

# Feed generation is usually not desired when developing
FEED_ALL_ATOM = 'rss.xml'
CATEGORY_FEED_ATOM = None
TRANSLATION_FEED_ATOM = None
AUTHOR_FEED_ATOM = None
AUTHOR_FEED_RSS = None

DELETE_OUTPUT_DIRECTORY = True

GOOGLE_ANALYTICS = 'UA-15957452-1'
