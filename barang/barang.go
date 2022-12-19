package barang

import (
	"database/sql"
	"errors"
	"log"
)

type Barang struct {
	ID        int
	Nama      string
	Deskripsi string
	Stok      int
	Owner     int
}

type BarangMenu struct {
	DB *sql.DB
}

func (am *BarangMenu) Insert(newBarang Barang) (int, error) {
	insertBarangQry, err := am.DB.Prepare("INSERT INTO barang (nama_barang, deskripsi, stok, id_pegawai) values (?,?,?,?)")
	if err != nil {
		log.Println("prepare insert barang ", err.Error())
		return 0, errors.New("prepare statement insert barang error")
	}

	res, err := insertBarangQry.Exec(newBarang.Nama, newBarang.Deskripsi, newBarang.Stok, newBarang.Owner)

	if err != nil {
		log.Println("insert barang ", err.Error())
		return 0, errors.New("insert barang error")
	}

	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after insert barang ", err.Error())
		return 0, errors.New("error setelah insert barang")
	}

	if affRows <= 0 {
		log.Println("no record affected")
		return 0, errors.New("no record")
	}

	id, _ := res.LastInsertId()

	return int(id), nil
}
