package main

type ServerMsg struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type ServerCode struct {
	Id   int    `json:"id"`
	Desc string `json:"Desc"`
}

var ServerCodeConfig = map[int]ServerCode{
	0: {Id: 0, Desc: "ok"},
	1: {Id: 1, Desc: "403 forbidden"},
	2: {Id: 2, Desc: "server error"},
}
