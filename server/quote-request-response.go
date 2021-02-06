package server

type RequestGetQuote struct {
	BathSize int      `json:"status"`
	Likes    []string `json:"likes"`
	Dislikes []string `json:"dislikes"`
}

type ResponseGetQuote struct {
	Quotes []Quote `json:"quotes"`
}

type Quote struct {
	Id       string   `json:"id"`
	Body     string   `json:"body"`
	Category []string `json:"category"`
	Source   string   `json:"source"`
	Likes    int64    `json:"likes"`
}

type Categories struct {
	List []string `json:"list"`
}

type RespError struct {
	Error string `json:"error"`
	Code  string `json:"code"`
}
