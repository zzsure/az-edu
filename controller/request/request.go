package request

type Echo struct {
	Data string `json:"data" binding:"required"`
}

type QuestionAdd struct {
	Title    string `required:"true" json:"title"`
	Content  string `json:"content"`
	LabelIDS []uint `json:"label_ids"`
}

type LabelAdd struct {
	Name string `json:"name"`
}
