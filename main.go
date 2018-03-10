package main

import (
    "fmt"

    "github.com/TuiBianWuLu/samplewechat/basic"
)

func main() {

    wechat := &basic.Wechat{
        AppId: "",
        Secret: "",
    }

    x, err := wechat.RefreshToken()

    fmt.Println(x, err)
}
