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

func (im *ItemMenu) Insert(newItem Item) (bool, error) {
	insertItemQry, err := im.DB.Prepare("INSERT INTO item (no_nota, id_barang, kuantitas) values (?,?,?)")
	if err != nil {
		log.Println("prepare insert item barang ", err.Error())
		return false, errors.New("prepare statement insert item barang error")
	}

	// menjalankan query prepare
	res, err := insertItemQry.Exec(newItem.NoNota, newItem.IdBarang, newItem.Kuantitas)

	if err != nil {
		log.Println("insert item barang ", err.Error())
		return false, errors.New("insert item barang error")
	}

	// mengecek jumlah baris yang terpengaruh query diatas
	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after insert barang ", err.Error())
		return false, errors.New("error setelah insert item barang")
	}

	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}

	return true, nil
}
