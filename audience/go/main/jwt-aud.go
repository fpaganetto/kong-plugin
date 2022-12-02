package main

import (
	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
)

type Config struct {
	Audience string
}

func New() interface{} {
	return &Config{}
}

func (config Config) Access(kong *pdk.PDK) {
	host, _ := kong.Request.GetHeader("Authorization")
	token := GetToken(host)
	aud := GetAudience(token)
	if !Contains(aud, config.Audience) {
		kong.Response.Exit(403, "Invalid audience", make(map[string][]string))
	}
}

func main() {
	Version := "1.1"
	Priority := 1000
	_ = server.StartServer(New, Version, Priority)
}
