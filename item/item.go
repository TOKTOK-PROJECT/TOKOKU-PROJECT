package item

import (
	"database/sql"
	"errors"
	"log"
)

type Item struct {
	NoNota    int
	IdBarang  int
	Kuantitas int
}

type ItemMenu struct {
	DB *sql.DB
}

func (it *ItemMenu) AddItem(newItem Item) (bool, error) {
	// menyiapakn query untuk insert
	tambahItemQry, err := it.DB.Prepare("INSERT INTO transaksi (id_barang, kuantitas) values (?,?)")
	if err != nil {
		log.Println("prepare insert konsumen ", err.Error())
		return false, errors.New("prepare statement insert konsumen error")
	}

	// menjalankan query prepare
	res, err := tambahItemQry.Exec(newItem.IdBarang, newItem.Kuantitas)
	if err != nil {
		log.Println("tambah item ", err.Error())
		return false, errors.New("tambah item error")
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

func (it *ItemMenu) DeleteTransaksi(DeleteItem Item) (bool, error) {

	deleteQry, err := it.DB.Prepare("DELETE FROM transaksi WHERE no_nota=?")
	if err != nil {
		log.Println("prepare delete transaksi ", err.Error())
		return false, errors.New("prepare statement delete transaksi error")
	}

	// menjalankan query prepare
	res, err := deleteQry.Exec(DeleteItem.NoNota)
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

func (it *ItemMenu) Show() (item []Item) {
	rows, e := it.DB.Query(
		`SELECT no_nota,
		id_barang,
		kuantitas
		FROM item;`)

	if e != nil {
		log.Println(e)
		return
	}

	item = make([]Item, 0)
	for rows.Next() {
		ite := Item{}
		rows.Scan(&ite.NoNota, &ite.IdBarang, &ite.Kuantitas)
		item = append(item, ite)
	}
	return item
}
