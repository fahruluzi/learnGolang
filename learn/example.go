package main

import (
	"fmt"
	"strconv"
)

// Variabel - Tipe Data : String
var pelajaran1 = "Go"
var pelajaran2 = "Dart"

// ! Golang tidak boleh memiliki variabel kosong -- tambahkanlan tipe data atau string kosong
var sekolah string // ? Seperti ini

var namaDepan, namaBelakang = "Fahrul", "Fauzi"
var namaLengkap = namaDepan + " " + namaBelakang

// Variabel - Tipe Data : Numerik = [Integer,Float]
var umur = 75
var jumlah int

// Variabel const (Data yang tidak bisa dirubah)
const gaji = 500000

func main() {
	// penulisan singkat variabel hanya bisa dipakai didalam fungsi
	sukamanah := "Sukamanah"

	fmt.Println("Saya belajar bahasa " + pelajaran1 + " Dan " + pelajaran2)
	fmt.Println("Nama Depan " + namaDepan)
	fmt.Println("Nama Saya " + namaLengkap)
	fmt.Println("Bersekolah di " + sukamanah)

	// Menyambungkan numerik ke string (Menggunakan koma) -- Cara simple
	fmt.Println("Umur saya ", umur, " tahun")

	// Operasi Matematika
	angkaPertama := 10
	angkaKedua := 20
	jumlah = angkaPertama + angkaKedua

	fmt.Println("Total ", jumlah)

	// gaji = 10 (Error)
	fmt.Println("Gaji Pasti ", gaji)

	// convert numerik kepada string (cek dokumentasi tentang strconv) -- Itoa
	gajiString := strconv.Itoa(gaji) // kebalikannya Atoi
	fmt.Println("Gaji Pasti " + gajiString)
}
