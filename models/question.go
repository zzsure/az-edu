package models

import (
	"az-edu/library/db"
	"az-edu/library/util"
)

type Quesiton struct {
	Model
	Title    string `gorm:"type:MEDIUMTEXT;comment:'标题，问题'" json:"title"`
	Content  string `gorm:"type:MEDIUMTEXT;comment:'内容，答案'" json:"content"`
	Weight   int    `gorm:"default:0;comment:'权重，见过的次数'" json:"weight"`
	TitleMd5 string `gorm:"type:varchar(32);unique;comment:'标题的md5，唯一'" json:"title"`
}

func (q *Quesiton) Save() error {
	return db.DB.Save(q).Error
}

func AddQuestion(title, content string) (*Quesiton, error) {
	titleMd5 := util.GetMd5Str(title)
	logger.Info("title md5 len:", len(titleMd5))
	q := &Quesiton{
		Title:    title,
		Content:  content,
		Weight:   0,
		TitleMd5: titleMd5,
	}
	err := q.Save()
	return q, err
}

type QuestionLabel struct {
	Model
	QuestionID uint `gorm:"comment:'question表id';not null" json:"question_id"`
	LabelID    uint `gorm:"comment:'label表id';not null" json:"label_id"`
}

func (ql *QuestionLabel) Save() error {
	return db.DB.Save(ql).Error
}

func AddQuestionLabel(questionID, labelID uint) error {
	ql := &QuestionLabel{
		QuestionID: questionID,
		LabelID:    labelID,
	}
	return ql.Save()
}
