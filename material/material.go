package material

import "github.com/TuiBianWuLu/samplewechat/config"

type Material struct {
    *config.Config
}

var m = new(Material)

func New(c *config.Config) *Material {
    m.Config = c
    return m
}