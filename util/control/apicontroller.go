package control

import (
	"github.com/astaxie/beego"
)

type APIController struct {
	beego.Controller
}

func (ctrl *APIController) Prepare() {
	output := ctrl.Ctx.Output

	output.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
	output.Header("Server", "kettle")
}

func (ctrl *APIController) ServeBinary(data []byte, filename string) {
	output := ctrl.Ctx.Output

	output.Header("Content-Description", "File Transfer")
	output.Header("Content-Type", "application/octet-stream")
	output.Header("Content-Disposition", "attachment; filename="+filename)
	output.Header("Content-Transfer-Encoding", "binary")
	output.Header("Expires", "0")
	output.Header("Cache-Control", "must-revalidate")
	output.Header("Pragma", "public")

	output.Body(data)
}

func (ctrl *APIController) Serve(err error, data interface{}) {

	if err != nil {
		ctrl.Ctx.Output.SetStatus(500)
		ctrl.Data["json"] = map[string]string{"Error": err.Error()}
	} else {
		ctrl.Data["json"] = map[string]interface{}{"Data": data}
	}

	ctrl.ServeJSON()
}
