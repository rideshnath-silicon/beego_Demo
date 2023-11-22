package models

import "github.com/astaxie/beego/orm"

func GetAllStudent() (interface{}, error) {
	o := orm.NewOrm()
	var student []orm.Params
	_, err := o.QueryTable(new(Students)).Values(&student, "first_name", "last_name", "email", "roll_number")
	if err != nil {
		return nil, err
	}
	return student, nil
}

func GetClassWiseStudent(id uint) (interface{}, error) {
	o := orm.NewOrm()
	var student []orm.Params
	_, err := o.QueryTable(new(Students)).Filter("class_id", id).Values(&student, "first_name", "last_name", "email", "roll_number")
	if err != nil {
		return nil, err
	}
	return student, nil
}
