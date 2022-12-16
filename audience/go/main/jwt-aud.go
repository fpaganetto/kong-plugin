package main

import (
	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
)

type Config struct {
}

func New() interface{} {
	return &Config{}
}

func (config Config) Access(kong *pdk.PDK) {
	authHeader, _ := kong.Request.GetHeader("Authorization")
	host, _ := kong.Request.GetHeader("Host")
	token := GetToken(authHeader)
	aud := GetAudience(token)
	if !Contains(aud, host) {
		kong.Response.Exit(403, "Invalid audience", make(map[string][]string))
	}
}

func main() {
	Version := "1.1"
	Priority := 1000
	_ = server.StartServer(New, Version, Priority)
}
