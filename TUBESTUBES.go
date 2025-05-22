package main

import "fmt"

const NMAX int = 10

type trashData struct {
	nama, jenis, daur string
	jumlah            int
	berat             float64
}
type ArrTrash [NMAX]trashData

func main() {
	var data ArrTrash
	var n int
	var cmd string
	var stop bool = false
	fmt.Println("Selamat Datang Di Aplikasi Pengolahan Data Sampah!")

	for !stop {
		Awal()
		fmt.Scan(&cmd)
		switch cmd {
		case "1":
			input(&data, &n)
		case "2":
			modify(&data, &n)
		case "3":
			search(&data, n)
		case "4":
			view(data, n)
		case "5":
			stat(&data, &n)
		case "Exit", "exit":
			stop = true
		default:
			fmt.Println("Error, Silahkan Coba lagi")
		}
	}
	fmt.Println("Terimakasih telah menggunkan kami!")
}

func Awal() {
	fmt.Println("Apa yang ingin anda lakukan?")
	fmt.Println("   1. Memasukkan Data Sampah")
	fmt.Println("   2. Memodifikasi Data Sampah")
	fmt.Println("   3. Mencari Data Sampah")
	fmt.Println("   4. Melihat Data Sampah")
	fmt.Println("   5. Menampilkan Statistik Data")
	fmt.Println("---------------------------------")
	fmt.Println("Exit")
	fmt.Print("Masukkan pilihan (1/2/3/4/5/Exit): ")
}

func input(A *ArrTrash, n *int) {
	var i int
	fmt.Print("Silahkan Input Berapa Banyak Data Yang Akan Dimasukkan : ")
	fmt.Scan(n)
	fmt.Println("Silahkan input dalam format [Nama/Jenis/Jumlah/Berat(kg)/Daur Ulang(yes/no)]")
	for i = 1; i <= *n; i++ {
		fmt.Scan(&(*A)[i].nama, &(*A)[i].jenis, &(*A)[i].jumlah, &(*A)[i].berat, &(*A)[i].daur)
	}
	fmt.Println("   Input Data Berhasil!")
	fmt.Println("Redirecting to main menu...")
}

func modify(A *ArrTrash, n *int) {
	var cmd string
	var stop bool = false

	for !stop {
		fmt.Println("Apa yang ingin anda lakukan?")
		fmt.Println("   1. Mengganti Data Sampah")
		fmt.Println("   2. Menambah Data Sampah")
		fmt.Println("   3. Menghapus Data Sampah")
		fmt.Println("   4. Mengurutkan Data Sampah")
		fmt.Println("---------------------------------")
		fmt.Println("Back")
		fmt.Print("Masukkan pilihan (1/2/3/4/Back): ")
		fmt.Scan(&cmd)
		switch cmd {
		case "1":
			ChData(A, n)
			stop = true
		case "2":
			add(A, n)
			stop = true
		case "3":
			delete(A, n)
			stop = true
		case "4":
			sort(A, n)
			stop = true
		case "back", "Back":
			fmt.Println("Redirecting to main menu...")
			stop = true
		default:
			fmt.Println("Error, Silahkan Coba lagi")

		}
	}
}

func delete(A *ArrTrash, n *int) {
	var idx, i int
	var stop bool

	for !stop {
		fmt.Print("Nomor Data Yang Akan Dihapus : ")
		fmt.Scan(&idx)
		if idx >= 1 && idx <= *n {
			for i = idx; i < *n; i++ {
				(*A)[i] = (*A)[i+1]

			}
			*n--
			fmt.Println("Data Berhasil DiHapus!")
			fmt.Println("Redirecting to main menu...")
			stop = true
		} else {
			fmt.Println("Error, Silahkan coba lagi")
		}
	}
}
func add(A *ArrTrash, n *int) {
	var i, plus int
	fmt.Print("Ingin menambah berapa data? : ")
	fmt.Scan(&plus)
	plus += *n
	fmt.Println("Silahkan input dalam format [Nama/Jenis/Jumlah/Berat(kg)/Daur Ulang(yes/no)]")
	for i = *n + 1; i <= plus; i++ {
		fmt.Scan(&(*A)[i].nama, &(*A)[i].jenis, &(*A)[i].jumlah, &(*A)[i].berat, &(*A)[i].daur)
	}
	fmt.Println("   Input Data Berhasil!")
	fmt.Println("Redirecting to main menu...")
	*n = plus

}

func view(A ArrTrash, n int) {
	var i int
	//fmt.Printf("%-2\n")
	for i = 1; i <= n; i++ {
		fmt.Println("----------------------------------------------------------")
		fmt.Printf("%-2d |%-20s |%-10s |%-5d |%-2.1f kg |%-3s |\n", i, A[i].nama, A[i].jenis, A[i].jumlah, A[i].berat, A[i].daur)
	}
	fmt.Println("----------------------------------------------------------")
}

func stat(A *ArrTrash, n *int) {
	var i int
	var ano, org int
	var daurY, daurN int
	var jumJum int
	var jumBer, ratBer float64
	for i = 1; i <= *n; i++ {
		if A[i].jenis == "Anorganik" || A[i].jenis == "anorganik" {
			ano++
		} else if A[i].jenis == "Organik" || A[i].jenis == "organik" {
			org++
		}
		if A[i].daur == "Yes" || A[i].daur == "yes" {
			daurY++
		} else if A[i].daur == "No" || A[i].daur == "no" {
			daurN++
		}
		jumBer += A[i].berat
		jumJum += A[i].jumlah
	}
	ratBer = jumBer / float64(i)
	fmt.Printf("Total jumlah data adalah : %d\n", i)
	fmt.Printf("Total jumlah data yang Anorganik adalah : %d\n", ano)
	fmt.Printf("Total jumlah data yang Organik adalah : %d\n", org)
	fmt.Printf("Total jumlah Sampah yang terkumpul adalah : %d\n", jumJum)
	fmt.Printf("Total jumlah Berat sampah yang terkumpul adalah : %.2f dan rata-ratanya adalah : %.2f", jumBer, ratBer)
	fmt.Printf("Total jumlah Sampah yang terolah : %d, dan yang tidak : %d\n", daurY, daurN)

}

func ChData(A *ArrTrash, n *int) {
	var idx int
	fmt.Println("Nomor Data yang ingin diubah : ")
	fmt.Scan(&idx)
	if idx >= 1 && idx <= *n {
		fmt.Scan(&A[idx].nama, &A[idx].jenis, &A[idx].jumlah, &A[idx].berat, &A[idx].daur)
	} else {
		fmt.Println("")
	}
}

func search(A *ArrTrash, n int) {
	var search, key string
	var keyIn, i, total int
	var keyFlo float64
	var ketemu bool = false
	var stop bool = false
	fmt.Print("Ingin Mensearch apa? : ")
	fmt.Scan(&search)

	for !stop {
		switch search {
		case "Nama", "nama":
			fmt.Print("Silahkan masukkan <nama> data yang akan dicari: ")
			fmt.Scan(&key)
			total = 0
			for i = 1; i <= n; i++ {
				if key == A[i].nama {
					fmt.Println("--------------------------------------------------------")
					fmt.Printf("|%-20s |%-10s |%-5d |%-4.1f kg |%-3s |\n", A[i].nama, A[i].jenis, A[i].jumlah, A[i].berat, A[i].daur)
					ketemu = true
					total++
				}
			}
			if !ketemu {
				fmt.Println("Tidak Ketemu")
			} else {
				fmt.Printf("Ada %d Data yang ditemukan\n", total)
			}
			fmt.Println("--------------------------------------------------------")
			stop = true
		case "Jenis", "jenis":
			fmt.Print("Silahkan masukkan <Jenis> data yang akan dicari: ")
			fmt.Scan(&key)
			total = 0
			for i = 1; i <= n; i++ {
				if key == A[i].jenis {
					fmt.Println("--------------------------------------------------------")
					fmt.Printf("|%-20s |%-10s |%-5d |%-2.1f kg |%-3s |\n", A[i].nama, A[i].jenis, A[i].jumlah, A[i].berat, A[i].daur)
					ketemu = true
					total++
				}
			}
			if !ketemu {
				fmt.Println("Tidak Ketemu")
			} else {
				fmt.Printf("Ada %d Data yang ditemukan\n", total)
			}
			fmt.Println("--------------------------------------------------------")
			stop = true
		case "Jumlah", "jumlah":
			fmt.Print("Silahkan masukkan <Jumlah> data yang akan dicari: ")
			fmt.Scan(&keyIn)
			total = 0
			for i = 1; i <= n; i++ {
				if keyIn == A[i].jumlah {
					fmt.Println("--------------------------------------------------------")
					fmt.Printf("|%-20s |%-10s |%-5d |%-2.1f kg |%-3s |\n", A[i].nama, A[i].jenis, A[i].jumlah, A[i].berat, A[i].daur)
					ketemu = true
					total++
				}
			}
			if !ketemu {
				fmt.Println("Tidak Ketemu")
			} else {
				fmt.Printf("Ada %d Data yang ditemukan\n", total)
			}
			fmt.Println("--------------------------------------------------------")
			stop = true
		case "Berat", "berat":
			fmt.Print("Silahkan masukkan <nama> data yang akan dicari: ")
			fmt.Scan(&keyIn)
			total = 0
			for i = 1; i <= n; i++ {
				if keyFlo == A[i].berat {
					fmt.Println("--------------------------------------------------------")
					fmt.Printf("|%-20s |%-10s |%-5d |%-2.1f kg |%-3s |\n", A[i].nama, A[i].jenis, A[i].jumlah, A[i].berat, A[i].daur)
					ketemu = true
					total++
				}
			}
			if !ketemu {
				fmt.Println("Tidak Ketemu")
			} else {
				fmt.Printf("Ada %d Data yang ditemukan\n", total)
			}
			fmt.Println("--------------------------------------------------------")
			stop = true
		case "Daur", "daur":
			fmt.Print("Silahkan masukkan <nama> data yang akan dicari: ")
			fmt.Scan(&key)
			total = 0
			for i = 1; i <= n; i++ {
				if key == A[i].daur {
					fmt.Println("--------------------------------------------------------")
					fmt.Printf("|%-20s |%-10s |%-5d |%-5.1f kg |%-3s |\n", A[i].nama, A[i].jenis, A[i].jumlah, A[i].berat, A[i].daur)
					ketemu = true
					total++
				}
			}
			if !ketemu {
				fmt.Println("Tidak Ketemu")
			} else {
				fmt.Printf("Ada %d Data yang ditemukan\n", total)
			}
			fmt.Println("--------------------------------------------------------")
			stop = true
		default:
			fmt.Println("Error silahkan coba lagi")
		}

	}
}
func menuSort() {
	fmt.Println("============================================")
	fmt.Println("           MENU SORTIR DATA SAMPAH          ")
	fmt.Println("============================================")
	fmt.Println("Pilih data yang ingin disortir:")
	fmt.Println("   1. Berat")
	fmt.Println("   2. Jumlah")
	fmt.Println("   3. Jenis")
	fmt.Println("   4. Daur Ulang")
	fmt.Println("---------------------------------------------")
	fmt.Println("Back")
	fmt.Print("Masukkan pilihan (1/2/3/4/Back): ")
}

func sort(A *ArrTrash, n *int) {
	var tipe, urutan string
	var stop bool = false

	for !stop {
		menuSort()
		fmt.Scan(&tipe)

		if tipe == "Back" || tipe == "back" {
			fmt.Println("Redirecting to main menu...")
			stop = true
		} else {
			switch tipe {
			case "1":
				fmt.Print("Masukkan urutan data (asc/dsc): ")
				fmt.Scan(&urutan)
				sortBerat(A, *n, urutan)
				stop = true
			case "2":
				fmt.Print("Masukkan urutan data (asc/dsc): ")
				fmt.Scan(&urutan)
				sortJumlah(A, *n, urutan)
				stop = true
			case "3":
				fmt.Print("Masukkan urutan data (asc/dsc): ")
				fmt.Scan(&urutan)
				sortJenis(A, *n, urutan)
				stop = true
			case "4":
				fmt.Print("Masukkan urutan data (asc/dsc): ")
				fmt.Scan(&urutan)
				sortDaur(A, *n, urutan)
				stop = true
			default:
				fmt.Println("Pilihan tidak invalid, silahkan input kembali !")
			}
		}
	}
	fmt.Println("Data berhasil disortir!")
}

func sortBerat(A *ArrTrash, n int, urutan string) {
	var i, j int
	var temp trashData
	for i = 2; i <= n; i++ {
		temp = A[i]
		j = i - 1

		if urutan == "asc" {
			for j >= 1 && A[j].berat > temp.berat {
				A[j+1] = A[j]
				j--
			}
		} else if urutan == "dsc" {
			for j >= 1 && A[j].berat < temp.berat {
				A[j+1] = A[j]
				j--
			}
		}

		A[j+1] = temp
	}
}

func sortJenis(A *ArrTrash, n int, urutan string) {
	var i, idx, j int
	var temp trashData
	for i = 1; i < n; i++ {
		idx = i
		for j = i + 1; j <= n; j++ {
			if (urutan == "asc" && A[j].jenis < A[idx].jenis) || (urutan == "dsc" && A[j].jenis > A[idx].jenis) {
				idx = j
			}
		}
		temp = A[i]
		A[i] = A[idx]
		A[idx] = temp
	}
}

func sortJumlah(A *ArrTrash, n int, urutan string) {
	var i, j int
	var temp trashData
	for i = 2; i <= n; i++ {
		temp = A[i]
		j = i - 1

		if urutan == "asc" {
			for j >= 1 && A[j].jumlah > temp.jumlah {
				A[j+1] = A[j]
				j--
			}
		}
		if urutan == "dsc" {
			for j >= 1 && A[j].jumlah < temp.jumlah {
				A[j+1] = A[j]
				j--
			}
		}
		A[j+1] = temp
	}
}

func sortDaur(A *ArrTrash, n int, urutan string) {
	var i, idx, j int
	var temp trashData
	for i = 1; i < n; i++ {
		idx = i
		for j = i + 1; j <= n; j++ {
			if (urutan == "asc" && A[j].daur < A[idx].daur) || (urutan == "dsc" && A[j].daur > A[idx].daur) {
				idx = j
			}
		}
		temp = A[i]
		A[i] = A[idx]
		A[idx] = temp
	}
}
