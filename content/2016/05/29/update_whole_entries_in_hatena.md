---
title: はてなブログの AtomPub API を使って記事を更新する
date: 2016-05-29T12:00:00+09:00
postCategory: Code
postTags: [AtomPub, Python]
---

## はじめに

こんばんは、 yosida95 です。
昨日お知らせとおり、ぼくのブログをはてなブログから yosida95.com へと移動しました。

これに伴って、ぼくのすべてのはてなブログエントリの本文を、対応する yosida95.com の新しい URL へと誘導する内容に差し替えました。
はてなブログでは138記事を公開していましたが、このすべての記事を差し替えることは人間のやる仕事ではありません。

幸いはてなブログでは [AtomPub に則った API](http://developer.hatena.ne.jp/ja/documents/blog/apis/atom) を公開してくれているので、これを使って自動で差し替えることができました。
雑に書いたコードですが、そのまま捨てることももったいないので、同じことをしたい人やはてなブログ AtomPub API の具体的な利用例になればと思い公開することにします。

## はてなブログのすべての記事を自動で更新する

{% raw %}
```python
# -*- coding: utf-8 -*-

'''
Copyright (c) 2016, Kohei YOSHIDA <https://yosida95.com/>. All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

    * Redistributions of source code must retain the above copyright
      notice, this list of conditions and the following disclaimer.
    * Redistributions in binary form must reproduce the above copyright
      notice, this list of conditions and the following disclaimer in the
      documentation and/or other materials provided with the distribution.
    * Neither the name of the copyright holder nor the names of its
      contributors may be used to endorse or promote products derived from
      this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
'''

import copy
import hashlib
import os
import re
from base64 import b64encode
from datetime import datetime
from urllib.parse import (
    urljoin,
    urlparse,
)
from xml.etree import ElementTree

import requests

ENDPOINT_PREFIX = 'https://blog.hatena.ne.jp/{はてなID}/{ブログID}/atom/'
YOUR_HATENA_ID = 'yosida95'
YOUR_API_KEY = 'dummyapikey'
PATH_TO_APACHE_REWRITE_RULES = '/var/www/blog.yosida95.com/.htaccess'
nsmap = dict(atom='http://www.w3.org/2005/Atom',
             app='http://www.w3.org/2007/app')


def build_rewrite_rules(filename):
    rewrite_rules = []
    with open(filename) as fh:
        for entry in fh:
            tokens = entry.split()
            if len(tokens) == 0 or tokens[0] != 'RewriteRule':
                continue
            source, dest = tokens[1:3]
            re_source = re.compile(source)
            rewrite_rules.append((re_source, dest.replace('$', '\\')))
    return rewrite_rules


def find_redirect_to(entry_url, rewrite_rules):
    entry_path = urlparse(entry_url).path
    redirect_to = None
    for (pattern, rewrite_to) in rewrite_rules:
        if not pattern.match(entry_path):
            continue
        redirect_to = pattern.sub(rewrite_to, entry_path)
        break
    return redirect_to


def make_new_content(entry_url):
    template = '''<p>この記事は yosida95.com に移動しました。<br>
新しい URL は <a href="{entry_url}">{entry_url}</a> です。</p>

<p>お手数をお掛けしますが、ブックマークの付け替えをお願いします。</p>

<!-- more -->
<script type="text/javascript">
    location.href = "{entry_url}" + location.search;
</script>'''
    return template.format(entry_url=entry_url)


def make_wsse_token(userid, password):
    nonce = hashlib.sha1(os.urandom(16)).digest()
    created = datetime.utcnow().strftime('%Y-%m-%dT%H:%M:%SZ')
    credential = b''.join((nonce, created.encode('ascii'),
                           password.encode('ascii')))
    digest = hashlib.sha1(credential).digest()

    tmpl = ('UsernameToken Username="{uname}",'
            'PasswordDigest="{digest}",'
            'Nonce="{nonce}",'
            'Created="{created}"')
    return tmpl.format(uname=userid,
                       digest=b64encode(digest).decode('ascii'),
                       nonce=b64encode(nonce).decode('ascii'),
                       created=created)


def update_entry_element(entry, new_content):
    inheritable_elements = ('{{{atom}}}title',
                            '{{{atom}}}category',
                            '{{{app}}}control')
    inheritable_elements = list(map(lambda name: name.format(**nsmap),
                                    inheritable_elements))
    content_tag = '{{{atom}}}content'.format(**nsmap)

    new_entry = ElementTree.Element(entry.tag)
    for node in entry:
        if node.tag == content_tag:
            node = copy.deepcopy(node)
            node.set('type', 'text/html')
            node.text = new_content
            new_entry.append(node)
        elif node.tag in inheritable_elements:
            new_entry.append(copy.deepcopy(node))
    return new_entry


def update_entry(entry, rewrite_rules, userid, password):
    is_draft = entry.find('app:control/app:draft', nsmap).text == 'yes'
    if is_draft:
        return

    entry_id = entry.find('atom:id', nsmap).text
    edit_url = entry.find('atom:link[@rel="edit"]', nsmap).attrib['href']

    entry_url = entry.find('atom:link[@rel="alternate"]', nsmap).attrib['href']
    redirect_to = find_redirect_to(entry_url, rewrite_rules)

    new_content = make_new_content(redirect_to)
    new_entry = update_entry_element(entry, new_content)

    headers = {'X-WSSE': make_wsse_token(userid, password),
               'Content-Type': 'application/atom+xml;type=entry'}
    body = ElementTree.tostring(new_entry, encoding='utf8')
    requests.put(edit_url, headers=headers, data=body)

    print(','.join((entry_id, entry_url, redirect_to)))


def iter_collections(userid, password):
    collection_url = urljoin(ENDPOINT_PREFIX, './entry')
    while collection_url:
        headers = {'X-WSSE': make_wsse_token(userid, password)}
        resp = requests.get(collection_url, headers=headers)
        root = ElementTree.fromstring(resp.text)
        yield root

        next_ = root.find('atom:link[@rel="next"]', nsmap)
        if next_ is None:
            break
        collection_url = next_.attrib['href']


def main(userid, password, rewrite_rule_file):
    rewrite_rules = build_rewrite_rules(rewrite_rule_file)

    for collection in iter_collections(userid, password):
        for entry in collection.findall('atom:entry', nsmap):
            update_entry(entry, rewrite_rules, userid, password)
            return


if __name__ == '__main__':
    main(YOUR_HATENA_ID, YOUR_API_KEY,
         PATH_TO_APACHE_REWRITE_RULES)
```
{% endraw %}

### 前提

1. Python 3 系でしか動きません
2. [requests](https://pypi.python.org/pypi/requests) のインストールが必要です

### 使い方

冒頭で定義されている以下の変数にそれぞれ適切な値を代入してください。

- ENDPOINT_PREFIX
- YOUR_HATENA_ID
- YOUR_API_KEY
- PATH_TO_APACHE_REWRITE_RULES

PATH_TO_APACHE_REWRITE_RULES には移転前のブログパスから移転後の URL へ転送する RewriteRule を対応づけた .htaccess へのパスを指定して下さい。
もっとも、新しいブログの URL を本文に記載する必要がない場合は、 `update_entry` 関数内の `make_new_content` の呼び出し周辺をいい感じに改変してあげて下さい。

本文は `make_new_content` 関数を改変することでいい感じになります。

### ライセンス

ソースコードに記載の通り修正 BSD ライセンスで提供するので、このライセンスが許す範囲内で自由に実行、改変、再配布して下さい。

## 最後に

今後ともよろしくお願いします。
