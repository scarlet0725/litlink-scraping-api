package model

type Log struct {
	Level string `json:"level"`
	Msg   string `json:"msg"`
	ip    string `json:"ip"`
}
