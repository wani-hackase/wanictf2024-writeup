---
title: devserv
level: 3
flag: FLAG{1d4p_s3rv3r_4nd_r3m0t3_p0rt_f0rw4rd1ng}
writer: ciffelia
badge: true
---

# devserv

> Sign in to get your flag. Source code is available on git!
>
> ログインするとフラグが手に入ります。ソースコードはGitで配布中！
>
> http://chal-lz56g6.wanictf.org:6867/
>
> ```
> git clone git@chal-lz56g6.wanictf.org:/devserv.git
> # Password: G1Ta34WdBLDd!w%Z
> ```

## Solution (English)

日本語版は英語版の後に記載しています。Japansese follows English.

In this challenge, there are three servers running: an LDAP server, an SSH server, and a web service called flagserv that provides the flag.

The flagserv connects to an LDAP server specified by the user and attempts to log in using the provided ID and password. If the login is successful, it returns the flag. However, the LDAP server that can be connected to is limited to pre-configured ones; specifying any other LDAP server will result in an error.

The process for determining whether an LDAP server is allowed is as follows:

```go
u, err := url.Parse(ldapUrl)
if err != nil {
  return "invalid ldap url"
}
if !slices.Contains(ldap_allow[:], u.Hostname()) {
  return "ldap url not allowed"
}
```

The URL provided by the user is parsed, and the hostname is checked against a pre-configured list (like `localhost`). However, the port number is not verified, allowing the specification of ports other than the default LDAP port `389`.

On the other hand, the SSH server is set up to provide a Git repository. The `git` user is configured as a connectable user, with the ID and password provided in the problem statement. While it might seem possible to execute arbitrary commands on the server by connecting with SSH, this is not actually possible. The login shell is set to `git-shell`, so only Git-related operations are allowed. This method is documented in the official Git documentation.

https://git-scm.com/book/en/v2/Git-on-the-Server-Setting-Up-the-Server

The documentation includes the following note:

> At this point, users are still able to use SSH port forwarding to access any host the git server is able to reach. If you want to prevent that, you can edit the `authorized_keys` file and prepend the following options to each key you’d like to restrict:
>
> ```
> no-port-forwarding,no-X11-forwarding,no-agent-forwarding,no-pty
> ```

This means that if not properly configured, the SSH port forwarding feature can be used. While the documentation mentions the risk of local port forwarding, remote port forwarding is also possible. This means an attacker can wait for connections on any port on the server and respond to connecting clients.

Based on the above, by exploiting the flagserv's port number verification oversight and the SSH server's configuration error, you can successfully log in by following these steps:

1. Set up an LDAP server on your local machine (let's use port 1234).
1. Use SSH remote port forwarding to forward port 5678 on the target server to port 1234 on your local machine.
1. Specify `ldap://localhost:5678` as the LDAP server and log in to flagserv.

Setting up an LDAP server for the first time can be a bit challenging, but it can be relatively easy by reusing a Dockerfile that can be obtained from a Git repository. The Dockerfile and commands used during the problem-checking process can be found in the [solver](solver) directory.

## Solution（日本語）


本問題のサーバーでは、LDAPサーバー、SSHサーバー、フラグを提供するWebサービス(flagserv)の3つが稼働しています。

flagservはユーザーが指定したLDAPサーバーに接続し、ユーザーから提供されたIDとパスワードでログインを試みます。ログインに成功するとフラグを返します。ただし、接続できるLDAPサーバーは事前に設定されたものに限定されており、それ以外のLDAPサーバーを指定するとエラーを返します。

ここで、接続が許可されているLDAPサーバーを判定する処理は次のようになっています。

```go
u, err := url.Parse(ldapUrl)
if err != nil {
  return "invalid ldap url"
}
if !slices.Contains(ldap_allow[:], u.Hostname()) {
  return "ldap url not allowed"
}
```

ユーザーから提示されたURLをパースし、そのホスト名が`localhost`など事前に設定されたリストに含まれているかをチェックしています。しかし、ここではポート番号が検証されていません。そのため、LDAPサーバーのデフォルトのポート番号`389`以外を指定することが可能になっています。

一方で、SSHサーバーはGitリポジトリを提供するために用意されています。`git`ユーザーが接続可能なユーザーとして設定されており、IDとパスワードも問題文で提示されています。これではSSHコマンドでサーバーに接続し任意のコマンドを実行できるように思えますが、実際にはそういったことはできません。ログインシェルが`git-shell`に設定されているため、Git関連の操作のみが可能になっています。この手法はGitの公式ドキュメントに記載されています。

https://git-scm.com/book/en/v2/Git-on-the-Server-Setting-Up-the-Server

このドキュメントを読むと、以下の記述が見つかります。

> At this point, users are still able to use SSH port forwarding to access any host the git server is able to reach. If you want to prevent that, you can edit the `authorized_keys` file and prepend the following options to each key you’d like to restrict:
>
> ```
> no-port-forwarding,no-X11-forwarding,no-agent-forwarding,no-pty
> ```

つまり、適切な設定をしなければSSHのポートフォワーディング機能が利用可能になるということです。このドキュメントではローカルポートフォワーディングのリスクに触れていますが、リモートポートフォワーディングも同様に可能となっています。すなわち、攻撃者はサーバー上の任意のポートで接続を待ち受け、接続してきたクライアントに応答を返すことが可能になってしまいます。

以上より、flagservのポート番号検証漏れとSSHサーバーの設定ミスを利用することで、以下の手順でログインを成功させることができます。

1. ローカルマシンにLDAPサーバーを構築する（仮にポート1234とする）
1. SSHリモートポートフォワーディングにより、攻撃対象サーバーのポート5678をローカルマシンのポート1234に転送する
1. LDAPサーバーとして`ldap://localhost:5678`を指定しflagservにログイン

LDAPサーバーを初めて構築するのは少し大変ですが、Gitリポジトリから取得できるDockerfileを再利用すると比較的簡単に構築できます。作問チェック時に利用したDockerfileやコマンドは[solver](solver)ディレクトリから確認できます。
