package main

import (
    "fmt"

    "github.com/TuiBianWuLu/samplewechat"
    "github.com/TuiBianWuLu/samplewechat/config"
    "github.com/TuiBianWuLu/samplewechat/menu"
    "github.com/TuiBianWuLu/samplewechat/util/cache"
    "github.com/go-redis/redis"
)

func main() {

    wechat := samplewechat.New(&config.Config{
        AppId:  "",
        Secret: "",
        Cache: cache.NewRedis(&redis.Options{
            Addr:     "localhost:6379",
            Password: "",
        }),
    })

    //token, err := wechat.AccessToken().AccessToken()
    //
    //fmt.Println(token, err)

    menuList := []menu.Item{
        {
            Name: "今晚吃鸡",
            Type: "click",
            SubButton: []menu.Item{
                {Name: "今晚不吃鸡吧", Type: "view", Url: "https://www.baidu.com"},
                {Name: "今晚不吃屎", Type: "view", Url: "https://www.baidu.com"},
            },
        },
        {
            Name: "点我一下",
            Url: "https://www.baidu.com/",
            Type: "view",
        },
    }

    res, err := wechat.Menu().CreateMenu(menuList)
    fmt.Println(res, err)
}
