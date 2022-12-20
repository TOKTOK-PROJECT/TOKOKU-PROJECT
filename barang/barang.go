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

func (bm *BarangMenu) Insert(newBarang Barang) (int, error) {
	insertBarangQry, err := bm.DB.Prepare("INSERT INTO barang (nama_barang, deskripsi, stok, id_pegawai) values (?,?,?,?)")
	if err != nil {
		log.Println("prepare insert barang ", err.Error())
		return 0, errors.New("prepare statement insert barang error")
	}

	// menjalankan query prepare
	res, err := insertBarangQry.Exec(newBarang.Nama, newBarang.Deskripsi, newBarang.Stok, newBarang.Owner)

	if err != nil {
		log.Println("insert barang ", err.Error())
		return 0, errors.New("insert barang error")
	}

	// mengecek jumlah baris yang terpengaruh query diatas
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

func (bm *BarangMenu) Show() (barang2 []Barang) {
	rows, e := bm.DB.Query(
		`SELECT id_barang,
		nama_barang,
		deskripsi,
		stok,
		id_pegawai
		FROM barang;`)

	if e != nil {
		log.Println(e)
		return
	}

	barang2 = make([]Barang, 0)
	for rows.Next() {
		bar := Barang{}
		rows.Scan(&bar.ID, &bar.Nama, &bar.Deskripsi, &bar.Stok, &bar.Owner)
		barang2 = append(barang2, bar)
	}
	return barang2
}

func (bm *BarangMenu) Delete(deleteBarang Barang) (bool, error) {

	registerQry, err := bm.DB.Prepare("DELETE FROM barang WHERE id_barang=?")
	if err != nil {
		log.Println("prepare delete barang ", err.Error())
		return false, errors.New("prepare statement delete barang error")
	}

	// menjalankan query prepare
	res, err := registerQry.Exec(deleteBarang.ID)
	if err != nil {
		log.Println("delete barang ", err.Error())
		return false, errors.New("delete barang error")
	}

	// mengecek jumlah baris yang terpengaruh query diatas
	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after Delete barang ", err.Error())
		return false, errors.New("error setelah delete barang")
	}

	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}

	return true, nil
}

func (bm *BarangMenu) Edit(editBarang Barang) (bool, error) {

	editQry, err := bm.DB.Prepare("UPDATE barang set deskripsi=?  where id_barang= ?")
	if err != nil {
		log.Println("Update barang prepare", err.Error())
		return false, errors.New("prepare Edit barang error")
	}

	// menjalankan query prepare
	res, err := editQry.Exec(editBarang.Deskripsi, editBarang.ID)
	if err != nil {
		log.Println("Update barang", err.Error())
		return false, errors.New("update barang error")
	}

	// mengecek jumlah baris yang terpengaruh query diatas
	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after Update Barang", err.Error())
		return false, errors.New("after Update Barang error")
	}

	if affRows <= 0 {
		log.Println("No record affected")
		return true, errors.New("no record")
	}

	return true, nil
}

func (bm *BarangMenu) UpdateStok(updateBarang Barang) (bool, error) {

	addQry, err := bm.DB.Prepare("UPDATE barang set stok=? where id_barang= ?")
	if err != nil {
		log.Println("Update stok barang prepare", err.Error())
		return false, errors.New("prepare update stok barang error")
	}

	// menjalankan query prepare
	res, err := addQry.Exec(updateBarang.Stok, updateBarang.ID)
	if err != nil {
		log.Println("Update barang", err.Error())
		return false, errors.New("update barang error")
	}

	// mengecek jumlah baris yang terpengaruh query diatas
	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after Update Barang", err.Error())
		return false, errors.New("after Update Barang error")
	}

	if affRows <= 0 {
		log.Println("No record affected")
		return true, errors.New("no record")
	}

	return true, nil
}
