package main

import (
	"flag"
	"math/rand"
	"strings"
	"time"

	"github.com/dylenfu/zion-meter/config"
	"github.com/dylenfu/zion-meter/core"
	"github.com/dylenfu/zion-meter/pkg/frame"
	"github.com/dylenfu/zion-meter/pkg/log"
)

var (
	loglevel   int    // log level [1: debug, 2: info]
	configpath string //config file
	Methods    string //Methods list in cmdline
	users, groups int
)

func init() {
	flag.StringVar(&configpath, "config", "config.json", "config path of palette-tool")
	flag.StringVar(&Methods, "t", "tps", "`methods` to run. use ',' to split methods")
	flag.IntVar(&groups, "group", 10, "`group` define user group number")
	flag.IntVar(&users, "user", 5, "`user` denote that user number per group")
	flag.IntVar(&loglevel, "loglevel", 2, "loglevel [1: debug, 2: info]")

	flag.Parse()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	defer time.Sleep(time.Second)

	log.InitLog(loglevel, log.Stdout)
	config.LoadConfig(configpath, groups, users)
	core.Endpoint()

	methods := make([]string, 0)
	if Methods != "" {
		methods = strings.Split(Methods, ",")
	}

	frame.Tool.Start(methods)
}
