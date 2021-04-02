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
1. Membuat Proc DelMK beserta nilai(sesuai pdf) (WIP)
2. Menampilkan daftar mahasiswa terurut berdasarkan nilai dan jumlah SKS yang diambil
3. Menampilkan IP tertinggi.
4. Menampilkan transkrip nilai(sepertinya sudah bisa) (DONE)

Hasil malam ini :
1. Perlu diubah: Minta nama diganti nim (DONE)
2. Perlu Ditambahkan: Poin D, prioritas utama adalah nilai. Jika nilai sama, maka urutkan dari jumlah SKS
3. Perlu diperbaiki: Jika NIM sama, maka diminta untuk mengisi NIM ulang. (DONE) (DONE)
4. Perlu diperbaiki: Mata Kuliah tidak boleh sama. (DONE)
5. Perlu Diperhatikan: Case-Sensitive, seharusnya ya atau Ya atau yA tidak ada perlakuan berbeda.  (DONE)

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

package main

import (
	"fmt"
	"strings"
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
var ArrTemp [MhsMax]string

// ArrTempMK adalah variabel global untuk mengarsip data mata kuliah dari data mahasiswa yang di delete.
var ArrTempMK [JumMK]string

// ArrMhs adalah Array utama
var ArrMhs mhsw

func main() {

	var (
		nim     string
		pilmenu int
		junk    string
	)

	Menu()               // Menampilkan Menu
	fmt.Scanln(&pilmenu) // Memilih Menu

	for pilmenu < 1 || pilmenu > 10 { // Validasi input menu dari user harus pada interval 1-7

		fmt.Println("\nMaaf, tidak ada pilihan dalam menu :)\n")
		Menu()
		fmt.Scanln(&pilmenu)
	}

	for pilmenu != 10 { // Mengakses pilihan menu

		fmt.Println()

		if pilmenu == 1 { // Input Data Mahasiswa

			CreateMhs(&ArrMhs)

		} else if pilmenu == 2 { // Input Data Mata Kuliah

			fmt.Print("Data mata kuliah siapa yang ingin anda tambahkan? Masukan NIM: ")
			fmt.Scanln(&nim)
			CreateMK(nim, &ArrMhs)

		} else if pilmenu == 3 { // Output Data Mahasiswa

			fmt.Print("Data Mahasiswa yang ingin dilihat? Masukan NIM: ")
			fmt.Scanln(&nim)
			ReadMhs(nim, ArrMhs)

		} else if pilmenu == 4 { // Output Data Mata Kuliah

			fmt.Print("Data Mahasiswa yang Mata kuliah nya akan dilihat? Masukan NIM: ")
			fmt.Scanln(&nim)
			ReadMK(nim, ArrMhs)

		} else if pilmenu == 5 { // Update Data Mahasiswa

			fmt.Print("Masukan NIM: ")
			fmt.Scanln(&nim)
			UpMhs(nim, &ArrMhs)

		} else if pilmenu == 6 { // Update Data Nilai Mata Kuliah

			fmt.Print("Masukan NIM: ")
			fmt.Scanln(&nim)
			UpMK(nim, &ArrMhs)

		} else if pilmenu == 7 { // Output Semua Data Mahasiswa

			ReadAllMhs(ArrMhs)

		} else if pilmenu == 8 { // Delete 1 Data Mahasiswa (semua data)

			fmt.Print("Data Mahasiswa yang datanya ingin Anda hapus, Masukan NIM: ")
			fmt.Scanln(&nim)
			DelMhs(nim, &ArrMhs)

		} else if pilmenu == 9 { // Delete 1 mata kuliah milik seorang mahasiswa

			fmt.Print("Data Mahasiswa yang mata kuliahnya ingin Anda hapus, Masukan NIM: ")
			fmt.Scanln(&nim)
			DelMK(nim, &ArrMhs)

		}

		/*
			fmt.Print("Ranking Berdasarkan IP: ")
			SortIP(&ArrMhs)

			k := 0
			for k < JumMhs {

				fmt.Println(ArrMhs[k].nama)
				fmt.Println(ArrMhs[k].nim)
				fmt.Println(ArrMhs[k].ip)
				k++
			}
		*/

		fmt.Println("")
		fmt.Print("Tekan Enter untuk kembali ke Main Menu ...")
		fmt.Scanln(&junk)

		Menu()               // Menampilkan Menu kembali setelah mengakses 1-9
		fmt.Scanln(&pilmenu) // Memilih menu kembali setelah mengakses 1-9
	}

	// Jika pilihan menu adalah 10, Keluar dari Aplikasi
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

	fmt.Println("\n.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.-._.\n")
	fmt.Println("Menu: ")
	fmt.Println("1. Input Data Mahasiswa.")
	fmt.Println("2. Input Data Mata Kuliah.")
	fmt.Println("3. Output Data Mahasiswa.")
	fmt.Println("4. Output Data Mata Kuliah.")
	fmt.Println("5. Update Data Mahasiswa.")
	fmt.Println("6. Update Nilai Suatu Mata Kuliah.")
	fmt.Println("7. Output Semua Nama Mahasiswa.")
	fmt.Println("8. Hapus Data Mahasiswa.")
	fmt.Println("9. Hapus Mata Kuliah Mahasiswa.")
	fmt.Println("10. Keluar dari Aplikasi.")

	fmt.Print("Menu Anda: ")

}

// CreateMhs adalah prosedur untuk menambah data mahasiswa
func CreateMhs(ArrMhs *mhsw) {

	var nim string
	var qna string
	var input string

	i := 0
	fmt.Println("\nKetik STOP untuk berhenti")
	fmt.Print("Input nama Mahasiswa ke-", JumMhs+1, ": ")
	fmt.Scan(&input)

	for (JumMhs < MhsMax) && (strings.ToLower(input) != "stop") {

		ArrMhs[JumMhs].nama = input

		fmt.Print("Input NIM: ")
		fmt.Scan(&nim)

		j := 0

		for j <= JumMhs && ArrMhs[j].nim != nim {

			j++
		}

		for ArrMhs[j].nim == nim { // Jika NIM sama, maka diminta untuk mengisi NIM ulang.

			fmt.Println("NIM sama telah ditemukan. ")
			fmt.Print("Input NIM: ")
			fmt.Scan(&nim)

			j = 0
			for j <= JumMhs && ArrMhs[j].nim != nim {

				j++
			}
		}

		fmt.Print("Input Mata Kuliah?(Ya/Tidak): ")
		fmt.Scan(&qna)

		if strings.ToLower(qna) == "ya" {

			CreateMK(ArrMhs[JumMhs].nim, ArrMhs)
		}

		ArrMhs[JumMhs].nim = nim

		JumMhs++
		i++

		fmt.Print("Input nama Mahasiswa ke-", JumMhs+1, ": ")
		fmt.Scan(&input)

	}
}

// CreateMK adalah prosedur untuk menambah data mata kuliah pada mahasiswa
func CreateMK(nim string, ArrMhs *mhsw) {

	var (
		i, j, k, sks            int
		clo1, clo2, clo3, nilai float64
		matkul                  string
	)

	j = 0
	i = 0

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
			fmt.Scan(&matkul)

			k = 0
			for k < JumMK-1 && strings.ToLower(ArrMhs[i].mk[k].nama) != strings.ToLower(matkul) {
				k++

			}

			for strings.ToLower(ArrMhs[i].mk[k].nama) == strings.ToLower(matkul) {

				fmt.Println("Mata Kuliah sama telah ditemukan.")
				fmt.Printf("Mata Kuliah ke %v: ", j+1)
				fmt.Scan(&matkul)

				k = 0
				for k < JumMK-1 && strings.ToLower(ArrMhs[i].mk[k].nama) != strings.ToLower(matkul) {

					k++
				}
			}

			for strings.ToLower(matkul) != "stop" && j < JumMK {

				ArrMhs[i].mk[j].nama = matkul
				fmt.Print("Jumlah SKS: ")
				fmt.Scan(&sks)

				for sks <= 0 {

					fmt.Print("Jumlah SKS: ")
					fmt.Scan(&sks)

				}

				ArrMhs[i].mk[j].sks = sks

				ArrMhs[i].jumsks = ArrMhs[i].jumsks + ArrMhs[i].mk[j].sks

				fmt.Print("NIlai CLO 1: ")
				fmt.Scan(&clo1)

				for clo1 < 0 || clo1 > 100 {
					fmt.Println("Nilai di luar rentang 0-100")
					fmt.Print("NIlai CLO 1: ")
					fmt.Scan(&clo1)
				}

				ArrMhs[i].mk[j].clo1 = clo1

				fmt.Print("NIlai CLO 2: ")
				fmt.Scan(&clo2)

				for clo2 < 0 || clo2 > 100 {
					fmt.Println("Nilai di luar rentang 0-100")
					fmt.Print("NIlai CLO 2: ")
					fmt.Scan(&clo2)
				}

				ArrMhs[i].mk[j].clo2 = clo2

				fmt.Print("NIlai CLO 3: ")
				fmt.Scan(&clo3)

				for clo3 < 0 || clo3 > 100 {
					fmt.Println("Nilai di luar rentang 0-100")
					fmt.Print("NIlai CLO 3: ")
					fmt.Scan(&clo3)
				}

				ArrMhs[i].mk[j].clo3 = clo3

				ArrMhs[i].mk[j].na = HitungNSM(clo1, clo2, clo3)
				nilai = ArrMhs[i].mk[j].na

				ArrMhs[i].mk[j].nmk = HitungNMK(i, j, nilai)

				j++

				if j < JumMK {

					fmt.Printf("Mata Kuliah ke %v: ", j+1)
					fmt.Scan(&matkul)

					k = 0
					for k < JumMK-1 && strings.ToLower(ArrMhs[i].mk[k].nama) != strings.ToLower(matkul) {
						k++

					}

					for strings.ToLower(ArrMhs[i].mk[k].nama) == strings.ToLower(matkul) {

						fmt.Println("Mata Kuliah sama telah ditemukan.")
						fmt.Printf("Mata Kuliah ke %v: ", j+1)
						fmt.Scan(&matkul)

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
		ArrMhs[i].ip = HitungIP(ArrMhs, i)

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
		fmt.Println("IP: ", ArrMhs[i].ip)
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

	} else {

		fmt.Println("DATA TIDAK DITEMUKAN.")
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
		NamaBaru, nimBaru string // variabel baru untuk menyimpan nama baru dan nim baru
	)

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
	var clo1, clo2, clo3, nilai float64
	var j, i int

	i = 0
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

			// Memanggil prosedur ReadMhs untuk output Daftar sebelumnya
			fmt.Println()

			fmt.Print("Mata Kuliah yang akan di-Update: ")
			fmt.Scanln(&matkul)

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

	}
}

// ReadAllMhs adalah prosedur untuk mengoutputkan semua data mahasiswa
func ReadAllMhs(ArrMhs mhsw) {

	var nim string

	for i := 0; i < JumMhs; i++ {

		if ArrMhs[i].nim != "ARCHIVED" {

			fmt.Println("Nama: ", ArrMhs[i].nama)
			fmt.Println("NIM: ", ArrMhs[i].nim)
			fmt.Println("Jumlah SKS: ", ArrMhs[i].jumsks)
			fmt.Println("IP: ", ArrMhs[i].ip)
			fmt.Println("Daftar Mata Kuliah: ")
			nim = ArrMhs[i].nim
			ReadMK(nim, ArrMhs)

		}

	}

}

// DelMhs adalah prosedur untuk menghapus data mahasiswa (hide)
func DelMhs(nim string, ArrMhs *mhsw) {

	i := 0
	for i < JumMhs && ArrMhs[i].nim != nim { // mencari NIM yang cocok

		i++
	}

	if ArrMhs[i].nim == nim { // jika nim sudah cocok

		ArrTemp[i] = ArrMhs[i].nim // Data dipindahkan ke array ArrTemp
		ArrMhs[i].nim = "ARCHIVED"
		fmt.Println("Data Mahasiswa ", ArrTemp[i], "terhapus.")

	} else {

		fmt.Println("DATA TIDAK DITEMUKAN.")
	}
}

// DelMK adalah prosedur untuk menghapus data mata kuliah dari mahasiswa
func DelMK(nim string, ArrMhs *mhsw) {

	var matkul string
	var j int

	i := 0
	for i < JumMhs && ArrMhs[i].nim != nim { // mencari NIM yang cocok

		i++

	}

	if ArrMhs[i].nim == nim { // jika nim sudah cocok

		fmt.Println("Mata Kuliah yang akan dihapus: ")
		fmt.Scan(&matkul)

		for j < JumMK-1 && strings.ToLower(matkul) != strings.ToLower(ArrMhs[i].mk[j].nama) {

			j++
		}

		if strings.ToLower(matkul) == strings.ToLower(ArrMhs[i].mk[j].nama) {

			ArrTempMK[i] = ArrMhs[i].mk[j].nama // Data dipindahkan ke array ArrTemp
			ArrMhs[i].mk[j].nama = "ARCHIVED"
			ArrMhs[i].jumsks = ArrMhs[i].jumsks - ArrMhs[i].mk[j].sks
			fmt.Println("Mata Kuliah ", ArrTempMK[i], " terhapus.")

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

// SortIP adalah prosedur untuk mengurutkan data dari IP dari yang terbesar dengan INSERTION
func SortIP(ArrMhs *mhsw) {

	var temp mahasiswa

	i := 1
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
}

// SortNIM adalah prosedur untuk mengurutkan NIM Mahasiswa dari NIM dari yang terkecil dengan SELECTION
func SortNIM(ArrMhs *mhsw) {

}

/*
func CreateMK(nim string, ArrMhs *mhsw) {

	var i, sks int
	var matkul string
	var clo1, clo2, clo3 float64

	i = 0
	for i < JumMhs && ArrMhs[i].nim != nim {

		i++
	}

	if ArrMhs[i].nim == nim {

		j := 0
		for ArrMhs[i].mk[j].nama != "" && j < JumMK-1 {

			j++
		}
		if j <= JumMK &&   (j==0 || ArrMhs[i].mk[j].nama == "") {
			fmt.Printf("Mata Kuliah ke %v(Isi STOP untuk berhenti) : ", j+1)
			fmt.Scanln(&matkul)
		} else {
			fmt.Println("Mata Kuliah sudah penuh terisi.")
		}


		if (ArrMhs[i].mk[j].nama != "") && (strings.ToLower(matkul) != "stop"  && j <= JumMK) {

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

			ArrMhs[i].mk[j].na = HitungNSM(clo1, clo2, clo3)
			ArrMhs[i].mk[j].nmk = HitungNMK(i, j, ArrMhs)
		}

		for strings.ToLower(matkul) != "stop" && j < JumMK {

			for ArrMhs[i].mk[j].nama != "" && j < JumMK-1 {

				j++
			}

			fmt.Printf("Mata Kuliah ke %v: ", j+1)
			fmt.Scanln(&matkul)

			if strings.ToLower(matkul) != "stop"  && j < JumMK {

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

				ArrMhs[i].mk[j].na = HitungNSM(clo1, clo2, clo3)
				ArrMhs[i].mk[j].nmk = HitungNMK(i, j, ArrMhs)
			}
			j++
		}

	} else {

		fmt.Println("DATA TIDAK DITEMUKAN.")
	}
}
*/
