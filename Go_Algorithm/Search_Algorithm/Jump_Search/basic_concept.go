package main

import (
	"fmt"
	"math"
)

/*
	JUMP SEARCH
		Sama seperti binary search bahwa jump search ini akan sangat fungsional ketika
		data yang diuji adalah data yang sudah di sorting.

		Konsep yang dipakai pada jump search adalah melakukan check seminimal mungkin dengan
		melakukan "loncatan" sampai elemen yang ada pada array n sama dengan atau lebih besar dari X

	>> CONTOH KASUS <<
		Asumsi jika kita mempunyai array Arr[] dengan ukuran n dan block yang akan menjadi loncatan sebanyak m
		Maka, dengan jump search kita akan melakukan check pada Arr[0], Arr[m], Arr[2m], ... Arr[km]


	!! WORST CASE !!
		Pada kasus terburuk jump search adalah dimana kita melakukan sebanyak n/m kali "loncatan"
		dan ketika element yang di dapat lebih besar nilainya daripada element yang menjadi target
		maka kita harus melakukan sebanyak m-1 perbandingan (Linear Search).
		Sehingga, total waktu yang ditempuh adalah ((n/m) + m-1)
*/

func jumpSearch(arr []int, target int) (int, bool) {
	// Pertama tentukan seberapa banyak block yang akan di langkahi
	arrSize := len(arr)
	m := int(math.Floor(math.Sqrt(float64(arrSize))))

	// Kemudian, temukan block dimana element target berada (Jika benar" berada)
	prevPosition := 0

	fmt.Println("Initial Step, prevPosition, ArrSize", m, ":", prevPosition, ":", arrSize)
	fmt.Println("Begin Skiping Element ...")
	fmt.Println("")

	for arr[int(math.Min(float64(m), float64(arrSize))-1)] < target {
		prevPosition = m
		m += int(math.Floor(math.Sqrt(float64(arrSize))))
		fmt.Println("PrevPosition :", prevPosition, " Next Step Index:", m)
		fmt.Println("Checking if PrevPosition > arrSize...")
		if prevPosition >= arrSize {
			fmt.Println("Condition True!! Can't find value", target, "in array!")
			return -1, false
		}
	}

	fmt.Println()
	fmt.Println("Value in array more than target! Do Linear Search!")
	fmt.Println()

	for arr[prevPosition] < target {
		prevPosition++
		fmt.Println("Element in Arr[", prevPosition, "]:", arr[prevPosition], "Target :", target)

		if prevPosition == int(math.Min(float64(m), float64(arrSize))) {
			fmt.Println("Condition True!! Can't find value", target, "in array!")
			return -1, false
		}
	}

	if arr[prevPosition] == target {
		fmt.Println("Condition True!! Found value", target, "in array at Index", prevPosition, "!")
		return prevPosition, true
	}
	return -1, false
}

func main() {
	// Sorted Array
	//dataToSearch := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610}

	// Non Sorted Array
	dataNonSorted := []int{10, 20, 80, 30, 60, 50, 110, 100, 170, 130}

	// Target Value
	X := 130

	resultIndex, cond := jumpSearch(dataNonSorted, X)

	if cond == true {
		fmt.Println("Found target value", X, "in array at Index:", resultIndex)
		return
	}
	fmt.Println("Condition True!! Can't find value", X, "in array!")
}
