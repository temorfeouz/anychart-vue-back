package main

import (
	"github.com/temorfeouz/anychart-vue-back/api"
	"github.com/temorfeouz/anychart-vue-back/config"
)

//go:generate /home/t/src/go/bin/go-bindata -o api/assets_gen.go -pkg api vue_src/dist/...

func main() {
	config.Init()

	api.Run(config.Get().Service.Host, config.Get().Service.Post)
}
