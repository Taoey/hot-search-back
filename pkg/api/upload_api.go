package api

import (
	. "github.com/Taoey/hot-search-back/pkg/entity"
	"github.com/Taoey/hot-search-back/pkg/service"
	"github.com/kataras/iris/v12"
	"io/ioutil"
)

func UploadAliBill(ctx iris.Context) {
	file, _, _ := ctx.FormFile("file")
	bytes, _ := ioutil.ReadAll(file)
	s := string(bytes)
	service.OnUploadAliBillPrint(s)

	result := Message{
		Code: MESSAGE_OK,
	}
	ctx.JSON(result)
}
