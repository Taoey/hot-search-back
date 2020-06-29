package pkg

import (
	"database/sql"
	"github.com/olebedev/config"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2"
)

// 全局变量

// 日志记录器
var LOG *zap.SugaredLogger

// 全局配置
var GCF *config.Config //global config

// Mongo连接
var MongoSession *mgo.Session

// mysql 连接（已经配置了连接池）
var MysqlSession *sql.DB
