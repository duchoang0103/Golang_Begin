package main

import (
	"fmt"
	"os"
)

func sum(n1 string, n2 string) string {
	var finalResult = ""
	var maxLen int

	if len(n1) > len(n2) { //kiem tra xem chuoi nao co do dai lon hon
		maxLen = len(n1)
	} else {
		maxLen = len(n2)
	}

	var index1 int // chỉ số của kí tự đang xét của chuỗi 1
	var index2 int // chỉ số của kí tự đang xét của chuỗi 2
	var c1 byte    // kí tự tại vị trí index1 của chuỗi s1
	var c2 byte    // kí tự tại vị trí index2 của chuỗi s2
	var d1 uint8   // kí số của c1
	var d2 uint8   // kí số của c2
	var t uint8    // tổng tạm của 2 kí số
	var mem uint8  // giá trị nhớ
	mem = 0
	for i := 0; i < maxLen; i++ {
		index1 = len(n1) - i - 1
		index2 = len(n2) - i - 1
		if index1 < 0 { //kiem tra index1 co be hon 0 neu be hon thi cho bang 0
			c1 = '0'
		} else {
			c1 = n1[index1]
		}
		if index2 < 0 { //kiem tra index2 co be hon 0 neu be hon thi cho bang 0
			c2 = '0'
		} else {
			c2 = n2[index2]
		}

		d1 = c1 - '0'
		d2 = c2 - '0'
		print("Lấy ", d1, " cộng ", d2, " => ")
		t = d1 + d2

		if mem < 1 {
			if t > 9 {
				print("Kết quả được ", t, ", lưu ", t%10, " nhớ ", t/10)
				mem = t / 10
				// Lấy hàng đơn vị của t ghép vào phía trước kết quả
				finalResult = fmt.Sprint(t%10, finalResult)
			} else {
				print("Kết quả được ", t)
				// Lấy hàng đơn vị của t ghép vào phía trước kết quả
				finalResult = fmt.Sprint(t%10, finalResult)
			}
		} else {
			if t > 8 {
				print("Kết quả được ", t, ", cộng với 1 được ", t+1, " lưu ", (t+1)%10, " nhớ ", (t+1)/10)
				// Lấy hàng đơn vị của t ghép vào phía trước kết quả
				mem = (t + 1) / 10
				finalResult = fmt.Sprint((t+1)%10, finalResult)
			} else {
				print("Kết quả được ", t, ", cộng với 1 được ", t+1)
				// Lấy hàng đơn vị của t ghép vào phía trước kết quả
				finalResult = fmt.Sprint((t+1)%10, finalResult)
			}
		}

		if mem > 0 && i == maxLen-1 {
			print("\n", "Hạ 1 xuống")

		}
		println()
	}

	if mem > 0 {
		finalResult = fmt.Sprint(mem, finalResult)
	}

	return finalResult
}

func main() {
	n := os.Args[1:]
	println("Kết qủa hiển thị:")
	println("Các bước thực hiện:")
	total := sum(n[1], n[2])
	println("Kết quả cuối cùng:", total)
}
