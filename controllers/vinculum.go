package controllers

import (
	"justment/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about object
type VinculumController struct {
	beego.Vinculum
}


// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (o *ObjectController) Get() {
	vinculumId := o.Ctx.Input.Param(":vinculumId")
	if vinculumId != "" {
		ob, err := models.GetOne(objectId)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
}