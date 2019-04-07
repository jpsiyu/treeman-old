package main

type ServerMsg struct {
	Ok       bool        `json:"ok"`
	ErrorMsg string      `json:"errorMsg"`
	Data     interface{} `json:"data"`
}
