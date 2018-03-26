package config

import "github.com/TuiBianWuLu/samplewechat/util/cache"

type Config struct {
    AppId                string
    Secret               string
    AESKey               string
    Token                string
    MchId                string
    MchKey               string
    Cache                cache.Cacher
    PrefixAccessTokenKey string // 指定自定义access toke key 缓存 不传则自己刷 传了不会在微信刷
}
