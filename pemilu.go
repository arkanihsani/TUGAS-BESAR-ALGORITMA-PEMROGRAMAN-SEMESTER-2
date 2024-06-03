package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// konstanta NMAX = 100
const NMAX int = 100
const NMAX2 int = 1000

// tipe bentukan struktur tCaleg dengan atribut nomor urut dan jumlah voting (integer), partai dan nama calon (string)
type tCaleg struct {
	no, vot      int
	partai, nama string
}

type tPemilih struct {
	nama string
	urut int
}

// Tipe alias listPemilu untuk array of tCaleg dengan ukuran NMAX
type listPemilu [NMAX]tCaleg
type namaPemilih [NMAX2]tPemilih

func main() {
	var milih int
	var pil, pil1, pil2 int
	var cek bool
	var lp listPemilu
	var la listPemilu
	var np namaPemilih
	lp[0].no, lp[0].partai, lp[0].nama, lp[0].vot = 1, "NasDem", "Anies", 21432
	lp[1].no, lp[1].partai, lp[1].nama, lp[1].vot = 2, "Gerindra", "Prabowo", 54362
	lp[2].no, lp[2].partai, lp[2].nama, lp[2].vot = 3, "PDIP", "Ganjar", 14533
	nPemilih := 0
	nCaleg := 3
	clearScreen()
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
					for {
						menu_daftar(&pil1)
						if pil1 == 1 {
							for {
								menu_pencarian(&pil2)
								if pil2 == 1 || pil2 == 2 {
									pencarian(nCaleg, pil2, lp)
								} else if pil2 == 4 {
									break
								} else if pil2 == 3 {
									pencarian_pemilih(nCaleg, nPemilih, lp, np)
								}
							}
						} else if pil1 == 2 {
							cetakData(nCaleg, lp)
						} else if pil1 == 3 {
							sortpartai(nCaleg, &lp, &la)
							cetakData(nCaleg, la)
						} else if pil1 == 4 {
							urutvoting(nCaleg, &lp, &la)
							cetakData(nCaleg, la)
						} else if pil1 == 5 {
							sortnama(nCaleg, &lp, &la)
							cetakData(nCaleg, la)
						} else if pil1 == 6 {
							break
						}
					}
				} else if pil == 3 {
					break
				} else if pil == 2 {
					menu_voting(&nCaleg, &nPemilih, &lp, &np)
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
		clearScreen()
		if user == "admin" && pw == "admin" {
			cek = true
		} else {
			pil = 0
			for pil != 1 && pil != 2 {
				fmt.Println("Username atau Password Salah! Coba Lagi?")
				fmt.Println("1. Yes")
				fmt.Println("2. No")
				fmt.Scan(&pil)
				clearScreen()
				if pil != 1 && pil != 2 {
					fmt.Println("Jawaban Tidak Valid!")
				}
			}
		}
		if pil == 2 {
			break
		}
	}
	clearScreen()
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
		clearScreen()
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
		fmt.Scan(p)
		clearScreen()
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
		clearScreen()
		if *p != 1 && *p != 2 && *p != 3 {
			fmt.Println("Jawaban Tidak Valid!")
		}
	}
	clearScreen()
}

// function untuk menampilkan menu daftar data caleg
func menu_daftar(pil *int) {
	*pil = 0
	for *pil != 1 && *pil != 2 && *pil != 3 && *pil != 4 && *pil != 5 && *pil != 6 {
		fmt.Println("MENU")
		fmt.Println("1. Pencarian Spesifik.")
		fmt.Println("2. Urutkan Berdasarkan Nomor Urut")
		fmt.Println("3. Urutkan Berdasarkan Partai")
		fmt.Println("4. Urutkan Berdasarkan Jumlah Pemungutan Suara")
		fmt.Println("5. Urutkan Berdasarkan Nama Calon")
		fmt.Println("6. Exit.")
		fmt.Println("Pilih 1/2/3/4/5/6? ")
		fmt.Scan(&*pil)
		clearScreen()
		if *pil != 1 && *pil != 2 && *pil != 3 && *pil != 4 && *pil != 5 && *pil != 6 {
			fmt.Println("Jawaban Tidak Valid!")
		}
	}
	clearScreen()
}

// function untuk menampilkan menu voting/pemilih
func menu_voting(nCal, nPem *int, li *listPemilu, pem *namaPemilih) {
	var bln, tgl, jam, min, pil int
	var name string
	fmt.Println("Inputkan Waktu Anda Sekarang (Bulan, Tanggal dan Waktu (Jam dan Menit)) dan Nama Anda (Voting Hanya Buka Pada Bulan Mei Tanggal 5-25 Pada Jam 8-18)")
	fmt.Scan(&bln, &tgl, &jam, &min, &name)
	clearScreen()
	if bln == 5 && tgl >= 5 && tgl <= 25 && jam >= 8 && jam <= 18 && min >= 0 && min <= 59 {
		cetakData(*nCal, *li)
		fmt.Println("Pilih Nomor Urut Yang Ingin Divote, Tulis 0 Jika Ingin Membatalkan.")
		fmt.Scan(&pil)
		clearScreen()
		if pil <= 0 || pil > 100 {
			fmt.Println("Voting Dibatalkan")
		} else if li[pil-1].vot != 0 {
			li[pil-1].vot++
			pem[*nPem].nama = name
			pem[*nPem].urut = pil
			*nPem++
			fmt.Println("Terimakasih Sudah Memvoting!")
		} else if li[pil-1].vot == 0 {
			fmt.Println("Nomor Urut Tidak Valid!")
		}
	} else if bln >= 1 && bln <= 12 && tgl <= 1 && tgl >= 4 && tgl <= 30 && tgl >= 26 && jam >= 1 && jam <= 7 && jam >= 19 && jam <= 24 && min >= 0 && min <= 59 {
		fmt.Println("Maaf, Waktu Voting Sudah Habis.")
	} else {
		fmt.Println("Input Tidak Valid!")
	}
}

// mencetak data caleg
func cetakData(nCal int, li listPemilu) {
	fmt.Println("Daftar Calon Legislatif")
	for i := 0; i < nCal; i++ {
		if li[i].no != 0 {
			fmt.Println(li[i].no, li[i].partai, li[i].nama, li[i].vot)
		}
	}
}

// function untuk menambah data caleg yang tersedia
func penambahan(nCal *int, li *listPemilu) {
	var pili int
	var cek bool
	var urut, vote int
	var party, name string
	pili = 0
	for pili != 1 && pili != 2 {
		fmt.Println("Lanjut Dengan Penambahan?")
		fmt.Println("1. Yes")
		fmt.Println("2. No")
		fmt.Scan(&pili)
		clearScreen()
		if pili != 1 && pili != 2 {
			fmt.Println("Jawaban Tidak Valid!")
		}
	}
	for pili != 2 {
		fmt.Println("Inputkan Nomor Urut, Nama Partai, Nama Calon. dan Jumlah Pemungutan Suara. (Minimum Jumlah Pemungutan Suara Adalah 1000)")
		fmt.Scan(&urut, &party, &name, &vote)
		clearScreen()
		cek = true
		for i := 0; i < *nCal; i++ {
			if li[i].no == urut {
				cek = false
			}
		}
		if urut <= 0 || urut > 100 {
			fmt.Println("Nomor Urut Tidak Valid! Penambahan Dibatalkan")
		} else if vote > 1000 && cek == true {
			li[urut-1].no, li[urut-1].partai, li[urut-1].nama, li[urut-1].vot = urut, party, name, vote
			if *nCal < urut {
				*nCal = urut
			}
		} else if vote < 1000 && cek == true {
			fmt.Println("Jumlah Pemungutan Suara Tidak Cukup!")
		} else if cek == false && urut > 0 {
			fmt.Println("Nomor Urut Sudah Terisi! Penambahan Dibatalkan")
		}
		pili = 0
		for pili != 1 && pili != 2 {
			fmt.Println("Lanjut Menambah?")
			fmt.Println("1. Yes")
			fmt.Println("2. No")
			fmt.Scan(&pili)
			clearScreen()
			if pili != 1 && pili != 2 {
				fmt.Println("Jawaban Tidak Valid!")
			}
		}
	}
	clearScreen()
}

// function untuk mengubah data caleg
func pengubahan(nCal *int, li *listPemilu) {
	var pil, pil1, pil2, pil3, nobaru, vote int
	for pil2 != 2 {
		cetakData(*nCal, *li)
		fmt.Println("Pilih Nomor Urut Yang Ingin Diubah.")
		fmt.Scan(&pil)
		clearScreen()
		if pil <= 0 || pil > 100 {
			fmt.Println("Nomor Urut Tidak Valid!")
		} else if li[pil-1].no == 0 {
			fmt.Println("Nomor Urut Tidak Valid!")
		} else if pil > 0 {
			for li[pil-1].no != 0 && pil > 0 {
				fmt.Println("Pilih Bagian Yang Ingin Diubah.")
				fmt.Println("1. Nomor Urut")
				fmt.Println("2. Partai")
				fmt.Println("3. Nama Calon")
				fmt.Println("4. Jumlah Pemungutan Suara")
				fmt.Println("5. Exit")
				fmt.Println("Pilih 1/2/3/4/5? ")
				fmt.Scan(&pil1)
				clearScreen()
				if pil1 == 1 {
					fmt.Println("Tuliskan Nomor Urut Yang Baru.")
					fmt.Scan(&nobaru)
					clearScreen()
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
					clearScreen()
				} else if pil1 == 3 {
					fmt.Println("Tuliskan Nama Calon Yang Baru.")
					fmt.Scan(&li[pil-1].nama)
					clearScreen()
				} else if pil1 == 4 {
					fmt.Println("Tuliskan Jumlah Pemungutan Suara Yang Baru.")
					fmt.Scan(&vote)
					clearScreen()
					if vote < 1000 {
						fmt.Println("Jumlah Pemungutan Suara Dibawah Threshold!")
					} else {
						li[pil-1].vot = vote
					}
				} else if pil1 == 5 {
					break
				} else {
					fmt.Println("Jawaban Tidak Valid!")
				}
				fmt.Println("Lanjut Dengan Pengubahan Calon Yang Dipilih?")
				fmt.Println("1. Yes")
				fmt.Println("2. No")
				fmt.Scan(&pil3)
				clearScreen()
				if pil3 == 2 {
					break
				}
			}
		}
		pil2 = 0
		for pil2 != 1 && pil2 != 2 {
			fmt.Println("Lanjut Pengubahan?")
			fmt.Println("1. Yes")
			fmt.Println("2. No")
			fmt.Scan(&pil2)
			clearScreen()
			if pil2 != 1 && pil2 != 2 {
				fmt.Println("Jawaban Tidak Valid!")
			}
		}
	}
}

// function untuk menghapus data caleg
func penghapusan(nCal *int, li *listPemilu) {
	var pil, pil1, pil2 int
	for pil2 != 2 {
		cetakData(*nCal, *li)

		fmt.Println("Pilih Nomor Urut Yang Ingin Dihapus.")
		fmt.Scan(&pil)
		clearScreen()

		if pil <= 0 || pil > 100 {
			fmt.Println("Nomor Urut Tidak Valid!")
		} else if li[pil-1].no != 0 {
			fmt.Println("Lanjut Dengan Penghapusan?")
			fmt.Println("1. Yes")
			fmt.Println("2. No")
			fmt.Scan(&pil1)
			clearScreen()
			if pil1 == 1 {
				li[pil-1].no = 0
			} else if pil1 == 2 {

			} else {
				fmt.Println("Jawaban Tidak Valid!")
			}
		} else if li[pil-1].no == 0 {
			fmt.Println("Nomor Urut Tidak Valid!")
		}
		pil2 = 0
		for pil2 != 1 && pil2 != 2 {
			fmt.Println("Lanjut Penghapusan?")
			fmt.Println("1. Yes")
			fmt.Println("2. No")
			fmt.Scan(&pil2)
			clearScreen()
			if pil2 != 1 && pil2 != 2 {
				fmt.Println("Jawaban Tidak Valid!")
			}
		}
	}
}

// function untuk mengurutkan data berdasarkan hasil voting
func urutvoting(nCal int, li, lu *listPemilu) {
	var pass, i, p int
	var temp tCaleg
	for i = 0; i < nCal; i++ {
		lu[i].no, lu[i].partai, lu[i].nama, lu[i].vot = li[i].no, li[i].partai, li[i].nama, li[i].vot
	}
	for p != 1 && p != 2 {
		fmt.Println("1. Menurun/2. Menaik? ")
		fmt.Scan(&p)
		clearScreen()
		if p != 1 && p != 2 {
			fmt.Println("Jawaban Tidak Valid!")
		}
	}
	if p == 1 {
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
	} else if p == 2 {
		for pass < nCal {
			i = pass
			temp = lu[pass]
			for i > 0 && temp.vot < lu[i-1].vot {
				lu[i] = lu[i-1]
				i--
			}
			lu[i] = temp
			pass++
		}
	}
}

// function untuk mengurutkan data berdasarkan nama partai
func sortpartai(nCal int, li, lu *listPemilu) {
	for i := 0; i < nCal; i++ {
		lu[i].no, lu[i].partai, lu[i].nama, lu[i].vot = li[i].no, li[i].partai, li[i].nama, li[i].vot
	}
	var p int
	for p != 1 && p != 2 {
		fmt.Println("1. Menurun/2. Menaik? ")
		fmt.Scan(&p)
		clearScreen()
		if p != 1 && p != 2 {
			fmt.Println("Jawaban Tidak Valid!")
		}
	}
	if p == 1 {
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
	} else if p == 2 {
		for i := 1; i < nCal; i++ {
			idxMin := i - 1
			for j := i; j < nCal; j++ {
				if lu[idxMin].partai < lu[j].partai {
					idxMin = j
				}
			}
			temp := lu[idxMin]
			lu[idxMin] = lu[i-1]
			lu[i-1] = temp
		}
	}
}

// function untuk mengurutkan data berdasarkan nama calon
func sortnama(nCal int, li, lu *listPemilu) {
	for i := 0; i < nCal; i++ {
		lu[i].no, lu[i].partai, lu[i].nama, lu[i].vot = li[i].no, li[i].partai, li[i].nama, li[i].vot
	}
	var p int
	for p != 1 && p != 2 {
		fmt.Println("1. Menurun/2. Menaik? ")
		fmt.Scan(&p)
		clearScreen()
		if p != 1 && p != 2 {
			fmt.Println("Jawaban Tidak Valid!")
		}
	}
	if p == 1 {
		for i := 1; i < nCal; i++ {
			idxMin := i - 1
			for j := i; j < nCal; j++ {
				if lu[idxMin].nama > lu[j].nama {
					idxMin = j
				}
			}
			temp := lu[idxMin]
			lu[idxMin] = lu[i-1]
			lu[i-1] = temp
		}
	} else if p == 2 {
		for i := 1; i < nCal; i++ {
			idxMin := i - 1
			for j := i; j < nCal; j++ {
				if lu[idxMin].nama < lu[j].nama {
					idxMin = j
				}
			}
			temp := lu[idxMin]
			lu[idxMin] = lu[i-1]
			lu[i-1] = temp
		}
	}
}

// function untuk menampilkan menu pencarian dan mengambil input
func menu_pencarian(pil *int) {
	*pil = 0
	for *pil != 1 && *pil != 2 && *pil != 3 && *pil != 4 {
		fmt.Println("MENU")
		fmt.Println("1. Pencarian Berdasarkan Partai")
		fmt.Println("2. Pencarian Berdasarkan Nama Calon")
		fmt.Println("3. Pencarian Data Pemilih.")
		fmt.Println("4. Exit")
		fmt.Println("Pilih 1/2/3/4? ")
		fmt.Scan(&*pil)
		clearScreen()
		if *pil != 1 && *pil != 2 && *pil != 3 && *pil != 4 {
			fmt.Println("Jawaban Tidak Valid!")
		}
	}
	clearScreen()
}

// function untuk mencari data berdasarkan nama partai atau calon
func pencarian(nCal, p int, li listPemilu) {
	var x string
	cek := -1
	i := 0
	if p == 1 {
		fmt.Println("Inputkan Nama Partai")
		fmt.Scan(&x)
		clearScreen()
		for i < nCal && cek == -1 {
			if li[i].partai == x {
				cek = i
			}
			i++
		}
	} else if p == 2 {
		fmt.Println("Inputkan Nama Calon")
		fmt.Scan(&x)
		clearScreen()
		for i < nCal && cek == -1 {
			if li[i].nama == x {
				cek = i
			}
			i++
		}
	}
	if cek == -1 {
		fmt.Println("Data Tidak Ditemukan!")
	} else {
		fmt.Println("Data Ditemukan!")
		fmt.Println(li[cek].no, li[cek].partai, li[cek].nama, li[cek].vot)
	}
}

func pencarian_pemilih(nCal, nPem int, li listPemilu, pem namaPemilih) {
	var x int
	var cek, cek2 bool
	fmt.Println("Inputkan Nomor Urut Yang Ingin Dicari!")
	fmt.Scan(&x)
	clearScreen()
	for i := 0; i < nCal; i++ {
		if li[i].no == x {
			cek = true
		}
	}
	if cek == true {
		fmt.Println("Pemilih Yang Memvoting Nomor Urut Tersebut :")
		for j := 0; j < nPem; j++ {
			if pem[j].urut == x {
				fmt.Println(pem[j].nama)
				cek2 = true
			}
		}
		if cek2 == false {
			fmt.Println("Data Kosong!")
		}
		fmt.Println("---------------------")
	} else if cek == false {
		fmt.Println("Nomor Urut Tidak Ditemukan!")
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
