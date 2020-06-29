# hot-search-back
热榜数据项目，后端


## 相关主要依赖库

- Mongo连接：gopkg.in/mgo.v2
- config配置读取 ：github.com/olebedev/config 
- 日志 ： zap + file-rotatelogs

## 跨域问题
- 详见：https://studyiris.com/example/exper/cors.html
- go get -u github.com/iris-contrib/middleware/...


## 开发环境
- go 1.14+
- mysql 5.7