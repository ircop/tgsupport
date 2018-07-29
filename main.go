package main

import (
	"flag"
	"ircop/tgsupport/cfg"
	"ircop/tgsupport/logger"
	"ircop/tgsupport/bot"
)

func main() {
	configPath := flag.String("c", "./tgsupport.toml", "Config file location")
	flag.Parse()

	err := cfg.NewConfig(*configPath)
	if nil != err {
		panic(err)
	}

	err = logger.SetPath(cfg.Config.LogPath);
	if nil != err {
		panic(err)
	}
	logger.SetDebug(cfg.Config.Debug)

	if err = bot.Init(); nil != err {
		panic(err)
	}
}
