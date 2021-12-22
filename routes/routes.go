package routes

// seluruh routes akan terdefinisi disini
import (
	"echo-rest/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// routes ini mengembalikan package controllers yg berisi FetchAllPegawai() yang akan mereturn JSON
	e.GET("/pegawai", controllers.FetchAllPegawai)
	
	// routes post
	// post: mengambil data dari resource luar menuju database kita
	e.POST("/pegawai", controllers.StorePegawai)

	// routes update
	e.PUT("/pegawai", controllers.UpdatePegawai)


	return e
}