package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// SOAL 1
	fmt.Println()
	var kata1 string = "Bootcamp "
	var kata2 = "Digital "
	kata3 := "Skill "
	kata4 := "Sanbercode "
	var kata5 string
	kata5 = "Golang"

	fmt.Println(kata1 + kata2 + kata3 + kata4 + kata5)

	// SOAL 2
	fmt.Println()
	halo := "Halo Dunia"
	find := "Dunia"
	var replaceWith = "Golang"

	var newHalo = strings.Replace(halo, find, replaceWith, 1)
	fmt.Println(newHalo)

	// SOAL 3
	fmt.Println()
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	var gabung = kataPertama + " " + strings.Title(kataKedua) + " " + kataKetiga[:6] + 
	strings.ToUpper(kataKetiga[6:]) + " " + strings.ToUpper(kataKeempat)

	fmt.Println(gabung)

	// SOAL 4
	fmt.Println()
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	var angka1, _ = strconv.Atoi(angkaPertama)
	var angka2, _ = strconv.Atoi(angkaKedua)
	var angka3, _ = strconv.Atoi(angkaKetiga)
	var angka4, _ = strconv.Atoi(angkaKeempat)

	fmt.Println(angka1 + angka2 + angka3 + angka4)

	// SOAL 5
	fmt.Println()
	kalimat := "halo halo bandung"
	angka := 2021

	kalimat = strings.Replace(kalimat, "halo", "Hi", -1)

	var kalimatAkhir = `"` + kalimat + `" -`

	fmt.Printf("%s%d", kalimatAkhir, angka)
}