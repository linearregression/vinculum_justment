package controllers

import (
	// "fmt"
	"justment/models"
	// "encoding/json"

	"github.com/astaxie/beego"
)

// Operations about object
type VinculumController struct {
	beego.Controller
}


// @Title Get
// @Description find object by objectid
// @Param	vinculumId		path 	string	true		"the vinculumId you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :vinculumId is empty
// @router /:vinculumId [get]
func (o *VinculumController) Get() {
	vinculumId := o.Ctx.Input.Param(":vinculumId")
	if vinculumId != "" {
		ob, err := models.GetOne(vinculumId)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
}