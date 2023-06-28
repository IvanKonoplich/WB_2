package controller

type response struct {
	Response string `json:"result"`
}

type responseError struct {
	Error string `json:"error"`
}

type incomingQuery struct {
	Date    string `json:"date"`
	UserID  int    `json:"user_id"`
	Message string `json:"message"`
}
