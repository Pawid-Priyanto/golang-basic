package main

import (
	"fmt"
	"time"

	// for generate id
	"github.com/rs/xid"
)

func generateIdParkir() (string, time.Time, string, string) {
	id := xid.New().String()
	time := time.Now()
	var (
		tipe, plat string
	)
	fmt.Print("Tipe : ")
	fmt.Scanf("%s", &tipe)
	fmt.Print("Plat Nomor : ")
	fmt.Scanf("%s", &plat)

	return id, time, tipe, plat
}

type Karcis struct {
	id   string
	plat string
	time time.Time
	tipe string
}

// fungsi untuk membaca inputan
func readLine(teks string) string {
	fmt.Println(teks)
	var input string
	fmt.Scanln(&input)
	return input
}

func ParkirMasuk() {

	var kendaraan []Karcis
	id, time, tipe, plat := generateIdParkir()
	data := Karcis{id, plat, time, tipe}

	kendaraan = append(kendaraan, data)
	for i := 0; i < len(kendaraan); i++ {
		fmt.Print("========================== \n")
		fmt.Println("ID Parkir :", kendaraan[i].id)
		fmt.Println("Lama Parkir :", kendaraan[i].time.Second(), "Detik")
		fmt.Println("Nomor Plat Parkir :", kendaraan[i].plat)
		fmt.Println("Jenis Kendaraan :", kendaraan[i].tipe)
		fmt.Print("========================== \n")
	}
}

func ParkirKeluar() {
	var kendaraan []Karcis
	var (
		idBaru        string
		platNomor     string
		tipeKendaraan string
	)
	now := time.Now()
	fmt.Print("========================== \n")
	fmt.Print("Tipe Kendaraan : ")
	fmt.Scan(&tipeKendaraan)
	fmt.Print("Plat Nomor : ")
	fmt.Scan(&platNomor)
	fmt.Print("ID Parkir : ")
	fmt.Scan(&idBaru)
	fmt.Print("========================== \n")

	for i := 0; i < len(kendaraan); i++ {
		if idBaru == kendaraan[i].id && tipeKendaraan == kendaraan[i].tipe && platNomor == kendaraan[i].plat {
			waktu := now.Sub(kendaraan[i].time).Seconds()
			fmt.Println(waktu)
			if kendaraan[i].tipe == "Motor" {
				if waktu > 1 {
					fmt.Print("Bayar parkir sebanyak : ")
					fmt.Println("Rp.", int((waktu-1)*3000+5000))
				} else {
					fmt.Print("Bayar parkir sebanyak")
					fmt.Println("Rp.", 5000)
				}
			} else if kendaraan[i].tipe == "Mobil" {
				if waktu > 1 {
					fmt.Print("Bayar parkir sebanyak : ")
					fmt.Println("Rp.", int((waktu-1)*2000+3000))
				} else {
					fmt.Print("Bayar parkir sebanyak")
					fmt.Println("Rp.", 3000)
				}
			}
			kendaraan = append(kendaraan[:i], kendaraan[i+1:]...)
		}

	}

}

func main() {

	menu := 0
	for menu != 99 {
		fmt.Println("Welcome to Aplikasi Parkir")
		pilihan := readLine(" \n 1. Parkir Masuk \n 2. Parkir Keluar \n 99. Keluar")
		fmt.Println("Pilhan menu: ")
		fmt.Println(pilihan)
		if pilihan == "1" {
			ParkirMasuk()
		} else if pilihan == "2" {
			ParkirKeluar()
		} else {
			fmt.Println("wrong type")
		}
	}

}
