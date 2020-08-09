package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type TiketMasuk2 struct {
	id   string
	time time.Time
}

type TiketMasuk3 struct {
	id    string
	plat  string
	tipe  string
	hasil int
}

var baseURL = "http://localhost:8080"

func getKendaraan() (TiketMasuk2, error) {
	var err error
	var client = &http.Client{}
	var data TiketMasuk2

	request, err := http.NewRequest("POST", baseURL+"/parkir-masuk", nil)

	if err != nil {
		return data, err
	}

	response, err := client.Do(request)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)

	if err != nil {
		return data, err
	}

	return data, nil
}

func bayarParkir(ID string, tipe string, plat string) (TiketMasuk3, error) {
	var err error
	var client = &http.Client{}
	var dataParkir TiketMasuk3

	var param = url.Values{}
	param.Set("id", ID)
	param.Set("tipe", tipe)
	param.Set("plat", plat)
	var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", baseURL+"/parkir-keluar", payload)

	if err != nil {
		return dataParkir, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		return dataParkir, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&dataParkir)

	if err != nil {
		return dataParkir, err
	}
	return dataParkir, nil

}

func main() {

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

			result, err := getKendaraan()
			if err != nil {
				fmt.Print("========================== \n")
				fmt.Println("Error \n", err.Error())
			}
			fmt.Print("========================== \n")
			fmt.Println("ID Parkir :", result.id)
			fmt.Print("========================== \n")

			break
		case 2:
			var (
				idBaru        string
				platNomor     string
				tipeKendaraan string
			)

			fmt.Print("========================== \n")
			fmt.Print("Tipe Kendaraan : ")
			fmt.Scan(&tipeKendaraan)
			fmt.Print("Plat Nomor : ")
			fmt.Scan(&platNomor)
			fmt.Print("ID Parkir : ")
			fmt.Scan(&idBaru)
			result1, err := bayarParkir(tipeKendaraan, platNomor, idBaru)
			if err != nil {
				fmt.Print("========================== \n")
				fmt.Println("Error \n", err.Error())
				return
			}
			fmt.Print("========================== \n")
			fmt.Println("Total biaya parkir :", result1.hasil)
			fmt.Print("========================== \n")

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
