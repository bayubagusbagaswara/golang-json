# Golang JSON

# Pengenalan JSON

- JSON singkatan dari JavaScript Object Notation, merupakan struktur format data yang bentuknya seperti Object di JavaScript
- JSON merupakan struktur format data yang paling banyak digunakan saat kita membuat RESTful API

# Package json

- Go-Lang sudah menyediakan package json, dimana kita bisa menggunakan package ini untuk melakukan konversi data ke JSON (encode) atau sebaliknya (decode)

# Encode JSON

- Go-Lang telah menyediakan function untuk melakukan `konversi data ke JSON`, yaitu menggunakan function `json.Marshal(interface{})`
- Karena parameter nya adalah `interface{}`, maka kita `bisa masukan tipe data apapun` ke dalam function Marshal

# JSON Object

- Sebelumnya kita melakukan encode data sepertin string, number, boolean, dan tipe data primitif lainnya
- Walaupun memang bisa dilakukan, karena sesuai dengan kontrak interface{}. Namun tidak sesuai dengan kontrak JSON
- Jika mengikuti kontrak json.org, data JSON bentuknya adalah Object dan Array
- Sedangkan value nya baru berupa string, number, boolean dan lain-lain

## Struct

- JSON Object di Go-Lang direpresentasikan dengan tipe data Struct
- Dimana tiap attribute di JSON Object merupakan attribute di Struct
