package models // jgn lupa package disamain dengan nama folder

type Response struct {
	Status  int         `json:"status"`  //mengembalikan status seperti 404, 500, dll
	Message string      `json:"message"` //digunakan untuk menampung pesan yg dikembalikan dari server
	Data    interface{} `json:"data"`
}
