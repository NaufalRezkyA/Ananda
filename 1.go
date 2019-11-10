package main

import "fmt"

func hitungSkor(soal *int, skor *int, name *string) {
	var waktu int
	var nama string

	fmt.Scan(&nama)
	*name = nama
	*skor = 0
	*soal = 0
	for i := 1; i <= 8; i++ {
		fmt.Scan(&waktu)
		if waktu != 301 {
			*skor = *skor + waktu
			*soal = *soal + 1

		}
	}
}

func main() {
	var skor1, skor2, soal1, soal2 int
	var nama1, nama2 string

	hitungSkor(&soal1, &skor1, &nama1)
	fmt.Println(soal1, skor1)

	hitungSkor(&soal2, &skor2, &nama2)
	fmt.Println(soal2, skor2)
	fmt.Println("Selesai")

	if soal1 > soal2 {
		fmt.Println("\n", nama1, soal1, skor1)
	} else if soal2 > soal1 {
		fmt.Println("\n", nama2, soal2, skor2)
	} else {
		if skor1 > skor2 {
			fmt.Println("\n", nama1, soal1, skor1)
		} else if skor2 > skor1 {
			fmt.Println("\n", nama2, soal2, skor2)
		} else {
			fmt.Println("\nTotal skor keduanya sama")
		}
	}
}
