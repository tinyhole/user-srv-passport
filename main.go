package main

import (
	"flag"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/sirupsen/logrus"
	"github.com/tinyhole/kit"
	"github.com/tinyhole/user-srv-passport/global"
	"github.com/tinyhole/user-srv-passport/dal/db"
	"github.com/tinyhole/user-srv-passport/handler"
	"github.com/tinyhole/user-srv-passport/idl/platform/user/srv/passport"
	"github.com/tinyhole/user-srv-passport/util"
	"os"
	"time"
)


var (
	GitTag    = "2000.01.01.release"
	BuildTime = "2000-01-01T00:00:00+0800"
)

func main(){
	//显示版本号信息
	version := flag.Bool("v", false, "version")
	flag.Parse()
	if *version {
		fmt.Println("Git Tag: " + GitTag)
		fmt.Println("Build Time: " + BuildTime)
		os.Exit(0)
	}

	kit.LoadConfig()
	svc := micro.NewService(
		micro.Name("platform.user.srv.passport"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Version("v1"),
		micro.Registry(consul.NewRegistry(registry.Addrs(kit.DefaultRegistryConf.Address))),
	)

	svc.Init()

	if err := passport.RegisterPassportHandler(svc.Server(), &handler.Handler{}); err != nil {
		logrus.Fatal(err)
	}

	global.Init(svc)
	db.Init()
	util.Init()

	if err := svc.Run(); err != nil {
		logrus.Fatalf("service run error: %v", err)
	}
}