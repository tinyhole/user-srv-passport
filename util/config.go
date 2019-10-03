package util

import (
	"github.com/micro/go-micro/config"
	"github.com/sirupsen/logrus"
)

type WeChatOpenConfig struct {
	URL string `json:"url"`
	Secrets map[string]string `json:"secrets"`
}

type Config struct {
	WeChatOpenConfig WeChatOpenConfig
}

var DefaultConfig *Config

func Init() {
	DefaultConfig = &Config{}

	err := config.Get("weChatOpen").Scan(&DefaultConfig.WeChatOpenConfig)
	if err != nil {
		logrus.Fatal(err, "Failed to get weChatOpen")
	}
}