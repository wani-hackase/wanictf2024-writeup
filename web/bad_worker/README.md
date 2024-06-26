---
title: Bad_Worker # (必須) 問題名
level: 1          # (必須) 難易度 1:beginner ~ 5:very hard
flag: FLAG{pr0gr3ssiv3_w3b_4pp_1s_us3fu1}  # (必須) フラグ
writer: kaki005   # (必須) 作問者
---

# 問題名
Bad_Worker
## 問題文
オフラインで動くウェブアプリをつくりました。
We created a web application that works offline.
https://web-bad-worker-lz56g6.wanictf.org


## 解法　
- このアプリは[ServiceWorker](https://developer.mozilla.org/ja/docs/Web/API/ServiceWorker)を用いてウェブサイトのリソースをキャッシュすることでオフラインでも動作するようにしています。
  - `src/service-worker.js`を見ると、FLAG.txtを取得しようとしたときに、DUMMY.txtを代わりに取得していることがわかります。
  - 開発者ツールでServiceworkerを無効にすることでFLAG.txtを取得できます。
- また curlコマンドでもFLAG.txtを直接取得することができます。

## Solution
- This application uses [ServiceWorker](https://developer.mozilla.org/en-US/docs/Web/API/ServiceWorker) to cache website resources so that it can work offline.
  - If you look at `src/service-worker.js`, you will find that when it tries to get FLAG.txt, it gets DUMMY.txt instead.
  - You can get FLAG.txt by disabling Serviceworker in the developer tools.
- Furthermore, you can also get FLAG.txt by executing curl command
