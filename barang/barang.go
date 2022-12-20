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

//Update db
//  func (bm *BarangMenu) Update(id int, stok int) () {
// 	stmt, e := bm.DB.Prepare("UPDATE barang SET stok=? where id=?")
// 	ErrorCheck(e)

// 	// execute
// 	res, e := stmt.Exec("This is post five", "5")
// 	ErrorCheck(e)

// 	a, e := res.RowsAffected()
// 	ErrorCheck(e)

// 	fmt.Println(a)
//  }

// query all data
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

// delete data
//  func (bm *BarangMenu) Delete(id int) (int, error) {
// 	stmt, e := bm.DB.Prepare("delete from posts where id=?")
// 	ErrorCheck(e)

// 	// delete 5th post
// 	res, e := stmt.Exec("5")
// 	ErrorCheck(e)

// 	// affected rows
// 	a, e := res.RowsAffected()
// 	ErrorCheck(e)

// 	fmt.Println(a) // 1
//    }

//    func ErrorCheck(err error) {
// 	if err != nil {
// 		panic(err.Error())
// 	}
//    }
