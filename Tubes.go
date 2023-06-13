package main

import "fmt"

const NMAX int = 100

type user struct {
	nama     string
	username string
	password string
}

type question struct {
	tag       string
	keluhan   string
	jumlahTag int
}

type arrUser [NMAX]user
type arrQuestion [NMAX]question

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
	var topic string

	fmt.Print("Silahkan masukkan keperluan anda: \n")

	fmt.Println("1. Registrasi")
	fmt.Println("2. Konsultasi")
	fmt.Println("3. Masuk sebagai dokter")
	fmt.Println("4. Lihat forum")
	fmt.Println("5. Cari pertanyaan")

	fmt.Scan(&choose)

	switch choose {
	case 1:
		register(T)

	case 2:
		if login(T) {
			postQuestion(&A, &n)
		} else {
			fmt.Println("Login gagal!")
			viewForum(A, n)
		}
	case 3:
		sortTag(&A, n)
		viewTag(A, n)
	case 4:
		viewForum(A, n)
	case 5:
		fmt.Print("Masukkan jenis pertanyaan yang akan anda cari: ")
		fmt.Scan(&topic)
		findTopic(A, topic, n)
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
		A[*n].jumlahTag = countTag(*A, tag, *n)
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
	if n == 0 {
		fmt.Println("Belum ada pertanyaan nih")
	}
	fmt.Println("=== Forum Pertanyaan ===")
	for i := 0; i < n; i++ {
		fmt.Printf("Pertanyaan %d:\n", i+1)
		fmt.Println("Keluhan:", A[i].keluhan)
		fmt.Println("Tag:", A[i].tag)

	}

}

func findTopic(A arrQuestion, topic string, n int) {
	var jumlah int = 0
	fmt.Println("Beberapa pertanyaan terkait: ")
	for i := 0; i < n; i++ {
		if A[i].tag == topic {

			fmt.Println(i+1, A[i].keluhan)
			jumlah++

		}

	}
	if jumlah == 0 {
		fmt.Println("Topik yang kamu cari belum ada ditanyakan nih,coba lagi nanti ya! ")
	}
}

func MenuDoctor() {
	fmt.Print("Selamat datang dokter")

}

func countTag(A arrQuestion, tag string, n int) int {
	var count int = 0
	for i := 0; i < n; i++ {
		if A[i].tag == tag {
			count++
		}
	}
	return count
}

func viewTag(A arrQuestion, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(A[i].tag, A[i].jumlahTag+1)
	}
}

func sortTag(A *arrQuestion, n int) {
	i := 1
	for i <= n-1 {
		j := i
		temp := A[j]
		for j > 0 && temp.jumlahTag > A[j-1].jumlahTag {
			A[j] = A[j-1]
			j--
		}
		A[j] = temp
		i++
	}
}
func main() {
	var T arrUser
	header(&T)

}
