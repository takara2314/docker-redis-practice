# Docker Redis Practice
takara2314 が Docker から Redis を触るトレーニングをするリポジトリです。

今年こそはDockerで開発できるように頑張ります・・・

（Docker使ってデプロイは日々やってるけど、開発環境にDocker使ったことはないため（環境めっちゃ整備されていた内定者アルバイトは除く））

Dockerでできたらヒーロー！！

（しまじろうのトイレでできたらパンツマン風）

## 僕へ

結構メモリ食うから、ChromeでYouTube見ながらコンテナ起動したら死ぬぞ

- Surface 6 Pro
- Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz 1.80 GHz
- RAM 8GB

## コンテナビルドして起動するコマンド
デーモン化したければ `-d` を入れるんだぞ…！

ビルドを強制したければ `--build` を入れるんだぞ…！

```bash
docker compose up
```

## サンプルで動かすAPIとは

僕の状態（`good`, `bad`, `normal`）がランダムで選出したのちRedisに格納され、
`/status`アクセスがあったときに、状態をJSONで返す。

## APIのアドレス

### Production

localhost:3000

### Development

Airによってファイル変更があればホットリロードするよ。

localhost:3001
