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

    WeChat := samplewechat.New(&config.Config{
        AppID:  "wxee829f0f8f8e310f",
        Secret: "d4624c36b6795d1d99dcf0547af5443d",
        Cache: cache.NewRedis(&redis.Options{
            Addr:     "localhost:6379",
            Password: "",
        }),
    })

    //token, err := wechat.AccessToken().AccessToken()
    //
    //fmt.Println(token, err)

    //创建自定义菜单
    //menuButtons := menu.CreateMenuButton{Buttons: []menu.Button{
    //    {Name: "今晚吃鸡", Type: "view", Url: "https://www.baidu.com"},
    //    {Name: "今晚吃屎", Type: "click", SubButton: []menu.Button{
    //        {Name: "今晚吃鸡吧", Type: "view", Url: "https://www.igetget.com/"},
    //    }},
    //}}

    //个性化菜单创建
    //menuButton := menu.CreateMenuButton{Buttons: []menu.Button{
    //    {Name: "我笑了", Type: "view", Url: "https://h5.sao.cn"},
    //    {Name: "我哭了", Type: "click", SubButton: []menu.Button{
    //        {Name: "你瞅啥", Type: "view", Url: "https://www.igetget.com/"},
    //    }},
    //},MatchRule : menu.MatchRule{Sex: 1, ClientPlatformType: 1, Country: "中国", Province: "北京", City: "通州", Language: "zh_CN"}}

    //测试个性化菜单匹配结果
    //menuButton := menu.TryMatch{UserID:"openid"}
    //res, err := WeChat.Menu().TestTryMatch(menuButton)

    menuButton := menu.DelMenu{MenuID: 863086243}


    res, err := WeChat.Menu().DelConditional(menuButton)



    //res, err := WeChat.Menu().CreateMenu(menuButtons)
    //
    //fmt.Println(res, err)

    //res, err := wechat.Menu().DelMenu()

    fmt.Printf("%+v ----- %s", res, err)
}
