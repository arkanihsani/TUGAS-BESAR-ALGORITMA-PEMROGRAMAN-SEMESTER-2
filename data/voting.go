package data

type Pemilih struct {
	Nama    string
	Pilihan *Caleg
}

type ListPemilih [100]Pemilih
