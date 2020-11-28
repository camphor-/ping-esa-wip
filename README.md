# ping-esa-wip

WIPのまま放置されている記事にメンション付きでコメントするBot

## Getting Started

### 1. Personal access tokenを発行する

1. https://camphor.esa.io/user/applications からPersonal access tokenを発行
    - 必要なスコープは `Read`, `Write`
1. 環境変数 `ESA_API_TOKEN` にセット

### 2. 実行

```shell script
go run main.go
```

#### ドライランで実行したい場合

```shell script
go run main.go --dry-run
```