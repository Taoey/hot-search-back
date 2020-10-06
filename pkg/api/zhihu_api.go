package api

import (
	"fmt"
	. "github.com/Taoey/hot-search-back/pkg"
	. "github.com/Taoey/hot-search-back/pkg/entity"
	. "github.com/Taoey/hot-search-back/pkg/service"
	"github.com/kataras/iris/v12"
)

func ApiZhihuQuery(ctx iris.Context) {
	params := map[string]interface{}{}

	if err := ctx.ReadJSON(&params); err != nil {
		ctx.JSON(BadResponse("read request data err"))
		LOG.Error(err)
		return
	}
	fmt.Println(params)

	res, resErr := OnZhiHuItemsQuery()

	if resErr != nil {
		ctx.JSON(BadResponse(resErr.Error()))
	} else {
		ctx.JSON(OKResponse(res))
	}
}
