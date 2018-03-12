package main

import (
    "fmt"

    "github.com/TuiBianWuLu/samplewechat/basic"
)

func main() {

    weChat := &basic.WeChat{
        AppId: "wx7d2495e65b96a40f",
        Secret: "0a8789cb6976f88fc698a1bc69c61b03",
    }

    x, err := weChat.RefreshToken()

    if err != nil {
        return
    }
    fmt.Println(x.AccessToken)
}
