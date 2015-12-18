Gehirn DNS API 仕様
===================

はじめに
--------

このドキュメントは、\ `ゲヒルン株式会社`_\ が Public Preview で提供している `Gehirn DNS`_ の API 仕様を説明するものです。
ただし、これはゲヒルン株式会社として提供する公式なドキュメントではなく、あくまでも個人として、公式のドキュメントが公開されるまでの期間、暫定的に公開するものです。

また、ここで説明する仕様は、\ `2015年12月21日17時に予定しているメンテナンス <http://support.gehirn.jp/information/maintenance/2015/12/17/1185/>`__\ が終了した後に有効となります。

API 仕様
--------

Root endpoint
   https://api.gis.gehirn.jp/dns/v1
Content-Type
   application/json
Authentication
   API キーを用いた Basic 認証

認証
~~~~

すべてのリクエストで認証が必要となります。
認証様式は、 API キーのトークンを usreid 、シークレットセールを password とする Basic 認証です。

API キーは\ `コントロールパネル`_\ から発行することができます。
また、リクエストに用いる API キーには適切な権限が設定されている必要があります。

Gehirn DNS の操作で必要になる権限は [DNS] -> [全般管理] -> [ゾーン] と [DNS] -> [ドメイン] -> [操作対象のゾーン] で、それぞれ読み取り以上が必要です。
読み取りは HTTP Verb の GET に、フルアクセスは HTTP Verb の GET に加え、 POST 、 PUT 、そして DELETE に対応します。

リクエスト
~~~~~~~~~~

API を用いて行える操作は以下に示すとおりです。
行いたい操作に合わせて、 Root endpoint と当該 Path からなるエンドポイントに対して、当該 HTTP Verb を用いたリクエストを行って下さい。

リクエストボディが必要な場合は、当該 JSON Schema に適合する形式の JSON を送信して下さい。
なお、この場合 `Content-Type: application/json` リクエストヘッダーが必要となります。

リクエスト例
^^^^^^^^^^^^

.. code-block:: http

   POST /dns/v1/zones HTTP/1.1
   Host: api.gis.gehirn.jp
   Content-Type: application/json
   Authorization: Basic dG9rZW46c2VjcmV0

   {
       "name": "example.com"
   }

ゾーン
~~~~~~

JSON Schema
^^^^^^^^^^^

.. code-block:: javascript

   {
     "type": "object",
     "properties": {
       "id": {
         "minLength": 36,
         "type": "string",
         "maxLength": 36
       },
       "name": {
         "minLength": 4,
         "type": "string",
         "maxLength": 256
       }
     }
   }


+------------+--------------------------+
| フィールド | 意味                     |
+============+==========================+
| id         | zone を特定する一意な ID |
+------------+--------------------------+
| name       | ドメイン名               |
+------------+--------------------------+

ゾーンの一覧の取得
^^^^^^^^^^^^^^^^^^

Path
   `/zones`
HTTP Verb
   GET
Request Body
   不要

ゾーンの作成
^^^^^^^^^^^^

Path
   `/zones`
HTTP Verb
   POST
Request Body
   必要

ゾーンの取得
^^^^^^^^^^^^

Path
   `/zones/:zone_id`
HTTP Verb
   GET
Request Body
   不要

ゾーンの削除
^^^^^^^^^^^^

Path
   `/zones/:zone_id`
HTTP Verb
   DELETE
Request Body
   不要

バージョン
~~~~~~~~~~

JSON Schema
^^^^^^^^^^^

.. code-block:: javascript

   {
     "type": "object",
     "properties": {
       "id": {
         "minLength": 36,
         "type": "string",
         "maxLength": 36
       },
       "name": {
         "minLength": 1,
         "type": "string",
         "maxLength": 255
       }
     }
   }

+------------+-------------------------------+
| フィールド | 意味                          |
+============+===============================+
| id         | バージョンを特定する一意な ID |
+------------+-------------------------------+
| name       | 任意のバージョン名            |
+------------+-------------------------------+

バージョン一覧の取得
^^^^^^^^^^^^^^^^^^^^

Path
   `/zones/:zone_id/versions`
HTTP Verb
   GET
Request Body
   不要

バージョンの作成
^^^^^^^^^^^^^^^^

Path
   `/zones/:zone_id/versions`
HTTP Verb
   POST
Request Body
   必要

バージョンの取得
^^^^^^^^^^^^^^^^

Path
   `/zones/:zone_id/versions/:version_id`
HTTP Verb
   GET
Request Body
   不要

バージョンの編集
^^^^^^^^^^^^^^^^

Path
   `/zones/:zone_id/versions/:version_id`
HTTP Verb
   PUT
Request Body
   必要

バージョンの削除
^^^^^^^^^^^^^^^^

Path
   `/zones/:zone_id/versions/:version_id`
HTTP Verb
   DELETE
Request Body
   不要

レコードセット
~~~~~~~~~~~~~~

JSON Schema
^^^^^^^^^^^

.. code-block:: javascript

   {
     "type": "object",
     "properties": {
       "id": {
         "minLength": 36,
         "type": "string",
         "maxLength": 36
       },
       "name": {
         "minLength": 1,
         "type": "string",
         "maxLength": 256
       },
       "type": {
         "minLength": 1,
         "type": "string",
         "maxLength": 5
       },
       "enable_alias": {
         "type": "boolean"
       },
       "alias_to": {
         "minLength": 1,
         "type": "string",
         "maxLength": 256
       },
       "ttl": {
         "minimum": 30,
         "type": "integer",
         "maximum": 2147483647
       },
       "records": {
         "type": "array",
         "minItems": 1,
         "items": {
           "type": "object",
           "properties": {
             "prio": {
               "minimum": 0,
               "type": "integer",
               "maximum": 32767
             },

             "address": {
               "minLength": 3,
               "type": "string",
               "maxLength": 39
             },

             "cname": {
               "minLength": 1,
               "type": "string",
               "maxLength": 256
             },

             "exchange": {
               "minLength": 1,
               "type": "string",
               "maxLength": 256
             },

             "nsdname": {
               "minLength": 1,
               "type": "string",
               "maxLength": 256
             },

             "target": {
               "minLength": 1,
               "type": "string",
               "maxLength": 256
             },
             "port": {
               "minimum": 0,
               "type": "integer",
               "maximum": 65535
             },
             "weight": {
               "minimum": 0,
               "type": "integer",
               "maximum": 65535
             },

             "data": {
               "minLength": 1,
               "type": "string",
               "maxLength": 64000
             }
           }
         }
       }
     }
   }

+------------------+---------------------------------------------------+--------------------------------+
| フィールド       | 意味                                              | リクエスト時要否               |
+==================+===================================================+================================+
| id               | レコードセットを特定する一意な ID                 | 不要                           |
+------------------+---------------------------------------------------+--------------------------------+
| name             | ホストネーム                                      | 必要                           |
+------------------+---------------------------------------------------+--------------------------------+
| type             | レコードタイプ (A, AAAA, CNAME, MX, NS, SRV, TXT) | 必要                           |
+------------------+---------------------------------------------------+--------------------------------+
| enable_alias     | エイリアス機能利用                                | 必要                           |
+------------------+---------------------------------------------------+--------------------------------+
| alias_to         | エイリアス先 (エイリアス機能利用時)               | enable_alias が true の時のみ  |
+------------------+---------------------------------------------------+--------------------------------+
| ttl              | TTL                                               | enable_alias が false の時のみ |
+------------------+---------------------------------------------------+--------------------------------+
| records          | レコードのリスト                                  | enable_alias が false の時のみ |
+------------------+---------------------------------------------------+--------------------------------+
| records.prio     | Priority                                          | type が MX または SRV の時のみ |
+------------------+---------------------------------------------------+--------------------------------+
| records.address  | IPv4 または IPv6 アドレス                         | type が A または AAAA の時のみ |
+------------------+---------------------------------------------------+--------------------------------+
| records.cname    | CNAME                                             | type が CNAME の時のみ         |
+------------------+---------------------------------------------------+--------------------------------+
| records.exchange | メールサーバーのドメインネーム                    | type が MX の時のみ            |
+------------------+---------------------------------------------------+--------------------------------+
| records.nsdname  | ネームサーバーのドメインネーム                    | type が NS の時のみ            |
+------------------+---------------------------------------------------+--------------------------------+
| records.target   | ターゲットのドメインネーム                        | type が SRV の時のみ           |
+------------------+---------------------------------------------------+--------------------------------+
| records.port     | ターゲットのポート番号                            | type が SRV の時のみ           |
+------------------+---------------------------------------------------+--------------------------------+
| records.weight   | ターゲットの重み                                  | type が SRV の時のみ           |
+------------------+---------------------------------------------------+--------------------------------+
| records.data     | TXT データ                                        | type が TXT の時のみ           |
+------------------+---------------------------------------------------+--------------------------------+

レコードセット一覧の取得
^^^^^^^^^^^^^^^^^^^^^^^^

Path
   `/zones/:zone_id/versions/:version_id/records`
HTTP Verb
   GET
Request Body
   不要

レコードセットの作成
^^^^^^^^^^^^^^^^^^^^

Path
   `/zones/:zone_id/versions/:version_id/records`
HTTP Verb
   POST
Request Body
   必要

レコードセットの取得
^^^^^^^^^^^^^^^^^^^^

Path
   `/zones/:zone_id/versions/:version_id/records/:record_id`
HTTP Verb
   GET
Request Body
   不要

レコードセットの編集
^^^^^^^^^^^^^^^^^^^^

Path
   `/zones/:zone_id/versions/:version_id/records/:record_id`
HTTP Verb
   PUT
Request Body
   必要

レコードセットの削除
^^^^^^^^^^^^^^^^^^^^

Path
   `/zones/:zone_id/versions/:version_id/records/:record_id`
HTTP Verb
   DELETE
Request Body
   不要

.. _`ゲヒルン株式会社`: http://www.gehirn.co.jp/
.. _`Gehirn DNS`: https://www.gehirn.jp/gis/dns.html
.. _`コントロールパネル`: https://gis.gehirn.jp/

.. author:: default
.. categories:: none
.. tags:: Gehirn
.. comments::
