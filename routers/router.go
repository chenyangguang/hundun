package routers

import (
	"github.com/chenyangguang/gosto/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
