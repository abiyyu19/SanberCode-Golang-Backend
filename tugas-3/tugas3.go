package main

import (
	"fmt"
	"strconv"
)

func main() {

	// SOAL 1
	fmt.Println("------Soal 1------")
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	intPanjangPersegiPanjang, _ := strconv.Atoi(panjangPersegiPanjang)
	intLebarPersegiPanjang, _ := strconv.Atoi(lebarPersegiPanjang)
	intAlasSegitiga, _ := strconv.Atoi(alasSegitiga)
	intTinggiSegitiga, _ := strconv.Atoi(tinggiSegitiga)

	var luasPersegiPanjang int
	var kelilingPersegiPanjang int
	var luasSegitiga int

	luasPersegiPanjang = intPanjangPersegiPanjang * intLebarPersegiPanjang
	kelilingPersegiPanjang = 2 * intPanjangPersegiPanjang + 2 * intLebarPersegiPanjang
	luasSegitiga = intAlasSegitiga * intTinggiSegitiga / 2

	fmt.Println("Panjang Persegi Panjang = ", intPanjangPersegiPanjang)
	fmt.Println("Lebar Persegi Panjang = ", intLebarPersegiPanjang)
	fmt.Println("\nAlas Segitiga = ", intAlasSegitiga)
	fmt.Println("Tinggi Segitiga = ", intTinggiSegitiga)
	fmt.Println("==========================")
	fmt.Println("Luas Persegi Panjang = ", luasPersegiPanjang)
	fmt.Println("Keliling Persegi Panjang = ", kelilingPersegiPanjang)
	fmt.Println("Luas Segitiga = ", luasSegitiga)


	// SOAL 2
	fmt.Println()
	
	fmt.Println("\n------Soal 2------")
	var nilaiJohn = 80
	var nilaiDoe = 50

	if nilaiJohn >= 80 {
		fmt.Println("Nilai John adalah", nilaiJohn, "dan berIndeks A")
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		fmt.Println("Nilai John adalah", nilaiJohn, "dan berIndeks B")
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		fmt.Println("Nilai John adalah", nilaiJohn, "dan berIndeks C")
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		fmt.Println("Nilai John adalah", nilaiJohn, "dan berIndeks D")
	} else if nilaiJohn < 50 {
		fmt.Println("Nilai John adalah", nilaiJohn, "dan berIndeks E")
	}

	if nilaiDoe >= 80 {
		fmt.Println("Nilai Doe adalah", nilaiDoe, "dan berIndeks A")
	} else if nilaiDoe >= 70 && nilaiDoe < 80 {
		fmt.Println("Nilai Doe adalah", nilaiDoe, "dan berIndeks B")
	} else if nilaiDoe >= 60 && nilaiDoe < 70 {
		fmt.Println("Nilai Doe adalah", nilaiDoe, "dan berIndeks C")
	} else if nilaiDoe >= 50 && nilaiDoe < 60 {
		fmt.Println("Nilai Doe adalah", nilaiDoe, "dan berIndeks D")
	} else if nilaiDoe < 50 {
		fmt.Println("Nilai Doe adalah", nilaiDoe, "dan berIndeks E")
	}

	// SOAL 3
	fmt.Println("\n------Soal 3------")
	var tanggal = 6
	var bulan = 8
	var tahun = 2001

	switch bulan {
	case 1:
		fmt.Println(strconv.Itoa(tanggal), "Januari", strconv.Itoa(tahun))
	case 2:
		fmt.Println(strconv.Itoa(tanggal), "Febuari", strconv.Itoa(tahun))
	case 3:
		fmt.Println(strconv.Itoa(tanggal), "Maret", strconv.Itoa(tahun))
	case 4:
		fmt.Println(strconv.Itoa(tanggal), "April", strconv.Itoa(tahun))
	case 5:
		fmt.Println(strconv.Itoa(tanggal), "Mei", strconv.Itoa(tahun))
	case 6:
		fmt.Println(strconv.Itoa(tanggal), "Juni", strconv.Itoa(tahun))
	case 7:
		fmt.Println(strconv.Itoa(tanggal), "Juli", strconv.Itoa(tahun))
	case 8:
		fmt.Println(strconv.Itoa(tanggal), "Agustus", strconv.Itoa(tahun))
	case 9:
		fmt.Println(strconv.Itoa(tanggal), "September", strconv.Itoa(tahun))
	case 10:
		fmt.Println(strconv.Itoa(tanggal), "Oktober", strconv.Itoa(tahun))
	case 11:
		fmt.Println(strconv.Itoa(tanggal), "November", strconv.Itoa(tahun))
	case 12:
		fmt.Println(strconv.Itoa(tanggal), "Desember", strconv.Itoa(tahun))
	}

	// SOAL 4
	fmt.Println("\n------Soal 4------")

	if tahun >= 1944 && tahun <= 1964 {
		fmt.Println("Tahun Kelahiran Anda Adalah", tahun, ", Anda adalah Baby Boomer")
	} else if tahun >= 1965 && tahun <= 1979 {
		fmt.Println("Tahun Kelahiran Anda Adalah", tahun, ", Anda adalah Generasi X")
	} else if tahun >= 1980 && tahun <= 1994 {
		fmt.Println("Tahun Kelahiran Anda Adalah", tahun, ", Anda adalah Generasi Y (Millenials)")
	} else if tahun >= 1995 && tahun <= 2015 {
		fmt.Println("Tahun Kelahiran Anda Adalah", tahun, ", Anda adalah Generasi Z")
	}
}