package main

import "fmt"

func cetakDeret(n *int) {
	var angka int

	fmt.Scanln(&angka)
	for angka != 1 {
		if angka%2 == 0 {
			angka = angka / 2
			*n = angka
		} else {
			angka = 3*angka + 1
			*n = angka
		}
		angka = *n
	}
}

func main() {
	var total int

	cetakDeret(&total)
	fmt.Print(total)

}
