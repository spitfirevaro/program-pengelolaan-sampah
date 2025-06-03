package main

import "fmt"

var sortedJumlah = false
var sortedBerat = false

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
		awal()
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

func awal() {
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
	fmt.Println("+-------------------------------------------------------------------------------+")
	fmt.Printf("|%-2s |%-25s |%-10s |%-7s |%-12s |%-12s |\n", "No", "Nama", "Jenis", "Jumlah", "Berat (kg)", "Daur (Y/N)")
	for i = 1; i <= n; i++ {
		fmt.Println("+-------------------------------------------------------------------------------+")
		fmt.Printf("|%-2d |%-25s |%-10s |%-7d |%-10.1f kg|%-12s |\n", i, A[i].nama, A[i].jenis, A[i].jumlah, A[i].berat, A[i].daur)
	}
	fmt.Println("+-------------------------------------------------------------------------------+")
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
	fmt.Printf("Total jumlah Berat sampah yang terkumpul adalah : %.2f dan rata-ratanya adalah : %.2f\n", jumBer, ratBer)
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
	var search string
	var stop bool = false

	for !stop {
		fmt.Print("Ingin Mencari apa? : ")
		fmt.Scan(&search)
		switch search {
		case "Nama", "nama":
			searchNama(A, n)
			stop = true
		case "Jenis", "jenis":
			searchJenis(A, n)
			stop = true
		case "Jumlah", "jumlah":
			searchJumlah(A, n)
			stop = true
		case "Berat", "berat":
			searchBerat(A, n)
			stop = true
		case "Daur", "daur":
			searchDaur(A, n)
			stop = true
		}
	}
}

func searchJumlah(A *ArrTrash, n int) {
	const ncopy int = 10
	type arrcopy [ncopy]trashData
	var B arrcopy
	var temp trashData
	var i, keyIn, j int
	var low, high, mid, idx int
	var ketemu bool = false
	var halt bool = false

	fmt.Print("Silahkan masukkan <Jumlah> data yang akan dicari: ")
	fmt.Scan(&keyIn)
	halt = false
	low = 1
	high = n
	idx = -1

	for i = 1; i <= n; i++ {
		B[i].nama = A[i].nama
		B[i].jenis = A[i].jenis
		B[i].jumlah = A[i].jumlah
		B[i].berat = A[i].berat
		B[i].daur = A[i].daur
	}
	if !sortedJumlah {
		for i = 2; i <= n; i++ {
			temp = B[i]
			j = i - 1
			for j >= 1 && B[j].jumlah > temp.jumlah {
				B[j+1] = B[j]
				j--
			}
			B[j+1] = temp
		}
	}

	for low <= high && !halt {
		mid = low + (high-low)/2
		if B[mid].jumlah < keyIn {
			low = mid + 1
		} else if B[mid].jumlah > keyIn {
			high = mid - 1
		} else {
			idx = mid
			ketemu = true
			halt = true
		}
	}

	if !ketemu {
		fmt.Println("+---------------------------------------------------------------------------+")
		fmt.Println("| Tidak Ketemu                                                              |")
	} else {
		fmt.Println("+---------------------------------------------------------------------------+")
		fmt.Printf("|%-25s |%-10s |%-7s |%-12s |%-12s |\n", "Nama", "Jenis", "Jumlah", "Berat (kg)", "Daur (Y/N)")
		fmt.Println("+---------------------------------------------------------------------------+")
		fmt.Printf("|%-25s |%-10s |%-7d |%-10.1f kg|%-12s |\n", B[idx].nama, B[idx].jenis, B[idx].jumlah, B[idx].berat, B[idx].daur)
	}

	fmt.Println("+---------------------------------------------------------------------------+")
}
func searchBerat(A *ArrTrash, n int) {
	const ncopy int = 10
	type arrcopy [ncopy]trashData
	var B arrcopy
	var temp trashData
	var i, j int
	var key float64
	var low, high, mid, idx int
	var ketemu bool = false
	var halt bool = false

	fmt.Print("Silahkan masukkan <Berat> data yang akan dicari: ")
	fmt.Scan(&key)
	halt = false
	low = 1
	high = n
	idx = -1

	for i = 1; i <= n; i++ {
		B[i].nama = A[i].nama
		B[i].jenis = A[i].jenis
		B[i].jumlah = A[i].jumlah
		B[i].berat = A[i].berat
		B[i].daur = A[i].daur
	}
	if !sortedBerat {
		for i = 2; i <= n; i++ {
			temp = B[i]
			j = i - 1
			for j >= 1 && B[j].berat > temp.berat {
				B[j+1] = B[j]
				j--
			}
			B[j+1] = temp
		}
	}

	for low <= high && !halt {
		mid = low + (high-low)/2
		if B[mid].berat < key {
			low = mid + 1
		} else if B[mid].berat > key {
			high = mid - 1
		} else {
			idx = mid
			ketemu = true
			halt = true
		}
	}

	if !ketemu {
		fmt.Println("+---------------------------------------------------------------------------+")
		fmt.Println("| Tidak Ketemu                                                              |")
	} else {
		fmt.Println("+---------------------------------------------------------------------------+")
		fmt.Printf("|%-25s |%-10s |%-7s |%-12s |%-12s |\n", "Nama", "Jenis", "Jumlah", "Berat (kg)", "Daur (Y/N)")
		fmt.Println("+---------------------------------------------------------------------------+")
		fmt.Printf("|%-25s |%-10s |%-7d |%-10.1f kg|%-12s |\n", B[idx].nama, B[idx].jenis, B[idx].jumlah, B[idx].berat, B[idx].daur)
	}

	fmt.Println("+---------------------------------------------------------------------------+")
}

func searchDaur(A *ArrTrash, n int) {
	var key string
	var total, i, first int
	var ketemu bool = false
	fmt.Print("Silahkan masukkan <Daur (Yes/No)> data yang akan dicari: ")
	fmt.Scan(&key)
	total = 0
	first = 0
	for i = 1; i <= n; i++ {
		if key == A[i].daur {
			if first == 0 {
				fmt.Println("+-------------------------------------------------------------------------------+")
				fmt.Printf("|%-2s |%-25s |%-10s |%-7s |%-12s |%-12s |\n", "No", "Nama", "Jenis", "Jumlah", "Berat (kg)", "Daur (Y/N)")
				first++
			}
			fmt.Println("+-------------------------------------------------------------------------------+")
			fmt.Printf("|%-2d |%-25s |%-10s |%-7d |%-10.1f kg|%-12s |\n", i, A[i].nama, A[i].jenis, A[i].jumlah, A[i].berat, A[i].daur)
			ketemu = true
			total++
		}
	}
	if !ketemu {
		fmt.Println("+-------------------------------------------------------------------------------+")
		fmt.Println("| Tidak Ketemu                                                                  |")
	} else {
		fmt.Println("+-------------------------------------------------------------------------------+")
		fmt.Printf("| Ada %d Data yang ditemukan                                                    |\n", total)

	}
	fmt.Println("+-------------------------------------------------------------------------------+")
}

func searchJenis(A *ArrTrash, n int) {
	var key string
	var total, i, first int
	var ketemu bool = false
	fmt.Print("Silahkan masukkan <Jenis> data yang akan dicari: ")
	fmt.Scan(&key)
	total = 0
	first = 0
	for i = 1; i <= n; i++ {
		if key == A[i].jenis {
			if first == 0 {
				fmt.Println("+-------------------------------------------------------------------------------+")
				fmt.Printf("|%-2s |%-25s |%-10s |%-7s |%-12s |%-12s |\n", "No", "Nama", "Jenis", "Jumlah", "Berat (kg)", "Daur (Y/N)")
				first++
			}
			fmt.Println("+-------------------------------------------------------------------------------+")
			fmt.Printf("|%-2d |%-25s |%-10s |%-7d |%-10.1f kg|%-12s |\n", i, A[i].nama, A[i].jenis, A[i].jumlah, A[i].berat, A[i].daur)
			ketemu = true
			total++
		}
	}
	if !ketemu {
		fmt.Println("+-------------------------------------------------------------------------------+")
		fmt.Println("| Tidak Ketemu                                                                  |")
	} else {
		fmt.Println("+-------------------------------------------------------------------------------+")
		fmt.Printf("| Ada %d Data yang ditemukan                                                     |\n", total)

	}
	fmt.Println("+-------------------------------------------------------------------------------+")
}

func searchNama(A *ArrTrash, n int) {
	var key string
	var total, i, first int
	var ketemu bool = false
	fmt.Print("Silahkan masukkan <nama> data yang akan dicari: ")
	fmt.Scan(&key)
	total = 0
	first = 0
	for i = 1; i <= n; i++ {
		if key == A[i].nama {
			if first == 0 {
				fmt.Println("+-------------------------------------------------------------------------------+")
				fmt.Printf("|%-2s |%-25s |%-10s |%-7s |%-12s |%-12s |\n", "No", "Nama", "Jenis", "Jumlah", "Berat (kg)", "Daur (Y/N)")
				first++
			}
			fmt.Println("+-------------------------------------------------------------------------------+")
			fmt.Printf("|%-2d |%-25s |%-10s |%-7d |%-10.1f kg|%-12s |\n", i, A[i].nama, A[i].jenis, A[i].jumlah, A[i].berat, A[i].daur)
			ketemu = true
			total++
		}
	}
	if !ketemu {
		fmt.Println("+-------------------------------------------------------------------------------+")
		fmt.Println("| Tidak Ketemu                                                                  |")
	} else {
		fmt.Println("+-------------------------------------------------------------------------------+")
		fmt.Printf("| Ada %d Data yang ditemukan                                                     |\n", total)

	}
	fmt.Println("+-------------------------------------------------------------------------------+")
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
	sortedBerat = true
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
	sortedBerat = false
	sortedJumlah = false
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
	sortedJumlah = true
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
	sortedJumlah = false
	sortedBerat = false
}
