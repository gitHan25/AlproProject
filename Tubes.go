package main

import "fmt"

type user struct {
	nama    string
	keluhan string
}

type dokter struct {
}

func header() {

	fmt.Println("------------------------------------------------")
	fmt.Println("       Selamat datang di aplikasi konsultasi")
	fmt.Println("              Kesehatan Online")
	fmt.Println("------------------------------------------------")
	fmt.Println()
	fmt.Println("                    - - - - - -")
	fmt.Println("     Kami di sini untuk menjaga kesehatan Anda!")
	fmt.Println("                    - - - - - -")
	fmt.Println()

	Menu()

}

func Menu() {
	var choose int
	fmt.Print("Silahkan masukkan keperluan anda: \n")

	fmt.Println("1. Registrasi")
	fmt.Println("2. Login")
	fmt.Println("3. Lihat Masalah")
	fmt.Println("4. Quit")

	fmt.Scan(&choose)

	switch choose {
	case 1:
		fmt.Println("Silahkan Registrasi terlebih dahulu")
	case 2:
		fmt.Println("Silahkan Login")
	case 3:
		fmt.Println("Sampai jumpa lagi")
	}

}
func main() {

	header()

}
