package boot

import (
	"github.com/qkgo/scaff/pkg/cfg"
	"github.com/qkgo/scaff/pkg/serialize"
)

func ConfigurationParam(args cfg.ParamLike) {
	cfg.ParseConfig(args)
}

/**
gorm 1
*/
func BootDBConnect(dialectName string, dbUrlName string) {
	dialect := cfg.ConfigParam.GetString(dialectName)
	url := cfg.ConfigParam.GetString(dbUrlName)
	serialize.ConfigDatabase(dialect, url)
}

/**
gorm 2
*/
func BootSecondDBConnect(dialectName string, dbUrlName string) {
	dialect := cfg.ConfigParam.GetString(dialectName)
	url := cfg.ConfigParam.GetString(dbUrlName)
	serialize.ConfigSecondDatabase(dialect, url)
}

/**
xorm 1
*/
func BootDBConnectX(dialectName string, dbUrlName string) {
	dialect := cfg.ConfigParam.GetString(dialectName)
	url := cfg.ConfigParam.GetString(dbUrlName)
	serialize.ConfigXDatabase(dialect, url)
}

/**
xorm 2
*/
func BootSecondDBConnectX(dialectName string, dbUrlName string) {
	dialect := cfg.ConfigParam.GetString(dialectName)
	url := cfg.ConfigParam.GetString(dbUrlName)
	serialize.ConfigSecondXDatabase(dialect, url)
}
