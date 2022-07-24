package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	arr := os.Args[1:]

	if len(arr) == 3 { //truong hop co ca hai tham so
		t3, _ := time.Parse("2006-01-02", arr[1])
		t4, _ := time.Parse("2006-01-02", arr[2])
		diff := t4.Sub(t3)
		fmt.Println("So ngay tu", arr[1], "den", arr[2], "la: ", int(diff.Hours()/24))
	} else { //truong hop khong co tham so thu 2
		t3, _ := time.Parse("2006-01-02", arr[1])
		date := time.Now()
		diff := date.Sub(t3)
		fmt.Println("So ngay tu", arr[1], "den ngay hom nay la:", int(diff.Hours()/24))
	}

}
