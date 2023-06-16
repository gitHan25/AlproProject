package main

import "fmt"

const NMAX int = 100

type user struct {
	nama     string
	username string
	password string
}

type question struct {
	tag           string
	keluhan       string
	jumlahTag     int
	balasanDokter string
}

type arrUser [NMAX]user
type arrQuestion [NMAX]question

func header(T *arrUser) {

	fmt.Println("------------------------------------------------")
	fmt.Println("       Selamat datang di aplikasi konsultasi")
	fmt.Println("              Kesehatan Online")
	fmt.Println("                 Created By :                                ")
	fmt.Println("           Muhammad Farhan Editya")
	fmt.Println("                 Rakha Bayu               ")
	fmt.Println("------------------------------------------------")
	fmt.Println()
	fmt.Println("                    - - - - - -")
	fmt.Println("     Kami di sini untuk menjaga kesehatan Anda!")
	fmt.Println("                    - - - - - -")
	fmt.Println()

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
	fmt.Print("Pilihan: ")

	fmt.Scan(&choose)

	switch choose {
	case 1:
		register(T)

	case 2:
		if login(*T) {
			postQuestion(&A, &n)
		} else {
			fmt.Println("Login gagal!")
			viewForum(A, n)
		}
	case 3:
		MenuDoctor(&A, n)

	case 4:
		viewForum(A, n)
	case 5:
		fmt.Print("Masukkan jenis pertanyaan yang akan anda cari: ")
		fmt.Scan(&topic)
		findTopic(A, topic, n)
	case 6:
		PrintUser(*T, A, n)
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

func login(T arrUser) bool {
	var username, password string

	fmt.Println("=== Login ===")
	fmt.Print("Masukkan username: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&password)

	for i := 0; i < len(T); i++ {
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

	newQuestion := question{
		keluhan:       pertanyaan,
		tag:           tag,
		balasanDokter: "",
	}

	for i := 0; i < len(*A); i++ {
		if A[i].keluhan == "" && A[i].tag == "" {
			A[i] = newQuestion
			*n++
			fmt.Println("Pertanyaan berhasil diposting!")
			return
		}
	}

	fmt.Println("Kesalahan: Kapasitas tanya sudah penuh!")
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
		if A[i].balasanDokter != "" {
			fmt.Println("Balasan Dokter:", A[i].balasanDokter)
		}
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

func MenuDoctor(A *arrQuestion, n int) {
	var choose int
	fmt.Println("Selamat datang dokter")
	fmt.Println("1.Lihat forum ")
	fmt.Println("2.Balas pertanyaan pasien")
	fmt.Println("3.Lihat tag penyakit")
	fmt.Print("Pilihan: ")
	fmt.Scan(&choose)

	switch choose {
	case 1:
		viewForum(*A, n)
		MenuDoctor(A, n)

	case 2:
		viewForum(*A, n)
		replyQuestion(A, n)
	case 3:
		sortTag(A, n)
		viewTag(*A, n)
	}

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
	fmt.Println("=== Jumlah Tag ===")

	for i := 0; i < n; i++ {

		tag := A[i].tag

		fmt.Printf("%s: %d\n", tag, countTag(A, A[i].tag, n))
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
func replyQuestion(A *arrQuestion, n int) {
	var reply string
	fmt.Print("Pilih nomor pertanyaan yang ingin Anda balas: ")
	var pilih int = 0
	fmt.Scan(&pilih)

	if pilih > 0 && pilih <= n {
		fmt.Println("Pertanyaan Pasien: ", A[pilih-1].keluhan)
		fmt.Print("Masukkan balasan Anda: ")
		fmt.Scan(&reply)
		A[pilih-1].balasanDokter = reply
		fmt.Println("Balasan berhasil ditambahkan!")
	} else {
		fmt.Println("Nomor pertanyaan tidak valid.")
	}
}
func PrintUser(T arrUser, A arrQuestion, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(T)
	}

	for i := 0; i < n; i++ {
		fmt.Println(A)
	}
}

func selectionSort(A *arrQuestion, n int) {
	var t question
	i := 0
	for i < n-1 {
		idx_min := i
		j := i + 1
		for j < n {
			if A[j].jumlahTag < A[idx_min].jumlahTag {
				idx_min = j
			}
			j = j + 1
		}
		t = A[idx_min]
		A[idx_min] = A[i]
		A[i] = t
		i = i + 1
	}
}

func main() {
	var T arrUser
	var A arrQuestion
	var n int
	header(&T)
	Menu(&T, A, n)

}
