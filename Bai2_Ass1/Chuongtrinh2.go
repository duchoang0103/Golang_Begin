package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func docso(n int) string {
	s := []string{}
	so := [10]string{"", "mot", "hai", "ba", "bon", "nam", "sau", "bay", "tam", "chin"}
	var donvi, chuc, tram int
	donvi = n % 10
	n /= 10
	chuc = n % 10
	tram = n / 10
	if tram > 0 {
		s = append(s, so[tram])
		s = append(s, " tram ")
	}
	if chuc > 0 {
		if chuc == 1 {
			s = append(s, " muoi ")
		} else {
			s = append(s, so[chuc])
			s = append(s, " muoi ")
		}
	}
	if donvi > 0 {
		if chuc == 0 && tram != 0 {
			s = append(s, "le ")
		}
		if donvi == 1 {
			s = append(s, "mot")
		} else if donvi == 5 && (chuc != 0 || tram != 0) {
			s = append(s, "lam")
		} else if donvi == 5 && (chuc == 0 || tram != 0) {
			s = append(s, "nam ")
		} else {
			s = append(s, so[donvi])
		}
	}
	js := strings.Join(s, "")
	//fmt.Println(js)
	return js
}
func main() {
	var n int
	var ngan, trieu, ty, donvi int
	fmt.Print("\nNhap so can doc: ")
	toGetAllArgs := os.Args[1:]
	//fmt.Scanf("%d", &n)
	n, _ = strconv.Atoi(toGetAllArgs[0])

	fmt.Print("So doc bang chu: ")
	if n == 0 {
		fmt.Printf("khong")
	} else {
		donvi = n % 1000
		n /= 1000
		ngan = n % 1000
		n /= 1000
		trieu = n % 1000
		ty = n / 1000

		if ty > 0 {
			fmt.Printf("%s ty ", docso(ty))
		}
		if trieu > 0 {
			fmt.Printf("%s trieu ", docso(trieu))
		}
		if ngan > 0 {
			fmt.Printf("%s ngan ", docso(ngan))
		}
		if donvi > 0 {
			fmt.Printf("%s", docso(donvi))
		}
	}
}
