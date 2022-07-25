package main

import "fmt"

// SOAL 1
func luasPersegiPanjang(panjang int, lebar int) int {
	return panjang * lebar
}
func kelilingPersegiPanjang(panjang int, lebar int) int {
	return 2*panjang + 2*lebar
}
func volumeBalok(panjang int, lebar int, tinggi int) int {
	return panjang * lebar * tinggi
}

// SOAL 2
func introduce(nama string, gender string, pekerjaan string, usia string) (gabung string) {
	if gender == "laki-laki" {
		gabung = "Pak " + nama + " adalah seorang " + pekerjaan + " yang berusia " + usia + " tahun"
		return
	} else if gender == "perempuan" {
		gabung = "Bu " + nama + " adalah seorang " + pekerjaan + " yang berusia " + usia + " tahun"
		return
	}
	return
}

// SOAL 3
func buahFavorit(nama string, buah ...string) (gabung string) {
	var namabuah string
	batas := 0
	for _, fruit := range buah {
		namabuah = namabuah + `"` + fruit + `"`
		batas++
		if batas < len(buah) {
			namabuah += ", "
		}
	}
	gabung = "halo nama saya " + nama + " dan buah favorit saya adalah " + namabuah
	return
}

func main() {
	// SOAL 1
	fmt.Println("------Soal 1------")
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)

	// SOAL 2
	fmt.Println("------Soal 2------")
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	// SOAL 3
	fmt.Println("------Soal 3------")
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"


	// SOAL 4
	fmt.Println("------Soal 4------")
	var dataFilm = []map[string]string{}
	// buatlah closure function disini
	var tambahDataFilm = func(judul string, durasi string, genre string, tahun string) {
		dataFilm = append(dataFilm, map[string]string{"genre": genre, "jam": durasi, "tahun": tahun, "title": judul})
	}


	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
	fmt.Println(item)
	}
}