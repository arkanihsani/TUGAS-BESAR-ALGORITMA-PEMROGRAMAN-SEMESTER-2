package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const NMAX int = 100

type caleg struct {
	no, vot      int
	partai, nama string
}

type listPemilu [NMAX]caleg

func main() {
	var milih int
	var pil, pil1 int
	var cek bool
	var lp listPemilu
	var la listPemilu
	lp[0].no, lp[0].partai, lp[0].nama, lp[0].vot = 1, "NasDem", "Anies", 21432
	lp[1].no, lp[1].partai, lp[1].nama, lp[1].vot = 2, "Gerindra", "Prabowo", 54362
	lp[2].no, lp[2].partai, lp[2].nama, lp[2].vot = 3, "PDIP", "Ganjar", 14533
	nCaleg := 3
	for milih != 3 {
		menu_utama(&milih)
		if milih == 1 {
			cek = menu_login()
			for cek == true {
				menu_petugas(&pil)
				if pil == 1 {
					penambahan(&nCaleg, &lp)
				} else if pil == 2 {
					pengubahan(&nCaleg, &lp)
				} else if pil == 3 {
					penghapusan(&nCaleg, &lp)
				} else if pil == 4 {
					break
				}
			}
		} else if milih == 2 {
			for {
				menu_pemilih(&pil)
				if pil == 1 {
					menu_daftar(&nCaleg, &lp, &pil1)
					if pil1 == 1 {
						sortpartai(nCaleg, &lp, &la)
						cetakData(nCaleg, la)
					} else if pil1 == 2 {
						cetakData(nCaleg, lp)
					} else if pil1 == 3 {
						urutturunvoting(nCaleg, &lp, &la)
						cetakData(nCaleg, la)
					}
				} else if pil == 3 {
					break
				} else if pil == 2 {
					menu_voting(&nCaleg, &lp)
				}
			}
		} else if milih == 3 {
			fmt.Println("Sampai Jumpa!")
		} else {
			fmt.Println("Jawaban Tidak Valid")
		}
	}
}

func menu_login() bool {
	var user, pw string
	var cek bool
	var pil int
	cek = false
	for cek == false {
		fmt.Println("Inputkan Username dan Password")
		fmt.Scan(&user, &pw)
		if user == "admin" && pw == "admin" {
			cek = true
		} else {
			fmt.Println("Username atau Passowrd Salah! Coba Lagi?")
			fmt.Println("1. Yes")
			fmt.Println("2. No")
			fmt.Scan(&pil)
		}
		if pil == 2 {
			break
		}
	}
	return cek
}

func menu_utama(p *int) {
	*p = 0
	for *p != 1 && *p != 2 && *p != 3 {
		fmt.Println("Login Sebagai Pemilih atau Petugas?")
		fmt.Println("1. Petugas")
		fmt.Println("2. Pemilih")
		fmt.Println("3. Exit")
		fmt.Print("Pilih 1/2/3? ")
		fmt.Scan(p)
		if *p != 1 && *p != 2 && *p != 3 {
			fmt.Println("Jawaban Tidak Valid!")
		}
	}
	clearScreen()
}

func menu_petugas(p *int) {
	*p = 0
	for *p != 1 && *p != 2 && *p != 3 && *p != 4 {
		fmt.Println("MENU")
		fmt.Println("1. Penambahan")
		fmt.Println("2. Pengubahan")
		fmt.Println("3. Penghapusan")
		fmt.Println("4. Exit")
		fmt.Print("Pilih 1/2/3/4? ")
		fmt.Scan(&*p)
		if *p != 1 && *p != 2 && *p != 3 && *p != 4 {
			fmt.Println("Jawaban Tidak Valid!")
		}
	}
	clearScreen()
}

func menu_pemilih(p *int) {
	*p = 0
	for *p != 1 && *p != 2 && *p != 3 {
		fmt.Println("MENU")
		fmt.Println("1. Daftar Calon Legislatif")
		fmt.Println("2. Halaman Voting")
		fmt.Println("3. Exit")
		fmt.Print("Pilih 1/2/3? ")
		fmt.Scan(p)
		if *p != 1 && *p != 2 && *p != 3 {
			fmt.Println("Jawaban Tidak Valid!")
		}
	}
	clearScreen()
}

func menu_daftar(nCal *int, li *listPemilu, pil *int) {
	fmt.Println("Daftar Calon Legislatif :")
	cetakData(*nCal, *li)
	fmt.Println("MENU")
	fmt.Println("1. Pencarian Spesifik.")
	fmt.Println("2. Urutkan Berdasarkan Nomor Urut")
	fmt.Println("3. Urutkan Berdasarkan Jumlah Pemungutan Suara")
	fmt.Println("Pilih 1/2/3? ")
	fmt.Scan(&*pil)
}

func menu_voting(nCal *int, li *listPemilu) {
	var bln, tgl, jam, min, pil int
	fmt.Println("Inputkan Tanggal dan Waktu Anda Sekarang. (Voting Hanya Buka Pada Bulan Mei Tanggal 5-25 Pada Jam 8-18)")
	fmt.Scan(&bln, &tgl, &jam, &min)
	clearScreen()
	if bln == 5 && tgl >= 5 && tgl <= 25 && jam >= 8 && jam <= 18 {
		fmt.Println("Daftar Calon Legislatif")
		cetakData(*nCal, *li)
		fmt.Println("Pilih Nomor Urut Yang Ingin Divote, Tulis 0 Jika Ingin Membatalkan.")
		fmt.Scan(&pil)
		clearScreen()
		if pil <= 0 {
			fmt.Println("Voting Dibatalkan")
		} else if li[pil-1].vot != 0 {
			li[pil-1].vot++
			fmt.Println("Terimakasih Sudah Memvoting!")
		} else if li[pil-1].vot == 0 {
			fmt.Println("Nomor Urut Tidak Valid!")
		}
	} else {
		fmt.Println("Maaf, Waktu Voting Sudah Habis.")
	}
}

func sortpartai(nCal int, li, lu *listPemilu) {
	for i := 0; i < nCal; i++ {
		lu[i].no, lu[i].partai, lu[i].nama, lu[i].vot = li[i].no, li[i].partai, li[i].nama, li[i].vot
	}
	for i := 1; i < nCal; i++ {
		idxMin := i - 1
		for j := i; j < nCal; j++ {
			if lu[idxMin].partai > lu[j].partai {
				idxMin = j
			}
		}
		temp := lu[idxMin]
		lu[idxMin] = lu[i-1]
		lu[i-1] = temp
	}
}

func cetakData(nCal int, li listPemilu) {
	for i := 0; i < nCal; i++ {
		if li[i].no != 0 {
			fmt.Println(li[i].no, li[i].partai, li[i].nama, li[i].vot)
		}
	}
}

func penambahan(nCal *int, li *listPemilu) {
	var pili int
	var cek bool
	var urut, vote int
	var party, name string
	fmt.Println("Lanjut Dengan Penambahan?")
	fmt.Println("1. Yes")
	fmt.Println("2. No")
	fmt.Scan(&pili)
	for pili != 2 {
		fmt.Println("Inputkan Nomor Urut, Nama Partai, Nama Calon. dan Jumlah Pemungutan Suara. (Minimum Jumlah Pemungutan Suara Adalah 1000")
		fmt.Scan(&urut, &party, &name, &vote)
		clearScreen()
		cek = true
		for i := 0; i < *nCal; i++ {
			if li[i].no == urut {
				cek = false
			}
		}
		if vote > 1000 && cek == true {
			li[urut-1].no, li[urut-1].partai, li[urut-1].nama, li[urut-1].vot = urut, party, name, vote
			if *nCal < urut {
				*nCal = urut
			}
		} else if vote < 1000 && cek == true {
			fmt.Println("Jumlah Pemungutan Suara Tidak Cukup!")
		} else if cek == false {
			fmt.Println("Nomor Urut Sudah Terisi!")
		}
		fmt.Println("Lanjut Menambah?")
		fmt.Println("1. Yes")
		fmt.Println("2. No")
		fmt.Scan(&pili)
		clearScreen()
	}
	clearScreen()
}

func pengubahan(nCal *int, li *listPemilu) {
	var pil, pil1, pil2, pil3, nobaru int
	for pil2 != 2 {
		fmt.Println("Daftar Calon Legislatif :")
		cetakData(*nCal, *li)
		fmt.Println("Pilih Nomor Urut Yang Ingin Diubah.")
		fmt.Scan(&pil)
		if li[pil-1].no == 0 {
			fmt.Println("Nomor Urut Tidak Valid!")
		}
		for pil1 != 5 && li[pil-1].no != 0 {
			fmt.Println("Pilih Bagian Yang Ingin Diubah.")
			fmt.Println("1. Nomor Urut")
			fmt.Println("2. Partai")
			fmt.Println("3. Nama Calon")
			fmt.Println("4. Jumlah Pemungutan Suara")
			fmt.Println("5. Exit")
			fmt.Println("Pilih 1/2/3/4/5? ")
			fmt.Scan(&pil1)
			if pil1 == 1 {
				fmt.Println("Tuliskan Nomor Urut Yang Baru.")
				fmt.Scan(&nobaru)
				if li[nobaru-1].no == 0 {
					li[nobaru-1].no = nobaru
					li[nobaru-1].partai = li[pil-1].partai
					li[nobaru-1].nama = li[pil-1].nama
					li[nobaru-1].vot = li[pil-1].vot
					li[pil-1].no, li[pil-1].partai, li[pil-1].nama, li[pil-1].vot = 0, "", "", 0
					pil = nobaru
					if nobaru > *nCal {
						*nCal = nobaru
					}
				} else if li[nobaru-1].no != 0 {
					fmt.Println("Nomor Urut ini Sudah Terisi! Pengubahan Dibatalkan!")
				}
			} else if pil1 == 2 {
				fmt.Println("Tuliskan Nama Partai Yang Baru.")
				fmt.Scan(&li[pil-1].partai)
			} else if pil1 == 3 {
				fmt.Println("Tuliskan Nama Calon Yang Baru.")
				fmt.Scan(&li[pil-1].nama)
			} else if pil1 == 4 {
				fmt.Println("Tuliskan Jumlah Pemungutan Suara Yang Baru.")
				fmt.Scan(&li[pil-1].vot)
			} else if pil1 == 5 {
				break
			}
			fmt.Println("Lanjut Dengan Pengubahan Calon Yang Dipilih?")
			fmt.Println("1. Yes")
			fmt.Println("2. No")
			fmt.Scan(&pil3)
			if pil3 == 2 {
				break
			}
		}
		fmt.Println("Lanjut Pengubahan?")
		fmt.Println("1. Yes")
		fmt.Println("2. No")
		fmt.Scan(&pil2)
	}
}

func penghapusan(nCal *int, li *listPemilu) {
	var pil, pil1, pil2 int
	for pil2 != 2 {
		fmt.Println("Daftar Calon Legislatif :")
		cetakData(*nCal, *li)

		fmt.Println("Pilih Nomor Urut Yang Ingin Dihapus.")
		fmt.Scan(&pil)
		if li[pil-1].no != 0 {
			fmt.Println("Lanjut Dengan Penghapusan?")
			fmt.Println("1. Yes")
			fmt.Println("2. No")
			fmt.Scan(&pil1)
			if pil1 == 1 {
				li[pil-1].no = 0
			}
		} else if li[pil-1].no == 0 {
			fmt.Println("Nomor Urut Tidak Valid!")
		}
		fmt.Println("Lanjut Penghapusan?")
		fmt.Println("1. Yes")
		fmt.Println("2. No")
		fmt.Scan(&pil2)
	}
}

func urutturunvoting(nCal int, li, lu *listPemilu) {
	var pass, i int
	var temp caleg
	for i = 0; i < nCal; i++ {
		lu[i].no, lu[i].partai, lu[i].nama, lu[i].vot = li[i].no, li[i].partai, li[i].nama, li[i].vot
	}
	for pass < nCal {
		i = pass
		temp = lu[pass]
		for i > 0 && temp.vot > lu[i-1].vot {
			lu[i] = lu[i-1]
			i--
		}
		lu[i] = temp
		pass++
	}
}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		cmd = exec.Command("clear")
	} else {
		fmt.Println("Unsupported platform")
		return
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
