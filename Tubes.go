package main

import "fmt"

const NMAX int = 1000

type user struct {
	nama     string
	username string
	password string
}

type question struct {
	tag             string
	keluhan         string
	jumlahTag       int
	balasanDokter   string
	TanggapanPasien string
}

type arrUser [NMAX]user
type arrQuestion [NMAX]question

func header(T *arrUser) {

	fmt.Println("------------------------------------------------")
	fmt.Println("       Selamat datang di aplikasi konsultasi")
	fmt.Println("              Kesehatan Online")
	fmt.Println("                 Created By :                                ")
	fmt.Println("           Muhammad Farhan Editya")
	fmt.Println("             Rakha Bayu Pratama              ")
	fmt.Println("------------------------------------------------")
	fmt.Println()
	fmt.Println("                    - - - - - -")
	fmt.Println("     Kami di sini untuk menjaga kesehatan Anda!")
	fmt.Println("                    - - - - - -")
	fmt.Println()

}

func Menu(T arrUser, A arrQuestion, n int, U int) {
	var choose int
	var topic string

	fmt.Print("Silahkan masukkan keperluan anda: \n")

	fmt.Println("1. Registrasi")
	fmt.Println("2. Konsultasi")
	fmt.Println("3. Masuk sebagai dokter")
	fmt.Println("4. Lihat forum")
	fmt.Println("5. Cari pertanyaan")
	fmt.Println("6. Beri Tanggapan")
	fmt.Print("Pilihan: ")

	fmt.Scan(&choose)

	switch choose {
	case 1:
		register(&T, &U)

	case 2:
		if login(T) {
			postQuestion(&A, &n)
		} else {
			fmt.Println("Login gagal!")
			viewForum(A, n)
		}
	case 3:
		MenuDoctor(&A, T, n, U)

	case 4:
		viewForum(A, n)
	case 5:
		fmt.Print("Masukkan jenis pertanyaan yang akan anda cari: ")
		fmt.Scan(&topic)
		findTopic(A, topic, n)
	case 6:
		viewForum(A, n)
		replyDoctor(&A, n)

	}
	Menu(T, A, n, U)
}

func register(T *arrUser, U *int) {
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
		if T[i].username == "" && T[i].password == "" && T[i].nama == "" {
			T[i] = newUser
			*U++
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
	index := cekTag(*A, tag, *n)
	if index != -1 {
		A[index].jumlahTag++
	}
	if index == -1 {
		newQuestion := question{
			keluhan:         pertanyaan,
			tag:             tag,
			jumlahTag:       1,
			balasanDokter:   "",
			TanggapanPasien: "",
		}

		for i := 0; i < len(*A); i++ {
			if A[i].keluhan == "" && A[i].tag == "" {

				A[i] = newQuestion
				*n++
				fmt.Println("Pertanyaan berhasil diposting!")
				return
			}
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
		if A[i].TanggapanPasien != "" {
			fmt.Println("Tanggapan: ", A[i].TanggapanPasien)
		}
	}
}

func findTopic(A arrQuestion, topic string, n int) {
	var jumlah int = 0
	fmt.Println("Beberapa pertanyaan terkait: ")
	for i := 0; i < n; i++ {
		if A[i].tag == topic {

			fmt.Println(i+1, A[i].keluhan)
			if A[i].balasanDokter != "" {
				fmt.Println("Balasan Dokter:", A[i].balasanDokter)
			}
			jumlah++

		}

	}
	if jumlah == 0 {
		fmt.Println("Topik yang kamu cari belum ada ditanyakan nih,coba lagi nanti ya! ")
	}
}

func MenuDoctor(A *arrQuestion, T arrUser, n, U int) {
	var choose int

	fmt.Println("Selamat datang dokter")
	fmt.Println("1.Lihat forum ")
	fmt.Println("2.Balas pertanyaan pasien")
	fmt.Println("3.Lihat tag penyakit")
	fmt.Println("4.Cari pasien ")
	fmt.Println("5.Tampilkan pasien ")
	fmt.Print("Pilihan: ")
	fmt.Scan(&choose)

	switch choose {
	case 1:
		viewForum(*A, n)
		MenuDoctor(A, T, n, U)

	case 2:
		viewForum(*A, n)
		replyQuestion(A, n)
	case 3:
		sortTag(A, n)
		viewTag(*A, n)
	case 4:
		var U int
		findUser(*A, T, U, n)
	case 5:
		selectionSort(&T, U)
		viewUser(T, U)

	}

}
func findUser(A arrQuestion, T arrUser, U int, n int) {
	selectionSort(&T, U)
	fmt.Print("Masukkan nama yang ingin dicari: ")
	var searchName string
	fmt.Scan(&searchName)

	index := binarySearch(T, U, searchName)
	if index != -1 {
		fmt.Println("Keluhan pasien: ", A[index].keluhan)
		fmt.Println("Jenis Keluhan: ", A[index].tag)
	} else {
		fmt.Println("Pasien tidak ada")
	}

}

func binarySearch(T arrUser, U int, s string) int {

	var found int = -1
	var med int
	var kiri int = 0
	var kanan int = U - 1
	for kiri <= kanan && found == -1 {
		med = (kiri + kanan) / 2
		if s < T[med].nama {
			kanan = med - 1
		} else if s > T[med].nama {
			kiri = med + 1

		} else {
			found = med
		}

	}
	return found
}
func viewTag(A arrQuestion, n int) {

	fmt.Println("=== Jumlah Tag ===")
	for i := 0; i < n; i++ {
		fmt.Println(A[i].tag, A[i].jumlahTag)

	}
}
func cekTag(A arrQuestion, tag string, n int) int {
	for i := 0; i < n; i++ {
		if A[i].tag == tag {
			return i
		}
	}
	return -1
}
func viewUser(T arrUser, U int) {
	for i := 0; i < U; i++ {
		fmt.Println(i+1, T[i].nama)
	}
}

func sortTag(A *arrQuestion, n int) { //insertion sort
	for i := 1; i < n; i++ {
		j := i
		temp := A[j]
		for j > 0 && temp.jumlahTag > A[j-1].jumlahTag {
			A[j] = A[j-1]
			j--
		}
		A[j] = temp
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

func replyDoctor(A *arrQuestion, n int) {
	var reply string

	fmt.Print("Pilih nomor pertanyaan yang ingin anda tanggapi: ")
	var pilih int
	fmt.Scan(&pilih)
	if pilih > 0 && pilih <= n {
		fmt.Println("Pertanyaan Pasien: ", A[pilih-1].keluhan)
		fmt.Println("Saran dokter: ", A[pilih-1].balasanDokter)
		fmt.Print("Tanggapan anda: ")
		fmt.Scan(&reply)
		A[pilih-1].TanggapanPasien = reply
		fmt.Println("Tanggapan berhasil diposting")
	} else {
		fmt.Println("Nomor tidak valid")

	}

}

func selectionSort(T *arrUser, U int) {
	for i := 0; i < U-1; i++ {
		maxIndex := i
		for j := i + 1; j < U; j++ {
			if T[j].nama < T[maxIndex].nama {
				maxIndex = j
			}
		}
		T[i], T[maxIndex] = T[maxIndex], T[i]
	}
}

func main() {
	var T arrUser
	var A arrQuestion
	var n, U int
	header(&T)
	Menu(T, A, n, U)

}
