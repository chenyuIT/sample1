package config

import (
	"github.com/chenyuIT/framework/facades"
)

func init() {
	config := facades.Config
	config.Add("captcha", map[string]any{
		"width":       config.Env("CAPTCHA_WIDTH", 40),
		"height":      config.Env("CAPTCHA_HEIGHT", 20),
		"color":       config.Env("CAPTCHA_COLOR", "#FF00FFFF"),
		"imageFormat": config.Env("CAPTCHA_IMAGEFORMAT", 0),
	})
}
