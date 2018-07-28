package cfg

import (
	"github.com/spf13/viper"
	"fmt"
)

type Cfg struct {
	Token		string
	LogPath		string
	WHUrl		string
	ListenAddr	string
	SupportChat	int64
	Debug		bool
}

var Config Cfg

func NewConfig(path string) error {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	Config.Token = viper.GetString("bot.token")
	Config.LogPath = viper.GetString("bot.log")
	Config.WHUrl = viper.GetString("bot.wh-url")
	Config.ListenAddr = viper.GetString("bot.listen")
	Config.SupportChat = viper.GetInt64("bot.support-chat")
	Config.Debug = viper.GetBool("bot.debug")

	return Config.checkParams()
}

func (c *Cfg) checkParams() error {
	if "" == c.Token {
		return fmt.Errorf("config: no bot token defined")
	}
	if "" == c.LogPath {
		return fmt.Errorf("config: no log path defined")
	}
	if "" == c.ListenAddr {
		return fmt.Errorf("config: no listen addr defined")
	}
	if 0 == c.SupportChat {
		return fmt.Errorf("config: no support chat ID defined")
	}
	return nil
}
