package main

import (
	"fmt"
	"pemilu/data"
	"pemilu/utils"
)

func main() {
	var caleg data.ListCaleg
	var pemilih data.ListPemilih
	var nCaleg, nPemilih int

	// Load existing data
	savedCaleg, count := utils.LoadCaleg()
	for i := 0; i < count; i++ {
		caleg[i] = savedCaleg[i]
	}
	nCaleg = count

	for {
		fmt.Println("====================================")
		fmt.Println("SISTEM PEMILU SEDERHANA")
		fmt.Println("====================================")
		fmt.Println("1. Voting")
		fmt.Println("2. Cetak Data Caleg")
		fmt.Println("3. Tambah Caleg")
		fmt.Println("4. Hapus Caleg")
		fmt.Println("0. Keluar")
		fmt.Print("Pilihan: ")

		var pilihan int
		fmt.Scan(&pilihan)
		utils.ClearScreen()

		switch pilihan {
		case 1:
			if nCaleg == 0 {
				fmt.Println("Belum ada caleg yang didaftarkan!")
				break
			}

			var nama string
			var noPilihan int

			fmt.Print("Masukkan nama pemilih: ")
			fmt.Scan(&nama)
			fmt.Println("Daftar Caleg:")
			for i := 0; i < nCaleg; i++ {
				fmt.Printf("%d. %s\n", caleg[i].No, caleg[i].Nama)
			}
			fmt.Print("Masukkan nomor caleg yang dipilih: ")
			fmt.Scan(&noPilihan)

			if noPilihan > 0 && noPilihan <= nCaleg {
				pemilih[nPemilih] = data.Pemilih{
					Nama:    nama,
					Pilihan: &caleg[noPilihan-1],
				}
				caleg[noPilihan-1].Vot++
				nPemilih++
				fmt.Println("Voting berhasil dicatat!")
				utils.SaveCaleg(caleg, nCaleg) // update file
			} else {
				fmt.Println("Nomor caleg tidak valid!")
			}

		case 2:
			if nCaleg == 0 {
				fmt.Println("Belum ada data caleg!")
			} else {
				fmt.Println("Daftar Caleg:")
				for i := 0; i < nCaleg; i++ {
					fmt.Printf("%d. %s - %d suara\n", caleg[i].No, caleg[i].Nama, caleg[i].Vot)
				}
			}

		case 3:
			var nama string
			fmt.Print("Masukkan Nama Caleg: ")
			fmt.Scan(&nama)

			nCaleg++
			caleg[nCaleg-1] = data.Caleg{
				No:   nCaleg,
				Nama: nama,
				Vot:  0,
			}
			fmt.Println("Caleg berhasil ditambahkan!")
			utils.SaveCaleg(caleg, nCaleg)

		case 4:
			if nCaleg == 0 {
				fmt.Println("Belum ada caleg yang didaftarkan!")
				break
			}

			var no int
			fmt.Println("Daftar Caleg:")
			for i := 0; i < nCaleg; i++ {
				fmt.Printf("%d. %s - %d suara\n", caleg[i].No, caleg[i].Nama, caleg[i].Vot)
			}
			fmt.Print("Masukkan nomor caleg yang ingin dihapus: ")
			fmt.Scan(&no)

			if no > 0 && no <= nCaleg {
				// Geser array ke kiri untuk hapus
				for i := no - 1; i < nCaleg-1; i++ {
					caleg[i] = caleg[i+1]
					caleg[i].No = i + 1 // update nomor
				}
				nCaleg--
				fmt.Println("Caleg berhasil dihapus!")
				utils.SaveCaleg(caleg, nCaleg)
			} else {
				fmt.Println("Nomor caleg tidak valid!")
			}

		case 0:
			utils.SaveCaleg(caleg, nCaleg)
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
