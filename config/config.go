package config

import (
    "sync"

    "github.com/TuiBianWuLu/samplewechat/util/cache"
)

type Config struct {
    AppID                string
    Secret               string
    AESKey               string
    Token                string
    MchID                string
    MchKey               string
    PrefixAccessTokenKey string // 指定自定义access toke key 缓存 不传则自己刷 传了不会在微信刷 开发环境不传就行
    Cache                cache.Cacher
    AccessTokenLock      sync.RWMutex
    JsTicket             sync.RWMutex
}
