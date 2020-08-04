package main

import (
	"fmt"
	"time"
)

func main() {

	// Inisiasi dari tahun bulan dan tanggal
	tahun := 2020
	bulan := []string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}
	hari := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	var haripertama int
	hariawalbulan := [12]int{}
	valKosong := "-- "

	// pada tahun kabisat
	if (tahun%4 == 0 && tahun%100 != 0) || tahun%400 == 0 {
		hari[1] = 29
	}

	// untuk mendapatkan hari pertama setiap bulan
	for i := 0; i < 12; i++ {
		hari1 := time.Date(tahun, time.Month(i+1), 1, 1, 1, 1, 1, time.UTC)

		switch hari1.Weekday() {
		case 1:
			haripertama = 1
		case 2:
			haripertama = 2
		case 3:
			haripertama = 3
		case 4:
			haripertama = 4
		case 5:
			haripertama = 5
		case 6:
			haripertama = 6
		case 7:
			haripertama = 7

		}
		hariawalbulan[i] = haripertama
	}

	fmt.Println("Welcome to KALENDER", tahun)
	for m := 0; m < 12; m++ {

		// Print awal pada bulan
		fmt.Print("======" + bulan[m])
		spasi := 14 - len(bulan[m])
		for aa := 0; aa < spasi; aa++ {
			fmt.Print("=")
		}
		fmt.Println("")
		fmt.Println("S  S  R  K  J  S  M ")

		// 	mencari start di hari apa
		if hariawalbulan[m] > 0 {
			for ii := 1; ii < hariawalbulan[m]; ii++ {
				fmt.Print(valKosong)
			}
		}

		// melooping tanggal
		for a, b := 1, hariawalbulan[m]; a <= hari[m]; a, b = a+1, b+1 {
			if a < 10 {
				fmt.Print(a, "  ")
			} else {
				fmt.Print(a, " ")
			}

			// untuk ganti baris jika sudah hari ke8
			if b%7 == 0 {
				fmt.Print("\n")
			}
		}
		// mencari sisa hari
		if hariawalbulan[m]+hari[m] <= 35 {
			sisa := 35 - (hariawalbulan[m] + hari[m])
			for j := 0; j <= sisa; j++ {
				fmt.Print(valKosong)
			}
		} else if hariawalbulan[m]+hari[m] > 36 {
			sisa := 42 - (hariawalbulan[m] + hari[m])
			for j := 0; j <= sisa; j++ {
				fmt.Print(valKosong)
			}

		}
		fmt.Println("")
	}
}
