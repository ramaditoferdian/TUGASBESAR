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
	menu()
	fmt.Print("Pilihan : ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		menu_1()
		fmt.Print("Pilihan : ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			//TAMBAH\\
			InputDataMahasiswa(data)
		}else if pilihan == 2 {
			//EDIT\\

		}else if pilihan == 3 {
			//HAPUS\\
		}
	}else if pilihan == 2 {

	}

}


//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//=================================================== `KUMPULAN FUNGSI` ======================================================================


func menu(){
	fmt.Println("1. Data Mahasiswa")
	fmt.Println("2. Daftar Program Studi")
}

func menu_1(){
	fmt.Println("1. Tambah")
	fmt.Println("2. Edit")
	fmt.Println("3. Hapus")
}

func menu_2(){

}

func InputDataMahasiswa(data CaMaba){
	var prodi int 

	fmt.Print("Nama	: ")
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
		fmt.Print("Nilai Tes Potensi Akademik : ")
		fmt.Scan(&data.arrNama[0].nilaiTPA)
		fmt.Print("Nilai Tes Bahasa Inggris : ")
		fmt.Scan(&data.arrNama[0].nilaiTBI)
	}else if prodi == 2 {
		fmt.Print("Nilai IPS : ")
		fmt.Scan(&data.arrNama[0].nilairapotnonteknik.IPS)
		fmt.Print("Nilai Matematika : ")
		fmt.Scan(&data.arrNama[0].nilairapotnonteknik.Matematika)
		fmt.Print("Nilai Bahasa Inggris : ")
		fmt.Scan(&data.arrNama[0].nilairapotnonteknik.Ingrris)
		fmt.Print("Nilai Tes Potensi Akademik : ")
		fmt.Scan(&data.arrNama[0].nilaiTPA)
		fmt.Print("Nilai Tes Bahasa Inggris : ")
		fmt.Scan(&data.arrNama[0].nilaiTBI)
	}

}

