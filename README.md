# Bytedance Open Platform SDK

[![Go Report Card](https://goreportcard.com/badge/github.com/houseme/bytedance)](https://goreportcard.com/report/github.com/houseme/bytedance)
[![Go Reference](https://pkg.go.dev/badge/github.com/houseme/bytedance.svg)](https://pkg.go.dev/github.com/houseme/bytedance)
[![Bytedance CI](https://github.com/houseme/bytedance/actions/workflows/go.yml/badge.svg)](https://github.com/houseme/bytedance/actions/workflows/go.yml)
[![License](https://img.shields.io/github/license/houseme/bytedance.svg?style=flat)](https://github.com/houseme/bytedance)
![GitHub go.mod Go version (branch)](https://img.shields.io/github/go-mod/go-version/houseme/bytedance/main)


Bytedance Mini Program Douyin Mini Program bytedance microapp SDK

### Installation

Enter your repo. directory and execute following command:
```bash
go get -u -v github.com/houseme/bytedance@main
```

### Limitation

```
golang version >= 1.20
```

### Usage

```go
package main

import (
    "context"
    "fmt"
    
    "github.com/houseme/bytedance"
    "github.com/houseme/bytedance/config"
    "github.com/houseme/bytedance/credential"
    "github.com/houseme/bytedance/utility/cache"
    "github.com/houseme/bytedance/utility/logger"
    "github.com/houseme/bytedance/utility/request"
)

func main() {
    var ctx = context.Background()
    wc := bytedance.New(ctx)
    cfg := config.New(
        ctx,
        config.WithClientKey(""),
        config.WithClientSecret(""),
        config.WithRedirectURL(""),
        config.WithLogger(logger.NewDefaultLogger()),
        config.WithRequest(request.NewDefaultRequest()),
        config.WithCache(cache.NewRedis(ctx, cache.NewDefaultRedisOpts())),
        config.WithScopes(""),
        config.WithSalt(""),
    )
    
    // 获取小程序实例
    miniProgram, err := wc.MiniProgram(ctx, cfg)
    if err != nil {
        panic(err)
    }
    auth := miniProgram.GetAuthorize()
    // 获取用户授权
    var accessToken credential.AccessToken
    if accessToken, err = auth.GetAccessToken(ctx, "code"); err != nil {
        panic(err)
    }
    fmt.Println(accessToken)
    
    // 创建二维码
    qrcode := miniProgram.GetQrcode()
    
    // 获取小程序码
    schema := miniProgram.GetSchema()
    
    // 获取链接
    link := miniProgram.GetLink()

}

```

### License

`Bytedance` is licensed under the [Apache License Version 2.0](LICENSE), 100% free and open-source, forever.