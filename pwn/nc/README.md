---
title: nc
level: 1
flag: FLAG{th3_b3ginning_0f_th3_r0ad_to_th3_pwn_p1ay3r}
writer: mn1
---

# nc

## 問題文
pwn問題はnc(net cat)コマンドを使って問題サーバに接続することがよくあります。ncの使い方を覚えておきましょう

下記コマンドをshellで実行することで問題サーバに接続することが出来ます。接続先で問題を解き、フラグを獲得してください　

Pwn challenges often require connecting to the challenge server using the nc (netcat) command. It's important to learn how to use nc.

You can connect to the challenge server by executing the following command in your shell. Solve the problem at the connection point and obtain the flag.

`nc chal-lz56g6.wanictf.org 9003`

---

NOTE: In case the server is down, try the following backup server. 上記のサーバーが正しく動作していない場合は、次のバックアップサーバーを使用してください。

`chal-ywu5dn.wanictf.org`

---

## 解法

ncして計算問題を解くとsystem(/bin/sh)が呼び出され、フラグを取得することが出来ます。
You can solve the math problem when you connect problem server using nc chal-lz56g6.wanictf.org 9003. After solve the problem, you can get FLAG.