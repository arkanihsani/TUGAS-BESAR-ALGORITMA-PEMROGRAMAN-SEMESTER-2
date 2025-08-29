package main

import (
	"fmt"
	"os"
	"os/exec"
)

const NMAX = 100

type caleg struct {
	no     int
	nama   string
	partai string
	vot    int
}

type pemilih struct {
	nama string
	urut int
}

type listPemilu [NMAX]caleg
type namaPemilih [NMAX]pemilih

// -------------------- MAIN --------------------

func main() {
	var (
		nCaleg, nPemilih int
		lp, la           listPemilu
		np               namaPemilih
		pilih            int
	)

	for {
		clearScreen()
		fmt.Println("====================================")
		fmt.Println("SISTEM PEMILU SEDERHANA")
		fmt.Println("====================================")
		fmt.Println("1. Voting")
		fmt.Println("2. Cetak Data Caleg")
		fmt.Println("3. Cari Caleg berdasarkan Partai")
		fmt.Println("4. Cari Caleg berdasarkan Nama")
		fmt.Println("5. Cari Pemilih")
		fmt.Println("6. Tambah Caleg")
		fmt.Println("7. Ubah Caleg")
		fmt.Println("8. Hapus Caleg")
		fmt.Println("9. Urutkan berdasarkan Partai")
		fmt.Println("10. Urutkan berdasarkan Nama")
		fmt.Println("11. Urutkan berdasarkan Jumlah Voting")
		fmt.Println("0. Keluar")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			menu_voting(&nCaleg, &nPemilih, &lp, &np)
		case 2:
			clearScreen()
			cetakData(nCaleg, lp)
			pause()
		case 3:
			clearScreen()
			pencarian(nCaleg, 1, lp)
			pause()
		case 4:
			clearScreen()
			pencarian(nCaleg, 2, lp)
			pause()
		case 5:
			clearScreen()
			pencarian_pemilih(nCaleg, nPemilih, lp, np)
			pause()
		case 6:
			penambahan(&nCaleg, &lp)
		case 7:
			pengubahan(&nCaleg, &lp)
		case 8:
			penghapusan(&nCaleg, &lp)
		case 9:
			sortpartai(nCaleg, &lp, &la)
			clearScreen()
			cetakData(nCaleg, la)
			pause()
		case 10:
			sortnama(nCaleg, &lp, &la)
			clearScreen()
			cetakData(nCaleg, la)
			pause()
		case 11:
			urutvoting(nCaleg, &lp, &la)
			clearScreen()
			cetakData(nCaleg, la)
			pause()
		case 0:
			fmt.Println("Terima kasih sudah menggunakan program ini!")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

// -------------------- UTILS --------------------

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls") // Windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func pause() {
	fmt.Println("\nTekan ENTER untuk kembali ke menu...")
	fmt.Scanln()
	fmt.Scanln()
}

// -------------------- VOTING & DISPLAY --------------------

func menu_voting(nCaleg, nPemilih *int, lp *listPemilu, np *namaPemilih) {
	clearScreen()
	var (
		nama    string
		no_urut int
		valid   bool
	)

	fmt.Print("Masukkan Nama Anda: ")
	fmt.Scan(&nama)

	// cek apakah sudah pernah memilih
	for i := 0; i < *nPemilih; i++ {
		if np[i].nama == nama {
			fmt.Println("Anda sudah pernah melakukan voting!")
			pause()
			return
		}
	}

	fmt.Println("\nDaftar Calon:")
	for i := 0; i < *nCaleg; i++ {
		fmt.Printf("%d. %s (%s)\n", lp[i].no, lp[i].nama, lp[i].partai)
	}

	valid = false
	for !valid {
		fmt.Print("Pilih nomor urut caleg: ")
		fmt.Scan(&no_urut)
		for i := 0; i < *nCaleg; i++ {
			if lp[i].no == no_urut {
				lp[i].vot++
				np[*nPemilih].nama = nama
				np[*nPemilih].urut = no_urut
				*nPemilih++
				fmt.Println("Voting berhasil!")
				valid = true
				break
			}
		}
		if !valid {
			fmt.Println("Nomor urut tidak valid!")
		}
	}
	pause()
}

func cetakData(nCaleg int, lp listPemilu) {
	fmt.Println("====================================")
	fmt.Println("Daftar Calon Legislatif")
	fmt.Println("====================================")
	for i := 0; i < nCaleg; i++ {
		fmt.Printf("%d. %-10s %-10s [Suara: %d]\n", lp[i].no, lp[i].nama, lp[i].partai, lp[i].vot)
	}
	fmt.Println("====================================")
}

func pencarian(nCaleg, pilihan int, lp listPemilu) {
	var cari string
	fmt.Print("Masukkan kata kunci pencarian: ")
	fmt.Scan(&cari)

	fmt.Println("Hasil Pencarian:")
	ditemukan := false
	for i := 0; i < nCaleg; i++ {
		if (pilihan == 1 && lp[i].partai == cari) || (pilihan == 2 && lp[i].nama == cari) {
			fmt.Printf("%d. %-10s %-10s [Suara: %d]\n", lp[i].no, lp[i].nama, lp[i].partai, lp[i].vot)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Data tidak ditemukan!")
	}
}

func pencarian_pemilih(nCaleg, nPemilih int, lp listPemilu, np namaPemilih) {
	var nama string
	fmt.Print("Masukkan nama pemilih: ")
	fmt.Scan(&nama)

	for i := 0; i < nPemilih; i++ {
		if np[i].nama == nama {
			no := np[i].urut
			for j := 0; j < nCaleg; j++ {
				if lp[j].no == no {
					fmt.Printf("Pemilih %s memilih %s (%s)\n", nama, lp[j].nama, lp[j].partai)
					return
				}
			}
		}
	}
	fmt.Println("Data pemilih tidak ditemukan!")
}

// -------------------- CRUD --------------------

func penambahan(nCaleg *int, lp *listPemilu) {
	clearScreen()
	var nama, partai string
	fmt.Print("Masukkan Nama Caleg: ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan Partai: ")
	fmt.Scan(&partai)

	*nCaleg++
	lp[*nCaleg-1].no = *nCaleg
	lp[*nCaleg-1].nama = nama
	lp[*nCaleg-1].partai = partai
	lp[*nCaleg-1].vot = 0

	fmt.Println("Caleg berhasil ditambahkan!")
	pause()
}

func pengubahan(nCaleg *int, lp *listPemilu) {
	clearScreen()
	var no int
	fmt.Print("Masukkan nomor urut caleg yang ingin diubah: ")
	fmt.Scan(&no)

	for i := 0; i < *nCaleg; i++ {
		if lp[i].no == no {
			fmt.Print("Masukkan nama baru: ")
			fmt.Scan(&lp[i].nama)
			fmt.Print("Masukkan partai baru: ")
			fmt.Scan(&lp[i].partai)
			fmt.Println("Data berhasil diubah!")
			pause()
			return
		}
	}
	fmt.Println("Caleg tidak ditemukan!")
	pause()
}

func penghapusan(nCaleg *int, lp *listPemilu) {
	clearScreen()
	var no int
	fmt.Print("Masukkan nomor urut caleg yang ingin dihapus: ")
	fmt.Scan(&no)

	for i := 0; i < *nCaleg; i++ {
		if lp[i].no == no {
			// geser elemen
			for j := i; j < *nCaleg-1; j++ {
				lp[j] = lp[j+1]
			}
			*nCaleg--
			fmt.Println("Caleg berhasil dihapus!")
			pause()
			return
		}
	}
	fmt.Println("Caleg tidak ditemukan!")
	pause()
}

// -------------------- SORTING --------------------

func sortpartai(nCaleg int, lp *listPemilu, la *listPemilu) {
	*la = *lp
	for i := 0; i < nCaleg-1; i++ {
		for j := i + 1; j < nCaleg; j++ {
			if la[i].partai > la[j].partai {
				la[i], la[j] = la[j], la[i]
			}
		}
	}
}

func sortnama(nCaleg int, lp *listPemilu, la *listPemilu) {
	*la = *lp
	for i := 0; i < nCaleg-1; i++ {
		for j := i + 1; j < nCaleg; j++ {
			if la[i].nama > la[j].nama {
				la[i], la[j] = la[j], la[i]
			}
		}
	}
}

func urutvoting(nCaleg int, lp *listPemilu, la *listPemilu) {
	*la = *lp
	for i := 0; i < nCaleg-1; i++ {
		for j := i + 1; j < nCaleg; j++ {
			if la[i].vot < la[j].vot { // descending
				la[i], la[j] = la[j], la[i]
			}
		}
	}
}
