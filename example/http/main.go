package main

import (
    "fmt"

    "github.com/TuiBianWuLu/samplewechat"
    "github.com/TuiBianWuLu/samplewechat/config"
    "github.com/TuiBianWuLu/samplewechat/util/cache"
    "github.com/go-redis/redis"
    "github.com/TuiBianWuLu/samplewechat/template"
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


    //msgData := message.News{
    //    Basic:message.Basic{ToUser:"oIVRjuA9Tle_WTi9Z_UkkSBXoYOU", MsgType: "news"},
    //    News: message.Info{Articles: []message.Info{
    //        {Title: "我笑了", Description: "感觉自己萌萌哒", Url: "https://h5.sao.cn/showcase/showlist", PicUrl: "https://piccdn.luojilab.com/store-pc/image/201803/1522336599895.jpg"},
    //        {Title: "我哭了", Description: "感觉自己呵呵哒", Url: "https://www.baidu.com/", PicUrl: "https://piccdn.luojilab.com/store-pc/image/201803/1522336599895.jpg"},
    //    }},
    //    CustomService: message.CService{KfAccount:"test@test"},
    //}


    //template := message.Template{
    //    ToUser: "oIVRjuA9Tle_WTi9Z_UkkSBXoYOU",
    //    TemplateID: "mT0utpTHr57bU3zIb_u56DMAk4cbIFK1jv0xZ_N-8i4",
    //    Data: map[string]message.Item{
    //        "first":{Value: "购买成功", Color: "#173177"},
    //        "keyword1":{Value: "巧克力", Color: "#173177"},
    //        "keyword2":{Value: "111111", Color: "#173177"},
    //        "keyword3":{Value: "108", Color: "#173177"},
    //        "keyword4":{Value: "2018年04月08日", Color: "#173177"},
    //        "remark":{Value: "欢迎再次购买", Color: "#173177"},
    //        },
    //}

    //industry := template.Industry{IndustryID1: 2, IndustryID2: 3}


    templateId := template.DelTemplateID{TemplateID: "Z9CfT-ylQ_Ww9i4D3hDTeDGlaG4xv__CY--oUlvMly8"}
    res, err := wechat.Template().DelTemplate(templateId)

    fmt.Printf("%+v ----- \n %s --- \n ", res, err)
}
