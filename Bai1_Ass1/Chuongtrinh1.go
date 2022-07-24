//Truong hop bai toan nay truyen vao hop le la day cac so nguyen
package main

import (
	"fmt"
	"os"
	"strconv"
)

func BubbleSort(slice []int) []int {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	return slice
}

func main() {
	toGetAllArgs := os.Args[1:]
	//fmt.Println(toGetAllArgs)
	var slice = []int{}

	for _, i := range toGetAllArgs {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		slice = append(slice, j)
	}
	//fmt.Println(slice)
	// soluong := 0
	// fmt.Print("Nhap so luong phan tu truyen vao: ")
	// fmt.Scanf("%d", &soluong)
	// slice := []int{}

	// var phantu int
	// fmt.Scanf("%d", &phantu)
	minnum := slice[0]
	maxnum := slice[0]
	mean := float64(slice[0])
	median := float64(slice[0])
	// slice = append(slice, phantu)

	for i := 1; i < len(toGetAllArgs); i++ {
		if minnum > slice[i] {
			minnum = slice[i]
		}
		if maxnum < slice[i] {
			maxnum = slice[i]
		}
		mean = float64(mean) + float64(slice[i])

	}

	mean = mean / float64(len(toGetAllArgs))

	BubbleSort(slice) //sap xep lai mang tu be den lon bang giai thuat noi bot
	//fmt.Println(slice)
	if len(slice)%2 == 0 {

		median = float64(slice[len(slice)/2]+slice[len(slice)/2-1]) / 2.0

	} else {
		median = float64(slice[len(slice)/2])
	}

	fmt.Println("Gia tri lon nhat: ", maxnum)
	fmt.Println("Gia tri nho nhat: ", minnum)
	fmt.Println("Gia tri trung binh: ", mean)
	fmt.Println("Gia tri trung vi: ", median)
}
