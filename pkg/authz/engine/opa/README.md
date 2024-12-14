# OPA 决策引擎

## [`policy.bindata.go`](policy.go)

此文件生成于 `policy/*.rego`.

运行以下命令生成：

```shell
go generate ./...
```

注意： 需要保证[go-bindata](github.com/kevinburke/go-bindata)被安装。

```shell
go install github.com/kevinburke/go-bindata/v4/...@latest

# for versions of Go < 1.11, or without module support, use:
go go install github.com/kevinburke/go-bindata/...
```
