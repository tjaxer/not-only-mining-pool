package main

import (
	"encoding/json"
	"flag"
	logging "github.com/ipfs/go-log"
	"github.com/mining-pool/go-pool-server/config"
	"github.com/mining-pool/go-pool-server/poolManager"
	"github.com/mining-pool/go-pool-server/utils"
	"os"
)

var log = logging.Logger("main")

const defaultConfigFileName = "config.json"

var configFileName = flag.String("c", defaultConfigFileName, "configuration file for pool")
var logLevel = flag.String("l", "info", "log level")

func main() {
	flag.Parse()

	lvl, err := logging.LevelFromString(*logLevel)
	if err != nil {
		panic(err)
	}
	logging.SetAllLoggers(lvl)

	var conf config.Options
	if !utils.FileExists(*configFileName) {
		log.Panic("the config file " + *configFileName + " does not exist")
	}

	f, err := os.Open(*configFileName)
	if err != nil {
		log.Panic(err)
	}

	err = json.NewDecoder(f).Decode(&conf)
	if err != nil {
		log.Panic(err)
	}

	p := poolManager.NewPool(&conf)
	p.Init()
	for {
		select {}
	}
}
