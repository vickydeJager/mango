package controllers

import (
	"github.com/louisevanderlith/mango/util"
	"github.com/louisevanderlith/mango/app/admin/logic"
	"github.com/louisevanderlith/mango/util/enums"
)

type CommsController struct {
	util.UIController
}

func init() {
	auths := make(util.ActionAuth)
	auths["GET"] = enums.Admin

	util.ProtectMethods(auths)
}

func (c *CommsController) Get() {
	c.Setup("comms")

	data, err := logic.GetCommsMessages()

	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["error"] = err
	} else {
		c.Data["data"] = data
	}
}
