package main

import "fmt"

type user struct {
	nama     string
	keluhan  string
	username string
	password string
}

type arrUser [100]user

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
	var users arrUser
	fmt.Print("Silahkan masukkan keperluan anda: \n")

	fmt.Println("1. Registrasi")
	fmt.Println("2. Login")
	fmt.Println("3. Lihat Masalah")
	fmt.Println("4. Quit")

	fmt.Scan(&choose)

	switch choose {
	case 1:
		register(&users)
	case 2:
		login(&users)
	case 3:
		fmt.Println("Sampai jumpa lagi")
	case 4:
		return
	default:
		fmt.Println("Pilihan tidak valid")
	}

	Menu()
}

func register(T *arrUser) {
	var username, password string

	fmt.Println("=== Registrasi ===")
	fmt.Print("Masukkan username: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&password)

	newUser := user{
		username: username,
		password: password,
	}

	for i := 0; i < len(*T); i++ {
		if T[i].username == "" && T[i].password == "" {
			T[i] = newUser
			fmt.Println("Registrasi berhasil!")
			return
		}

	}

	fmt.Println("Tidak dapat melakukan registrasi. Data sudah penuh.")
}

func login(T *arrUser) {
	var username, password string

	fmt.Println("=== Login ===")
	fmt.Print("Masukkan username: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&password)

	for i := 0; i < len(*T); i++ {
		if T[i].username == username && T[i].password == password {
			fmt.Println("Login berhasil!")
			return
		}
	}

	fmt.Println("Login gagal!")
}

func main() {
	header()
}
