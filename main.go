package main

import "fmt"

func main() {

	type Makanan struct {
		id    int
		nama  string
		harga int
	}

	total := 0
	
	var makanan = []Makanan{
		{id: 1, nama: "Keripik Kentang", harga: 3000},
		{id: 2, nama: "Pisang Goreng", harga: 1000},
		{id: 3, nama: "Biskuat", harga: 500},
		{id: 4, nama: "Beng-Beng", harga: 2000},
		{id: 5, nama: "Indomie", harga: 3000},
	}

	for {
		fmt.Print("Halo, selamat datang. Silahkan pilih barang yang ingin anda beli \n")

		for _, food := range makanan {
			fmt.Println(food.id, ".", food.nama, "-", food.harga)
		}

		var id_Makanan int
		var quantity int

		fmt.Print("Pilih nomor makanan yang akan anda dibeli : ")
		fmt.Scan(&id_Makanan)

		if id_Makanan > len(makanan) {
			fmt.Println("Nomor makanan yang anda dipilih kosong !")
		} else {
			fmt.Println("Kamu akan membeli ", makanan[id_Makanan-1].nama, "dengan harga ", makanan[id_Makanan-1].harga)
			fmt.Println("Masukan jumlah makanan yang akan dibeli : ")
			fmt.Scan(&quantity)
			subtotal := makanan[id_Makanan-1].harga * quantity
			fmt.Println("Pembelian sukses, subtotal : ", subtotal)
			total += subtotal

			var option string
			fmt.Print("Mau beli lagi ? (Y/N) : ")
			fmt.Scan(&option)
			if option == "N" {
				break
			}

		}
	}

	var uang int
	fmt.Println("Total belanjaan anda adalah :", total)
	fmt.Print("Masukan total uang anda :")
	fmt.Scan(&uang)
	kembalian := uang - total
	if uang < total {
		fmt.Print("Uang anda tidaak mencukupi, pembayaran gagal")
	} else if uang == total {
		fmt.Print("Pembayaran sukses, uang anda pas")
	} else if uang > total {
		fmt.Print("Pembayaran sukses, uang kembalian anda ", kembalian)
	}

}
