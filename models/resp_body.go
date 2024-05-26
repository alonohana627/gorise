package models

type ErrorJSON struct {
	Message  string `json:"message"`
	ErrorNum int    `json:"error_num"`
}

type SuccessJSON struct {
	Message string `json:"message"`
}

type AlreadyExistJSON struct {
	Message string
	Contact
}
