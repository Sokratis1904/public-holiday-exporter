package config

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type (
	Opts struct {
		// logger
		Logger struct {
			Debug   bool `           long:"debug"        env:"DEBUG"    description:"debug mode"`
			Verbose bool `short:"v"  long:"verbose"      env:"VERBOSE"  description:"verbose mode"`
			LogJson bool `           long:"log.json"     env:"LOG_JSON" description:"Switch log output to json format"`
		}

		Api struct {
			BaseUrl string  `long:"api.baseurl"    env:"API_BASEURL"      description:"API base url" default:"https://date.nager.at/api/v2/publicholidays/"`
			Proxy   *string `long:"api.proxy"      env:"API_PROXY"        description:"API proxy url"`
		}

		Cache struct {
			Path string `long:"cache.path"    env:"CACHE_PATH"      description:"Cache path" default:"cache.json"`
		}

		// tasks
		Preload bool `long:"preload" description:"Do cache preload and exit"`

		// general options
		ServerBind string `long:"bind" env:"SERVER_BIND"   description:"Server address" default:":8080"`

		ConfigPath string `long:"config" short:"c"  env:"CONFIG"   description:"Config path" required:"true"`

		// exporter settings
		ExporterDaysToFetchNewYear int64 `env:"DAYS_TO_NEXT_YEAR"   description:"days to next year to fetch also next years data" default:"30"`
	}
)

func (o *Opts) GetJson() []byte {
	jsonBytes, err := json.Marshal(o)
	if err != nil {
		log.Panic(err)
	}
	return jsonBytes
}
