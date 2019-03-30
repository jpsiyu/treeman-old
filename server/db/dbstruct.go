package db

type DBPerson struct {
	Name   string `json:"name"`
	Gender int    `json:"gender"`
	Age    int    `json:"age"`
}

type DBRecord struct {
	Timestamp int64  `json:"timestamp"`
	Detail    string `json:"detail"`
	Comment   string `json:"comment"`
}
