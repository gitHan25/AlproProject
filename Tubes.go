package main

import "fmt"

type user struct {
	nama     string
	username string
	password string
}

type question struct {
	tag     string
	keluhan string
}

type arrUser [100]user
type arrQuestion [100]question

func header(T *arrUser) {
	var A arrQuestion
	var n int
	fmt.Println("------------------------------------------------")
	fmt.Println("       Selamat datang di aplikasi konsultasi")
	fmt.Println("              Kesehatan Online")
	fmt.Println("------------------------------------------------")
	fmt.Println()
	fmt.Println("                    - - - - - -")
	fmt.Println("     Kami di sini untuk menjaga kesehatan Anda!")
	fmt.Println("                    - - - - - -")
	fmt.Println()

	Menu(T, A, n)
}

func Menu(T *arrUser, A arrQuestion, n int) {
	var choose int

	fmt.Print("Silahkan masukkan keperluan anda: \n")

	fmt.Println("1. Registrasi")
	fmt.Println("2. Konsultasi")
	fmt.Println("3. Masuk sebagai dokter")
	fmt.Println("4. Lihat forum")

	fmt.Scan(&choose)

	switch choose {
	case 1:
		register(T)

	case 2:
		if login(T) {
			postQuestion(&A, &n)
		} else {
			viewForum(A, n)
		}
	case 3:
		fmt.Println("aku dokter")
	case 4:
		viewForum(A, n)
	default:
		fmt.Println("Pilihan tidak valid")
	}
	Menu(T, A, n)
}

func register(T *arrUser) {
	var username, password, nama string

	fmt.Println("=== Registrasi ===")
	fmt.Println("Masukkan nama anda: ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan username: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&password)

	newUser := user{
		nama:     nama,
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

func login(T *arrUser) bool {
	var username, password string

	fmt.Println("=== Login ===")
	fmt.Print("Masukkan username: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&password)

	for i := 0; i < len(*T); i++ {
		if T[i].username == username && T[i].password == password {
			fmt.Println("Login berhasil! ")
			return true
		}
	}

	fmt.Println("Login gagal!")

	return false
}

func postQuestion(A *arrQuestion, n *int) {
	var pertanyaan, tag string
	fmt.Print("Mau konsultasi apa? (ketik 'cukup' untuk menghentikan): ")
	fmt.Scan(&pertanyaan)
	fmt.Print("Masukkan jenis keluhannya: ")
	fmt.Scan(&tag)
	*n = 0
	for pertanyaan != "cukup" && tag != "cukup" {

		A[*n].keluhan = pertanyaan
		A[*n].tag = tag
		*n++

		fmt.Print("Mau konsultasi apa? (ketik 'cukup' untuk menghentikan): ")
		fmt.Scan(&pertanyaan)
		fmt.Print("Masukkan jenis keluhannya: ")
		fmt.Scan(&tag)
	}
	fmt.Println("Pertanyaan berhasil diposting!")
}
func viewForum(A arrQuestion, n int) {
	fmt.Println("=== Forum Pertanyaan ===")
	for i := 0; i < n; i++ {
		fmt.Printf("Pertanyaan %d:\n", i+1)
		fmt.Println("Keluhan:", A[i].keluhan)
		fmt.Println("Tag:", A[i].tag)
	}
}

func main() {
	var T arrUser
	header(&T)
}
