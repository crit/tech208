package main

import (
	"encoding/json"
)

type Person struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (Person) TableName() string {
	return "people"
}

func PersonList() []Person {
	list := make([]Person, 0)

	cache := cacher.Get("people")

	if err := json.Unmarshal([]byte(cache), &list); cache != "" && err == nil {
		return list
	}

	db.Order("id desc").Find(&list)

	if len(list) > 0 {
		data, _ := json.Marshal(list)
		cacher.Set("people", string(data))
	}

	return list
}

func PersonCreate(name, email string) error {
	person := Person{Name: name, Email: email}
	err := db.Create(&person).Error

	if err == nil {
		cacher.Delete("people")
	}

	return err
}
