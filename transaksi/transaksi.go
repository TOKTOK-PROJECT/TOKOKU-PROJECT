package transaksi

import (
	"database/sql"
	"errors"
	"log"
)

type Transaksi struct {
	NoNota      int
	IdPegawai   int
	HpKonsumen  string
	Tanggal     string
	NamaPembeli string
	NamaPegawai string
	NamaBarang  string
	Kuantitas   int
}

type TransaksiMenu struct {
	DB *sql.DB
}

func (tm *TransaksiMenu) AddTransaksi(newTransaksi Transaksi) (bool, error) {
	// menyiapakn query untuk insert
	tambahTransaksiQry, err := tm.DB.Prepare("INSERT INTO transaksi (id_pegawai, hp_konsumen) values (?,?)")
	if err != nil {
		log.Println("prepare insert konsumen ", err.Error())
		return false, errors.New("prepare statement insert konsumen error")
	}

	// menjalankan query prepare
	res, err := tambahTransaksiQry.Exec(newTransaksi.IdPegawai, newTransaksi.HpKonsumen)
	if err != nil {
		log.Println("tambah transaksi ", err.Error())
		return false, errors.New("tambah transaksi error")
	}
	// mengecek jumlah baris yang terpengaruh query diatas
	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after tambah transaksi ", err.Error())
		return false, errors.New("error setelah tambah transaksi")
	}

	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}

	return true, nil
}

func (tm *TransaksiMenu) Show() (transaksi []Transaksi) {
	rows, e := tm.DB.Query(
		`SELECT no_nota,
		id_pegawai,
		hp_konsumen,
		tanggal_cetak
		FROM transaksi;`)

	if e != nil {
		log.Println(e)
		return
	}

	transaksi = make([]Transaksi, 0)
	for rows.Next() {
		trans := Transaksi{}
		rows.Scan(&trans.NoNota, &trans.IdPegawai, &trans.HpKonsumen, &trans.Tanggal)
		transaksi = append(transaksi, trans)
	}
	return transaksi
}

func (tm *TransaksiMenu) DeleteTransaksi(DeleteTransaksi Transaksi) (bool, error) {

	deleteQry, err := tm.DB.Prepare("DELETE FROM transaksi WHERE no_nota=?")
	if err != nil {
		log.Println("prepare delete transaksi ", err.Error())
		return false, errors.New("prepare statement delete transaksi error")
	}

	// menjalankan query prepare
	res, err := deleteQry.Exec(DeleteTransaksi.NoNota)
	if err != nil {
		log.Println("delete transaksi ", err.Error())
		return false, errors.New("delete transaksi error")
	}
	// Cek jumlah baris yang terpengaruh query diatas
	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after delete transaksi", err.Error())
		return false, errors.New("error setelah delete transaksi")
	}

	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}

	return true, nil
}

func (tm *TransaksiMenu) Cetak(newCetak Transaksi) ([]Transaksi, error) {
	// menyiapakn query untuk insert
	cetakTransaksiQry, err := tm.DB.Prepare(
		`SELECT t.tanggal_cetak,  k.nama, p.nama, b.nama_barang, i.kuantitas 
		FROM item i 
		JOIN barang b  on b.id_barang  = i.id_barang  
		JOIN transaksi t  on t.no_nota  = i.no_nota  
		JOIN pegawai p  on p.id_pegawai  = t.id_pegawai 
		JOIN konsumen k on k.hp_konsumen = t.hp_konsumen
		where i.no_nota = ?;`)
	if err != nil {
		log.Println("prepare cetak transaksi ", err.Error())
		return nil, errors.New("select cetak transaksi error")
	}

	// menjalankan query prepare
	rows, err := cetakTransaksiQry.Query(newCetak.NoNota)
	if err != nil {
		log.Println("cetak transaksi ", err.Error())
		return nil, errors.New("select cetak transaksi error")
	}

	transaksi := []Transaksi{}
	for rows.Next() {
		trans := Transaksi{}
		err = rows.Scan(&trans.Tanggal, &trans.NamaPembeli, &trans.NamaPegawai, &trans.NamaBarang, &trans.Kuantitas)
		if err != nil {
			log.Println("error Loop baris untuk memasukkan data ke struct", err.Error())
			return transaksi, err
		}
		transaksi = append(transaksi, trans)
	}
	return transaksi, nil
}

// func (tm *TransaksiMenu) Pembeli(newCetak Transaksi) (transaksi []Transaksi) {
// 	// menyiapakn query untuk insert
// 	cetakTransaksiQry, err := tm.DB.Prepare(
// 		`SELECT k.nama
// 		FROM transaksi t
// 		join konsumen k on t.hp_konsumen = k.hp_konsumen
// 		WHERE t.no_nota = ?;`)
// 	if err != nil {
// 		log.Println("prepare insert konsumen ", err.Error())
// 		return
// 	}

// 	// menjalankan query prepare
// 	rows, err := cetakTransaksiQry.Query(newCetak.NoNota)
// 	if err != nil {
// 		log.Println("cetak transaksi ", err.Error())
// 		return
// 	}

// 	transaksi = make([]Transaksi, 0)
// 	for rows.Next() {
// 		trans := Transaksi{}
// 		rows.Scan(&trans.NamaPembeli)
// 		transaksi = append(transaksi, trans)
// 	}
// 	return transaksi
// }

// func (tm *TransaksiMenu) Kasir(newCetak Transaksi) (transaksi []Transaksi) {
// 	// menyiapakn query untuk insert
// 	cetakTransaksiQry, err := tm.DB.Prepare(
// 		`SELECT p.nama
// 		FROM transaksi t
// 		join pegawai p on t.id_pegawai = p.id_pegawai
// 		where t.no_nota = ?;`)
// 	if err != nil {
// 		log.Println("prepare insert konsumen ", err.Error())
// 		return
// 	}

// 	// menjalankan query prepare
// 	rows, err := cetakTransaksiQry.Query(newCetak.NoNota)
// 	if err != nil {
// 		log.Println("cetak transaksi ", err.Error())
// 		return
// 	}

// 	for rows.Next() {
// 		trans := Transaksi{}
// 		rows.Scan(&trans.NamaPegawai)
// 		transaksi = append(transaksi, trans)
// 	}
// 	return transaksi
// }
