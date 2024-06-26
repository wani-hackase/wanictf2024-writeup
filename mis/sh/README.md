---
title: sh
level: 3
flag: FLAG{use_she11check_0r_7he_unexpec7ed_h4ppens}
writer: ciffelia
---

# sh

> Guess?
>
> ```sh
> nc chal-lz56g6.wanictf.org 7580
> ```

## Solution

日本語版は英語版の後に記載しています。Japansese follows English.

You can obtain the flag with the following input:

```sh
[%1234567890]*
```

There are two key points to solving this problem:

The first is the command `if [[ $r == $i ]]`. In a command like `[[ $r = $i ]]`, if the right-hand side `$i` is not enclosed in double quotes, it will be interpreted as a glob pattern. For example, if `$i` is `*`, this expression will be true regardless of the value of `$r`.

The second is the command `printf $i | grep -e [^0-9]`. This command checks if `$i` contains any non-numeric characters. If non-numeric characters are found, `grep` returns an exit code of 0, making the if statement true and causing the script to exit.

The `pipefail` option is enabled in this script. Therefore, if `printf` exits with a non-zero code, the if statement will not be satisfied. By specifying an invalid escape sequence like `%z`, `printf` will fail, allowing us to bypass the check.

From this, we can infer that by providing a string that makes `printf $i` fail and matches any digit as a glob pattern, we can obtain the flag. The input `[%1234567890]*` meets these conditions.

Note that the behavior might slightly differ depending on the shell. This challenge uses ash found in Alpine Linux as `/bin/sh`.

If Bash is used, RCE could be achieved using the `-v` option in the built-in `printf` command. However, the built-in `printf` in ash does not support the `-v` option, so this attack method is not applicable.

```sh
printf -va[1$(cat${IFS}flag.txt>&2)]
```

[ShellCheck](https://www.shellcheck.net/) is a well-known tool for detecting bugs in shell scripts. The bugs highlighted in this problem can be detected with ShellCheck, providing hints for solving the problem.

## 解法

次の入力によりフラグを得ることができます。

```sh
[%1234567890]*
```

本問題を解くポイントは2つあります。

1つ目は`if [[ $r == $i ]]`です。`[[ $r = $i ]]`のようなコマンドでは、右辺`$i`がダブルクオーテーションで囲われていないと`$i`がGlobパターンとして解釈されてしまいます。例えば`$i`が`*`であれば、この式は`$r`の値にかかわらず真となります。

2つ目は`printf $i | grep -e [^0-9]`です。これらのコマンドでは`$i`に数字以外の文字が含まれるかをチェックしています。数字以外の文字が含まれると`grep`が終了コード0を返し、if文の条件を満たすためスクリプトが終了します。

このスクリプトでは`pipefail`オプションが有効になっています。そのため、`printf`が0以外の終了コードで終了するとif文の条件を満たさなくなります。例えば不正なエスケープシーケンス`%z`などを指定すれば、`printf`が異常終了しチェックを回避できます。

以上より、`printf $i`を異常終了させ、かつGlobパターンとして任意の数字にマッチする文字列を入力すればフラグが得られます。`[%1234567890]*`などがこの条件に当てはまる入力です。

なお、これらの挙動はシェルによって多少異なります。本問ではAlpine Linuxに`/bin/sh`として入っているashを使用しています。

もし実行環境がBashであった場合、組み込みの`printf`に存在する`-v`オプションを利用して次のようにRCEが可能です。Ashに組み込みの`printf`では`-v`オプションは使えないため、この攻撃は成立しません。

```sh
printf -va[1$(cat${IFS}flag.txt>&2)]
```

こういったシェルスクリプトにおけるバグを検出するツールとして[ShellCheck](https://www.shellcheck.net/)が有名です。本問題で取り上げたバグもShellCheckで検出することができ、問題を解くヒントとなります。
