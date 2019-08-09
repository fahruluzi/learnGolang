package main

import (
	"fmt"
)

// Variabel - Tipe Data : String
var pelajaran1 = "Go"
var pelajaran2 = "Dart"

// ! Golang tidak boleh memiliki variabel kosong -- tambahkanlan tipe data atau string kosong
var sekolah string // ? Seperti ini

var namaDepan, namaBelakang = "Fahrul", "Fauzi"
var namaLengkap = namaDepan + " " + namaBelakang

func main() {
	fmt.Println("Saya belajar bahasa " + pelajaran1 + " Dan " + pelajaran2)
	fmt.Println("Nama Depan " + namaDepan)
	fmt.Println("Nama Saya " + namaLengkap)
}
