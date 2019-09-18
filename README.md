# MackerelGoScrapingMetal
金とプラチナの価格をスクレイピングしてMackerelのカスタムメトリクスとして可視化する

■ デプロイ
```
- Goコンパイル
  - $ make build

- ServerlessFrameworkでデプロイ
  - $ sls deploy --aws-profile <PROFILE> --mkrkey <MACKEREL_API_KEY>
```


`$ make help`
```
build:             Build binaries
build-deps:        Setup build
deps:              Install dependencies
devel-deps:        Setup development
help:              Show help
lint:              Lin
```
