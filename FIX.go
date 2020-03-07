/* 
	KELOMPOK 15
	NAMA 	:	BAGJA 9102 KURNIAWAN
	NIM		:	1301194020
	NAMA	:	MUHAMMAD KAMIL HASAN
	NIM		:	1301190420
*/
package main

import (	"fmt"
			"strings"
			"os"
			"os/exec"
		)

const T = 20
const N = 4
const M = 4
const DATA = 2019

type kendaraan struct{
	nopol	string
	jenis	string
	bulan	int
	tanggal int
	tahun	int
	jam		int
	menit	int
}
	
type motor [T]kendaraan
type mobil [N][M]kendaraan
type tampung [DATA]kendaraan

var kosong = kendaraan {
	nopol	:	"",
	jenis	:	"",
	bulan	:	0,
	tanggal	:	0,
	tahun	:	0,
	jam		:	0,
	menit	:	0 }
	
var nMotor,nMobil,nTruk, sMotor, sMobil, total int
var nomor,baris,kolom int
var mot motor
var mob mobil
var temp tampung

func isiParkir(){
/*	I.S. Array parkir kosong
	F.S. Array parkir terisi sesuai keinginan */
	var k kendaraan
	i := 0
	
	fmt.Println("Masukkan data kendaraan")
	fmt.Print("Masukkan jenis kendaraan : ")
	fmt.Scanln(&k.jenis)
	k.jenis = strings.ToLower(k.jenis)
	for k.jenis != "motor" && k.jenis != "mobil" && k.jenis != "truk" {
		fmt.Println("Mohon Maaf Kendaraan Anda Tidak Valid, silakan masukkan ulang jenis kendaraan anda :")
		fmt.Scanln(&k.jenis)
		k.jenis = strings.ToLower(k.jenis)
	}
	fmt.Print("Nomor Polisi: ")
	fmt.Scanln(&k.nopol)
	fmt.Print("Tangggal Masuk(DD MM YY): ")
	fmt.Scanln(&k.tanggal, &k.bulan, &k.tahun)
	for k.tanggal > 31 || k.tanggal <= 0 || k.bulan > 12 || k.bulan <= 0 || (k.bulan == 2 && k.tanggal>28) {
		fmt.Print("tanggal masuk tidak valid, mohon input ulang : ")
		fmt.Scan(&k.tanggal, &k.bulan, &k.tahun)
	}
	fmt.Print("Waktu Masuk(09:00 - 22:00): ")
	fmt.Scanln(&k.jam, &k.menit)
	for (k.jam > 22 || k.jam < 9) || (k.menit >= 60 || k.menit < 0) {
		fmt.Print("waktu masuk tidak valid, mohon input ulang : ")
		fmt.Scan(&k.jam, &k.menit)
	}
	if k.jenis == "motor" {
		fmt.Print("Masukkan nomor parkir yang diinginkan :")
		fmt.Scan(&nomor)
		for nomor < 0 || nomor >= T || mot[nomor].nopol != "" {
			if 	nomor < 0 || nomor >= T{
				fmt.Println("MAAF LOKASI PARKIR TIDAK VALID")
				fmt.Print("Silahkan masukkan lokasi parkir: ")
				fmt.Scanln(&nomor)
			}else if mot[nomor].nopol != ""{
				fmt.Println("PARKIR SUDAH TERISI")
				fmt.Print("Silahkan masukkan lokasi parkir: ")
				fmt.Scan(&nomor)
			}
		}
		mot[nomor] = k
		temp[i] = k
		nMotor++
		sMotor++
	}else if k.jenis == "truk" {
		fmt.Print("Masukkan baris dan kolom parkir yang diinginkan :")
		fmt.Scan(&baris, &kolom)
		for baris < 0 || kolom < 0 || baris >= N || kolom >= M-1 || mob[baris][kolom].nopol != "" {
			if baris < 0 || kolom < 0 || baris >= N || kolom >= M-1 {
				fmt.Println("MAAF LOKASI PARKIR TIDAK VALID")
				fmt.Print("Silahkan masukkan lokasi parkir: ")
				fmt.Scanln(&baris, &kolom)
			}else if mob[baris][kolom].nopol != "" || mob[baris][kolom+1].nopol != "" {
				fmt.Println("Maaf parkir sudah terisi")
				fmt.Print("Masukkan lokasi parkir: ")
				fmt.Scanln(&baris, &kolom)
			}
		}
		if (baris+1) > N && (kolom+1) > M {
			baris = 0
			mob[baris][kolom] = k
			mob[baris][kolom+1] = k
			temp[i] = k
			nTruk++
			sMobil++
		} else {
			mob[baris][kolom] = k
			mob[baris][kolom+1] = k
			temp[i] = k
			nTruk++
			sMobil++
		}
	}else if k.jenis == "mobil" {
		fmt.Print("Masukkan baris dan kolom parkir yang diinginkan :")
		fmt.Scan(&baris, &kolom)
		for baris < 0 || kolom < 0 || baris >= N || kolom >= M || mob[baris][kolom].nopol != ""{
			if baris < 0 || kolom < 0 ||baris >= N || kolom >= M{
				fmt.Println("MAAF LOKASI PARKIR TIDAK VALID")
				fmt.Print("Silahkan masukkan lokasi parkir: ")
				fmt.Scanln(&baris, &kolom)
			}else if mob[baris][kolom].nopol != ""{
				fmt.Println("Maaf parkir sudah terisi")
				fmt.Print("silahkan masukkan lokasi parkir: ")
				fmt.Scanln(&baris, &kolom)
			}
		}
		mob[baris][kolom] = k
		temp[i] = k
		nMobil++
		sMobil++
	}
	fmt.Println("=========================================================================")
	fmt.Println("=			Mall Grand Sukabirus Square			=")
	fmt.Println("=========================================================================")
	fmt.Println("Data Parkir Kendaraan : ")
	fmt.Println("Jenis Kendaraan : ",k.jenis)
	fmt.Println("Nomor Polisi : ",k.nopol)
	fmt.Println("Waktu Masuk : ",k.jam,":",k.menit)
	fmt.Println("Tanggal Masuk : ",k.tanggal,"-",k.bulan,"-",k.tahun)
	fmt.Println("=========================================================================")
	i++
}

func jumlahkendaraan() {
/*	I.S. Berisi data kendaraan didalam tempat parkir
	F.S. Menampilkan total kendaraan yang masuk pada hari ini */
	jumlah := nMotor+nMobil+nTruk
	fmt.Println("jumlah kendaraan yang masuk pada hari ini ada :",jumlah)
	fmt.Println("jumlah motor :",nMotor,"jumlah mobil :",nMobil,"jumlah truk :",nTruk)
}

func okupansi() {
/*	I.S. Berisi data kendaraan didalam tempat parkir
	F.S. Menampilkan total kendaraan di dalam area parkir */
	fmt.Println("okupansi parkir kendaraan pada hari ini ada :")
	fmt.Println("motor :",((float64(sMotor)/T)*100),"%, mobil/roda4++ :",((float64(sMobil)/(N*M))*100),"%")
}

func pendapatan() {
/*	I.S. 
	F.S. Menampilkan hasil pendapatan pada hari ini */
	fmt.Println("pendapatan pada hari ini : ", total)	
}

func cari() {
/*	I.S. Berisi data kendaraan didalam tempat parkir
	F.S. Menampilkan detail dari kendaraan tertentu */
	var nopol string
	var found bool = false
	fmt.Print("Masukkan nomor kendaraan: ")
	fmt.Scanln(&nopol)
	i := 0
	j := 0
	for i < T && found == false {
		if nopol == mot[i].nopol {
			fmt.Println("Kendaraan ditemukan")
			fmt.Println("Data Kendaraan : ")
			fmt.Println("Jenis Kendaraan : ",mot[i].jenis)
			fmt.Println("Nomor Polisi : ",mot[i].nopol)
			fmt.Println("Waktu Masuk : ",mot[i].jam,":",mot[i].menit)
			fmt.Println("Tanggal Masuk : ",mot[i].tanggal,"-",mot[i].bulan,"-",mot[i].tahun)
			found = true
		}
		i++
	}
	for j < N && found == false {
		for k := 0; k < M; k++ {
			if nopol == mob[j][k].nopol {
				fmt.Println("Kendaraan ditemukan")
				fmt.Println("Data Kendaraan : ")
				fmt.Println("Jenis Kendaraan : ",mob[j][k].jenis)
				fmt.Println("Nomor Polisi : ",mob[j][k].nopol)
				fmt.Println("Waktu Masuk : ",mob[j][k].jam,":",mob[j][k].menit)
				fmt.Println("Tanggal Masuk : ",mob[j][k].tanggal,"-",mob[j][k].bulan,"-",mob[j][k].tahun)
				found = true
			}
		}
		j++
	}
	if found == false {
		fmt.Println("Maaf kendaraan yang anda cari tidak ditemukan")
	}

}

func rapiinkuy(){
/*	I.S. Data kendaraan yang ada di parkiran
	F.S. kendaraan terurut agar tidak ada jarak */
	var j, i int
	for i<T{
		j = i+1
		for j < T {
			if mot[i].nopol == "" && mot[j].nopol != "" {
				mot[i] = mot[j]
				mot[j] = kosong
			}
			j++
			}
		i++	
	}
}

func cariDataKendaraan(nomor string, kend *kendaraan, x *int, y *int) {
/*	I.S. 
	F.S. Menampilkan data kendaraan */
	*kend = kosong
	*x = -1
	*y = -1

	found := false

	for i := 0; i < T && !found; i++ {
		if nomor == mot[i].nopol {
			*kend = mot[i]
			*x = i

			found = true
		}
	
	}

	for j := 0; j < N; j++ {
		for k := 0; k < M && !found; k++ {
			if nomor == mob[j][k].nopol {
				*kend = mob[j][k]
				*x = j
				*y = k

				found = true
			}
		}
	}
}

func tarifPrintKeluar(){
/*	I.S. Data kendaraan dalam area parkir
	F.S. Data kendaraan dalam area parkir setelah mengeluarkan suatu kendaraan */
	var k, keluar kendaraan
	var nopolis string
	var durasi, tarif, jamKeluar, menitKeluar, i, j int
	fmt.Println("=========================================================================")
	fmt.Println("=			Mall Grand Sukabirus Square			=")
	fmt.Println("=========================================================================")
	fmt.Print("Masukkan nomor kendaraan yang keluar: ")
	fmt.Scanln(&nopolis)

	cariDataKendaraan(nopolis, &k, &i, &j)

	if k.jenis == "motor"{
		mot[i] = kosong
		sMotor = sMotor - 1
	}else if k.jenis == "mobil"{
		mob[i][j] = kosong
		sMobil = sMobil - 1
	}else if k.jenis == "truk" {
		mob[i][j] = kosong
		mob[i][j+1] = kosong
		sMobil = sMobil - 1
	}
	fmt.Print("masukkan tarif parkir: ")
	fmt.Scanln(&tarif)
	fmt.Print("jam keluar : ")
	fmt.Scanln(&keluar.jam, &keluar.menit)
	
	jamKeluar = (keluar.jam * 60) - (k.jam * 60)
	menitKeluar = keluar.menit - k.menit
	
	durasi = jamKeluar + menitKeluar
	jam := durasi / 60
	menit := durasi % 60
	if menit <= 10 {
		tarif = tarif * jam
	} else {
		tarif = tarif * (jam + 1)
	}
	fmt.Println("Durasi : ", jam,"jam", menit,"menit")
	fmt.Println("tarif parkir : Rp.", tarif)
	total = total + tarif
	printtxt(k.nopol, k.jam, k.menit, keluar.jam, keluar.menit)
}

func cetak(){
/*	I.S. 
	F.S.  */
	var j, min, min2 int
	var t kendaraan
	for i := 0; i < T; i++ {
		min = i
		j = 1
		for j < T {
			if mot[j].nopol < mot[min].nopol && mot[j].nopol != "" {
				min = j
			}
			j++
		}
		t = mot[i]
		mot[i] = mot[min]
		mot[min] = t
	}
	fmt.Println("Parkir Motor: ")
	for k := 0; k < T; k++ {
		fmt.Println(mot[k].jenis)
		fmt.Println(mot[k].nopol)
		fmt.Println(mot[k].jam,":",mot[k].menit)
		fmt.Println(mot[k].tanggal,"-",mot[k].bulan,"-",mot[k].tahun)
		}
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			min = i
			min2 = j
			for m := 0; m < N; m++ {
				for n := 0; n < M; n++ {
					if mob[m][n].nopol < mob[min][min2].nopol {
						min = m
						min2 = n
					}
				}
			}
			t = mob[i][j]
			mob[i][j] = mob[min][min2]
			mob[min][min2] = t
		}
	}
	fmt.Println("Parkir Mobil: ")
	for k := 0; k < N; k++ {
		for l := 0; l < M; l++ {
			fmt.Println(mob[k][l].jenis)
			fmt.Println(mob[k][l].nopol)
			fmt.Println(mob[k][l].jam,":",mob[k][l].menit)
			fmt.Println(mob[k][l].tanggal,"-",mob[k][l].bulan,"-",mob[k][l].tahun)
		}
	}
}

func main () {
	var pilih int
	var bebas, no string
	var mj,mm,kj,km int
	for pilih != 11 {
		fmt.Println("=========================================================================")
		fmt.Println("=			Mall Grand Sukabirus Square			=")
		fmt.Println("=========================================================================")
		fmt.Println("1. Kendaraan Masuk")
		fmt.Println("2. Keluar Parkir")
		fmt.Println("3. Rapikan Motor")
		fmt.Println("4. Mencari Kendaraan dan Output Data")
		fmt.Println("5. Statistik Harian")
		fmt.Println("6. Okupansi Parkir Saat Ini")
		fmt.Println("7. Cetak Nomor Polisi")
		fmt.Println("8. Pendapatan Harian")
		fmt.Println("9. Urutkan berdasarkan jenis dan nopol")
		fmt.Println("10. Print to txt")
		fmt.Println("11. exit")
		fmt.Print("PILIH MENU: ")
		fmt.Scanln(&pilih)

		switch {
		case pilih == 1:
			hapus()
			isiParkir()
		case pilih == 2:
			hapus()
			tarifPrintKeluar()
		case pilih == 3:
			hapus()
			fmt.Println("MERAPIHKAN......")
			rapiinkuy()
		case pilih == 4:
			hapus()
			cari()
		case pilih == 5:
			hapus()
			jumlahkendaraan()
		case pilih == 6:
			hapus()
			okupansi()
		case pilih == 7:
			hapus()
			cetak()
		case pilih == 8:
			hapus()
			pendapatan()
		case pilih == 9:
			hapus()
			printmerge()
		case pilih == 10:
			hapus()
			fmt.Println("PRINTING...... DATA_PARKIR_HARIAN.txt")
			printtxt(no, mj, mm, kj, km)
		default :
			return
		}
		fmt.Print("\nPRESS ANY KEY AND ENTER TO CONTINUE. . . . .")
		fmt.Scan(&bebas)
		main()
		hapus()
	}
}

func printtxt(nopol string, jam, menit, jamKeluar, menitKeluar int) {
/*	I.S. 
	F.S. Mencatat nomor kendaraan, jam masuk dan keluar pada sebuah file txt */
	f, err := os.OpenFile("DATA_PARKIR_HARIAN.txt", os.O_APPEND|os.O_CREATE, 0600) //if FOUND = append file. if !FOUND create file.

	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return
	}
	defer f.Close()
	for i := 0; i < 1; i++ { // Generating...
		_, err = f.WriteString(fmt.Sprintf("nomor kendaraan = %s\njam masuk kendaraan = %02d : %02d\njam keluar kendaraan = %02d : %02d \n\n", nopol, jam, menit, jamKeluar, menitKeluar)) // writing...
		if err != nil {
			fmt.Printf("error writing string: %v", err)
		}
	}
}

func hapus() {
/*	I.S. Layar terminal
	F.S. Menghapus layar terminal */
	a:=exec.Command("cmd","/c","cls")
	a.Stdout=os.Stdout
	a.Run()
}

func printmerge() {
/*	I.S. Berisi data kendaraan didalam tempat parkir
	F.S. Kendaraan terurut berdasarkan jenis dan nomor polisinya */
	for i:=0 ; i<M ; i++ {
		min:=i
		for j:=i+1 ; j<M ; j++ {
			if temp[j].nopol < temp[min].nopol {
				min=j
			}
		}
		tempat:=temp[i]
		temp[i]=temp[min]
		temp[min]=tempat
	}
	for i:=0 ; i<T ; i++ {
		if mot[i].jenis == "motor" && mot[i] != kosong {
			fmt.Println(mot[i]," ")
		}
	}
	for j := 0; j < N ; j++ {
		for k := 0; k < M; k++ {
			if mob[j][k].jenis == "mobil" && mob[j][k] != kosong {
				fmt.Print(mob[j][k]," ")
			} else {
				fmt.Print(mob[j][k]," ")
				k++
			}
		}
	}
	fmt.Println()
} 