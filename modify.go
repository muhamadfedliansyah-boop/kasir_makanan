package main

import "fmt"

func modifyOrder() {
	for {
		var isOrderOK string
		fmt.Print("\nApakah Anda ingin mengubah pesanan? [y/n]: ")
		fmt.Scan(&isOrderOK)
		if isOrderOK != "y" && isOrderOK != "Y" {
			return
		}
		var serialNo uint
		var modifyType uint

		fmt.Println("\nPilih tindakan:")
		fmt.Println("1. Ubah jumlah porsi menu")
		fmt.Println("2. Hapus menu dari pesanan")
		fmt.Println("3. Tambah menu baru ke pesanan")
		fmt.Println("Pilihan lainnya untuk selesai.")
		fmt.Print("Masukkan pilihan [1/2/3]: ")
		fmt.Scan(&modifyType)

		switch modifyType {
		case 1:
			printMenu()
			fmt.Print("Masukkan nomor menu yang ingin diubah: ")
			fmt.Scan(&serialNo)
			updateQuantity(serialNo)

		case 2:
			printMenu()
			fmt.Print("Masukkan nomor menu yang ingin dihapus: ")
			fmt.Scan(&serialNo)
			delFromOrder(serialNo)

		case 3:
			insertIntoOrder()
		default:
			return
		}
		displayGeneratingBill()
		generateBill()
	}
}

func updateQuantity(serialNo uint) {
	var newQuantity uint

	for _, targetItem := range menu {
		if serialNo == targetItem.itemNo {
			itemName := targetItem.itemName
			oldQuantity := customerOrder[itemName]

			fmt.Printf("Jumlah pesanan %v saat ini: %v\n", itemName, oldQuantity)
			fmt.Printf("Masukkan jumlah porsi baru untuk %v: ", itemName)
			fmt.Scan(&newQuantity)

			if newQuantity == 0 {
				delFromOrder(serialNo)
				return
			}

			customerOrder[targetItem.itemName] = newQuantity
			fmt.Printf("Jumlah pesanan %v berhasil diubah dari %v menjadi %v.\n", itemName, oldQuantity, newQuantity)

			// Update total tagihan
			subTotalBill -= float64(oldQuantity) * float64(targetItem.itemPrice)
			subTotalBill += float64(newQuantity) * float64(targetItem.itemPrice)
			break
		}
	}
}

func delFromOrder(serialNo uint) {
	for _, targetItem := range menu {
		if serialNo == targetItem.itemNo {
			itemName := targetItem.itemName
			subTotalBill -= float64(customerOrder[itemName]) * float64(targetItem.itemPrice)
			delete(customerOrder, itemName)
			fmt.Printf("%v berhasil dihapus dari daftar pesanan.\n", itemName)
			break
		}
	}
}

func insertIntoOrder() {
	orderItems()
}
