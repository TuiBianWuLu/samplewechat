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

    //res, err := wechat.Menu().CreateMenu(menu.CreateMenuButton{
    //    Buttons: []menu.Button{
    //        {Type: "view", Name: "修改化菜单1", Url: "https://igetget.com"},
    //        {Type: "view", Name: "今晚在北京", Url: "https://www.baidu.com"},
    //        {Type: "click", Name: "明晚在东京", Url: "https://igetget.com", SubButton:[]menu.Button{
    //            {Type: "view", Name: "哈哈", Url: "https://www.baidu.com"},
    //        }},
    //    },
    //    MatchRule: menu.MatchRule{
    //        Country: "中国",
    //        Province: "北京",
    //        City: "通州",
    //    },
    //})

    // 863665865
    //res, err := wechat.Menu().DelConditional(menu.DelMenu{MenuID: 863665865})

    //res, err := wechat.Menu().QueryMenu()

    res, err := wechat.Menu().TryMatch(menu.TryMatchUser{UserId: "huangjinlong19920805"})

    fmt.Printf("%+v ----- \n %s --- \n ", res, err)
}
