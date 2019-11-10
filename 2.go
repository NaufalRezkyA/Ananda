package main

import "fmt"

func fact(n int) int {
	jumlah := 1
	i := 1
	for n >= i {
		jumlah = jumlah * n
		n = n - 1
	}
	return jumlah

}

func main() {
	var angka, batas, number int

	batas = 0
	fmt.Scanln(&angka)
	fmt.Println(fact(angka))
	for angka > batas {
		fmt.Scanln(&number)
		fmt.Println(fact(number))
		batas = batas + number
	}
}
