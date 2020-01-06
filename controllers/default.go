package controllers

import (
	"github.com/astaxie/beego"
)

type HeathController struct {
	beego.Controller
}

func (c *HeathController) Health() {
	c.Ctx.WriteString("200")
}
