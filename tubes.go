package main

import "fmt"

const NMAX = 9999

type makan struct {
	namaMakanan, resep string
}

type Makanan [NMAX]makan

var daftarResep Makanan
var n int

func dashboard() {

	fmt.Println("Pilihan")
	fmt.Println("1. Tambah Resep")
	fmt.Println("2. Lihat Resep Makanan")
	fmt.Println("3. Hapus Resep")
	fmt.Println("4. Keluar")
}

func tambahResep() {
	var namaBaru, resepBaru string

	fmt.Println("Masukkan Nama Makanan")
	fmt.Scan(&namaBaru)
	fmt.Println("Masukan Resep")
	fmt.Scan(&resepBaru)

	posisi := sorting(namaBaru)

	j := n - 1
	for j >= posisi {
		daftarResep[j+1] = daftarResep[j]
		j--
	}
	daftarResep[posisi] = makan{namaBaru, resepBaru}
	n += 1

}

func sorting(namaBaru string) int { //mecari posisi atau sorting
	i := 0
	for i < n {
		if namaBaru < daftarResep[i].namaMakanan {
			return i
		}
		i++
	}
	return n
}

func tampilkanResep() { //menampilkan semua resep

	if n == 0 {
		fmt.Println("Belum Ada Data Makanan/Resep")
		return
	}

	i := 0
	for i < n {
		fmt.Println((i + 1), daftarResep[i].namaMakanan)
		i += 1
	}
}

func detailResep() {
	var nomor int
	if n == 0 {
		fmt.Println("Belum Ada Data Makanan/Resep")
	}

	tampilkanResep()
	fmt.Scan(&nomor)

	if nomor < 1 || nomor > n {
		fmt.Println("TIdak Valid")
		return
	}

	idx := nomor - 1
	fmt.Println(daftarResep[idx].namaMakanan)
	fmt.Println(daftarResep[idx].resep)
}

func hapusResep() {
	var nomor int

	if n == 0 {
		fmt.Println("Belum Ada Data Makanan/Resep")
	}

	tampilkanResep()
	fmt.Scan(&nomor)

	if nomor < 1 || nomor > n {
		fmt.Println("TIdak Valid")
		return
	}

	idx := nomor - 1
	i := idx
	for i < n-1 {
		daftarResep[i] = daftarResep[i+1]
		i += 1
	}
	n -= 1
}

func main() {
	var pilihan int

	for {
		dashboard()
		fmt.Scan(&pilihan)

		switch {
		case pilihan == 1:
			tambahResep()
		case pilihan == 2:
			tampilkanResep()
			detailResep()
		case pilihan == 3:
			hapusResep()
		default:
			return
		}
	}
}
