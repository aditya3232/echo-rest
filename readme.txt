1.  folder config akan diisi dengan konfigurasi dari golang, seperti konfigurasi database
2.  folder controllers akan handle semua file controller
3.  folder db akan diisi dengan konfiguasi yang berhubungan dengan database, misal mengembalikan instance, atau menginisialisasi instance dari database
4.  folder models akan berisi models
5.  folder routes berhubungan dengan endpoint yang diakses
6.  untuk mengatur dependency seperti npm/composer, maka golang membutuhkan golang/dep. namun saat ini windows masih belum bisa. tidak usah sebenernya tidak apa
7.  install echo framework dengan cara
    - go mod init myapp
    - go get github.com/labstack/echo/v4
8.  

