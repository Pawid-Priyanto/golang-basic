package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	// for generate id
	"github.com/rs/xid"
)

type TiketMasuk struct {
	Id   string
	Time time.Time
}

type TiketMasuk1 struct {
	id    string
	plat  string
	tipe  string
	hasil int
}

type harga struct {
	hrg int
}

func generateIdParkir() (string, time.Time) {
	id := xid.New().String()
	time := time.Now()

	return id, time
}

func generateIdParkir1(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application-json")

	id, time := generateIdParkir()
	data := TiketMasuk{id, time}

	kendaraan = append(kendaraan, data)

	if r.Method == "POST" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func bayar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		now := time.Now()
		var id = r.FormValue("id")
		var tipeKendaraan = r.FormValue("tipe")
		var platNomor = r.FormValue("plat")
		var result []byte
		var guest TiketMasuk1
		var err error

		for _, each := range kendaraan {
			if each.Id == id {
				waktu := int(now.Sub(each.Time).Seconds())
				fmt.Println("Kamu parkr selama:", waktu, "detik")

				if tipeKendaraan == "Mobil" {
					if waktu > 1 {
						jumlahbayar := (waktu-1)*3000 + 5000
						fmt.Println("ID Parkir: ", id)
						fmt.Println("tipe kendaraan: ", tipeKendaraan)
						fmt.Println("Plat nomor: ", platNomor)
						fmt.Println("Total Biaya: ", jumlahbayar)
						guest = TiketMasuk1{id, tipeKendaraan, platNomor, jumlahbayar}
						result, err = json.Marshal(guest)
						w.Write(result)
						return
					} else {
						jumlahbayar := 5000
						fmt.Println("ID Parkir: ", id)
						fmt.Println("tipe kendaraan: ", tipeKendaraan)
						fmt.Println("Plat nomor: ", platNomor)
						fmt.Println("Total Biaya: ", jumlahbayar)
						guest = TiketMasuk1{id, tipeKendaraan, platNomor, jumlahbayar}
						result, err = json.Marshal(guest)
						w.Write(result)
						return
					}
				} else if tipeKendaraan == "Motor" {
					if waktu > 1 {
						jumlahbayar := (waktu-1)*3000 + 3000
						fmt.Println("ID Parkir: ", id)
						fmt.Println("tipe kendaraan: ", tipeKendaraan)
						fmt.Println("Plat nomor: ", platNomor)
						fmt.Println("Total Biaya: ", jumlahbayar)
						guest = TiketMasuk1{id, tipeKendaraan, platNomor, jumlahbayar}
						result, err = json.Marshal(guest)
						w.Write(result)
						return
					} else {
						jumlahbayar := 3000
						fmt.Println("ID Parkir: ", id)
						fmt.Println("tipe kendaraan: ", tipeKendaraan)
						fmt.Println("Plat nomor: ", platNomor)
						fmt.Println("Total Biaya: ", jumlahbayar)
						guest = TiketMasuk1{id, tipeKendaraan, platNomor, jumlahbayar}
						result, err = json.Marshal(guest)
						w.Write(result)
						return
					}
				}
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				return
			}
		}
		http.Error(w, "Belum ada kendaraan", http.StatusBadRequest)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

var kendaraan = []TiketMasuk{}

func main() {

	http.HandleFunc("/parkir-masuk", generateIdParkir1)
	http.HandleFunc("/parkir-keluar", bayar)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)

}
