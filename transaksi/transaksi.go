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
	Tanggal    string
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

func (tm *TransaksiMenu) CetakNota(newCetak Transaksi) (bool, error) {
	addQry, err := tm.DB.Prepare("SELECT item.Kuantitas, item.IdBarang,  item.NoNota, konsumen.HP as konsumen, transaksi.tanggal_cetak,transaksi.id_pegawai,user.nama as user FROM transaksi INNER JOIN konsumen on konsumen.HP = transaksi.NoNota_konsumen INNER JOIN pegawai user on transaksi.id_pegawai = user.id Left join item on item.transaksi_NoNota = transaksi.NoNota WHERE transaksi.HP_konsumen = ?")
	if err != nil {
		log.Println("Select Cetak prepare", err.Error())
		return false, errors.New("prepare Select Cetak error")
	}

	res, err := addQry.Exec(newCetak.HpKonsumen)
	if err != nil {
		log.Println("Select cetak", err.Error())
		return false, errors.New("Select Cetak error")
	}

	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after Select Cetak", err.Error())
		return false, errors.New("after Select cetak error")
	}

	if affRows <= 0 {
		log.Println("No record affected")
		return true, errors.New("No record")
	}

	return true, nil
}
