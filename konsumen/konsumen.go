package konsumen

import (
	"database/sql"
	"errors"
	"log"
)

type Konsumen struct {
	HP        string
	Nama      string
	IdPegawai int
}

type KonsumMenu struct {
	DB *sql.DB
}

func (km *KonsumMenu) DuplicateKonsumen(name string) bool {
	res := km.DB.QueryRow("SELECT hp_konsumen FROM konsumen where nama = ?", name)
	var idExist int
	err := res.Scan(&idExist)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			log.Println("Result scan error", err.Error())
		}
		return false
	}
	return true
}

func (km *KonsumMenu) RegistKonsumen(newKonsumen Konsumen) (bool, error) {
	// menyiapakn query untuk insert
	registKonsumenQry, err := km.DB.Prepare("INSERT INTO konsumen (hp_konsumen, nama, id_pegawai) values (?,?,?)")
	if err != nil {
		log.Println("prepare insert konsumen ", err.Error())
		return false, errors.New("prepare statement insert konsumen error")
	}

	if km.DuplicateKonsumen(newKonsumen.Nama) {
		log.Println("duplicated information")
		return false, errors.New("nama sudah digunakan")
	}

	// menjalankan query prepare
	res, err := registKonsumenQry.Exec(newKonsumen.HP, newKonsumen.Nama, newKonsumen.IdPegawai)
	if err != nil {
		log.Println("insert konsumen ", err.Error())
		return false, errors.New("insert konsumen error")
	}
	// mengecek jumlah baris yang terpengaruh query diatas
	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after insert konsumen ", err.Error())
		return false, errors.New("error setelah insert konsumen")
	}

	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}

	return true, nil
}

func (km *KonsumMenu) DeleteKonsumen(deleteKonsumen Konsumen) (bool, error) {

	deleteQry, err := km.DB.Prepare("DELETE FROM konsumen WHERE hp_konsumen=?")
	if err != nil {
		log.Println("prepare delete konsumen ", err.Error())
		return false, errors.New("prepare statement delete konsumen error")
	}

	// menjalankan query prepare
	res, err := deleteQry.Exec(deleteKonsumen.HP)
	if err != nil {
		log.Println("delete konsumen ", err.Error())
		return false, errors.New("delete konsumen error")
	}
	// Cek jumlah baris yang terpengaruh query diatas
	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after delete konsumen", err.Error())
		return false, errors.New("error setelah delete konsumen")
	}

	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}

	return true, nil
}

func (km *KonsumMenu) Show() (pelanggan []Konsumen) {
	rows, e := km.DB.Query(
		`SELECT hp_konsumen,
		nama
		FROM konsumen;`)

	if e != nil {
		log.Println(e)
		return
	}

	pelanggan = make([]Konsumen, 0)
	for rows.Next() {
		konsumen := Konsumen{}
		rows.Scan(&konsumen.HP, &konsumen.Nama)
		pelanggan = append(pelanggan, konsumen)
	}
	return pelanggan
}
