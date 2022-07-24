package main

import (
	//"crypto/rand"

	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//Hien thuc ham sinh chuoi string ngau nhien
var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-"

func RandString(length int) string {
	ll := len(chars)
	b := make([]byte, length)
	rand.Read(b) // generates len(b) random bytes
	for i := 0; i < length; i++ {
		b[i] = chars[int(b[i])%ll]
	}
	return string(b)
}

//Hien thuc ham sinh so nguyen ngau nhien
func RandInt() int {
	min := 1
	max := 1000000
	rand.Seed(time.Now().UnixNano())
	a := (rand.Intn(max-min) + min)
	return a
}

//Hien thuc ham sinh so thuc ngau nhien
func RandFloat() float64 {
	rand.Seed(time.Now().UnixNano())
	a := ((rand.Float64() * 100) + 100)
	return a
}

func RandTime() time.Time {
	min := time.Date(1770, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

//Chuong trinh chinh
func main() {

	// fmt.Print("Cac kieu du lieu ho tro: string,int,float,time. ")
	// fmt.Println("Vi du chuoi nhap: 3 ten1 string ten2 int ten3 float ten4 time")
	a := os.Args[1:]

	// consoleReader := bufio.NewReader(os.Stdin) //nhap chuoi nhan space
	// input, _ := consoleReader.ReadString('\n')
	// a := strings.Fields(input) // chuyen chuoi vao mang
	//fmt.Println(a)
	sodong, _ := strconv.Atoi(a[0])

	for i := 0; i < sodong; i++ { //in ra bang
		for j := 1; j < len(a); j++ {
			if j%2 == 0 {
				if a[j] == "string" {
					fmt.Print("\t", RandString(10), "\t")
				}
				if a[j] == "int" {
					fmt.Printf("\t%d\t", RandInt())
				}
				if a[j] == "float" {
					fmt.Printf("\t%0.5f\t", RandFloat())
				}
				if a[j] == "time" {
					fmt.Print("\t", RandTime(), "\t")
				}
			}
		}
		fmt.Println()
	}
}
