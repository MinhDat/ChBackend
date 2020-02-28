package models

type Response struct {
	Status string
	Data   interface{}
	Error  error
}
