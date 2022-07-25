package main

import "fmt"

func main() {
	// SOAL 1
	fmt.Println("\n------Soal 1------")

	for angka := 1; angka <= 20; angka++ {
		if angka % 2 == 1 {
			if angka % 3 == 0 {
				fmt.Println(angka, "- I Love Coding")
			} else {
				fmt.Println(angka, "- Santai")
			}
		} else if angka % 2 == 0 {
			fmt.Println(angka, "- Berkualitas")
		}
	}

	// SOAL 2
	fmt.Println("\n------Soal 2------")

	for i := 0; i <= 7; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}

	// SOAL 3
	fmt.Println("\n------Soal 3------")

	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}

	fmt.Println(kalimat[2:])


	// SOAL 4
	fmt.Println("\n------Soal 4------")

	var sayuran = []string{}

	sayuran = append(sayuran, "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")

	for i, vegetable := range sayuran {
		fmt.Printf("%d. %s\n", i+1, vegetable)
	}

	// SOAL 5
	fmt.Println("\n------Soal 5------")

	var satuan = map[string]int{
	"panjang": 7,
	"lebar":   4,
	"tinggi":  6,
	}

	for key, value := range satuan {
		fmt.Println(key, "=", value)
	}
	fmt.Println("Volume Balok =", satuan["panjang"]*satuan["lebar"]*satuan["tinggi"])

	// Other ways
	volumebalok := 1
	batas := 0

	for i, satuan2 := range satuan {
		fmt.Println(i, "=", satuan2)
		volumebalok *= satuan2
		batas++
		if batas == len(satuan) {
			fmt.Println("volume Balok =", volumebalok)
		}
	}

}