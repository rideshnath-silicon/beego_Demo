package controllers

import (
	"CrudDemo/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

type StudentController struct {
	beego.Controller
}

func (c *StudentController) Get() {
	data, err := models.GetAllStudent()
	if err != nil {
		c.Data["json"] = err
		c.ServeJSON()
		return
	}
	c.Data["json"] = data
	c.ServeJSON()
}

func (c *StudentController) GetClassWise() {
	var student models.ClassWiseStudent
	bodyData := c.Ctx.Input.RequestBody
	err := json.Unmarshal(bodyData, &student)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	output, err := models.GetClassWiseStudent(uint(student.ClassId))
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	c.Data["json"] = output
	c.ServeJSON()
}
