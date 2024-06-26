---
title: JQ Playground
level: 2
flag: FLAG{jqj6jqjqjqjqjqj6jqjqjqjqj6jqjqjq}
writer: hi120ki
---

# JQ Playground

## 問題文

Let's use JQ!

JQを使いこなそう！

http://chal-lz56g6.wanictf.org:8000/

---

NOTE: In case the server is down, try the following backup server. 上記のサーバーが正しく動作していない場合は、次のバックアップサーバーを使用してください。

`chal-ywu5dn.wanictf.org`

---

## 解法

<https://gtfobins.github.io/gtfobins/jq/> こちらの記事から着想を得た問題です。

<https://gtfobins.github.io/gtfobins/jq/> This problem was inspired by this article.

jqコマンドは`R` `r`オプションを指定した上でフィルターとして`.`を指定し、読み込むファイルを指定することでファイルの中身を取得することが可能

Jq command can read the contents of a file by specifying the `R` and `r` options, setting the filter to `.`, and specifying the file to be read.

```
jq ''-Rr . /flag'' test.json
```

```
'-Rr . /flag'
```

さらにワイルドカードを使い`/flag`を`/*`に置き換えることで、より文字列を短くすることができます

Furthermore, you can shorten the string more by replacing  `/flag` with `/*`.

```
'-Rr . /*'
```

さらにrオプションを削除することで、より短くすることができます

Furthermore, by removing the `r` option, You can shorten it even more.

```
' -R /*'
```
