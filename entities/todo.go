package entities

type Todo struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  int    `json:"status"`
	UserId  int64  `json:"userId"`
}
