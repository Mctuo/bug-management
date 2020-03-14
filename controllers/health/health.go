package health

import (
	"bug-management/models"
	"github.com/astaxie/beego"
	"time"
)

type HealthController struct {
	beego.Controller
}

func (c *HealthController) Health() {
	c.Ctx.WriteString("200")
}

func (c *HealthController)GetNowTimes(){
	nowtimes := time.Now().UnixNano()/1e6
	models.HandleError(models.Success,models.GetErrMsg(models.Success,""),nowtimes,c.Ctx)
	return
}

