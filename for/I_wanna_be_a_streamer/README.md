---
title: "I_wanna_be_a_streamer"
level: 2
flag: "FLAG{Th4nk_y0u_f0r_W4tching}"
writer: "kaki005"
---

# I_wanna_be_a_streamer
## 問題文
母ちゃんごめん、俺配信者として生きていくよ。
たまには配信に遊び来てな。
(動画のエンコーディングにはH.264が使われています。)

Sorry Mom, I'll work as a streamer.
Watch my stream once in a while.
(H.264 is used for video encoding.)

## 解法(日本語)
配布ファイルは配信をキャプチャしたものであり、RTSP(Real Time Streaming Protocol)とRTP(Realtime Transport Protocol)が配信に利用されている。<br/>
そこでpcapファイルからRTPを解析して配信の映像を復元するのが今回の目標である。<br/>

1. WireSharkのデフォルトの設定ではRTPを解析できない場合があるので、以下の設定を行う。<br/>
https://fumimaker.net/entry/2021/03/17/215110<br/>
RTPパケットを見るとPayload Typeが96と確認できるので、H.264のプロトコル画面を開きペイロードタイプを96に設定する。
2. 音声であればWireshark上で復元して再生する機能があるが、動画にはない。<br/>
そこでRTPから動画を復元する[プラグイン](https://github.com/volvet/h264extractor)をWireSharkに導入する。
3. プラグインを実行すると動画ファイル(.264ファイル)が抽出される。<br/>
ただしH264のPayLoad Typeを96番に設定しないとプラグインが正しく動かないことに注意。
4. .264から.mp4へ変換することで動画が再生でき、動画からフラグを得られる。<br/>

### 補足
- ちなみに今回のpcapファイル作成に関して、送信側にHappytimeSoftの[RTSPサーバ](https://www.happytimesoft.com/products/rtsp-server/index.html)、受信側に[VLC](https://www.videolan.org/vlc/index.ja.html)を使用しました。
- Writeupを書いていただきありがとうございます。楽しく拝見させていただいております。想定解とは違う方法で解かれていた方がいらっしゃいますので、writeupを紹介させていただきます。
    - https://qiita.com/kusano_k/items/59bba1527a83dcd124d4#i_wanna_be_a_streamer-easy
    - https://zenn.dev/asusn/articles/200feee34f5186#i_wanna_be_a_streamer-(169pt-144-solves-easy)%E2%9C%85
- 多くの方から面白かった、勉強になったという感想を頂きまして、大変嬉しかったです。

## Solution(English)
The distributed file captures the streaming, and RTSP (Real Time Streaming Protocol) and RTP (Realtime Transport Protocol) are used for the streaming.<br/>
Therefore, the goal of this problem is to recover the video by analyzing the RTP from the file.<br/>

1. Since WireShark's default settings may not be able to analyze RTP, make the following settings.<br/>
https://fumimaker.net/entry/2021/03/17/215110<br/>
You can see that the Payload Type is 96 in RTP Packet, so open the H.264 protocol setting and set the Payload Type to 96.
2. Wireshark has a function to restore audio from RTP Packet and play it, but not for video.<br/>
so you introduce a [plug-in](https://github.com/volvet/h264extractor) that restores video from RTP in WireShark.
3. When you execute the plug-in, the video file (.264 file) is extracted.<br/>
Note that the plug-in will not work properly unless the PayLoad Type for H264 is set to 96.<br/>
4. By converting from .264 to .mp4, you can play the video and get the flag from the video.


### Supplement
- To create pcap file, I used HappytimeSoft's RTSP server(https://www.happytimesoft.com/products/rtsp-server/index.html) and VLC(https://www.videolan.org/vlc/index.en.html) as client side.
- Thank you for writeup. I enjoyed readiing them. Some people solved the chall using a method different from the expected solution, so I would like to introduce　their writeup.
    - https://qiita.com/kusano_k/items/59bba1527a83dcd124d4#i_wanna_be_a_streamer-easy
    - https://zenn.dev/asusn/articles/200feee34f5186#i_wanna_be_a_streamer-(169pt-144-solves-easy)%E2%9C%85
- I'm very happy to hear that many people said that it was interesting and learned a lot.
