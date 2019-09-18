# MackerelGoScrapingMetal
金とプラチナの価格をスクレイピングしてMackerelのカスタムメトリクスとして可視化する

■ デプロイ

- Mackerelにサービスを登録する
```bash
export MKRKEY=XXX

curl -X POST https://api.mackerelio.com/api/v0/services \
    -H "X-Api-Key: ${MKRKEY}" \
    -H "Content-Type: application/json" \
    -d '{"name": "Metal", "memo": "metal"}'

make build
sls deploy --aws-profile <PROFILE> --mkrkey ${MKRKEY}
```


`$ make help`
```
build:             Build binaries
build-deps:        Setup build
deps:              Install dependencies
devel-deps:        Setup development
help:              Show help
lint:              Lint
```
