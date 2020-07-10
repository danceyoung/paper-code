package cfg

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type configM struct {
	envmode string
	currcfg map[string]interface{}
}

var cfgm configM

func init() {
	filebytes, err := ioutil.ReadFile("./config/conf.yml")
	if err != nil {
		panic(err)
	}
	out := make(map[string]interface{})
	err = yaml.Unmarshal(filebytes, &out)
	if err != nil {
		panic(err)
	}

	cfgm = configM{envmode: "dev"}
	if temp := os.Getenv("ENV_MODE"); temp != "" {
		cfgm.envmode = temp
	}
	cfgm.currcfg = out[cfgm.envmode].(map[string]interface{})
	log.Println("config is ", cfgm)
}

func String(key string) string {
	if v, ok := cfgm.currcfg[key]; ok {
		return v.(string)
	}

	return ""
}
