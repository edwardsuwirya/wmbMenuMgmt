package delivery

import (
	"github.com/edwardsuwirya/wmbMenuMgmt/delivery/appresponse"
	"github.com/edwardsuwirya/wmbMenuMgmt/usecase"
	"github.com/gin-gonic/gin"
)

type MenuApi struct {
	usecase     usecase.IMenuUseCase
	publicRoute *gin.RouterGroup
}

func NewMenuApi(publicRoute *gin.RouterGroup, usecase usecase.IMenuUseCase) *MenuApi {
	MenuApi := MenuApi{
		usecase:     usecase,
		publicRoute: publicRoute,
	}
	MenuApi.initRouter()
	return &MenuApi
}
func (api *MenuApi) initRouter() {
	userRoute := api.publicRoute.Group("/menu")
	userRoute.GET("", api.getMenuList)
}
func (api *MenuApi) getMenuList(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")
	if id != "" {
		menu, err := api.usecase.SearchMenuById(id)
		if err != nil {
			appresponse.NewJsonResponse(c).SendError(appresponse.NewInternalServerError(err, "Failed Get Menu"))
			return
		}
		appresponse.NewJsonResponse(c).SendData(appresponse.NewResponseMessage("SUCCESS", "Menu By ID", menu))
		return
	}

	if name != "" {
		menulist, err := api.usecase.SearchMenuByName(name)
		if err != nil {
			appresponse.NewJsonResponse(c).SendError(appresponse.NewInternalServerError(err, "Failed Get Menu"))
			return
		}
		appresponse.NewJsonResponse(c).SendData(appresponse.NewResponseMessage("SUCCESS", "Menu List by name", menulist))
		return
	}

	menulist, err := api.usecase.GetAllMenu()
	if err != nil {
		appresponse.NewJsonResponse(c).SendError(appresponse.NewInternalServerError(err, "Failed Get Menu"))
		return
	}

	appresponse.NewJsonResponse(c).SendData(appresponse.NewResponseMessage("SUCCESS", "Menu List", menulist))
}
