package main

import "fmt"

// JumMK adalah jumlah maksimal mata kuliah
const JumMK int = 15

// MhsMax adalah jumlah maksimal mahasiswa
const MhsMax int = 100

// matakuliah adalah tipe data baru untuk nama Matakuliah dan Nilai nya
type matakuliah struct {
	nama                 string
	sks                  int
	na, clo1, clo2, clo3 float64
}

// mahasiswa adalah tipe data baru untuk mendata mahasiswa, nim, dan matkul yang diambil
type mahasiswa struct {
	nama, nim string
	mk        [JumMK]matakuliah
	ipk       float64
}

//mhsw adalah tipe data untuk membentuk var dengan tipe array(tipe alias)
type mhsw [MhsMax]mahasiswa

// JumMhs adalah variabel global
var JumMhs int

func main() {

	var (
		JumInput int
		ArrMhs   mhsw // Array utama
		NamaMhs  string
		nim      string
		pilmenu  int
	)

	Menu()               // Menampilkan Menu
	fmt.Scanln(&pilmenu) // Memilih Menu

	for pilmenu < 1 || pilmenu > 7 { // Validasi input menu dari user harus pada interval 1-7

		fmt.Println("\nMaaf, tidak ada pilihan dalam menu :)\n")
		Menu()
		fmt.Scanln(&pilmenu)
	}

	for pilmenu != 6 { // Mengakses pilihan menu

		if pilmenu == 1 { // Input Data Mahasiswa

			fmt.Println()
			fmt.Print("Berapa jumlah mahasiswa yang ingin anda simpan? ")
			fmt.Scanln(&JumInput)
			CreateMhs(JumInput, &ArrMhs)

		} else if pilmenu == 2 { // Input Data Mata Kuliah

			fmt.Println()
			fmt.Print("Data mata kuliah siapa yang ingin anda tambahkan? ")
			fmt.Scanln(&NamaMhs)
			CreateMK(NamaMhs, &ArrMhs)

		} else if pilmenu == 3 { // Output Data Mahasiswa

			fmt.Print("Data diri siapa yang ingin anda lihat?: ")
			fmt.Scanln(&NamaMhs)
			ReadMhs(NamaMhs, ArrMhs)

		} else if pilmenu == 4 { // Output Data Mata Kuliah

			fmt.Print("Mata kuliah siapa yang ingin anda lihat?: ")
			fmt.Scanln(&NamaMhs)
			ReadMK(NamaMhs, ArrMhs)

		} else if pilmenu == 5 { // Update Data Mahasiswa

			fmt.Print("Masukan NIM: ")
			fmt.Scanln(&nim)
			UpMhs(NamaMhs, nim, &ArrMhs)

		} else if pilmenu == 6 { // Update Data Mahasiswa

			fmt.Print("Masukan NIM: ")
			fmt.Scanln(&nim)
			UpMK(nim, &ArrMhs)

		}

		Menu()               // Menampilkan Menu kembali setelah mengakses 1-6
		fmt.Scanln(&pilmenu) // Memilih menu kembali setelah mengakses 1-6
	}

	// Jika pilihan menu adalah 7, Keluar dari Aplikasi
	// Jika aplikasi ditutup akan menampilkan ini
	fmt.Println("Terimakasih telah menggunakan aplikasi ini.")
	fmt.Println("Tertanda, Adit dan Faris.")

	// NOTE
	/*	INI UNTUK NGECEK DATA TERSIMPAN DENGAN BAIK ATAU TIDAK!!

		for i:=0;i<JumMhs;i++ {

			fmt.Println(ArrMhs[i].nama)
		}
	*/
}

// Menu adalah prosedur untuk menampilkan halaman awal program
func Menu() {

	fmt.Println(".-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.\n")
	fmt.Println("Menu: ")
	fmt.Println("1. Input Data Mahasiswa")
	fmt.Println("2. Input Data Mata Kuliah")
	fmt.Println("3. Output Data Mahasiswa")
	fmt.Println("4. Output Data Mata Kuliah")
	fmt.Println("5. Update Data Mahasiswa")
	fmt.Println("6. Update Nilai Suatu Mata Kuliah")
	fmt.Println("7. Keluar dari Aplikasi.")
	fmt.Print("Menu Anda: ")

}

// CreateMhs adalah prosedur untuk menambah data mahasiswa
func CreateMhs(JumLoc int, ArrMhs *mhsw) {

	var qna string

	i := 0
	for (JumMhs < MhsMax) && (i < JumLoc) {

		fmt.Print("Input nama Mahasiswa ke-", JumMhs+1, ": ")
		fmt.Scanln(&ArrMhs[JumMhs].nama)
		fmt.Print("Input NIM: ")
		fmt.Scanln(&ArrMhs[JumMhs].nim)
		fmt.Print("Input Mata Kuliah?(Ya/Tidak): ")
		fmt.Scanln(&qna)

		if qna == "Ya" || qna == "ya" {

			CreateMK(ArrMhs[JumMhs].nama, ArrMhs)
		}

		fmt.Println()

		JumMhs++
		i++
	}
}

// CreateMK adalah prosedur untuk menambah data mata kuliah pada mahasiswa
func CreateMK(NamaMhs string, ArrMhs *mhsw) {

	var i, sks int
	var matkul string
	var clo1, clo2, clo3 float64

	i = 0
	for i < JumMhs && ArrMhs[i].nama != NamaMhs {

		i++
	}

	if ArrMhs[i].nama == NamaMhs {

		j := 0
		for ArrMhs[i].mk[j].nama != "" {

			j++
		}

		fmt.Printf("Mata Kuliah ke %v(Isi STOP untuk berhenti) : ", j+1)
		fmt.Scanln(&matkul)

		if matkul != "STOP" && matkul != "stop" {

			ArrMhs[i].mk[j].nama = matkul
			fmt.Print("Jumlah SKS: ")
			fmt.Scanln(&sks)
			ArrMhs[i].mk[j].sks = sks

			fmt.Print("NIlai CLO 1: ")
			fmt.Scanln(&clo1)
			ArrMhs[i].mk[j].clo1 = clo1
			fmt.Print("NIlai CLO 2: ")
			fmt.Scanln(&clo2)
			ArrMhs[i].mk[j].clo2 = clo2
			fmt.Print("NIlai CLO 3: ")
			fmt.Scanln(&clo3)
			ArrMhs[i].mk[j].clo3 = clo3

			ArrMhs[i].mk[j].na = HitungNA(clo1, clo2, clo3)
		}

		for matkul != "STOP" && matkul != "stop" && j < JumMK {

			for ArrMhs[i].mk[j].nama != "" {

				j++
			}

			fmt.Printf("Mata Kuliah ke %v: ", j+1)
			fmt.Scanln(&matkul)

			if matkul != "STOP" && matkul != "stop" {

				ArrMhs[i].mk[j].nama = matkul
				fmt.Print("Jumlah SKS: ")
				fmt.Scanln(&sks)
				ArrMhs[i].mk[j].sks = sks

				fmt.Print("NIlai CLO 1: ")
				fmt.Scanln(&clo1)
				ArrMhs[i].mk[j].clo1 = clo1
				fmt.Print("NIlai CLO 2: ")
				fmt.Scanln(&clo2)
				ArrMhs[i].mk[j].clo2 = clo2
				fmt.Print("NIlai CLO 3: ")
				fmt.Scanln(&clo3)
				ArrMhs[i].mk[j].clo3 = clo3

				ArrMhs[i].mk[j].na = HitungNA(clo1, clo2, clo3)
			}
		}

	} else {

		fmt.Println("DATA TIDAK DITEMUKAN.")
	}
}

// ReadMhs adalah prosedur untuk menampilkan data seorang mahasiswa
func ReadMhs(NamaMhs string, ArrMhs mhsw) {

	var i int

	i = 0
	for i < JumMhs && ArrMhs[i].nama != NamaMhs {

		i++
	}

	fmt.Println()

	if ArrMhs[i].nama == NamaMhs {

		fmt.Println("Nama: ", ArrMhs[i].nama)
		fmt.Println("NIM: ", ArrMhs[i].nim)
		fmt.Println("Daftar Mata Kuliah: ")
		ReadMK(NamaMhs, ArrMhs)

	} else {

		fmt.Println("DATA TIDAK DITEMUKAN.")
	}
}

// ReadMK adalah prosedur untuk melihat data Matakuliah dari seorang mahasiswa
func ReadMK(NamaMhs string, ArrMhs mhsw) {

	var i int

	i = 0
	for i < JumMhs && ArrMhs[i].nama != NamaMhs {

		i++
	}

	fmt.Println()

	if ArrMhs[i].nama == NamaMhs {

		for j := 0; j < len(ArrMhs[i].mk); j++ {

			if ArrMhs[i].mk[j].nama != "" {

				fmt.Printf("Mata Kuliah ke %v : %v \n", j+1, ArrMhs[i].mk[j].nama)
				fmt.Printf("SKS : %v", ArrMhs[i].mk[j].sks)
			}
		}

	} else {

		fmt.Println("DATA TIDAK DITEMUKAN.")
	}

}

// HitungNA adalah fungsi untuk menghitung Nilai Akhir per mata kuliah dengan CLO1 20%, CLO2 35%, CLO3 45%
func HitungNA(clo1, clo2, clo3 float64) float64 {

	var Total float64
	Total = (clo1 * 20 / 100) + (clo2 * 35 / 100) + (clo3 * 45 / 100)

	return Total

}

// UpMhs adalah prosedur untuk meng-Update nama atau nim dari seorang mahasiswa
func UpMhs(NamaMhs string, nim string, ArrMhs *mhsw) {

	var (
		NamaBaru, nimBaru string // variabel baru untuk menyimpan nama baru dan nim baru
	)

	i := 0
	for i < JumMhs && ArrMhs[i].nim != nim { // mencari NIM yang cocok

		i++
	}

	fmt.Println()

	if ArrMhs[i].nim == nim { // Jika nim sudah cocok

		fmt.Println("Data sebelumnya adalah: ") // menampilkan terlebih dahulu data sebelum akan di update
		fmt.Println("Nama: ", ArrMhs[i].nama)
		fmt.Println("NIM: ", ArrMhs[i].nim)

		fmt.Println("-------------------------------------") // Meng-Update nama dan nim mahasiswa
		fmt.Print("Update Nama Mahasiswa: ")
		fmt.Scanln(&NamaBaru)
		fmt.Print("Input NIM: ")
		fmt.Scanln(&nimBaru)

		// menyimpan nama dan nim pada array utama
		ArrMhs[i].nama = NamaBaru
		ArrMhs[i].nim = nimBaru

		fmt.Println("DATA BERHASIL DIUPDATE!")

	} else {

		fmt.Println("DATA TIDAK DITEMUKAN.")
	}
}

// UpMK adalah prosedur untuk meng-Update nilai suatu mata kuliah dari seorang mahasiswa
func UpMK(nim string, ArrMhs *mhsw) {

	var matkul string
	var clo1, clo2, clo3 float64

	i := 0
	for i < JumMhs && ArrMhs[i].nim != nim { // mencari NIM yang cocok

		i++
	}

	fmt.Println()

	if ArrMhs[i].nim == nim { // Jika nim sudah cocok

		fmt.Println("Data sebelumnya adalah: ") // menampilkan terlebih dahulu data sebelum akan di update
		fmt.Println("Nama: ", ArrMhs[i].nama)
		fmt.Println("NIM: ", ArrMhs[i].nim)
		fmt.Println("Daftar Mata Kuliah: ")

		for j := 0; j < len(ArrMhs[i].mk); j++ {

			if ArrMhs[i].mk[j].nama != "" {

				fmt.Printf("Mata Kuliah ke %v : %v \n", j+1, ArrMhs[i].mk[j].nama)
				fmt.Printf("SKS : %v", ArrMhs[i].mk[j].sks)
			}
		}

		fmt.Print("Mata Kuliah yang akan di-Update: ")
		fmt.Scan(&matkul)

		for j := 0; j < len(ArrMhs[i].mk); j++ {

			if ArrMhs[i].mk[j].nama == matkul {

				fmt.Print("NIlai CLO 1: ")
				fmt.Scanln(&clo1)
				ArrMhs[i].mk[j].clo1 = clo1
				fmt.Print("NIlai CLO 2: ")
				fmt.Scanln(&clo2)
				ArrMhs[i].mk[j].clo2 = clo2
				fmt.Print("NIlai CLO 3: ")
				fmt.Scanln(&clo3)
				ArrMhs[i].mk[j].clo3 = clo3

				ArrMhs[i].mk[j].na = HitungNA(clo1, clo2, clo3)

			}
		}

		fmt.Println("DATA BERHASIL DIUPDATE!")

	} else {

		fmt.Println("DATA TIDAK DITEMUKAN.")
	}
}
