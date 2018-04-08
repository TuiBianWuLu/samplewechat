package main

import (
    "fmt"

    "github.com/TuiBianWuLu/samplewechat"
    "github.com/TuiBianWuLu/samplewechat/config"
    "github.com/TuiBianWuLu/samplewechat/util/cache"
    "github.com/go-redis/redis"
    "github.com/TuiBianWuLu/samplewechat/message"
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

    //res, err := wechat.Menu().QueryMenu()

    //msgData := message.Text{
    //    Basic:message.Basic{ToUser:"oIVRjuA9Tle_WTi9Z_UkkSBXoYOU", MsgType:"text"},
    //    Text: message.Info{Content:"呵1111呵哒"}}


    msgData := message.News{
        Basic:message.Basic{ToUser:"oIVRjuA9Tle_WTi9Z_UkkSBXoYOU", MsgType: "news"},
        News: message.Info{Articles: []message.Info{
            {Title: "我笑了", Description: "感觉自己萌萌哒", Url: "https://h5.sao.cn/showcase/showlist", PicUrl: "https://piccdn.luojilab.com/store-pc/image/201803/1522336599895.jpg"},
            {Title: "我哭了", Description: "感觉自己呵呵哒", Url: "https://www.baidu.com/", PicUrl: "https://piccdn.luojilab.com/store-pc/image/201803/1522336599895.jpg"},
        }},
    }

    res, err := wechat.Message().SendCustom(msgData)

    fmt.Printf("%+v ----- \n %s --- \n ", res, err)
}
