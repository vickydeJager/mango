package controllers

import (
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/db/things"
)

type ModelController struct{
	util.SecureController
}

func (req *ModelController) Get(){
	var results []*things.Model
	mdl := things.Model{}
	err := things.Ctx.Model.Read(mdl, &results)

	if err != nil {
		req.Ctx.Output.SetStatus(500)
		req.Data["json"] = map[string]string{"Error": err.Error()}
	} else {
		req.Data["json"] = map[string]interface{}{"Data": results}
	}

	req.ServeJSON()
}