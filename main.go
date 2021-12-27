package main

import "github.com/edwardsuwirya/wmbMenuMgmt/config"

func main() {
	appConfig := config.NewConfig()
	appConfig.RunMigration()
	appConfig.StartEngine()
}
