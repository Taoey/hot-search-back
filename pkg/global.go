package pkg

import (
	"github.com/olebedev/config"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2"
)

// 全局变量
var LOG *zap.SugaredLogger
var GCF *config.Config //global config
var MongoSession *mgo.Session
