package main

import (
    "fmt"

    "github.com/TuiBianWuLu/samplewechat"
    "github.com/TuiBianWuLu/samplewechat/config"
    "github.com/TuiBianWuLu/samplewechat/util/cache"
    "github.com/go-redis/redis"
)

func main() {

    wechat := samplewechat.New(&config.Config{
        AppId:                "",
        Secret:               "",
        Cache: cache.NewRedis(&redis.Options{
            Addr:     "localhost:6379",
            Password: "",
        }),
    })

    token, err := wechat.AccessToken().AccessToken()

    fmt.Println(token, err)
}
