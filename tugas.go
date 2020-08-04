package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	fmt.Println("Welcome to Go calculator!")
	perintah := readLine("masukan perintah: \n 1. tambah \n 2. kurang \n 3. bagi \n 4. kali \n 5. akar \n 6. pangkat \n 8. luas persegi \n 9. luas lingkaran \n 10. volume tabung \n 11. volume balok \n 12. volume prisma")
	fmt.Println(perintah)

	if perintah == "1" {
		num1, num2 := getNumbers()
		result := num1 + num2
		fmt.Println("hasil penambahan = ", result)
	} else if perintah == "2" {
		num1, num2 := getNumbers()
		result := num1 - num2
		fmt.Println("hasil pengurangan = ", result)
	} else if perintah == "3" {
		num1, num2 := getNumbers()
		result := num1 / num2
		fmt.Println("hasil pembagian = ", result)
	} else if perintah == "4" {
		num1, num2 := getNumbers()
		result := num1 * num2
		fmt.Println("hasil perkalian =", result)
	} else if perintah == "5" {
		fmt.Println("\n==== akar ====")
		num1, num2 := getNumbers()

		helper := helper{float64(num1), float64(num2)}
		helper.Akar()
	} else if perintah == "6" {
		fmt.Println("\n==== pangkat ====")
		num1, num2 := getNumbers()

		helper := helper{float64(num1), float64(num2)}
		helper.Kuadrat()
	} else if perintah == "7" {
		fmt.Println("opps.. its a trap !!")
	} else if perintah == "8" {
		Luas_Persegi()
	} else if perintah == "9" {
		Luas_Lingkaran()
	} else if perintah == "10" {
		Volume_Tabung()
	} else if perintah == "11" {
		Volume_Balok()
	} else if perintah == "12" {
		Volume_Prisma()
	} else {
		fmt.Println("Invalid input")
	}

	readLine("press 'anything' to exit")
}

// fungsi untuk membaca inputan
func readLine(teks string) string {
	fmt.Println(teks)
	var input string
	fmt.Scanln(&input)
	return input
}

// fungsi untuk mengconvert string ke
func getNumbers() (int, int) {
	StringNum1 := readLine("angka pertama: ")
	num1, _ := strconv.Atoi(StringNum1)
	StringNum2 := readLine("angka kedua: ")
	num2, _ := strconv.Atoi(StringNum2)
	return num1, num2
}

func Luas_Persegi() {
	var sisi float64

	fmt.Printf("Masukkan Panjang Sisi = ")
	fmt.Scanf("%f", &sisi)
	fmt.Println("Hasil =", sisi*sisi)
}

func Luas_Lingkaran() {
	var jari float64

	fmt.Printf("Masukkan Panjang jari = ")
	fmt.Scanf("%f", &jari)
	fmt.Println("Hasil =", 3.14*jari*jari)
}

func Volume_Tabung() {
	var jari float64
	var tinggi float64

	fmt.Printf("Masukkan Panjang jari = ")
	fmt.Scanf("%f", &jari)
	fmt.Printf("Masukkan Tinggi = ")

	fmt.Scanf("%f", &tinggi)
	fmt.Println("Hasil =", 3.14*jari*jari*tinggi)
}

func Volume_Balok() {
	var panjang float64
	var lebar float64
	var tinggi float64

	fmt.Printf("Masukkan Panjang = ")
	fmt.Scanf("%f", &panjang)
	fmt.Printf("Masukkan Lebar = ")
	fmt.Scanf("%f", &lebar)
	fmt.Printf("Masukkan Tinggi = ")
	fmt.Scanf("%f", &tinggi)
	fmt.Println("Hasil =", panjang*lebar*tinggi)
}

func Volume_Prisma() {
	var a float64
	var b float64
	var tinggi float64

	fmt.Printf("Masukkan sisi a = ")
	fmt.Scanf("%f", &a)
	fmt.Printf("Masukkan sisi b = ")
	fmt.Scanf("%f", &b)
	fmt.Printf("Masukkan Tinggi = ")
	fmt.Scanf("%f", &tinggi)
	fmt.Println("Hasil =", 0.5*a*b*tinggi)
}

type helper struct {
	Angka1 float64
	Angka2 float64
}

func (h helper) Akar() {
	fmt.Println("Akar angka 1:", math.Sqrt(h.Angka1))

}

func (h helper) Kuadrat() {
	fmt.Println("Kuadrat angka 1:", h.Angka1*h.Angka1)

}
