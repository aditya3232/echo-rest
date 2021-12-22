package controllers

import (
	"echo-rest/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// namanya sama antara func model dan controller pada task yang sama
// disini berfungsi menghandle result ataupun error yang masuk dari package models. Disini berfungsi menampilkan JSON result ataupun error
// kemudian controller akan dipanggil didalam routes. karena dia langsung mengembalikan JSON
func FetchAllPegawai(c echo.Context) error {
	result, err := models.FetchAllPegawai()
	
	//ketika ada error maka return JSON
	//keynya message: errornya yg direturn dari model 
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	// ketika tidak ada error maka mereturn JSON. hasilnya direturn dari model
	return c.JSON(http.StatusOK, result)
}

func StorePegawai(c echo.Context) error {
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	telepon := c.FormValue("telepon")

	result, err := models.StorePegawai(nama, alamat, telepon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}