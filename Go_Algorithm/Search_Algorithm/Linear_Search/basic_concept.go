package main

import (
	"fmt"
)

/*
	Basic Problem :
		Diberikan array dengan n element di dalamnya
		Tulislah sebuah program yang dapat melakukan pencarian sebuah element x
		di dalam array

	?! BASIC SOLUTION !?
		Linear Search :
			Pendekatan :
			1. Dimulai dari indeks array paling kiri kemudian membandingkan satu per satu element dengan x
			2. Jika nilai x sesuai dengan nilai element di array pada indeks tertentu, maka kembalikan indeksnya
			3. Jika nilai x tidak ditemukan maka kembalikan nilai -1

	!! WORST CASE !!
		1. Jika element yang ditemukan berada pada indeks terakhir maka kompleksitas waktu nya O(n) to O(1)
		Keterangan :
			n = adalah banyak data dari array. sehingga, disimpulkan bahwa waktu pencarian bergantung pada banyak data

	-- IMPROVEMENT --
		1. Menggunakan 2 Pointer yang akan mengindeks element paling kiri (awal) dan satu pointer lagi mengindeks
			Element paling akhir (kanan)
		2. Dengan begitu, Jika Element ditemukan di sisi kanan atau kiri, maka kembalikan indeks
		3. Jika posisi pointer kiri dan kanan sama atau overlaping. maka kembalikan -1
		4. Kompleksitas waktu untuk algoritma ini O(1)
*/

func linearSearch(arr []int, x int) (int, bool) {
	for index, element := range arr {
		if element == x {
			return index, true
		}
	}
	return -1, false
}

func linearSearchImprovement(arr []int, x int) (int, bool) {
	// Siapkan pointer
	leftPointer := 0
	rightPointer := len(arr) - 1
	for leftPointer = 0; leftPointer <= rightPointer; {

		// Check apakah element x ditemukan di pointer kiri
		if x == arr[leftPointer] {
			return leftPointer, true
		}

		if x == arr[rightPointer] {
			return rightPointer, true
		}

		leftPointer++
		rightPointer--
	}

	return -1, false
}

func main() {

	// Array
	dataToSearch := []int{10, 20, 80, 30, 60, 50, 110, 100, 130, 170}

	// Given Element X
	x := 50

	// Call linearSearch here
	//result, cond := linearSearch(dataToSearch, x)
	result, cond := linearSearchImprovement(dataToSearch, x)

	if cond == true {
		fmt.Println("Data Found in Array!! Index:", result, "and Condition:", cond)
		return
	}
	fmt.Println("Data Not Found in Array!! Index:", result, "and Condition:", cond)
}
