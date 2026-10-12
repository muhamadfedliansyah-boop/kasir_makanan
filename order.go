package main

import (
	"fmt"
	"strings"
)

// Fungsi untuk memesan menu
func orderItems() {
	printMenu()
	var itemNumber uint
	var noOfPlates uint

	for {
		fmt.Println()
		fmt.Println("Ketik '0' jika sudah selesai memesan.")
		fmt.Print("Masukkan nomor menu yang ingin dipesan: ")

		fmt.Scan(&itemNumber)
		if itemNumber == 0 {
			break
		}

		var choiceName string
		var itemPrice float64

		for index, item := range menu {
			if index+1 == int(itemNumber) {
				choiceName = item.itemName
				itemPrice = item.itemPrice
				break
			}
		}

		if choiceName == "" {
			fmt.Println("Nomor menu tidak valid, silakan coba lagi.")
			continue
		}

		fmt.Printf("Berapa porsi %v yang ingin dipesan: ", choiceName)
		fmt.Scan(&noOfPlates)
		if noOfPlates == 0 {
			continue
		} else {
			for index := range menu {
				if index+1 == int(itemNumber) {
					customerOrder[choiceName] += noOfPlates
					subTotalBill += itemPrice * float64(noOfPlates)
					break
				}
			}
			fmt.Printf("\nAnda memesan %v porsi %v (Total: Rp%.0f).\n", noOfPlates, choiceName, itemPrice*float64(noOfPlates))
			orderTillNow()
		}
		fmt.Println()
	}
}

// Menampilkan daftar pesanan yang sudah dipilih
func orderTillNow() {
	fmt.Println("\nPesanan Anda saat ini:")
	fmt.Printf("%s\n", strings.Repeat("-", 35))
	fmt.Printf(" %-10s %s\n", "Jumlah", "Menu")
	fmt.Printf("%s\n", strings.Repeat("-", 35))

	for i := range customerOrder {
		fmt.Printf(" %-10v %v\n", customerOrder[i], i)
	}

	fmt.Printf("%s\n", strings.Repeat("-", 35))
}
