package main

import (
	"fmt"
)

func greet(customerName string) {
	fmt.Printf("\nHalo, %s!\n", customerName)
	fmt.Println("=== Selamat Datang di Rumah Makan Nusantara! ===")
	fmt.Println()
}

func sayTata(customerName string) {
	fmt.Println()
	fmt.Printf("=== Terima kasih %v telah berkunjung ke Rumah Makan Nusantara! ===\n", customerName)
	fmt.Println("Kami berharap dapat melayani Anda kembali!")
	fmt.Println("Semoga hari Anda menyenangkan!")
}
