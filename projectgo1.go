package main

import (
	"log"
	"net/http"
	pnamhome "projectgo1/pNamHome"
)

func main() {
	wadweb := http.NewServeMux()

	wadweb.HandleFunc("/", pnamhome.HHome) //http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {fmt.Fprintln(rw, "Ini Halaman Utama")})
	wadweb.HandleFunc("/login", pnamhome.HLogin)
	wadweb.HandleFunc("/Koin", pnamhome.HlKoin)
	wadweb.HandleFunc("/Wallet", pnamhome.HWallet)
	wadweb.HandleFunc("/Data", pnamhome.HData)
	wadweb.HandleFunc("/TradingView", pnamhome.HTrading)
	wadweb.HandleFunc("/Histori", pnamhome.HHistori)
	wadweb.HandleFunc("/Laporan_Keuangan", pnamhome.HLap)
	wadweb.HandleFunc("/terimakasih", pnamhome.Kirtrim)
	wadweb.HandleFunc("/formulir", pnamhome.Formulir)
	wadweb.HandleFunc("/olah", pnamhome.Olah)

	filserver := http.FileServer(http.Dir("azzetz"))                  //load folder css
	wadweb.Handle("/cicing/", http.StripPrefix("/cicing", filserver)) //mengarahkan ke folder static

	log.Println("Menjalankan web di 8080")
	err := http.ListenAndServe(":8080", wadweb) //http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

/*deklarasi
type wltKoin struct {
	//koinS
	whKoin, wSaldo   int32
	wSAkoin, wSBkoin float32
}

type dkoins struct {
	tgl, bln, tahun string
	//koinS
	priksa                                                                bool
	stus                                                                  string
	djml, dpersen                                                         float32
	dhrgBeliIDR, dvIDR1, dhrgTblIDR, dvIDR2, dhrgJualIDR, dvIDR3, dPropit int32
}

var waktu time.Time
var k koinS
var w wltKoin
var d dkoins
var jdltblData = []string{"No", "Kode Koin", "Tanggal", "Jumlah Koin", "Harga Beli", "Nilai Rupiah Beli (IDR)", "Harga Tabel", "Nilai Rupiah Tabel (IDR)", "Harga Jual", "Nilai Rupiah Jual (IDR)", "Profit", "Persentase", "Status", "Periksa"}
var jdltblKoin = []string{jdltblData[0], jdltblData[1], "Nama Koin", "Market", "Harga Koin"}
var jdltblDompet = []string{jdltblKoin[0], jdltblKoin[1], jdltblKoin[2], jdltblKoin[3], jdltblKoin[4], "Saldo Aktif", "Saldo Beku", "Nilai Saldo"}
*/
