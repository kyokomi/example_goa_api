example_goa_api
==========================

[![wercker status](https://app.wercker.com/status/d4e641fa622bd1e5cacf50b6dc9e1332/s/master "wercker status")](https://app.wercker.com/project/byKey/d4e641fa622bd1e5cacf50b6dc9e1332)

It is an example of api using goa

## goaのgenerate手順

1. _design下を編集する
1. ルートディレクトリでgo generateする
1. 生成された `gen/main.go` から追加したリソースのMountのコードをルートディレクトリの `main.go` にコピペする
1. `gen/` 下に生成されたリソースのControllerを `application/controllers`下にコピーしてimport等を直す（リソース追加ではなく編集の場合は、必要なMethodだけコピペする）

## License

[MIT](/LICENSE)
