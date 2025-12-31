---
title: yosida95 の 2025 年まとめ
postCategory: Diary
postTags: [Year in Review]
date: 2025-12-31T12:00:00+09:00
---

こんばんは、 yosida95 です。1年ぶり16回目のまとめをやります。人生の半分以上がこのブログにまとめられている。

## ことし買ったもの

ことしはパソコンを3台買った。

### GMKTec NucBox G3 Plus

いわゆるミニ PC を初めて買ったみた。 Intel N150 と NVMe SSD が思いのほか快適で、 Ubuntu のインストールはトイレに行っている間に終わるし、 Docker も問題なく動く。

おもちゃとして買った当初の目論見から外れ、 Unbound や Grafana 、 Tailscale の Subnet Router や Exit Node なんかを 24/7 で動かしている。 2018 年からこの用途で使っていた Raspberry Pi 2 Modal B は退役させた。

<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">my new gear… / おもちゃとして買ってみたけれど、 Intel N150 が思っていたよりだいぶ速いし NVMe SSD も載っていて、トイレに行って戻ってくる間に Ubuntu のインストールが終わっていた <a href="https://t.co/rf6KggLTv8">pic.twitter.com/rf6KggLTv8</a></p>&mdash; よしだ (@yosida95) <a href="https://twitter.com/yosida95/status/1963541585762623630?ref_src=twsrc%5Etfw">September 4, 2025</a></blockquote>

### Lenovo IdeaPad Flex 3 Chrome

ChromeOS デバイスに触ってみたくて買った。 ChromeOS は2012年くらいからなんとなく気になっていたけれど、 GIGA スクール構想で広く使われるようになって興味を増した。

買うなら Chromebox がよかったんだけれどもはや流通しておらず、 Chromebook の新品をヨドバシカメラの店頭で眺めたところ思いのほか値が張ったので、初めて中古のパソコンを買った。

<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">Windows 10 のサポート終了で Chromebook に乗り換えるお客さんが多いらしく、わたしもそういう体で説明された</p>&mdash; よしだ (@yosida95) <a href="https://twitter.com/yosida95/status/1925822238650962323?ref_src=twsrc%5Etfw">May 23, 2025</a></blockquote>

この記事を書いているときに型番で検索したら 180° 開いている写真を見つけ、試してみたら本当に開いてしまった。ただの液晶だと思ったディスプレイがタッチパネルだったことにも気づいた。そんな感じでまだおもちゃにしきれていない。

KVM ベースの Linux VM が使える [Crostini](https://www.chromium.org/chromium-os/developer-library/guides/containers/containers-and-vms/) は魅力的で、普段のソフトウェア開発では Linux ボックスに SSH して Vim でコードを書いているわたしにとって、 Chromebook 1台ですべてを済ませられるのではという誘惑がある。

でも、そこまでうまい話はなかった。 Intel N100 / 4 GB / eMMC の構成で VM を快適に使うのは難しいし、ターミナルエミュレータの自由度や好みの問題もある。そもそも今のスタイルに落ち着いたのは、 OS X 上の VirtualBox で Linux を動かしても性能が出ないし、 `/dev/random` のエントロピーもなかなかたまらなくて実機を求めた経緯がある。 Crostini は UI との融合が見事で、ワンクリックで Linux VM が立ち上がってくる利便性は高いのだけれど。

起動・終了の速さや筐体の頑丈さ、ユーザーに不用意に変なことをさせない堅牢さなどは、なるほど中高生に持たせるのに向いているし、外回りをする営業職や Web ベースの定型作業、コールセンターなどで重宝がられることが理解できる。安いし。 Chromebook には夢と実用性がある。

<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">E3-1240 は当時高校生のわたしにとっては強い CPU のつもりだったのだけれど、 GIGA スクールで高校生がみんな持っている Chromebook にも載っている N100 の方が PassMark のスコアが上回っていてすごくいい時代</p>&mdash; よしだ (@yosida95) <a href="https://twitter.com/yosida95/status/1967127425202217132?ref_src=twsrc%5Etfw">September 14, 2025</a></blockquote>

### ThinkCentre neo

日常のソフトウェア開発のために買ったけれどほとんど触れてない。手癖で Ubuntu Server を入れたけれど、私物なんだから雑に壊してもよく、 Arch Linux にしておけばよかったと後悔している。 Arch Linux を入れ直すとして、慣れている GRUB 2 にするか systemd-boot にするか悩むけれど、自宅から持ち出さないしひとり暮らしだし別にセキュアブートはいらないよね、みたいなところで放置している。

とりあえず PiKVM デバイスを繋いでいるけれど、 ThinkCentre の筐体から電源を引き出せないし、そもそも外出しないしで無用の長物と化している。

### NETGEAR WAX 210

10年以上使っていたコンシューマ向けの NEC Aterm Wi-Fi AP を NETGEAR Bussiness の WAX 210 に入れ替えた。スペックと値段だけ見て決めたけれど、明確に失敗だった。まったく安定せず、多いと日に2度くらい、不定期にクラッシュする。

コイル鳴きもひどいんだけれど、 YouTube やdアニメストアで動画をシークする度に音がなるのが少しおもしろい。

<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">…… Wi-Fi AP 壁面設置ヨシ！ <a href="https://t.co/L8G9vDCLhQ">pic.twitter.com/L8G9vDCLhQ</a></p>&mdash; よしだ (@yosida95) <a href="https://twitter.com/yosida95/status/1962393553914806317?ref_src=twsrc%5Etfw">September 1, 2025</a></blockquote>
<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">NETGEAR WAX210 の WPA3-Personal Transition Mode と iOS 18.6 、ひょっとして相性問題がないだろうか　WPA2 で運用していた AP を置き換えて SSID はそのまま PSK を変更した上で Transition Mode を有効化、 iOS の Wi-Fi 設定を削除して再設定したら不安定になった　WPA2 固定で安定する</p>&mdash; よしだ (@yosida95) <a href="https://twitter.com/yosida95/status/1967802229643219040?ref_src=twsrc%5Etfw">September 16, 2025</a></blockquote>
<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">NETGEAR の Wi-Fi AP のファームウェア、ログにタイポがあるわね <a href="https://t.co/okojtIChBp">pic.twitter.com/okojtIChBp</a></p>&mdash; よしだ (@yosida95) <a href="https://twitter.com/yosida95/status/1999422859752669642?ref_src=twsrc%5Etfw">December 12, 2025</a></blockquote>

### そのほか

ほかにも Dyson V12 や空気清浄機、布団掃除機、サーキュレーターなどを買ったけれど、買い替えだったり日常の買い物という感じだったりで特筆することはない。

<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">空気清浄機を12年ぶりに買い替えたぞ　加湿器は別途あるから加湿機能のないものを選んだらコンパクトでよい　古い機種では垂直に設置されていたフィルターが、新しい機種では水平に置かれていて高さが抑えられている <a href="https://t.co/JCoMDfphl7">pic.twitter.com/JCoMDfphl7</a></p>&mdash; よしだ (@yosida95) <a href="https://twitter.com/yosida95/status/1962407219502452901?ref_src=twsrc%5Etfw">September 1, 2025</a></blockquote>

## イベント

2024年のまとめで、いつまで心身が健康でいられるか分からないのだから行きたいイベントには行こうみたいな話をした。おととし低温障害型感音難聴を発症したこともあるし。

そんなことで1月4日から『HAYAMI SAORI Orchestra Concert 2025』に行ったり、竹達彩奈さんの体調不良で振替になった petit milady 『le diner de cons～淑女たちの秘密の打ち上げ～』に行ったりした。

<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">HAYAMI SAORI Orchestra Concert 2025 (@ カルッツかわさき (川崎市スポーツ・文化総合センター) in 川崎市, 神奈川県) <a href="https://t.co/55MCJDT4eA">https://t.co/55MCJDT4eA</a></p>&mdash; よしだ (@yosida95) <a href="https://twitter.com/yosida95/status/1875472465264308372?ref_src=twsrc%5Etfw">January 4, 2025</a></blockquote>
<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">『le diner de cons～淑女たちの秘密の打ち上げ～』 <a href="https://twitter.com/hashtag/petitmilady?src=hash&amp;ref_src=twsrc%5Etfw">#petitmilady</a> (@ 関内ホール 大ホール in 横浜市, 神奈川県) <a href="https://t.co/2lYihVPgMj">https://t.co/2lYihVPgMj</a></p>&mdash; よしだ (@yosida95) <a href="https://twitter.com/yosida95/status/1888526699647910157?ref_src=twsrc%5Etfw">February 9, 2025</a></blockquote>

でもそれきりだった。喉元をすぎて熱さを忘れたというのもあるかもしれないし、3月に超!A&G+が "拡充" して行きたいイベントがパタリと途絶えたからというのもあるかもしれない。

鷲崎健さんのことが好きで、鷲崎さんがことし始めたライブハウス・ラジオスタジオ『[トカトントン](https://tokatonton.bitfan.id/)』にも興味があるけれど、始まったばかりということだったりキャパシティの関係による換気といった部分だったりに不安がある。あと内輪感が強くて尻込みしている部分もある。とりあえずファンクラブには開設当日に入会した。

<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">めちゃかわ梱包の大ボリュームアルバムが届きました。5月中に届いていたけれどきょうになって開梱した <a href="https://twitter.com/hashtag/%E7%AC%AC%E4%BA%8C%E6%AC%A1%E9%B7%B2%E5%B4%8E%E5%A4%A7%E4%BD%9C%E6%88%A6?src=hash&amp;ref_src=twsrc%5Etfw">#第二次鷲崎大作戦</a> <a href="https://twitter.com/hashtag/%E3%81%88%E3%81%8F%E3%81%BC%E3%83%B6%E5%8E%9F%E9%A3%84%E5%A4%A2%E8%AD%9A?src=hash&amp;ref_src=twsrc%5Etfw">#えくぼヶ原飄夢譚</a> <a href="https://t.co/NgApPMRHgu">pic.twitter.com/NgApPMRHgu</a></p>&mdash; よしだ (@yosida95) <a href="https://twitter.com/yosida95/status/1934485922692907246?ref_src=twsrc%5Etfw">June 16, 2025</a></blockquote>

## 仕事

<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">登庁 (4年9ヶ月ぶり) (@ ゲヒルン 司令室 in 千代田区, 東京都) <a href="https://t.co/sy4YLHy0UY">https://t.co/sy4YLHy0UY</a></p>&mdash; よしだ (@yosida95) <a href="https://twitter.com/yosida95/status/1955915536967328027?ref_src=twsrc%5Etfw">August 14, 2025</a></blockquote>

感染症が広まって以降、ずっとリモートワークを続けていたけれど、8月、ついに物理出社した。

担当外だけれどリモートワークが寂しくてミーティングに混ざって Google Meet で茶々を入れていたら、なんやかんやあって担当者になってネットワーク機器のキッティングやセットアップの仕事が生まれた。それ以来、月に2度くらい出社している。

出社していなかった間に静脈認証器の更改があって、社歴が1番古いのにオフィスに入れない状態もついに終わりを迎えた。

パソコンを3台買ったのも自宅のネットワーク機器を入れ替えたのもことし後半だけれど、これは物理出社で刺激を受けて興味が外に向いたからというのが大きいはず。出社するとしないとでは情報量が大違いで、人が集まるオフィスそのものが福利厚生という感じがする。

ソフトウェア開発では、相変わらず完全仮想化、コンテナ仮想化、 DNS 、SMTP 、 Web PKI 、 OAuth 、 WebAuthn 、 HTTP 、 gRPC なんかの周りでパソコンカタカタをしている。

ただ、リリース以来、ずっと開発していた[レンタルサーバーサービスが終了](https://support.gehirn.jp/information/2025/01/15/rs2plus-sunset/)することが決まって、その対応にも迫られている。個人でも社内でもちゃんとドッグフーディングしていたので、何度も移行作業を繰り返している。このブログも移行しないといけないので寂しい。

## ことし寄付した先

ことし寄付した先は次の 4 団体です。

### がん研究会

[「グエー死んだンゴ」](https://x.com/nkym7856/status/1978053179700060502)ニキを[きっかけとするムーブメント](https://www.asahi.com/articles/ASTBQ1SXYTBQULLI00XM.html)に乗っかった。こういう祭りでは踊らないと損だと思っている。

寄付先の候補はいくつかあったけれど、江東区に住んでいて住民税で少し優遇されそうだったのでがん研究会を選んだ。それから叔母が以前ここのお世話になったと聞いた気がする。

<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">公益財団法人がん研究会のサーバーの時計が1分弱進んでいることがわかりました <a href="https://t.co/oVDSIgqmnE">https://t.co/oVDSIgqmnE</a> <a href="https://t.co/LCirtseoG2">pic.twitter.com/LCirtseoG2</a></p>&mdash; よしだ (@yosida95) <a href="https://twitter.com/yosida95/status/1979824849003630932?ref_src=twsrc%5Etfw">October 19, 2025</a></blockquote>

### 日本赤十字社

ことしも日本赤十字社に運営費を通常寄付した。2022年から毎年続けていたところ特別社員の称号を賜った。

<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">日本赤十字社特別社員3日目です <a href="https://t.co/N3INGzr21E">https://t.co/N3INGzr21E</a> <a href="https://t.co/KYxai6Em6M">pic.twitter.com/KYxai6Em6M</a></p>&mdash; よしだ (@yosida95) <a href="https://twitter.com/yosida95/status/1951212849457668126?ref_src=twsrc%5Etfw">August 1, 2025</a></blockquote>

### Mozilla Foundation

2024年のまとめで書いたとおり毎月 1,000 JPY ずつの定期寄付を続けて合計 12,000 JPY 寄付した。

### Wikimedia Foundation

Wikimedia Foundation への寄付は 2014 年から始め、 2018 年 7 月からは毎月の定額寄付を行っており、ことしもこれを継続しました。

## これまでのまとめ

- [2010 年](/2010/12/31/115758.html)
- [2011年が終わるね！！](/2011/12/31/235927.html)
- [2012 年にぼくがさせていただいたこと](/2013/01/01/005050.html)
- [yosida95 の2013年を振り返る](/2013/12/31/111207.html)
- [yosida95 の 2014 年まとめ](/2014/12/29/130000.html)
- [yosida95 の 2015 年まとめ](/2015/12/31/yearly_report.html)
- [yosida95 の 2016 年まとめ](/2016/12/31/yearly_report.html)
- [yosida95 の 2017 年まとめ](/2017/12/31/greetings.html)
- [yosida95 の 2018 年まとめ](/2018/12/31/year-in-review.html)
- [yosida95 の 2019 年まとめ](/2019/12/31/year-in-review.html)
- [yosida95 の 2020 年まとめ](/2020/12/31/year-in-review.html)
- [yosida95 の 2021 年まとめ](/2021/12/31/year-in-review.html)
- [yosida95 の 2022 年まとめ](/2022/12/31/year-in-review.html)
- [yosida95 の 2023 年まとめ](/2023/12/31/year-in-review.html)
- [yosida95 の 2024 年まとめ](/2024/12/31/year-in-review.html)
