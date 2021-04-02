/*
APLIKASI NIM MAHASISWA BY ADITYA & FARIS
UPDATE KE : 9
UPDATE TERAKHIR: 29/11/2019
Detail Update :

1. Perbaikan UpdateMK.
2. Efisiensi UpdateMK, ketika output data sebelumnya langsung panggil command ReadMK saja.
3. Sukses membuat proc DelMhs untuk menghapus data mahasiswa(timpa array).
4. Membuat Proc ReadAllMhs untuk output semua data mahasiswa.
5. Semua sudah input NIM, bukan nama lagi.
6. NIM sama, diminta untuk mengisi ulang.
7. DelMK memiliki menu.
8. Case Sensitive. (LATEST)
9. Input jumlah mahasiswa menggunakan marker
10. Data sebelumnya pada UpMK (copy paste kode dari prosedurnya)
11. ReadAllMhs menampilkan semua nama mahasiswa sekaligus mata kuliah yang diambil
12. HitungIP sudah benar dengan skala penilaian 4
13. Merubah nama fungsi HitungNA menjadi HitungNSM
14. "Nilai Akhir" diubah menjadi "Nilai Skor Mata Kuliah"
15. Nilai Skor Mata Kuliah diubah menjadi skala 100 (awalnya 4), namun IP tetap skala 4

Harus Dikerjakan :
1. Membuat Proc DelMK beserta nilai(sesuai pdf) (DONE)
2. Menampilkan daftar mahasiswa terurut berdasarkan nilai dan jumlah SKS yang diambil (DONE)
3. Menampilkan IP tertinggi. (DONE)
4. Menampilkan transkrip nilai(sepertinya sudah bisa) (DONE)

Hasil malam ini :
1. Perlu diubah: Minta nama diganti nim (DONE)
2. Perlu diperbaiki: Jika NIM sama, maka diminta untuk mengisi NIM ulang. (DONE) (DONE)
3. Perlu diperbaiki: Mata Kuliah tidak boleh sama. (DONE)
4. Perlu Diperhatikan: Case-Sensitive, seharusnya ya atau Ya atau yA tidak ada perlakuan berbeda.  (DONE)

BUG :
1. Mata Kuliah Ke-N bug. (TIDAK LAGI DITEMUKAN)
2. NIM sama masih diterima. (DONE)
3. Matkul sama masih diterima. (DONE)
4. DelMK belum memiliki pilihan pada Menu. (DONE)
5. tidak bisa UpMK (DONE)
6. UpMK tidak ada data sebelumnya (DONE)
7. CreateMK masih bisa keduplikat jika huruf kecil besarnya beda. (DONE)
8. HitungIP untuk cara penghitungan nilainya masih salah, silahkan riset dahulu. (DONE)

UPDATE 22/11/19 :
1. Input jumlah mahasiswa menggunakan marker
2. Data sebelumnya pada UpMK (copy paste kode dari prosedurnya)
3. ReadAllMhs menampilkan semua nama mahasiswa sekaligus mata kuliah yang diambil
4. Pembuatan Func HitungIP untuk menghitung IP Keseluruhan.
5. Perbaikan mekanisme ReadAllMhs, langsung panggil func ReadMhs saja.
6. Penambahan atribut jumsks pada type Mahasiswa untuk menghitung jumlah keseluruhan SKS.
*/

/*

BUG YANG BELUM BISA TERATASI:
1. Menu jika di input int(spasi)int(spasi)int akan memproses ketika int tersebut dan mengulang kalimat error sebanyak int
2. (SDA) namun dengan string dengan input stringstringstring

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv" // dibutuhkan untuk mengubah string ke int pada fungsi isNumeric dengan syntax strconv.ParseFloat(s, 64)
	"strings" // dibutuhkan untuk mengatasi case sensitive
)

// JumMK adalah jumlah maksimal mata kuliah
const JumMK int = 1000

// MhsMax adalah jumlah maksimal mahasiswa
const MhsMax int = 1500

// matakuliah adalah tipe data baru untuk nama Matakuliah dan Nilai nya
type matakuliah struct {
	nama, nmk            string
	sks                  int
	na, clo1, clo2, clo3 float64
}

// mahasiswa adalah tipe data baru untuk mendata mahasiswa, nim, dan matkul yang diambil
type mahasiswa struct {
	nama, nim string
	mk        [JumMK]matakuliah
	ip        float64
	jumsks    int
}

//mhsw adalah tipe data untuk membentuk var dengan tipe array(tipe alias)
type mhsw [MhsMax]mahasiswa

// JumMhs adalah variabel global
var JumMhs int

// ArrTemp adalah variabel global untuk mengarsip data mahasiswa yang di delete(hanya nama).
var ArrTemp [MhsMax]mahasiswa

// ArrMhs adalah Array utama
var ArrMhs mhsw

func main() {

	var (
		nim, matkul string
		pilmenu     int
		junk        string // untuk melakukan 'tekan enter'
		sort        string
		min         float64
	)

	Menu()               // Menampilkan Menu
	fmt.Scanln(&pilmenu) // Memilih Menu

	for pilmenu < 1 || pilmenu > 6 { // Validasi input menu dari user harus pada interval 1-6

		// jika inputan user lebih dari 6 atau kurang dari 1, menu akan terus diulang hingga inputan valid
		fmt.Println("\nMaaf, tidak ada pilihan dalam menu :)\n")
		Menu()
		fmt.Scanln(&pilmenu)
	}

	for pilmenu != 6 { // Mengakses pilihan menu

		nim = "" // debug supaya tidak otomatis terisi
		fmt.Println()

		if pilmenu == 1 { // Input Data

			pilmenu = 0
			InputMenu() // menampikan menu input

			fmt.Scanln(&pilmenu)
			for pilmenu < 1 || pilmenu > 3 { // Validasi input menu dari user harus pada interval 1-3

				// menu akan terus diulang hingga inputan valid
				fmt.Println("\nMaaf, tidak ada pilihan dalam menu :)\n")
				InputMenu()
				fmt.Scanln(&pilmenu)
			}
			if pilmenu == 1 { // Input Data Mahasiswa

				CreateMhs(&ArrMhs)

			} else if pilmenu == 2 { // Input Data Mata Kuliah

				fmt.Print("Masukan NIM: ")
				fmt.Scanln(&nim)

				for isNumeric(nim) == false { // melakukan validasi dengan memanggil fungsi isNumeric

					fmt.Println("NIM harus angka.")
					fmt.Print("Masukan NIM: ")
					fmt.Scanln(&nim)
				}

				CreateMK(nim, &ArrMhs) // memanggil prosedur CreateMK untuk menginput data kuliah
			}

		} else if pilmenu == 2 { // Output Data

			pilmenu = 0  // debug supaya tidak otomatis terinput
			OutputMenu() // menampilkan menu output
			fmt.Scanln(&pilmenu)

			for pilmenu < 1 || pilmenu > 6 { // Validasi input menu dari user harus pada interval 1-6

				// menu akan terus diulang hingga inputan valid
				fmt.Println("\nMaaf, tidak ada pilihan dalam menu :)\n")
				OutputMenu()
				fmt.Scanln(&pilmenu)
			}

			if pilmenu == 1 { // output seluruh data mahasiswa

				ReadAllMhs(ArrMhs) // memanggil prosedur ReadAllMhs
			} else if pilmenu == 2 { // output satu data mahasiswa

				fmt.Print("Masukan NIM: ")
				fmt.Scanln(&nim)

				for isNumeric(nim) == false { // melakukan validasi dengan memanggil fungsi isNumeric

					fmt.Println("NIM harus angka.")
					fmt.Print("Masukan NIM: ")
					fmt.Scanln(&nim)
				}

				ReadMhs(nim, ArrMhs)

			} else if pilmenu == 3 { // output nama mahasiswa dengan membatasi minimal IP

				fmt.Print("IP Minimum (Bilangan Bulat): ")
				fmt.Scanln(&min) // menginputkan ip, jika inputan adalah sebagai contoh 3,5 3,6 3,7 maka akan tetap dianggap 3

				RentangIP(&ArrMhs, min) // memanggil prosedur rentangIP (Binary Search)

			} else if pilmenu == 4 { // output data mata kuliah dari satu mahasiswa

				fmt.Print("Masukan NIM: ")
				fmt.Scanln(&nim)

				for isNumeric(nim) == false { // melakukan validasi dengan memanggil fungsi isNumeric

					fmt.Println("NIM harus angka.")
					fmt.Print("Masukan NIM: ")
					fmt.Scanln(&nim)
				}

				ReadMK(nim, ArrMhs) // memanggil prosedur ReadMK untuk menampilkan data mata kuliah yang diambil
			} else if pilmenu == 5 { // output mahasiswa yang mengambil suatu mata kuliah

				fmt.Print("Masukan Nama Mata Kuliah: ")
				scanner := bufio.NewScanner(os.Stdin) // variabel scanner dideklarasiin sebagai scannernya dari file import bufio
				scanner.Scan()                        // melakukan permintaan input (Scanln) dan menyimpannya di variabel scanner
				matkul = scanner.Text()               // string versi scanner

				ReadByMK(matkul, ArrMhs)
			}

		} else if pilmenu == 3 { // Update mahasiswa

			pilmenu = 0          // debug
			UpdateMenu()         // menampilkan menu update
			fmt.Scanln(&pilmenu) // memilih menu

			for pilmenu < 1 || pilmenu > 3 { // Validasi input menu dari user harus pada interval 1-3

				// menu akan terus diulang hingga inputan valid
				fmt.Println("\nMaaf, tidak ada pilihan dalam menu :)\n")
				UpdateMenu()
				fmt.Scanln(&pilmenu)

			}

			if pilmenu == 1 { // Update Data Mahasiswa

				fmt.Print("Masukan NIM: ")
				fmt.Scanln(&nim)

				for isNumeric(nim) == false { // validasi nim harus angka

					fmt.Println("NIM harus angka.")
					fmt.Print("Masukan NIM: ")
					fmt.Scanln(&nim)
				}

				UpMhs(nim, &ArrMhs) // memanggil prosedut UpMhs

			} else if pilmenu == 2 { // Update nilai Mata Kuliah

				fmt.Print("Masukan NIM: ")
				fmt.Scanln(&nim)

				for isNumeric(nim) == false { // validasi nim harus angka

					fmt.Println("NIM harus angka.")
					fmt.Print("Masukan NIM: ")
					fmt.Scanln(&nim)
				}

				UpMK(nim, &ArrMhs) // memanggil prosedur UpMK
			}
		} else if pilmenu == 4 { // Hapus Data

			pilmenu = 0 // debug
			HapusMenu() // menampilkan menu hapus data
			fmt.Scanln(&pilmenu)

			for pilmenu < 1 || pilmenu > 3 { // Validasi input menu dari user harus pada interval 1-3

				fmt.Println("\nMaaf, tidak ada pilihan dalam menu :)\n")
				HapusMenu()
				fmt.Scanln(&pilmenu)
			}

			if pilmenu == 1 { // Hapus Data Mahasiswa

				fmt.Print("Masukan NIM: ")
				fmt.Scanln(&nim)

				for isNumeric(nim) == false { // validasi nim harus angka

					fmt.Println("NIM harus angka.")
					fmt.Print("Masukan NIM: ")
					fmt.Scanln(&nim)
				}

				DelMhs(nim, &ArrMhs) // memanggil prosedur DelMhs

			} else if pilmenu == 2 { // Hapus Mata Kuliah

				fmt.Print("Masukan NIM: ")
				fmt.Scanln(&nim)

				for isNumeric(nim) == false { // validasi nim harus angka

					fmt.Println("NIM harus angka.")
					fmt.Print("Masukan NIM: ")
					fmt.Scanln(&nim)
				}

				DelMK(nim, &ArrMhs) // memanggil prosedur DelMK
			}
		} else if pilmenu == 5 { // Urutkan Data

			pilmenu = 0 // debug
			UrutMenu()  // menampilkan menu pilihan mengurutkan
			fmt.Scanln(&pilmenu)

			for pilmenu < 1 || pilmenu > 5 { // Validasi input menu dari user harus pada interval 1-5

				// menu akan terus diulang hingga inputan valid
				fmt.Println("\nMaaf, tidak ada pilihan dalam menu :)\n")
				UrutMenu()
				fmt.Scanln(&pilmenu)
			}
			if pilmenu == 1 { //

				SortIPBK(&ArrMhs) // insertion
				sort = "IP terbesar - terkecil"

			} else if pilmenu == 2 {

				SortIPKB(&ArrMhs) // insertion
				sort = "IP terkecil - terbesar"

			} else if pilmenu == 3 {

				SortSKSBK(&ArrMhs) // selection
				sort = "SKS terbesar - terkecil"

			} else if pilmenu == 4 {

				SortSKSKB(&ArrMhs) // selection
				sort = "SKS terkecil - terbesar"

			}

			if JumMhs > 0 {

				fmt.Println("Ranking Berdasarkan ", sort, ": ")
				fmt.Println("Proses . . .")
				fmt.Println("Berhasil.")
				fmt.Println()

				ReadAllMhs(ArrMhs)
			}
		}

		fmt.Println("")
		fmt.Print("Tekan Enter untuk kembali ke Main Menu ...")
		fmt.Scanln(&junk)

		Menu() // Menampilkan Menu kembali setelah mengakses 1-6

		pilmenu = 0 // debug

		fmt.Scanln(&pilmenu)             // Memilih menu kembali setelah mengakses 1-6
		for pilmenu < 1 || pilmenu > 6 { // Validasi input menu dari user harus pada interval 1-6

			// menu akan terus diulang hingga inputan valid
			fmt.Println("\nMaaf, tidak ada pilihan dalam menu :)\n")
			Menu()
			fmt.Scanln(&pilmenu)
		}
	}
	// Jika pilihan menu adalah 6, Keluar dari Aplikasi

	// Jika aplikasi ditutup akan menampilkan ini
	fmt.Println()
	fmt.Println("Terimakasih telah menggunakan aplikasi ini.")
	fmt.Println("Tertanda, Adit dan Faris.")
	fmt.Scanln()

	// NOTE
	/*	INI UNTUK NGECEK DATA TERSIMPAN DENGAN BAIK ATAU TIDAK!!

		for i:=0;i<JumMhs;i++ {

			fmt.Println(ArrMhs[i].nama)
		}
	*/
}

// Menu adalah prosedur untuk menampilkan halaman awal program
func Menu() {

	fmt.Println("\n.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.\n")
	fmt.Println("Menu: ")
	fmt.Println("1. Input Data.")
	fmt.Println("2. Output Data.")
	fmt.Println("3. Update Data.")
	fmt.Println("4. Hapus Data.")
	fmt.Println("5. Urutkan Data.")
	fmt.Println("6. Keluar dari Aplikasi.")
	fmt.Print("Menu Anda: ")

}

//InputMenu adalah prosedur untuk menampilkan pilihan input
func InputMenu() {

	fmt.Println("Menu Input: ")
	fmt.Println("1. Input Data Mahasiswa.")
	fmt.Println("2. Input Data Mata Kuliah.")
	fmt.Println("3. Kembali.")
	fmt.Print("Menu Anda: ")

}

//OutputMenu adalah prosedur untuk menampilkan pilihan output
func OutputMenu() {

	fmt.Println("Menu Output ")
	fmt.Println("1. Output Seluruh Data Mahasiswa.")
	fmt.Println("2. Output Data Mahasiswa.")
	fmt.Println("3. Output Nama Mahasiswa dalam rentang IP tertentu.")
	fmt.Println("4. Output Data Mata Kuliah.")
	fmt.Println("5. Output Mahasiswa yang mengambil suatu mata kuliah.")
	fmt.Println("6. Kembali.")
	fmt.Print("Menu Anda: ")

}

//UpdateMenu adalah prosedur untuk menampilkan pilihan Update
func UpdateMenu() {

	fmt.Println("Menu Update: ")
	fmt.Println("1. Update Data Mahasiswa.")
	fmt.Println("2. Update Nilai Mata Kuliah.")
	fmt.Println("3. Kembali.")
	fmt.Print("Menu Anda: ")

}

//HapusMenu adalah prosedur untuk menampilkan pilihan Hapus
func HapusMenu() {

	fmt.Println("Menu Hapus: ")
	fmt.Println("1. Hapus Data Mahasiswa.")
	fmt.Println("2. Hapus Data Mata Kuliah.")
	fmt.Println("3. Kembali.")
	fmt.Print("Menu Anda: ")
}

//UrutMenu adalah prosedur untuk menampilkan pilihan Urut
func UrutMenu() {

	fmt.Println("Menu Urutkan: ")
	fmt.Println("1. Urutkan Berdasarkan IP(terbesar - terkecil).")
	fmt.Println("2. Urutkan Berdasarkan IP(terkecil - terbesar).")
	fmt.Println("3. Urutkan Berdasarkan Jumlah SKS(terbesar - terkecil).")
	fmt.Println("4. Urutkan Berdasarkan Jumlah SKS(terkecil - terbesar).")
	fmt.Println("5. Kembali.")
	fmt.Print("Menu Anda: ")
}

// CreateMhs adalah prosedur untuk menambah data mahasiswa
func CreateMhs(ArrMhs *mhsw) {

	var nim string
	var qna string
	var input string

	i := 0
	fmt.Println("Ketik STOP untuk berhenti")
	fmt.Print("Input nama Mahasiswa ke-", JumMhs+1, ": ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()

	for input == "" {

		fmt.Println("Nama tidak boleh kosong.")
		fmt.Print("Input nama Mahasiswa ke-", JumMhs+1, ": ")
		scanner.Scan()
		input = scanner.Text()
	}

	for (JumMhs < MhsMax) && (strings.ToLower(input) != "stop") {

		ArrMhs[JumMhs].nama = input

		fmt.Print("Input NIM: ")
		fmt.Scanln(&nim)

		for isNumeric(nim) == false {
			fmt.Println("NIM harus angka dan positif.")
			fmt.Print("Input NIM: ")
			fmt.Scanln(&nim)
		}

		j := 0

		for j <= JumMhs && ArrMhs[j].nim != nim {

			j++
		}

		for ArrMhs[j].nim == nim { // Jika NIM sama, maka diminta untuk mengisi NIM ulang.

			fmt.Println("NIM sama telah ditemukan. ")
			fmt.Print("Input NIM: ")
			fmt.Scanln(&nim)

			for isNumeric(nim) == false {

				fmt.Println("NIM harus angka dan positif.")
				fmt.Print("Input NIM: ")
				fmt.Scanln(&nim)
			}

			j = 0
			for j <= JumMhs && ArrMhs[j].nim != nim {

				j++
			}
		}

		ArrMhs[JumMhs].nim = nim

		fmt.Print("Input Mata Kuliah?(Ya/Tidak): ")
		fmt.Scanln(&qna)

		for strings.ToLower(qna) != "ya" && strings.ToLower(qna) != "tidak" {

			fmt.Println("Input tidak valid.")
			fmt.Print("Input Mata Kuliah?(Ya/Tidak): ")
			fmt.Scanln(&qna)
		}

		if strings.ToLower(qna) == "ya" {

			CreateMK(ArrMhs[JumMhs].nim, ArrMhs)

		} else if strings.ToLower(qna) == "tidak" {

		}

		JumMhs++
		i++

		fmt.Print("Input nama Mahasiswa ke-", JumMhs+1, ": ")
		scanner.Scan()
		input = scanner.Text()

		for input == "" {

			fmt.Println("Nama tidak boleh kosong.")
			fmt.Print("Input nama Mahasiswa ke-", JumMhs+1, ": ")
			scanner.Scan()
			input = scanner.Text()
		}
	}
}

// CreateMK adalah prosedur untuk menambah data mata kuliah pada mahasiswa
func CreateMK(nim string, ArrMhs *mhsw) {

	var (
		i, j, k       int
		nilai         float64
		matkul, input string
	)

	j = 0
	i = 0
	scanner := bufio.NewScanner(os.Stdin)

	for i < JumMhs && ArrMhs[i].nim != nim {

		i++
	}

	if ArrMhs[i].nim == nim {

		j = 0
		for j < JumMK-1 && ArrMhs[i].mk[j].nama != "" {

			j++

		}

		if ArrMhs[i].mk[j].nama == "" {

			fmt.Println("\nInput STOP untuk berhenti.\n")
			fmt.Printf("Mata Kuliah ke %v: ", j+1)
			scanner.Scan()
			matkul = scanner.Text()

			for matkul == "" {

				fmt.Println("Nama mata kuliah tidak boleh kosong.")
				fmt.Printf("Mata Kuliah ke %v: ", j+1)
				scanner.Scan()
				matkul = scanner.Text()
			}

			k = 0
			for k < JumMK-1 && strings.ToLower(ArrMhs[i].mk[k].nama) != strings.ToLower(matkul) {

				k++

			}

			for strings.ToLower(ArrMhs[i].mk[k].nama) == strings.ToLower(matkul) {

				fmt.Println("Mata Kuliah sama telah ditemukan.")
				fmt.Printf("Mata Kuliah ke %v: ", j+1)
				scanner.Scan()
				matkul = scanner.Text()

				for matkul == "" {

					fmt.Println("Nama mata kuliah tidak boleh kosong.")
					fmt.Printf("Mata Kuliah ke %v: ", j+1)
					scanner.Scan()
					matkul = scanner.Text()
				}

				k = 0
				for k < JumMK-1 && strings.ToLower(ArrMhs[i].mk[k].nama) != strings.ToLower(matkul) {

					k++
				}
			}

			for strings.ToLower(matkul) != "stop" && j < JumMK {

				ArrMhs[i].mk[j].nama = matkul
				fmt.Print("Jumlah SKS: ")
				input = "0"
				fmt.Scanln(&input)
				sks, err := strconv.ParseFloat(input, 64)

				for err != nil {

					fmt.Println("SKS harus angka.")
					input = "0"
					fmt.Print("Jumlah SKS: ")
					fmt.Scanln(&input)
					sks, err = strconv.ParseFloat(input, 64)
				}

				for sks <= 0 {

					fmt.Println("SKS tidak boleh negatif.")
					fmt.Print("Jumlah SKS: ")
					input = "0"
					fmt.Scanln(&input)
					sks, err = strconv.ParseFloat(input, 64)

					for err != nil {

						fmt.Println("SKS harus angka.")
						input = "0"
						fmt.Print("Jumlah SKS: ")
						fmt.Scanln(&input)
						sks, err = strconv.ParseFloat(input, 64)
					}

				}

				ArrMhs[i].mk[j].sks = int(sks)

				ArrMhs[i].jumsks = ArrMhs[i].jumsks + ArrMhs[i].mk[j].sks

				fmt.Print("NIlai CLO 1: ")
				input = "0"
				fmt.Scanln(&input)
				clo1, err := strconv.ParseFloat(input, 64)

				for err != nil {

					fmt.Println("CLO 1 harus angka.")
					fmt.Print("NIlai CLO 1: ")

					input = "0"
					fmt.Scanln(&input)
					clo1, err = strconv.ParseFloat(input, 64)
				}

				for clo1 < 0 || clo1 > 100 {

					fmt.Println("Nilai di luar rentang 0-100")
					fmt.Print("NIlai CLO 1: ")

					input = "0"
					fmt.Scanln(&input)
					clo1, err = strconv.ParseFloat(input, 64)

					for err != nil {
						fmt.Println("CLO 1 harus angka.")
						fmt.Print("NIlai CLO 1: ")

						input = "0"
						fmt.Scanln(&input)
						clo1, err = strconv.ParseFloat(input, 64)
					}

				}

				ArrMhs[i].mk[j].clo1 = clo1

				fmt.Print("NIlai CLO 2: ")

				input = "0"
				fmt.Scanln(&input)
				clo2, err := strconv.ParseFloat(input, 64)

				for err != nil {

					fmt.Println("CLO 2 harus angka.")
					fmt.Print("NIlai CLO 2: ")

					input = "0"
					fmt.Scanln(&input)
					clo2, err = strconv.ParseFloat(input, 64)
				}

				for clo2 < 0 || clo2 > 100 {

					fmt.Println("Nilai di luar rentang 0-100")
					fmt.Print("NIlai CLO 2: ")

					input = "0"
					fmt.Scanln(&input)
					clo2, err = strconv.ParseFloat(input, 64)

					for err != nil {

						fmt.Println("CLO 2 harus angka.")
						fmt.Print("NIlai CLO 2: ")

						input = "0"
						fmt.Scanln(&input)
						clo2, err = strconv.ParseFloat(input, 64)
					}
				}

				ArrMhs[i].mk[j].clo2 = clo2

				fmt.Print("NIlai CLO 3: ")

				input = "0"
				fmt.Scanln(&input)
				clo3, err := strconv.ParseFloat(input, 64)

				for err != nil {

					fmt.Println("CLO 3 harus angka.")
					fmt.Print("NIlai CLO 3: ")

					input = "0"
					fmt.Scanln(&input)
					clo3, err = strconv.ParseFloat(input, 64)
				}

				for clo3 < 0 || clo3 > 100 {

					fmt.Println("Nilai di luar rentang 0-100")
					fmt.Print("NIlai CLO 3: ")

					input = "0"
					fmt.Scanln(&input)
					clo3, err = strconv.ParseFloat(input, 64)

					for err != nil {

						fmt.Println("CLO 3 harus angka.")
						fmt.Print("NIlai CLO 3: ")

						input = "0"
						fmt.Scanln(&input)
						clo3, err = strconv.ParseFloat(input, 64)
					}
				}

				ArrMhs[i].mk[j].clo3 = clo3

				ArrMhs[i].mk[j].na = HitungNSM(clo1, clo2, clo3)
				nilai = ArrMhs[i].mk[j].na

				ArrMhs[i].mk[j].nmk = HitungNMK(i, j, nilai)

				j++

				if j < JumMK {

					fmt.Printf("Mata Kuliah ke %v: ", j+1)
					scanner.Scan()
					matkul = scanner.Text()

					for matkul == "" {
						fmt.Println("Nama mata kuliah tidak boleh kosong.")
						fmt.Printf("Mata Kuliah ke %v: ", j+1)
						scanner.Scan()
						matkul = scanner.Text()
					}

					k = 0
					for k < JumMK-1 && strings.ToLower(ArrMhs[i].mk[k].nama) != strings.ToLower(matkul) {
						k++

					}

					for strings.ToLower(ArrMhs[i].mk[k].nama) == strings.ToLower(matkul) {

						fmt.Println("Mata Kuliah sama telah ditemukan.")
						fmt.Printf("Mata Kuliah ke %v: ", j+1)
						scanner.Scan()
						matkul = scanner.Text()

						for matkul == "" {

							fmt.Println("Nama mata kuliah tidak boleh kosong.")
							fmt.Printf("Mata Kuliah ke %v: ", j+1)
							scanner.Scan()
							matkul = scanner.Text()
						}

						k = 0
						for k < JumMK-1 && strings.ToLower(ArrMhs[i].mk[k].nama) != strings.ToLower(matkul) {

							k++
						}
					}
				}

			}

		} else {

			fmt.Println("Mata Kuliah sudah terisi penuh. ")
		}

		if ArrMhs[i].jumsks > 0 {

			ArrMhs[i].ip = HitungIP(ArrMhs, i)
		}

	} else {

		fmt.Println("DATA TIDAK DITEMUKAN.")
	}
}

// ReadMhs adalah prosedur untuk menampilkan data seorang mahasiswa
func ReadMhs(nim string, ArrMhs mhsw) {

	var i int

	i = 0
	for i < JumMhs && ArrMhs[i].nim != nim {

		i++
	}

	if ArrMhs[i].nim == nim {

		fmt.Println("Nama: ", ArrMhs[i].nama)
		fmt.Println("NIM: ", ArrMhs[i].nim)
		fmt.Println("Jumlah SKS: ", ArrMhs[i].jumsks)
		fmt.Printf("IP: %.2f\n", ArrMhs[i].ip)
		fmt.Println("Daftar Mata Kuliah: ")
		ReadMK(nim, ArrMhs)

	} else {

		fmt.Println("DATA TIDAK DITEMUKAN.")
	}
}

// ReadMK adalah prosedur untuk melihat data Matakuliah dari seorang mahasiswa
func ReadMK(nim string, ArrMhs mhsw) {

	var i, num int

	i = 0
	num = 0

	for i < JumMhs && ArrMhs[i].nim != nim {

		i++
	}

	if ArrMhs[i].nim == nim {

		if ArrMhs[i].mk[0].nama != "" {

			if ArrMhs[i].mk[0].nama != "ARCHIVED" {

				fmt.Printf("Mata Kuliah ke %v : %v \n", num+1, ArrMhs[i].mk[0].nama)
				fmt.Printf("SKS : %v\n", ArrMhs[i].mk[0].sks)
				fmt.Printf("CLO1 : %v \n", ArrMhs[i].mk[0].clo1)
				fmt.Printf("CLO2 : %v \n", ArrMhs[i].mk[0].clo2)
				fmt.Printf("CLO3 : %v \n", ArrMhs[i].mk[0].clo3)
				fmt.Printf("Nilai Skor Mata Kuliah : %v \n", ArrMhs[i].mk[0].na)
				fmt.Printf("Nilai Mata Kuliah : %v \n", ArrMhs[i].mk[0].nmk)
				fmt.Println()
			}

		} else {

			fmt.Println("Tidak ada mata kuliah yang dapat ditampilkan.")
		}

		num = 1

		for j := 1; j < len(ArrMhs[i].mk); j++ {

			if ArrMhs[i].mk[j].nama != "" {

				fmt.Printf("Mata Kuliah ke %v : %v \n", num+1, ArrMhs[i].mk[j].nama)
				fmt.Printf("SKS : %v\n", ArrMhs[i].mk[j].sks)
				fmt.Printf("CLO1 : %v \n", ArrMhs[i].mk[j].clo1)
				fmt.Printf("CLO2 : %v \n", ArrMhs[i].mk[j].clo2)
				fmt.Printf("CLO3 : %v \n", ArrMhs[i].mk[j].clo3)
				fmt.Printf("Nilai Skor Mata Kuliah : %v \n", ArrMhs[i].mk[j].na)
				fmt.Printf("Nilai Mata Kuliah : %v \n", ArrMhs[i].mk[j].nmk)
				fmt.Println()
				num++
			}

		}

	} else {

		fmt.Println("DATA TIDAK DITEMUKAN.")
	}

}

//ReadByMK adalah prosedur untuk mencari data mahasiswa mana saja yang mengambil suatu mata kuliah
func ReadByMK(NamaMK string, ArrMhs mhsw) {

	c := 0

	for i := 0; i < JumMhs; i++ {

		for j := 0; j < JumMK; j++ {

			if strings.ToLower(ArrMhs[i].mk[j].nama) == strings.ToLower(NamaMK) {

				c++
				fmt.Println("Nama: ", ArrMhs[i].nama)
				fmt.Println("NIM: ", ArrMhs[i].nim)
			}
		}
	}

	if c == 0 {

		fmt.Println("Tidak ada data yang cocok.")

	} else {

		fmt.Println("Banyak Mahasiswa: ", c)
	}
}

// HitungNSM adalah fungsi untuk menghitung Nilai Skor per mata kuliah dengan CLO1 20%, CLO2 35%, CLO3 45%
func HitungNSM(clo1, clo2, clo3 float64) float64 {

	var Total float64
	Total = (clo1 * 20 / 100) + (clo2 * 35 / 100) + (clo3 * 45 / 100)

	return Total

}

// UpMhs adalah prosedur untuk meng-Update nama atau nim dari seorang mahasiswa
func UpMhs(nim string, ArrMhs *mhsw) {

	var (
		NamaBaru, NimBaru string // variabel baru untuk menyimpan nama baru dan nim baru
	)

	scanner := bufio.NewScanner(os.Stdin)
	i := 0
	for i < JumMhs && ArrMhs[i].nim != nim { // mencari NIM yang cocok

		i++
	}

	if ArrMhs[i].nim == nim { // Jika nim sudah cocok

		fmt.Println("Data sebelumnya adalah: ") // menampilkan terlebih dahulu data sebelum akan di update
		fmt.Println("Nama: ", ArrMhs[i].nama)
		fmt.Println("NIM: ", ArrMhs[i].nim)

		fmt.Println("-------------------------------------") // Meng-Update nama dan nim mahasiswa
		fmt.Print("Update Nama Mahasiswa: ")
		scanner.Scan()
		NamaBaru = scanner.Text()

		for NamaBaru == "" {

			fmt.Println("Nama tidak boleh kosong.")
			fmt.Print("Update Nama Mahasiswa: ")
			scanner.Scan()
			NamaBaru = scanner.Text()
		}

		fmt.Print("Input NIM: ")
		fmt.Scanln(&NimBaru)

		for isNumeric(NimBaru) == false {
			fmt.Println("NIM harus angka dan positif.")
			fmt.Print("Input NIM: ")
			fmt.Scanln(&NimBaru)
		}

		j := 0

		for j <= JumMhs && ArrMhs[j].nim != NimBaru {

			j++
		}

		for ArrMhs[j].nim == NimBaru { // Jika NIM sama, maka diminta untuk mengisi NIM ulang.

			fmt.Println("NIM sama telah ditemukan. ")
			fmt.Print("Input NIM: ")
			fmt.Scanln(&NimBaru)

			for isNumeric(NimBaru) == false {

				fmt.Println("NIM harus angka dan positif.")
				fmt.Print("Input NIM: ")
				fmt.Scanln(&NimBaru)
			}

			j = 0
			for j <= JumMhs && ArrMhs[j].nim != NimBaru {

				j++
			}
		}

		// menyimpan nama dan nim pada array utama
		ArrMhs[i].nama = NamaBaru
		ArrMhs[i].nim = NimBaru

		fmt.Println("DATA BERHASIL DIUPDATE!")

	} else {

		fmt.Println("DATA TIDAK DITEMUKAN.")
	}
}

// UpMK adalah prosedur untuk meng-Update nilai suatu mata kuliah dari seorang mahasiswa
func UpMK(nim string, ArrMhs *mhsw) {

	var matkul string
	var clo1, clo2, clo3, nilai float64
	var j, i int

	i = 0
	scanner := bufio.NewScanner(os.Stdin)

	for i < JumMhs && ArrMhs[i].nim != nim { // mencari NIM yang cocok

		i++
	}

	if ArrMhs[i].nim == nim { // Jika nim sudah cocok

		fmt.Println("Data sebelumnya adalah: ") // menampilkan terlebih dahulu data sebelum akan di update

		fmt.Println("Nama: ", ArrMhs[i].nama)
		fmt.Println("NIM: ", ArrMhs[i].nim)
		fmt.Println("Daftar Mata Kuliah: ")

		num := 0

		if ArrMhs[i].mk[0].nama != "" {

			if ArrMhs[i].mk[0].nama != "ARCHIVED" {

				fmt.Printf("Mata Kuliah ke %v : %v \n", num+1, ArrMhs[i].mk[0].nama)
				fmt.Printf("SKS : %v\n", ArrMhs[i].mk[0].sks)
				fmt.Printf("CLO1 : %v \n", ArrMhs[i].mk[0].clo1)
				fmt.Printf("CLO2 : %v \n", ArrMhs[i].mk[0].clo2)
				fmt.Printf("CLO3 : %v \n", ArrMhs[i].mk[0].clo3)
				fmt.Printf("Nilai Skor Mata Kuliah : %v \n", ArrMhs[i].mk[0].na)
				fmt.Printf("Nilai Mata Kuliah : %v \n", ArrMhs[i].mk[0].nmk)
				fmt.Println()

			}

			num = 1

			for j := 1; j < len(ArrMhs[i].mk); j++ {

				if ArrMhs[i].mk[j].nama != "" {

					if ArrMhs[i].mk[j].nama != "ARCHIVED" {

						fmt.Printf("Mata Kuliah ke %v : %v \n", num+1, ArrMhs[i].mk[j].nama)
						fmt.Printf("SKS : %v\n", ArrMhs[i].mk[j].sks)
						fmt.Printf("CLO1 : %v \n", ArrMhs[i].mk[j].clo1)
						fmt.Printf("CLO2 : %v \n", ArrMhs[i].mk[j].clo2)
						fmt.Printf("CLO3 : %v \n", ArrMhs[i].mk[j].clo3)
						fmt.Printf("Nilai Skor Mata Kuliah : %v \n", ArrMhs[i].mk[j].na)
						fmt.Printf("Nilai Mata Kuliah : %v \n", ArrMhs[i].mk[j].nmk)
						fmt.Println()
						num++
					}

				}

			}

			fmt.Println()

			fmt.Print("Mata Kuliah yang akan di-Update: ")
			scanner.Scan()
			matkul = scanner.Text()

			j = 0

			for j < JumMK-2 && strings.ToLower(ArrMhs[i].mk[j].nama) != strings.ToLower(matkul) {

				j++
			}

			if strings.ToLower(ArrMhs[i].mk[j].nama) == strings.ToLower(matkul) {

				fmt.Print("NIlai CLO 1: ")
				fmt.Scanln(&clo1)

				for clo1 < 0 || clo1 > 100 {

					fmt.Println("Nilai di luar rentang 0-100")
					fmt.Print("NIlai CLO 1: ")
					fmt.Scanln(&clo1)
				}

				ArrMhs[i].mk[j].clo1 = clo1

				fmt.Print("NIlai CLO 2: ")
				fmt.Scanln(&clo2)

				for clo2 < 0 || clo2 > 100 {

					fmt.Println("Nilai di luar rentang 0-100")
					fmt.Print("NIlai CLO 2: ")
					fmt.Scanln(&clo2)
				}

				ArrMhs[i].mk[j].clo2 = clo2

				fmt.Print("NIlai CLO 3: ")
				fmt.Scanln(&clo3)

				for clo3 < 0 || clo3 > 100 {

					fmt.Println("Nilai di luar rentang 0-100")
					fmt.Print("NIlai CLO 3: ")
					fmt.Scanln(&clo3)
				}

				ArrMhs[i].mk[j].clo3 = clo3

				ArrMhs[i].mk[j].na = HitungNSM(clo1, clo2, clo3)
				nilai = ArrMhs[i].mk[j].na

				ArrMhs[i].mk[j].nmk = HitungNMK(i, j, nilai)

				fmt.Println("DATA BERHASIL DIUPDATE!")

			} else {

				fmt.Println("DATA TIDAK DITEMUKAN.")
			}

		} else {

			fmt.Println("Tidak ada")
			fmt.Println("")
			fmt.Println("Tidak bisa memperbarui nilai karena tidak ada mata kuliah yang diambil")
		}

	} else {

		fmt.Println("DATA TIDAK DITEMUKAN.")
	}
}

// ReadAllMhs adalah prosedur untuk mengoutputkan semua data mahasiswa
func ReadAllMhs(ArrMhs mhsw) {

	var nim string
	if JumMhs > 0 {

		for i := 0; i < JumMhs; i++ {

			if ArrMhs[i].nim != "" {

				fmt.Println("Nama: ", ArrMhs[i].nama)
				fmt.Println("NIM: ", ArrMhs[i].nim)
				fmt.Println("Jumlah SKS: ", ArrMhs[i].jumsks)
				fmt.Printf("IP: %.2f\n", ArrMhs[i].ip)
				fmt.Println("Daftar Mata Kuliah: ")
				nim = ArrMhs[i].nim
				ReadMK(nim, ArrMhs)

			}

		}

	} else {

		fmt.Println("Data masih kosong.")
	}

}

// DelMhs adalah prosedur untuk menghapus data mahasiswa (hide)
func DelMhs(nim string, ArrMhs *mhsw) {
	var n int

	n = JumMhs
	i := 0
	j := 0

	for i < JumMhs && ArrMhs[i].nim != nim { // mencari NIM yang cocok

		i++
	}

	if ArrMhs[i].nim == nim { // jika nim sudah cocok

		ArrTemp[i] = ArrMhs[i] // Data dipindahkan ke array ArrTemp

		if i < JumMhs {

			n = n - 1
			for j = i; j < n; j++ {

				ArrMhs[j] = ArrMhs[j+1]
			}

			ArrMhs[j].nama = ""
			ArrMhs[j].nim = ""
			ArrMhs[j].jumsks = 0
			ArrMhs[j].ip = 0

			for k := 0; k < JumMK; k++ {

				ArrMhs[j].mk[k].nama = ""
				ArrMhs[j].mk[k].nmk = ""
				ArrMhs[j].mk[k].sks = 0
				ArrMhs[j].mk[k].na = 0
				ArrMhs[j].mk[k].clo1 = 0
				ArrMhs[j].mk[k].clo2 = 0
				ArrMhs[j].mk[k].clo3 = 0
			}

		}

		fmt.Println("Data Mahasiswa ", ArrTemp[i].nim, "terhapus.")
		JumMhs--

	} else {

		fmt.Println("DATA TIDAK DITEMUKAN.")
	}
}

// DelMK adalah prosedur untuk menghapus data mata kuliah dari mahasiswa
func DelMK(nim string, ArrMhs *mhsw) {

	var matkul string
	var j, n int

	n = JumMK
	i := 0
	scanner := bufio.NewScanner(os.Stdin)

	for i < JumMhs && ArrMhs[i].nim != nim { // mencari NIM yang cocok

		i++
	}

	if ArrMhs[i].nim == nim { // jika nim sudah cocok

		fmt.Println("Mata Kuliah yang akan dihapus: ")
		scanner.Scan()
		matkul = scanner.Text()

		for j < JumMK-1 && strings.ToLower(matkul) != strings.ToLower(ArrMhs[i].mk[j].nama) {

			j++
		}

		if strings.ToLower(matkul) == strings.ToLower(ArrMhs[i].mk[j].nama) {

			fmt.Println("Mata Kuliah ", ArrMhs[i].mk[j].nama, " terhapus.")

			ArrTemp[i].mk[j] = ArrMhs[i].mk[j] // Data dipindahkan ke array ArrTemp
			ArrMhs[i].jumsks = ArrMhs[i].jumsks - ArrTemp[i].mk[j].sks

			if j < JumMK {

				n = n - 1
				for j < n {

					ArrMhs[i].mk[j] = ArrMhs[i].mk[j+1]
					j++
				}

				ArrMhs[i].mk[j].nama = ""
				ArrMhs[i].mk[j].nmk = ""
				ArrMhs[i].mk[j].sks = 0
				ArrMhs[i].mk[j].na = 0
				ArrMhs[i].mk[j].clo1 = 0
				ArrMhs[i].mk[j].clo2 = 0
				ArrMhs[i].mk[j].clo3 = 0

			}

			if ArrMhs[i].jumsks > 0 {

				ArrMhs[i].ip = HitungIP(ArrMhs, i)
			}

			if ArrMhs[i].jumsks == 0 {

				ArrMhs[i].ip = 0
			}

		} else {

			fmt.Println("Mata Kuliah tidak ditemukan.")
		}

	} else {

		fmt.Println("DATA TIDAK DITEMUKAN.")
	}
}

// HitungIP adalah fungsi yang mereturn perhitungan ip mahasiswa
func HitungIP(ArrMhs *mhsw, i int) float64 {

	var Total float64
	var Hasil float64

	Total = 0

	for j := 0; ArrMhs[i].mk[j].nama != ""; j++ {

		if ArrMhs[i].mk[j].nama != "ARCHIVED" {

			Hasil = float64(ArrMhs[i].mk[j].sks) * ArrMhs[i].mk[j].na
		}

		Total = Hasil + Total
	}

	Total = Total / float64(ArrMhs[i].jumsks)
	Total = Total / 100 * 4

	return Total
}

// HitungNMK adalah fungsi untuk mereturn Nilai Mata Kuliah
func HitungNMK(i int, j int, nilai float64) string {

	switch {
	case nilai > 80:
		return "A"
	case nilai > 70:
		return "AB"
	case nilai > 65:
		return "B"
	case nilai > 60:
		return "BC"
	case nilai > 50:
		return "C"
	case nilai > 40:
		return "D"
	default:
		return "E"
	}

}

// SortIPBK adalah prosedur untuk mengurutkan data dari IP dari yang terbesar dengan INSERTION
func SortIPBK(ArrMhs *mhsw) {

	var temp mahasiswa

	i := 1
	if JumMhs > 0 {

		for i < JumMhs {

			j := i - 1
			temp.nama = ArrMhs[i].nama
			temp.ip = ArrMhs[i].ip
			temp.nim = ArrMhs[i].nim
			temp.mk = ArrMhs[i].mk
			temp.jumsks = ArrMhs[i].jumsks

			for j >= 0 && ArrMhs[j].ip < temp.ip {

				ArrMhs[j+1].nama = ArrMhs[j].nama
				ArrMhs[j+1].nim = ArrMhs[j].nim
				ArrMhs[j+1].mk = ArrMhs[j].mk
				ArrMhs[j+1].jumsks = ArrMhs[j].jumsks
				ArrMhs[j+1].ip = ArrMhs[j].ip
				j--
			}

			ArrMhs[j+1].nama = temp.nama
			ArrMhs[j+1].nim = temp.nim
			ArrMhs[j+1].mk = temp.mk
			ArrMhs[j+1].jumsks = temp.jumsks
			ArrMhs[j+1].ip = temp.ip
			i++
		}

	} else {

		fmt.Println("Data masih kosong.")
	}
}

// SortIPKB adalah prosedur untuk mengurutkan data dari IP dari yang terkecil dengan INSERTION
func SortIPKB(ArrMhs *mhsw) {

	var temp mahasiswa

	i := 1
	if JumMhs > 0 {

		for i < JumMhs {

			j := i - 1
			temp.nama = ArrMhs[i].nama
			temp.ip = ArrMhs[i].ip
			temp.nim = ArrMhs[i].nim
			temp.mk = ArrMhs[i].mk
			temp.jumsks = ArrMhs[i].jumsks

			for j >= 0 && ArrMhs[j].ip > temp.ip {

				ArrMhs[j+1].nama = ArrMhs[j].nama
				ArrMhs[j+1].nim = ArrMhs[j].nim
				ArrMhs[j+1].mk = ArrMhs[j].mk
				ArrMhs[j+1].jumsks = ArrMhs[j].jumsks
				ArrMhs[j+1].ip = ArrMhs[j].ip
				j--
			}

			ArrMhs[j+1].nama = temp.nama
			ArrMhs[j+1].nim = temp.nim
			ArrMhs[j+1].mk = temp.mk
			ArrMhs[j+1].jumsks = temp.jumsks
			ArrMhs[j+1].ip = temp.ip
			i++
		}

	} else {

		fmt.Println("Data masih kosong.")
	}
}

// SortSKSBK adalah prosedur untuk mengurutkan data SKS dari yang terbesar dengan SELECTION
func SortSKSBK(ArrMhs *mhsw) {

	var temp mahasiswa
	var max int

	i := 0
	if JumMhs > 0 {

		for i < JumMhs {

			j := i + 1
			max = i

			for j < JumMhs {

				if ArrMhs[j].jumsks > ArrMhs[max].jumsks {

					max = j
				}

				j++
			}

			temp.nama = ArrMhs[max].nama
			temp.ip = ArrMhs[max].ip
			temp.nim = ArrMhs[max].nim
			temp.mk = ArrMhs[max].mk
			temp.jumsks = ArrMhs[max].jumsks

			ArrMhs[max].nama = ArrMhs[i].nama
			ArrMhs[max].ip = ArrMhs[i].ip
			ArrMhs[max].nim = ArrMhs[i].nim
			ArrMhs[max].mk = ArrMhs[i].mk
			ArrMhs[max].jumsks = ArrMhs[i].jumsks

			ArrMhs[i].nama = temp.nama
			ArrMhs[i].ip = temp.ip
			ArrMhs[i].nim = temp.nim
			ArrMhs[i].mk = temp.mk
			ArrMhs[i].jumsks = temp.jumsks

			i++
		}

	} else {

		fmt.Println("Data masih kosong.")
	}
}

// SortSKSKB adalah prosedur untuk mengurutkan data SKS dari yang terkecil dengan SELECTION
func SortSKSKB(ArrMhs *mhsw) {

	var temp mahasiswa
	var min int

	i := 0
	if JumMhs > 0 {
		for i < JumMhs {

			j := i + 1
			min = i

			for j < JumMhs {

				if ArrMhs[j].jumsks < ArrMhs[min].jumsks {

					min = j
				}

				j++
			}

			temp.nama = ArrMhs[min].nama
			temp.ip = ArrMhs[min].ip
			temp.nim = ArrMhs[min].nim
			temp.mk = ArrMhs[min].mk
			temp.jumsks = ArrMhs[min].jumsks

			ArrMhs[min].nama = ArrMhs[i].nama
			ArrMhs[min].ip = ArrMhs[i].ip
			ArrMhs[min].nim = ArrMhs[i].nim
			ArrMhs[min].mk = ArrMhs[i].mk
			ArrMhs[min].jumsks = ArrMhs[i].jumsks

			ArrMhs[i].nama = temp.nama
			ArrMhs[i].ip = temp.ip
			ArrMhs[i].nim = temp.nim
			ArrMhs[i].mk = temp.mk
			ArrMhs[i].jumsks = temp.jumsks

			i++
		}

	} else {

		fmt.Println("Data masih kosong.")
	}
}

//RentangIP adalah prosedur yang akan memeriksa rentan IP
func RentangIP(ArrMhs *mhsw, min float64) {

	var awal, tengah, akhir int
	c := 0

	SortIPKB(ArrMhs)

	awal = 0
	akhir = JumMhs
	tengah = (awal + akhir) / 2

	for awal < akhir && int(min) != int(ArrMhs[tengah].ip) {

		if min > ArrMhs[tengah].ip {

			awal = tengah + 1
		} else if int(min) < int(ArrMhs[tengah].ip) {

			akhir = tengah - 1
		}

		tengah = (awal + akhir) / 2
	}

	if int(ArrMhs[tengah].ip) >= int(min) {

		for tengah < JumMhs {
			fmt.Println("Nama: ", ArrMhs[tengah].nama)
			fmt.Println("NIM: ", ArrMhs[tengah].nim)
			fmt.Printf("IP: %.2f\n", ArrMhs[tengah].ip)
			fmt.Println()
			c++
			tengah++
		}
	}

	if c == 0 {

		fmt.Println("Tidak ada data yang cocok.")

	} else {
		fmt.Println("Banyak Mahasiswa: ", c)
	}

}

//isNumeric adalah fungsi untuk mengubah data string nim menjadi float64, sekaligus mengecek nim harus numerik
func isNumeric(s string) bool {

	var temp float64

	temp, err := strconv.ParseFloat(s, 64)
	return (err == nil) && (temp >= 0)
}
