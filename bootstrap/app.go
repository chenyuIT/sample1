package bootstrap

import (
	"github.com/chenyuIT/framework/foundation"

	"chenyuIT/config"
)

func Boot() {
	app := foundation.Application{}

	//Bootstrap the application
	app.Boot()

	//Bootstrap the config.
	config.Boot()
}
