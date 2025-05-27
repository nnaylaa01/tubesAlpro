package main

import "fmt"

const nMAX = 2026

type tabPolusi [nMAX]Polusi
type tabWaktu [nMAX]Waktu

type Polusi struct {
	lokasi        string
	tingkatPolusi float64
	sumber        string
	status		  string
}

type Waktu struct {
	tanggal, bulan, tahun int
}

func bacaData(A *tabPolusi, B *tabWaktu, n *int) {
	var i int
	fmt.Scan(n)
	for i = 1; i <= *n; i++ {
		fmt.Scan(&A[i].lokasi, &A[i].tingkatPolusi, &A[i].sumber, &B[i].tanggal, &B[i].bulan, &B[i].tahun)
	}
}

func cetakData(A *tabPolusi, B *tabWaktu, n int) {
	var i int
	fmt.Printf("%-20s %-10.1f %-20s %02d/%02d/%04d %-15s\n", "Sumber", "Waktu", "Status")
	for i = 1; i <= n; i++ {
		fmt.Printf("%-20s %-10.1f %-20s %02d/%02d/%04d %-15s\n", A[i].lokasi, A[i].tingkatPolusi, A[i].sumber, B[i].tanggal, B[i].bulan, B[i].tahun, A[i].status)
	}
}

func selectionSort(B *tabWaktu, A *tabPolusi, n int) {
	var pass, i, idx int
	var temp Waktu
	var temp2 Polusi
	for pass = 1; pass < n-1; pass++ {
		idx = pass
		for i = pass + 1; i < n; i++ {
			for i = pass + 1; i < n; i++ {
				if B[i].tahun > B[idx].tahun || (B[i].tahun == B[idx].tahun && B[i].bulan > B[idx].bulan) || (B[i].tahun == B[idx].tahun && B[i].bulan == B[idx].bulan && B[i].tanggal > B[idx].tanggal) {
					idx = i
				}
			}
			temp = B[pass]
			B[pass] = B[idx]
			B[idx] = temp

			temp2 = A[pass]
			A[pass] = A[idx]
			A[idx] = temp2
		}
	}
}

func InsertionSort(A *tabPolusi, B *tabWaktu, n int) {
	var pass, i int
	var temp Polusi
	var temp2 Waktu
	for pass = 1; pass <= n-1; pass++ {
		i = pass
		temp = A[pass]
		temp2 = B[pass]
		for i > 0 && temp.tingkatPolusi > A[i-1].tingkatPolusi {
			A[i] = A[i-1]
			B[i] = B[i-1]
			i = i - 1
		}
		A[i] = temp
		B[i] = temp2
	}
}

func kategoriPolusi(idxPolusi float64) string {
	if idxPolusi <= 2.0 {
		return "Baik"
	} else if idxPolusi <= 5.0 {
		return "Sedang"
	} else if idxPolusi <= 8.0 {
		return "Tidak sehat"
	} else {
		return "Berbahaya"
	}
}

func seqSearch(A tabPolusi, n int, x string) int{
	var i int
	
	for i = 1; i <= n; i++ {
		if A[i].lokasi == x {
			return i
		}
	}
	return -1
}

func binarySearch( A tabPolusi, n int, x string) int{
	var left, right, mid int
	
	left = 1
	right = n
	
	for left <= right{
		mid = (left + right) / 2
		if A[mid].lokasi == x{
			return mid
		}else if x < A[mid].lokasi{
			right = mid
		}else {
			left = mid
		}
	}
	return -1
}

func cariPolusi(A tabPolusi, B tabWaktu, n int) int{
	var Lokasi string
	var idx int

	fmt.Println("Masukkan lokasi kota atau wilayah yang ingin dicari : ")
	fmt.Scan(&Lokasi)
	
	idx = seqSearch(A, n, Lokasi)
	
	if idx == -1 {
		fmt.Println("Data tidak ditemukan")
	}else {
	fmt.Println("Data ditemukan")
	fmt.Printf("%-20s %-10.1f %-20s %02d/%02d/%04d %-15s\n", A[idx].lokasi, A[idx].tingkatPolusi, A[idx].sumber, B[idx].tanggal, B[idx].bulan, B[idx].tahun, A[idx].status)
	}
	return idx
}

func hapusData(A *tabPolusi, B *tabWaktu, n *int){	
	var i, idx int
	
	idx = cariPolusi(*A, *B, *n)
	
	if idx == -1 {
		return
	}	
	
	for i = idx; i < *n; i++ {
		(*A)[i] = (*A)[i+1]
		(*B)[i] = (*B)[i+1]
	}
	*n = *n - 1
	fmt.Println("Data berhasil dihapus")
}

func main() {
	var polusi tabPolusi
	var wkt tabWaktu
	var nPolusi, x, sub int
	var isRun bool
	isRun = true
	for isRun {
		fmt.Println("Selamat Datang di Aplikasi Manajemen dan Pemantauan Polusi Udara Lokal: ")
		fmt.Println("1. Input Data")
		fmt.Println("2. Mencari Data")
		fmt.Println("3. Mengurutkan Data")
		fmt.Println("4. Menghapus Data")
		fmt.Println("5. Out")
		fmt.Println("----------------------")
		fmt.Println("Pilih 1/2/3/4/5")
		fmt.Scan(&x)

		if x == 1 {
			fmt.Println("catatan : input Tanggal Pengukuran dengan format d/m/yyyy")
			bacaData(&polusi, &wkt, &nPolusi)
		} else if x == 2 {
			cariPolusi(polusi, wkt, nPolusi)
		} else if x == 3 {
			fmt.Println("1. Waktu terbaru")
			fmt.Println("2. Polusi Tertinggi")
			fmt.Println("Pilih 1/2")
			fmt.Scan(&sub)

			if sub == 1 {
				fmt.Println("Berdasarkan pengukuran terbaru: ")
				selectionSort(&wkt, &polusi, nPolusi)
				cetakData(&polusi, &wkt, nPolusi)
			} else if sub == 2 {
				fmt.Println("Berdasarkan polusi tertinggi: ")
				InsertionSort(&polusi, &wkt, nPolusi)
				cetakData(&polusi, &wkt, nPolusi)
			} else {
				fmt.Println("input tidak valid")
			}
		
		}else if x == 4{
			hapusData(&polusi, &wkt, &nPolusi)
		} else if x == 5 {
			fmt.Println("Terimakasih sudah menggunakan Aplikasi Manajemen dan Pemantauan Polusi Udara Lokal")
			fmt.Println("Have a Nice Day:))")
			isRun = false
		} else {
			fmt.Println("input tidak valid")
		}
		fmt.Println()
	}
}