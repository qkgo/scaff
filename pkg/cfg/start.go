package cfg

import (
	"github.com/qkgo/scaff/pkg/util/system"
	"log"
)

type ParamLike struct {
	Config  string
	Env     string
	Mapping string `arg:"--mapping"`
	Verbose bool   `arg:"-v" help:"verbosity level"`
}

func ParseConfigEx(args ParamLike) map[string]interface{} {
	var path string
	if args.Env == "test" {
		path = "config/test.yml"
	}
	if args.Env == "prd" {
		path = "config/prd.yml"
	}
	if args.Env == "dev" {
		path = "config/dev.yml"
	}
	if path != "" {
		log.Println("local specify config: ", path)
		Configuration = ReadYamlFile(path)
		env := Configuration["env"].(string)
		if env == "" {
			log.Println("local specify config failed")
			system.Exit(-2)
			return nil
		}
		return Configuration
	}
	if args.Config != "" {
		log.Println("local specify config: ", args.Config)
		Configuration = ReadYamlFile(args.Config)
		env := Configuration["env"].(string)
		if env == "" {
			log.Println("local specify config failed")
			system.Exit(-4)
			return nil
		} else {
			log.Println("local specify config finish,current: ", env)
			return Configuration
		}
	} else {
		log.Println("local specify config no found")
		system.Exit(-3)
		return nil
	}
}

func ParseConfig(args ParamLike) {
	var path string
	if args.Env == "test" {
		path = "config/test.yml"
	}
	if args.Env == "prd" {
		path = "config/prd.yml"
	}
	if args.Env == "dev" {
		path = "config/dev.yml"
	}
	if loadConfigToGlobal(path) {
		if env := ConfigParam.GetString("env"); env == "" {
			log.Println("local specify config failed")
		} else {
			log.Println("load config finish, current: ", env)
			return
		}
	}
	if args.Config != "" {
		if loadConfigToGlobal(args.Config) {
			if env := ConfigParam.GetString("env"); env == "" {
				log.Println("local specify config failed")
				system.Exit(-1)
			} else {
				log.Println("load config finish, current: ", env)
				return
			}
			return
		} else {
			system.Exit(-2)
			return
		}
	} else {
		log.Println("local specify config no found")
		system.Exit(-3)
		return
	}
}

func loadConfigToGlobal(path string) bool {
	log.Println("local specify config: ", path)
	if path == "" {
		return false
	}
	err := ConfigParam.Load(path)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}
