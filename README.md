# Redis Practice
takara2314 が Docker から Redis を触るトレーニングをするリポジトリです。

今年こそは一人でDockerできるように頑張ります・・・

## 僕へ

結構メモリ食うから、ChromeでYouTube見ながらコンテナ起動したら死ぬぞ

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
