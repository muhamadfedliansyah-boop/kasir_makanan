package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func displayGeneratingBill() {
	fmt.Println()
	billDisplayText := "************************************* Memproses Tagihan *************************************"
	for _, element := range billDisplayText {
		fmt.Printf("%c", element)
		time.Sleep(time.Millisecond * 10)
	}
	fmt.Println()
}

func generateBill() {
	fmt.Println()
	fmt.Printf("+%s+\n", strings.Repeat("-", 76))
	fmt.Printf(" %-26s %-16s %-12s %-16s\n", "Nama Menu", "Harga Satuan", "Jumlah", "Total")
	fmt.Printf("+%s+\n", strings.Repeat("-", 76))

	printOrderData()

	fmt.Printf("+%s+\n", strings.Repeat("-", 76))

	fmt.Printf("%42s", " ")
	for _, element := range "Sub Total (tanpa PPN): Rp" {
		fmt.Printf("%c", element)
		time.Sleep(time.Millisecond * 20)
	}
	fmt.Printf("%.0f\n", subTotalBill)
}

func printOrderData() {
	for key := range customerOrder {
		for _, element := range menu {
			if key == element.itemName {
				totalCostOfItem := float64(customerOrder[key]) * element.itemPrice
				fmt.Printf(" %-26s Rp%-14.0f %-12v Rp%-14.0f\n", key, element.itemPrice, customerOrder[key], totalCostOfItem)
			}
		}
	}
}

func printFinalBill() {
	for _, element := range "Berikut adalah rincian tagihan akhir Anda:-" {
		fmt.Printf("%c", element)
		time.Sleep(time.Millisecond * 25)
	}
	fmt.Println()

	fmt.Printf("\n%48s\n", "RUMAH MAKAN NUSANTARA")
	time.Sleep(time.Millisecond * 150)
	fmt.Printf("%s\n", strings.Repeat("*", 78))
	time.Sleep(time.Millisecond * 150)
	fmt.Printf("%62s\n", "Jl. Merdeka No. 45, Jakarta Pusat, Indonesia")
	time.Sleep(time.Millisecond * 150)
	fmt.Printf("%48s\n", "Telp: 0812-3456-7890")
	fmt.Printf("%54s\n\n", "Email: rmnusantara@example.com")
	time.Sleep(time.Millisecond * 150)
	fmt.Printf("%s", strings.Repeat("-", 30))
	fmt.Printf("%s", " NOTA PEMBAYARAN ")
	fmt.Printf("%s\n", strings.Repeat("-", 30))
	time.Sleep(time.Millisecond * 150)

	rand.Seed(time.Now().UnixNano())
	fmt.Printf(" No. Antrean: %d\n", rand.Intn(550)+1)
	fmt.Printf(" Tanggal    : %v\n", time.Now().Local().Format("02-Jan-2006"))
	fmt.Printf(" Waktu      : %v\n", time.Now().Local().Format("15:04 WIB"))
	time.Sleep(time.Millisecond * 150)

	generateBill()

	tax := 11 * subTotalBill / 100 // PPN 11%
	grandTotal := subTotalBill + tax

	time.Sleep(time.Millisecond * 150)
	fmt.Printf("%58s: Rp%.0f\n", "PPN (11%)", tax)
	fmt.Printf("+%s+\n", strings.Repeat("-", 76))
	time.Sleep(time.Millisecond * 150)
	fmt.Printf("%58s: Rp%.0f\n", "Total Pembayaran", grandTotal)
	fmt.Printf("+%s+\n", strings.Repeat("-", 76))
}
