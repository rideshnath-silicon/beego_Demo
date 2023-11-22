package controllers

import (
	"CrudDemo/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	data, err := models.GetAllBooks()
	if err != nil {
		c.Data["json"] = err
		c.ServeJSON()
		return
	}
	c.Data["json"] = data
	c.ServeJSON()
}

func (c *MainController) CreateNew() {
	var book models.InsertBook
	json.Unmarshal(c.Ctx.Input.RequestBody, &book)
	data, err := models.InsertnewBook(book)
	if err != nil {
		c.Data["json"] = err
		c.ServeJSON()
		return
	}
	c.Data["json"] = data
	c.ServeJSON()
}

func (c *MainController) GetSingleBook() {
	var data models.GetBooks
	bodyData := c.Ctx.Input.RequestBody
	// Unmarshal JSON data into the map variable
	err := json.Unmarshal(bodyData, &data)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	output, err := models.GetSingleBook(data)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"data": output, "messge": "successfull"}
	c.Data["message"] = "succssfull"
	c.ServeJSON()
}

func (c *MainController) UpdateBook() {
	var data models.UpdateBooks
	bodyData := c.Ctx.Input.RequestBody
	err := json.Unmarshal(bodyData, &data)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	output, err := models.UpdateBook(data)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	c.Data["json"] = output
	c.ServeJSON()
}

func (c *MainController) DeleteBook() {
	var book models.GetBooks
	bodyData := c.Ctx.Input.RequestBody
	err := json.Unmarshal(bodyData, &book)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	output, err := models.DeleteBook(book.Id)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	c.Data["json"] = output
	c.ServeJSON()
}
