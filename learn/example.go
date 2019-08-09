package main

import (
	"fmt"
)

// Variabel - Tipe Data : String
var pelajaran1 = "Go"
var pelajaran2 = "Dart"

var namaDepan = "Fahrul"
var namaBelakang = "Fauzi"
var namaLengkap = namaDepan + " " + namaBelakang

func main() {
	fmt.Println("Saya belajar bahasa " + pelajaran1 + " Dan " + pelajaran2)
	fmt.Println("Nama Saya " + namaLengkap)
}
