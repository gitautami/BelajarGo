package main

import "fmt"

// Struct
type Mahasiswa struct {
    Nama   string
	Kelas   string
    NPM    string
    Jurusan string
}

// Fungsi untuk menampilkan data mahasiswa
func TampilkanData(m Mahasiswa) {
    fmt.Println("Nama:", m.Nama)
	 fmt.Println("Kelas:", m.Kelas)
    fmt.Println("NPM:", m.NPM)
    fmt.Println("Jurusan:", m.Jurusan)
}

func main() {
    mhs := Mahasiswa{
        Nama:   "Resqi Aulia Gita Utami",
		Kelas:  "2A",
        NPM:    "714230003",
        Jurusan: "D4 Teknik Informatika",
    }

    TampilkanData(mhs)
}
