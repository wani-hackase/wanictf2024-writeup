---
title: cached hash
level: 2
flag: FLAG{r3m07E_bU1ld_C4ch3_suPp0r7_1s_4dded_Las7_y34r}
writer: ciffelia
---

# cached hash

> If you add sensitive information to a container image, it seems that it will remain in the intermediate layers even if you delete it later.
>
> But we’re using multi-stage builds this time, so it should be okay, right?
>
> コンテナイメージに機密情報を追加すると、あとから削除しても途中のレイヤーに残ってしまうらしい。
>
> 今回はマルチステージビルドを使ってるから大丈夫だよね……？
>
> http://chal-lz56g6.wanictf.org:5089/

## Solution

日本語版は英語版の後に記載しています。Japansese follows English.

By looking at the `docker-compose.yaml`, we can see that an image has been pushed to Amazon ECR Public, AWS' container registry.

Images stored in Amazon ECR Public can be viewed in detail via the [Amazon ECR Public Gallery](https://gallery.ecr.aws/). For the image in question, besides the `latest` tag, there is also an image tagged `cache`.

https://gallery.ecr.aws/s8s0z7v7/r39cvpwh

For images built using multi-stage Dockerfiles, only the final stage is pushed to the registry. However, BuildKit has a feature that allows intermediate stages to be output as separate images for caching purposes. It can be inferred that this feature was utilized in this case. Indeed, when checking the "Type" field of the image in the ECR Public Gallery, it shows `application/vnd.buildkit.cacheconfig.v0`.

https://github.com/moby/buildkit#export-cache

Since these cache images differ from regular container images, they cannot be retrieved with `docker pull`. Instead, the tool [skopeo](https://github.com/containers/skopeo), which allows various operations on container repositories, can be used to download the image.

```sh
skopeo copy docker://public.ecr.aws/s8s0z7v7/r39cvpwh:cache dir:./cache-image
```

Within the downloaded directory, there are several files. By using the `file` command, we can see that they are tar archives, so we extract all of them and use the `find` command to search for `flag.txt`.

```sh
cd cache-image
for x in *; do mkdir -p out/$x; tar xf $x -C out/$x; done
find out -name flag.txt
```

It was only last year that it became possible to push such cache images to Amazon ECR.

https://aws.amazon.com/jp/blogs/news/announcing-remote-cache-support-in-amazon-ecr-for-buildkit-clients/

These images were pushed to ECR Public using GitHub Actions. The workflow used is [`container-image.yaml`](src/.github/workflows/container-image.yaml). Of course, it is also possible to push from a local machine using the `docker build` command.

## 解法

`docker-compose.yaml`を見ると、AWSのコンテナレジストリAmazon ECR Publicにイメージがプッシュされていることがわかります。

Amazon ECR Publicに存在するイメージは、[Amazon ECR Public Gallery](https://gallery.ecr.aws/)から詳細を確認することができます。今回のイメージを確認すると、`latest`以外に`cache`というタグのイメージが存在することがわかります。

https://gallery.ecr.aws/s8s0z7v7/r39cvpwh

マルチステージのDockerfileにより構築されたイメージでは、レジストリにプッシュされるのは最終ステージのみです。しかしBuildKitには途中のステージをキャッシュとして別イメージに出力する機能があります。今回はこの機能を利用したのではないかと推測できます。実際、ECR Public Galleryで当該イメージのType欄を確認すると`application/vnd.buildkit.cacheconfig.v0`と表示されています。

https://github.com/moby/buildkit#export-cache

このようなキャッシュイメージは通常のコンテナイメージとは異なるため、`docker pull`で取得することはできません。代わりに、コンテナリポジトリに対して様々な操作を行えるツール[skopeo](https://github.com/containers/skopeo)を使用してイメージをダウンロードできます。

```sh
skopeo copy docker://public.ecr.aws/s8s0z7v7/r39cvpwh:cache dir:./cache-image
```

ダウンロードしたディレクトリ内にはいくつかのファイルが存在します。`file`コマンドで中身を調べるとtar形式のアーカイブであることがわかるため、すべて展開して`find`コマンドで`flag.txt`を探します。

```sh
cd cache-image
for x in *; do mkdir -p out/$x; tar xf $x -C out/$x; done
find out -name flag.txt
```

ちなみに、このようなキャッシュイメージをAmazon ECRにプッシュできるようになったのは昨年のことだそうです。

https://aws.amazon.com/jp/blogs/news/announcing-remote-cache-support-in-amazon-ecr-for-buildkit-clients/

これらのイメージはGitHub Actionsを用いてECR Publicにプッシュしました。使用したワークフローは[`container-image.yaml`](src/.github/workflows/container-image.yaml)です。もちろん、手元のPCから`docker build`コマンドを用いてプッシュすることも可能です。
