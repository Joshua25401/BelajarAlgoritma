package main

import "fmt"

/*
	Binary Search
		BS biasanya digunakan untuk melakukan pencarian pada array yang sudah diurutkan sebelumnya
		Konsep dari BS sendiri yaitu membagi ukuran array menjadi setengah begitu seterusnya
		hingga ukuran menjadi lebih kecil secara rekursif dan mendapatkan nilai dari element x

	!! ALGORITHM !!
	1. Dimulai dari melihat kedalam keseluruhan array
	2. Jika nilai dari kata kunci pencarian yaitu x lebih kecil dari nilai item yang berada di tengah
		dari interval pencarian maka cari nilai middle yang baru yang lebih kecil.
	3. Jika sebaliknya maka cari nilai middle baru yang lebih besar.
	4. Ulangi langkah tersebut hingga selesai.
*/

func binarySearch(arr []int, leftIndex int, rightIndex int, x int) (int, bool) {
	if rightIndex >= 1 {
		if rightIndex < leftIndex {
			return -1, false
		}

		middleIndex := (leftIndex + rightIndex) / 2 // use this to avoid ArrayIndexOutOfBound Exception
		// Base Case
		// Jika array pada index mid tertentu sama dengan nilai x maka hentikan program dan kembalikan nilai
		if arr[middleIndex] == x {
			return middleIndex, true
		}

		// Jika element yang ada pada variabel array di index mid lebih besar dari X
		// Maka lakukan pencarian pada bagian kiri array saja
		if arr[middleIndex] > x {
			return binarySearch(arr, leftIndex, middleIndex-1, x)
		}

		// Jika ternyata element pada variabel array di index mid lebih kecil dari X
		// Maka lakukan pencarian pada bagian kanan array dengan memindahkan leftIndex menjadi middleIndex + 1
		return binarySearch(arr, middleIndex+1, rightIndex, x)
	}

	return -1, false
}

func main() {
	dataToSearch := []int{10, 20, 30, 50, 60, 80, 110, 130, 140, 170}

	// Nilai X
	x := 999

	resultIndex, cond := binarySearch(dataToSearch, 0, len(dataToSearch)-1, x)

	if cond == true {
		fmt.Println("Value", x, "Found in Index", resultIndex, "Within the Array")
		return
	}
	fmt.Println("Value", x, "Not Found Within the Array")
}
