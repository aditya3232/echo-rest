package controllers

import (
	"echo-rest/models"
	"net/http"
	"strconv"

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

func UpdatePegawai(c echo.Context) error {
	id := c.FormValue("id") // menentukan row mana yang akan kita update // id harus diubah ke int dulu
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	telepon := c.FormValue("telepon")

	// konversi id dari string jadi int
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// menampung parmeter yg akan dibawa ke models
	result, err := models.UpdatePegawai(conv_id, nama, alamat, telepon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// jika udah ke models akan menampilkan JSON
	return c.JSON(http.StatusOK, result)



	
}