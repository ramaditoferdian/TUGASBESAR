package main
import (
	"fmt"
)

//////////////////////////////////////////////////////  ` VARIABLE dan TIPE DATA BENTUKAN`  \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\

type (
	CaMaba struct{
		arrNama[10] DataDiri
	}

	DataDiri struct{
		nama   string
		gender string
		nilaiTPA int
		nilaiTBI int
		nilairapotteknik Nilaiteknik
		nilairapotnonteknik Nilainonteknik
	}

	Nilaiteknik struct {
		IPA int
		Ingrris int
		Matematika int
	}

	Nilainonteknik struct{
		IPS int
		Ingrris int
		Matematika int
	}
)

var (
	
)

//  ////////////////////////////////////////////////////////////////\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
// |||||||||||||||||||||||||||||||||||||||||||||||||||||||| `FUNGSI UTAMA`  |||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
//  \\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\//////////////////////////////////////////////////////////////////////////

func main(){
	var (
		data CaMaba
		pilihan int
	)
	menu_1()
	fmt.Print("Pilihan : ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		menu_1_1()
		fmt.Print("Pilihan : ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			InputDataMahasiswa(data)
		}
	}

}


//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//=================================================== `KUMPULAN FUNGSI` ======================================================================


func menu_1(){
	fmt.Println("1.Data Mahasiswa")
}

func menu_1_1(){
	fmt.Println("1.Tambah")
	fmt.Println("2.Edit")
	fmt.Println("3.Hapus")
}

func InputDataMahasiswa(data CaMaba){
	var prodi int 

	fmt.Print("Nama  : ")
	fmt.Scan(&data.arrNama[0].nama)
	fmt.Println("1. Teknik ")
	fmt.Println("2. Non-Teknik ")
	fmt.Print("Pilih : ")
	fmt.Scan(&prodi)
	if prodi == 1 {
		fmt.Print("Nilai IPA : ")
		fmt.Scan(&data.arrNama[0].nilairapotteknik.IPA)
		fmt.Print("Nilai Matematika : ")
		fmt.Scan(&data.arrNama[0].nilairapotteknik.Matematika)
		fmt.Print("Nilai Bahasa Inggris : ")
		fmt.Scan(&data.arrNama[0].nilairapotteknik.Ingrris)
	}else if prodi == 2 {
		fmt.Print("Nilai IPS : ")
		fmt.Scan(&data.arrNama[0].nilairapotnonteknik.IPS)
		fmt.Print("Nilai Matematika : ")
		fmt.Scan(&data.arrNama[0].nilairapotnonteknik.Matematika)
		fmt.Print("Nilai Bahasa Inggris : ")
		fmt.Scan(&data.arrNama[0].nilairapotnonteknik.Ingrris)
	}

}

