package sysinit

import (
	. "github.com/Taoey/hot-search-back/pkg"
	"gopkg.in/mgo.v2"
)

func InitMongo() {
	MGO_URL := GCF.UString("mongodb.url")
	MongoSession, _ = mgo.Dial(MGO_URL)
	MongoSession.SetMode(mgo.Monotonic, true)                 //连接模式设置
	MongoSession.SetPoolLimit(GCF.UInt("mongodb.pool_limit")) // 设置连接池数量
}
