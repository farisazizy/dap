package main

import "fmt"

var jumMhs int = 1

const N, NMAX int = 50, 10000

type matakuliah struct { // tipe data baru untuk data suatu mata kuliah

	nama                  string
	UTS, UAS, QUIZ, TOTAL int
	grade                 string
}

type mahasiswa struct { // tipe data baru untuk data seorang mahasiswa

	nama, nim string
	matkul    [N]matakuliah
}

var mhs [NMAX]mahasiswa

func main() { // fungsi UTAMA

	var choose, ch2 int

	menu()
	fmt.Scanln(&choose)

	for choose < 0 || choose > 5 {

		fmt.Println()
		menu()
		fmt.Scanln(&choose)

	}

	for choose != 5 {

		fmt.Println()
		mhsMk()
		fmt.Scanln(&ch2)

		switch {
		case choose == 1 && ch2 == 1:
			createMhs()
		case choose == 1 && ch2 == 2:
			create_mk()
		case choose == 2 && ch2 == 1:
			read_mhs()
		case choose == 2 && ch2 == 2:
			read_mk()
		case choose == 3 && ch2 == 1:
			update_mhs()
		case choose == 3 && ch2 == 2:
			update_mk()
		case choose == 4 && ch2 == 1:
			delete_mhs()
		case choose == 4 && ch2 == 2:
			delete_mk()

		}

		fmt.Println()
		menu()
		fmt.Println()
		fmt.Scanln(&choose)

	}

	fmt.Println("\n\n\nTerimakasih telah menggunakan aplikasi ini.")
	fmt.Println("Tertanda, Adit dan Faris.\n")
	fmt.Printf("====================================================================================================\n\n")

}

func menu() { // Untuk menampilkan tampilan awal program

	fmt.Printf("====================================================================================================\n")
	fmt.Printf("Silahkan pilih menu untuk mulai menggunakan aplikasi ..\n")
	fmt.Printf("====================================================================================================\n\n")
	fmt.Printf("1. Input data baru.\n2. Lihat data lama.\n3. Update data lama.\n4. Delete data lama.\n5. Keluar dari aplikasi.\n\nPilih Menu: ")

}

func mhsMk() { // Untuk menampilkan

	fmt.Printf("1.Mahasiswa\n2.Mata Kuliah\n\nPilih menu: ")

}

func createMhs() {

	var jum_loc int

	fmt.Println()
	fmt.Print("Berapa jumlah mahasiswa yang ingin anda simpan?")
	fmt.Scanln(&jum_loc)

	i := 0

	for (jumMhs < NMAX) && (i < jum_loc) {

		fmt.Print("Input nama Mahasiswa ke-", jumMhs, ": ")
		fmt.Scanln(&mhs[jumMhs].nama)
		fmt.Print("Input NIM: ")
		fmt.Scanln(&mhs[jumMhs].nim)
		fmt.Println()

		jumMhs++
		i++
	}
}

func read_mhs() {

	var jum_loc int

	fmt.Println()
	fmt.Print("Berapa jumlah mahasiswa yang ingin anda lihat datanya?: ")
	fmt.Scanln(&jum_loc)
	fmt.Println()

	for i := 1; i <= jum_loc; i++ {

		fmt.Printf("Nama: %v \n", mhs[i].nama)
		fmt.Printf("NIM: %v \n", mhs[i].nim)
		fmt.Println()

	}

	fmt.Scanln()

}

func update_mhs() {

}

func delete_mhs() {

}

func create_mk() {

}

func read_mk() {

}

func update_mk() {

}

func delete_mk() {

}
