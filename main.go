package main

import (
	"fmt"
)

// subTotalBill adalah total tagihan sebelum pajak.
var subTotalBill float64

// customerOrder menyimpan nama menu sebagai "key" dan jumlah porsi sebagai "value".
var customerOrder = make(map[string]uint, 0)

func main() {
	var customerName string
	fmt.Print("Masukkan nama Anda: ")
	fmt.Scan(&customerName)

	greet(customerName)
	orderItems()
	displayGeneratingBill()
	generateBill()
	modifyOrder()
	printFinalBill()
	sayTata(customerName)
}
