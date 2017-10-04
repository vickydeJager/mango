package main

import (
	"log"

	"github.com/louisevanderlith/mango/util/enums"

	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
	_ "github.com/louisevanderlith/mango/app/classifieds/routers"
	"github.com/louisevanderlith/mango/db/classifieds"
	"github.com/louisevanderlith/mango/util"
)

var instanceKey string

func main() {
	// Register with router
	srv := util.Service{
		Environment: enums.GetEnvironment(beego.AppConfig.String("runmode")),
		Name:        beego.AppConfig.String("appname"),
		Type:        enums.APP}

	discURL := beego.AppConfig.String("discovery")
	key, err := util.Register(srv, discURL)

	if err != nil {
		log.Panic(err)
	} else {
		instanceKey = key
		classifieds.NewDatabase(instanceKey, discURL)
		beego.Run()
	}
}
