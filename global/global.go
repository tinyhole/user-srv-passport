package global

import "github.com/micro/go-micro"

var PassportSvc micro.Service

func Init(passportSvc micro.Service) {
	PassportSvc = passportSvc
}
