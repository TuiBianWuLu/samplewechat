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
        AppID:  "",
        Secret: "",
        Cache: cache.NewRedis(&redis.Options{
            Addr:     "localhost:6379",
            Password: "",
        }),
    })

    //token, err := wechat.AccessToken().AccessToken()
    //
    //fmt.Println(token, err)

    //menuButtons := menu.CreateMenuButton{Buttons: []menu.Button{
    //    {Name: "今晚", Type: "view", Url: "https://www.baidu.com"},
    //    {Name: "今晚吃屎", Type: "click", SubButton: []menu.Button{
    //        {Name: "今晚吃鸡吧", Type: "view", Url: "https://www.igetget.com/"},
    //    }},
    //}, MatchRule: menu.MatchRule{Sex:"1"}}
    //
    //res, err := wechat.Menu().CreateMenu(menuButtons)

    //fmt.Println(res, err)

    res, err := wechat.Menu().TryMatch(menu.TryMatchUser{UserId:"huangjinlong19920805"})

    fmt.Printf("%+v ----- %s", res, err)
}
