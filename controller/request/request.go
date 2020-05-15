package request

type Echo struct {
	Data string `json:"data" binding:"required"`
}

type QuestionAdd struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
