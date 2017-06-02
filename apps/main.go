package main

import (
	"Accountkit/config"
	"log"
	"Accountkit/factory"
	"Accountkit/context"
	"github.com/golang/glog"
	"fmt"
	"Accountkit/constants"
	"os"
	"flag"
)

func usage() {
	// -logtostderr=true -v=2
	// -stderrthreshold=FATAL -log_dir=./log -v=2
	fmt.Fprintf(os.Stderr, "usage: example -stderrthreshold=[INFO|WARN|FATAL] -log_dir=[string]\n", )
	flag.PrintDefaults()
	os.Exit(2)
}

func init() {
	flag.Usage = usage
	flag.Parse()
}

func main() {
	var appConfig *config.Config
	var err error

	appConfig, err = config.NewAppSettingsFromFile(constants.ConfigFile)
	handleConfigErrors(err)

	dbFactory := factory.NewPostgresDatabaseFactory(appConfig)
	dbFactory.InitializeDatabaseConnection()
	defer dbFactory.CloseConnection()

	appContext := context.NewContext(dbFactory.DBConnection())
	glog.Infoln(appContext)
}

func handleConfigErrors(err error)  {
	if err != nil {
		message := fmt.Sprintf("error initialising config: %s", err)
		glog.Fatalln(message)
		log.Fatalf(message)
	}
}
