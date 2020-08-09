package main

import (
	"fmt"
	"time"

	// for generate id
	"github.com/rs/xid"
)

type TiketMasuk struct {
	id   string
	time time.Time
}

func generateIdParkir() (string, time.Time) {
	id := xid.New().String()
	time := time.Now()

	return id, time
}

func main() {
	var kendaraan []TiketMasuk
	input := 0

	for input != 3 {
		fmt.Println("Welcome to Aplikasi Parkir")
		fmt.Println("1. Parkir Masuk ")
		fmt.Println("2. Parkir Keluar")
		fmt.Println("3. Keluar")
		fmt.Println("Masukkan Angka pilihan: ")

		fmt.Scan(&input)

		switch input {
		case 1:

			id, time := generateIdParkir()
			data := TiketMasuk{id, time}

			kendaraan = append(kendaraan, data)
			for i := 0; i < len(kendaraan); i++ {
				fmt.Print("========================== \n")
				fmt.Println("ID Parkir :", kendaraan[i].id)
				fmt.Println("Waktu Masuk :", kendaraan[i].time)
				fmt.Print("========================== \n")
			}
			break
		case 2:
			var (
				idBaru        string
				platNomor     string
				tipeKendaraan string
			)
			now := time.Now()
			fmt.Print("========================== \n")
			fmt.Print("Tipe Kendaraan (HURUF CAPITAL): ")
			fmt.Scan(&tipeKendaraan)
			fmt.Print("Plat Nomor : ")
			fmt.Scan(&platNomor)
			fmt.Print("ID Parkir : ")
			fmt.Scan(&idBaru)

			for i := 0; i < len(kendaraan); i++ {
				if idBaru == kendaraan[i].id {
					waktu := int(now.Sub(kendaraan[i].time).Seconds())
					fmt.Println("Lama parkir anda :", waktu, "Detik")
					if tipeKendaraan == "MOBIL" {
						if waktu > 1 {
							fmt.Print("Total Biaya Parkir : ")
							fmt.Println("Rp.", int((waktu-1)*3000+5000))
							fmt.Print("========================== \n")
						} else {
							fmt.Print("Total Biaya Parkir")
							fmt.Println("Rp.", 5000)
							fmt.Print("========================== \n")
						}
					} else if tipeKendaraan == "MOTOR" {
						if waktu > 1 {
							fmt.Print("Total Biaya Parkir : ")
							fmt.Println("Rp.", int((waktu-1)*2000+3000))
							fmt.Print("========================== \n")
						} else {
							fmt.Print("Total Biaya Parkir : ")
							fmt.Println("Rp.", 3000)
							fmt.Print("========================== \n")
						}
					} else {
						fmt.Println("Kendaraan yg anda masukan salah")
					}
				} else {
					fmt.Println("Id yg anda masukan salah")
				}
			}

			break
		case 3:
			fmt.Println("Terimakasih")
			break
		default:
			fmt.Println("Maaf salah inputan")
			break
		}

	}

}
