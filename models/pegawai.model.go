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
	// defer rows.Close() // digunakan untuk mengakhirkan eksekusi sebuah statement tepat sebelum blok fungsi selesai

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

// fungsi store/post/insert
// disini ada beberapa parameter inputnya, sesuaikan dengan database kita. id tidak perlu, karena auto increment
// kemudian akan mereturn struct response & error
func StorePegawai(nama string, alamat string, telepon string) (Response, error) {
	var res Response

	con := db.CreateCon()
	
	// di sqlstatement ini ada variabel ? yg akan menyimpan value 
	sqlStatement := "INSERT pegawai (nama, alamat, telepon) VALUES (?, ?, ?)"

	// menampung eror statement
	// Prepare() akan mempersiapkan sql statement kita agar bisa dapat dieksekusi, dan akhirnya objek dapat digunakan berkali2
	stmt, err := con.Prepare(sqlStatement)

	// cek ada eror, jika ada eror return aja
	if err != nil {
		return res, err
	}

	// menampung return value dari stmt
	// nanti ini akan mengsisi VALUES (?, ?, ?)
	// urutan disini samakan dengan sqlStatement
	result, err := stmt.Exec(nama, alamat, telepon)
	if err != nil {
		return res, err
	}

	// variabel ini akan terisi dengan Last id didalam database
	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id" : lastInsertedId,
	}

	return res, nil
}

func UpdatePegawai(id int, nama string, alamat string, telepon string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE pegawai SET nama = ?, alamat = ?, telepon = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	// urutan disini samakan dengan di sqlStatements
	result, err := stmt.Exec(nama, alamat, telepon, id)
	if err != nil {
		return res, err
	}

	// kalau di store kita butuh lastinsertid, kalau disini kita butuh row affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	// setup response
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected" : rowsAffected,
	}

	return res, nil


}

func DeletePegawai(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM pegawai WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	// tampung berapa banyak row yg diupdate (row yang terhapus)
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	// atur response yang akan dilempar kedalam controller
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected" : rowAffected,
	}

	return res, nil

}