# -*- coding: utf-8 -*-

import os

from fabric.api import (
    cd,
    env,
    run,
    task,
)
from fabric.contrib.project import rsync_project

HERE = os.path.dirname(os.path.abspath(__file__))


@task
def production():
    env.user = 'rs2p'
    env.hosts = ['8b66b542-8b34-476b-808a-68d85be08243.gehirn.ne.jp:22329']


@task
def deploy():
    rsync_project(local_dir=os.path.join(HERE, 'blog/html/'),
                  remote_dir='/var/www/yosida95.com/html/',
                  exclude=['.DS_Store'],
                  delete=True)
    with cd('/var/www/yosida95.com/html'):
        run('ln -s rss.html rss.xml')
        run('ln -s _static/googlecad1c35a95af6e0c.html .')
        run('ln -s _static/robots.txt .')
