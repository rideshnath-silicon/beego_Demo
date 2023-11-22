package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=postgres password=root dbname=mydb sslmode=disable")
	orm.RegisterModel(new(Book), new(Students), new(Users))
	orm.RunSyncdb("default", false, true)
}

func GetAllBooks() ([]Book, error) {
	o := orm.NewOrm()
	var bokks []Book
	_, err := o.QueryTable(new(Book)).All(&bokks)
	if err != nil {
		return nil, err
	}
	return bokks, nil

}
// dsfsfsfsf
func InsertnewBook(data InsertBook) (interface{}, error) {
	o := orm.NewOrm()
	book := Book{
		Name:      data.Name,
		Author:    data.Author,
		CreatedAt: time.Now(),
	}

	_, err := o.Insert(&book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func GetSingleBook(data GetBooks) (interface{}, error) {
	o := orm.NewOrm()
	book := Book{Id: data.Id}
	err := o.Read(&book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func UpdateBook(data UpdateBooks) (interface{}, error) {
	o := orm.NewOrm()
	book := Book{Id: data.Id, Author: data.Author, Name: data.Name}
	_, err := o.Update(&book, "author", "name")
	if err != nil {
		return nil, err
	}
	return book, nil
}

func DeleteBook(id int) (interface{}, error) {
	o := orm.NewOrm()
	book := Book{Id: id}
	orm.Debug = true
	_, err := o.Delete(&book)
	if err != nil {
		return nil, err
	}
	return book, nil
}
