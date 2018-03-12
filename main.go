package main

import (
    "fmt"

    "github.com/TuiBianWuLu/samplewechat/basic"
)

func main() {

    weChat := &basic.WeChat{
        AppId: "wxee829f0f8f8e310f",
        Secret: "d4624c36b6795d1d99dcf0547af5443d",
    }

    x, err := weChat.RefreshToken()

    if err != nil {
        return
    }
    fmt.Println(x.AccessToken)
}
