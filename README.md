# ping-esa-wip

WIP のまま放置されている記事にメンション付きでコメントする Bot

## Getting Started

### 1. Personal access token を発行する

1. https://camphor.esa.io/user/applications から Personal access token を発行
   - 必要なスコープは `Read`, `Write`
1. 環境変数 `ESA_API_TOKEN` にセット

### 2. 実行

```shell script
go run main.go --team=camphor
```

#### ドライランで実行したい場合

```shell script
go run main.go --dry-run --team=camphor
```
