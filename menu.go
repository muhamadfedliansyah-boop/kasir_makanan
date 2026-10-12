package main

import (
	"fmt"
	"strings"
)

type Menu struct {
	itemNo    uint
	itemName  string
	itemPrice float64
}

var menu = []Menu{
	//itemNo  itemName   itemPrice
	{1, "Teh Tarik", 12000},
	{2, "Kopi Tubruk", 10000},
	{3, "Es Teh Manis", 5000},
	{4, "Es Jeruk", 7000},
	{5, "Jus Alpukat", 15000},
	{6, "Nasi Goreng Spesial", 25000},
	{7, "Ayam Bakar Madu", 30000},
	{8, "Mie Goreng Jawa", 22000},
	{9, "Soto Ayam Lamongan", 20000},
	{10, "Roti Bakar Cokelat", 15000},
	{11, "Indomie Kuah Telur", 13000},
	{12, "Bakwan Jagung (3 pcs)", 10000},
	{13, "Pisang Keju Crispy", 14000},
}

func printMenu() {
	fmt.Println()
	fmt.Printf("%25s\n", "DAFTAR MENU")
	fmt.Printf("%s\n", strings.Repeat("-", 45))
	fmt.Printf("%-6s %-24s %-12s\n", "No.", "Nama Menu", "Harga")
	fmt.Printf("%s\n", strings.Repeat("-", 45))

	for _, element := range menu {
		fmt.Printf(" %-5d %-24s Rp%-10.0f\n", element.itemNo, element.itemName, element.itemPrice)
	}
	fmt.Printf("%s\n", strings.Repeat("-", 45))
}
