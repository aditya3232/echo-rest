package models

import (
	"echo-rest/db"
	"net/http"
)

// disamakan dengan database
type Pegawai struct {
	Id      int    `json:"id"`
	Nama    string `json:"nama"`
	Alamat  string `json:"alamat"`
	Telepon string `json:"telepon"`
}

// fungsi mengambil data dari database
// disini mengembalikan respinse.go dan error
func FetchAllPegawai() (Response, error) {
	var obj Pegawai //obj berisi instance dari Pegawai struct
	var arrobj []Pegawai //array objek, karena bisa jadi di data pegawai kita gk cuma 1 value yg ada pada tabel, tapi bisa lebih dari 1 data pegawai
	var res Response  

	// membuat koneksi 
	con := db.CreateCon()

	// sql statement 
	sqlStatement := "SELECT * FROM pegawai"

	// variabel rows digunakan untuk menampung banyak data pegawai yang ada di sqlStatement
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	// jika ada error di Query(sqlStatement) maka return response, dan errornya
	// error akan diatur di controller
	if err != nil {
		return res, err
	}

	// ketika tidak ada error maka dihandle datanya
	// ketika ada datanya maka kita akan scan, namun sebelumnya ditampung dulu error nya
	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Nama, &obj.Alamat, &obj.Telepon) //masing2 kolom dimasukkan ke obj diatas. ini harus sesuai urutan 
		// jika ada error maka tampilkan res, dan error. controller yang akan menghandle errornya
		if err != nil {
			return res, err
		}

		// kalau tidak ada error, maka kita akan isi array objek kita 
		// isinya kita akan append langsung arrayobjek tsb dengan objek yang sudah kita scan diatas	 
		arrobj = append(arrobj, obj)
	}

	// ketika sudah selesai (berhasil mendapatkan array bojek) kita akan coba setup returnnya 
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	// di pegawai.model.go setiap return harus ada res & nil karena memang diatas kita define 2 value (Response, error)
	return res, nil
}
