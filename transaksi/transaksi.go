package transaksi

import (
	"database/sql"
	"errors"
	"log"
)

type Transaksi struct {
	NoNota     int
	IdPegawai  int
	HpKonsumen string
	IdBarang   int
	Kuantitas  int
	Tanggal    string
}

type TransaksiMenu struct {
	DB *sql.DB
}

func (tm *TransaksiMenu) AddTransaksi(newTransaksi Transaksi) (bool, error) {
	// menyiapakn query untuk insert
	tambahTransaksiQry, err := tm.DB.Prepare("INSERT INTO transaksi (id_pegawai, hp_konsumen, id_barang, kuantitas) values (?,?,?,?)")
	if err != nil {
		log.Println("prepare insert konsumen ", err.Error())
		return false, errors.New("prepare statement insert konsumen error")
	}

	// menjalankan query prepare
	res, err := tambahTransaksiQry.Exec(newTransaksi.IdPegawai, newTransaksi.HpKonsumen, newTransaksi.IdBarang, newTransaksi.Kuantitas)
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

func (tm *TransaksiMenu) Show() (transaksi []Transaksi) {
	rows, e := tm.DB.Query(
		`SELECT no_nota,
		id_pegawai,
		hp_konsumen,
		id_barang,
		kuantitas,
		tanggal_cetak
		FROM transaksi;`)

	if e != nil {
		log.Println(e)
		return
	}

	transaksi = make([]Transaksi, 0)
	for rows.Next() {
		trans := Transaksi{}
		rows.Scan(&trans.NoNota, &trans.IdPegawai, &trans.HpKonsumen, &trans.IdBarang, &trans.Kuantitas, &trans.Tanggal)
		transaksi = append(transaksi, trans)
	}
	return transaksi
}
