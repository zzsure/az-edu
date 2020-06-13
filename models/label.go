package models

import (
	"az-edu/library/db"
	"errors"
)

type Label struct {
	Model
	Name string `gorm:"type:varchar(32);unique;comment:'标签名称，唯一'" json:"name"`
}

func (l *Label) Save() error {
	return db.DB.Save(l).Error
}

func AddLabel(name string) error {
	if name == "" {
		return errors.New("label name not null")
	}
	l := &Label{
		Name: name,
	}
	return l.Save()
}