---
title: "tiny_usb"
level: 1
flag: "FLAG{hey_i_just_bought_a_usb}"
writer: "Mikka"
---

## 問題文

USB が狭い

What a small usb!

## 解法

iso ファイルが渡されるので，任意の方法でファイルを取り出す．

mount や foremost，バイナリから 0 でない箇所を保存，その他様々な方法が使用できる．

WSL や Linux を使用する場合，foremost コマンドを使用するのが一番早い．

```
foremost chal_tiny_usb.iso
```

An ISO file is provided, and files need to be extracted using any method.

Options such as mounting, using the 'foremost' command, saving non-zero sections from the binary, among other various methods can be used.

If using WSL or Linux, the 'foremost' command is the quickest way.

```
foremost chal_tiny_usb.iso
```
