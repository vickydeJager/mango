package routers

import (
	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/app/admin/controllers"
	"github.com/louisevanderlith/mango/util/control"
)

func init() {
	setupMapping()

	beego.Router("/", &controllers.DefaultController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/comms", &controllers.CommsController{})
	beego.Router("/manufacturer", &controllers.ManufacturerController{})
	beego.Router("/model", &controllers.ModelController{})
	beego.Router("/subcategory", &controllers.SubCategoryController{})
	beego.Router("/site", &controllers.SiteController{})
	beego.Router("/site/:id([0-9]+)", &controllers.SiteController{}, "get:GetEdit")
}

func setupMapping() {
	uploadMap := make(control.MethodMap)
	//uploadMap["POST"] = enums.Admin
	//uploadMap["GET"] = enums.Admin

	control.AddControllerMap("/", uploadMap)
	control.AddControllerMap("/category", uploadMap)
	control.AddControllerMap("/comms", uploadMap)
	control.AddControllerMap("/manufacturer", uploadMap)
	control.AddControllerMap("/model", uploadMap)
	control.AddControllerMap("/subcategory", uploadMap)
	control.AddControllerMap("/site", uploadMap)

	beego.InsertFilter("/*", beego.BeforeRouter, control.FilterUI)
}
