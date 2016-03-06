Gehirn DNS API 仕様
===================

追記
----

`ゲヒルン株式会社`_\ として公式に Gehirn DNS を含む `Gehrin Web Services API Documentation`_ を公開したので、今後はそちらを参照して下さい。
今後、この非公式ドキュメントはメンテナンスしません。

.. _`Gehirn Web Services API Documentation`: https://support.gehirn.jp/apidocs/

はじめに
--------

このドキュメントは、\ `ゲヒルン株式会社`_\ が Public Preview で提供している `Gehirn DNS`_ の API 仕様を説明するものです。
ただし、これはゲヒルン株式会社として提供する公式なドキュメントではなく、あくまでも個人として、公式のドキュメントが公開されるまでの期間、暫定的に公開するものです。

また、ここで説明する仕様は、\ `2015年12月21日17時に予定しているメンテナンス <http://support.gehirn.jp/information/maintenance/2015/12/17/1185/>`__\ が終了した後に有効となります。

.. more::

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
       },
       "current_version_id": {
         "minLength": 36,
         "type": "string",
         "maxLength": 36
       },
       "current_version": {
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
     }
   }


+--------------------+---------------------------------+------------------+
| フィールド         | 意味                            | リクエスト時要否 |
+====================+=================================+==================+
| id                 | ゾーンを特定する一意な ID       | 不要             |
+--------------------+---------------------------------+------------------+
| name               | ドメイン名                      | 必要             |
+--------------------+---------------------------------+------------------+
| current_version_id | 現在アクティブなバージョンの ID | 不要             |
+--------------------+---------------------------------+------------------+
| current_version    | 現在アクティブなバージョン      | 不要             |
+--------------------+---------------------------------+------------------+

ゾーンの作成
^^^^^^^^^^^^

Path
   `/zones`
HTTP Verb
   POST
Request Body
   必要

**リクエスト例**

.. code-block:: http

   POST /dns/v1/zones HTTP/1.1
   Host: api.gis.gehirn.jp
   Content-Type: application/json
   Authorization: Basic dG9rZW46c2VjcmV0

   {
       "name": "yaml.jp"
   }

**レスポンス例**

.. code-block:: http

   HTTP/1.1 200 OK
   Server: nginx
   Date: Fri, 18 Dec 2015 10:41:01 GMT
   Content-Type: application/json; charset=UTF-8
   Content-Length: 388

   {
     "id": "92e52aab-81ac-4c87-b659-b7b36e05cb7f",
     "name": "yaml.jp",
     "current_version_id": "234b6f0e-8b64-4cd9-8647-16cd26133266",
     "current_version": {
       "id": "234b6f0e-8b64-4cd9-8647-16cd26133266",
       "editable": true,
       "name": "\u6700\u521d\u306e\u30d0\u30fc\u30b8\u30e7\u30f3",
       "created_at": "2015-03-05T10:49:04Z",
       "last_modified_at": "2015-03-05T10:49:04Z"
     }
   }

ゾーンのリストの取得
^^^^^^^^^^^^^^^^^^^^

Path
   `/zones`
HTTP Verb
   GET
Request Body
   不要

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

+------------------+-------------------------------+------------------+
| フィールド       | 意味                          | リクエスト時要否 |
+==================+===============================+==================+
| id               | バージョンを特定する一意な ID | 不要             |
+------------------+-------------------------------+------------------+
| name             | 任意のバージョン名            | 必要             |
+------------------+-------------------------------+------------------+
| editable         | 編集可否                      | 不要             |
+------------------+-------------------------------+------------------+
| created_at       | バージョン作成時刻            | 不要             |
+------------------+-------------------------------+------------------+
| last_modified_at | バージョン最終更新時刻        | 不要             |
+------------------+-------------------------------+------------------+


バージョンの作成
^^^^^^^^^^^^^^^^

Path
   `/zones/:zone_id/versions`
HTTP Verb
   POST
Request Body
   必要

**リクエスト例**

.. code-block:: http

   POST /dns/v1/zones/234b6f0e-8b64-4cd9-8647-16cd26133266/versions HTTP/1.1
   Host: api.gis.gehirn.jp
   Content-Type: application/json
   Authorization: Basic dG9rZW46c2VjcmV0

   {
       "name": "新しいバージョン"
   }

**レスポンス例**

.. code-block:: http

   HTTP/1.1 200 OK
   Server: nginx
   Date: Fri, 18 Dec 2015 10:41:01 GMT
   Content-Type: application/json; charset=UTF-8
   Content-Length: 218

   {
     "id": "f66504b0-bb65-4766-9d7c-18c4e8406071",
     "editable": true,
     "name": "\u65b0\u3057\u3044\u30d0\u30fc\u30b8\u30e7\u30f3",
     "created_at": "2015-12-18T10:49:13Z",
     "last_modified_at": "2015-12-18T10:49:13Z"
   }

バージョンリストの取得
^^^^^^^^^^^^^^^^^^^^^^

Path
   `/zones/:zone_id/versions`
HTTP Verb
   GET
Request Body
   不要

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

レコードセットの作成
^^^^^^^^^^^^^^^^^^^^

Path
   `/zones/:zone_id/versions/:version_id/records`
HTTP Verb
   POST
Request Body
   必要

**リクエスト例**

.. code-block:: http

   POST /dns/v1/zones/234b6f0e-8b64-4cd9-8647-16cd26133266/versions/f66504b0-bb65-4766-9d7c-18c4e8406071/records HTTP/1.1
   Host: api.gis.gehirn.jp
   Content-Type: application/json
   Authorization: Basic dG9rZW46c2VjcmV0

   {
     "name": "yaml.jp.",
     "ttl": 300,
     "type": "A",
     "enable_alias": false,
     "records": [
       {
         "address":"192.0.2.10"
       },
       {
         "address":"192.0.2.11"
       }
     ]
   }

**レスポンス例**

.. code-block:: http

   HTTP/1.1 200 OK
   Server: nginx
   Date: Fri, 18 Dec 2015 10:41:01 GMT
   Content-Type: application/json; charset=UTF-8
   Content-Length: 218

   {
     "id": "e590d62a-3676-4b08-832a-a1fdd6dfefdf",
     "name": "yaml.jp.",
     "type": "A",
     "enable_alias": false,
     "ttl": 300,
     "records": [
       {
         "address": "192.0.2.10"
       },
       {
         "address": "192.0.2.11"
       }
     ]
   }

レコードセットリストの取得
^^^^^^^^^^^^^^^^^^^^^^^^^^

Path
   `/zones/:zone_id/versions/:version_id/records`
HTTP Verb
   GET
Request Body
   不要

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
