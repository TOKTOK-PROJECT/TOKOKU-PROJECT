package user

import (
	"database/sql"
	"errors"
	"log"
)

type User struct {
	ID       int
	Nama     string
	Password string
}

type AuthMenu struct {
	DB *sql.DB
}

func (am *AuthMenu) Duplicate(name string) bool {
	res := am.DB.QueryRow("SELECT id_pegawai FROM pegawai where nama = ?", name)
	var idExist int
	err := res.Scan(&idExist)
	if err != nil {
		log.Println("Result scan error", err.Error())
		return false
	}
	return true
}

func (am *AuthMenu) Register(newUser User) (bool, error) {
	// menyiapakn query untuk insert
	registerQry, err := am.DB.Prepare("INSERT INTO pegawai (nama, password) values (?,?)")
	if err != nil {
		log.Println("prepare insert user ", err.Error())
		return false, errors.New("prepare statement insert user error")
	}

	if am.Duplicate(newUser.Nama) {
		log.Println("duplicated information")
		return false, errors.New("nama sudah digunakan")
	}

	// menjalankan query dengan parameter tertentu
	res, err := registerQry.Exec(newUser.Nama, newUser.Password)
	if err != nil {
		log.Println("insert user ", err.Error())
		return false, errors.New("insert user error")
	}
	// Cek berapa baris yang terpengaruh query diatas
	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after insert user ", err.Error())
		return false, errors.New("error setelah insert")
	}

	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}

	return true, nil
}

func (am *AuthMenu) Login(nama string, password string) (User, error) {
	loginQry, err := am.DB.Prepare("SELECT id_pegawai FROM pegawai WHERE nama = ? AND password = ?")
	if err != nil {
		log.Println("prepare insert user ", err.Error())
		return User{}, errors.New("prepare statement insert user error")
	}

	row := loginQry.QueryRow(nama, password)

	if row.Err() != nil {
		log.Println("login query ", row.Err().Error())
		return User{}, errors.New("tidak bisa login, data tidak ditemukan")
	}
	res := User{}
	err = row.Scan(&res.ID)

	if err != nil {
		log.Println("after login query ", err.Error())
		return User{}, errors.New("tidak bisa login, kesalahan setelah error")
	}

	res.Nama = nama

	return res, nil
}

func (am *AuthMenu) Delete(deleteUser User) (bool, error) {

	deleteQry, err := am.DB.Prepare("DELETE FROM pegawai WHERE id_pegawai=?")
	if err != nil {
		log.Println("prepare delete pegawai ", err.Error())
		return false, errors.New("prepare statement delete pegawai error")
	}

	// menjalankan query prepare
	res, err := deleteQry.Exec(deleteUser.ID)
	if err != nil {
		log.Println("delete pegawai ", err.Error())
		return false, errors.New("delete pegawai error")
	}
	// Cek jumlah baris yang terpengaruh query diatas
	affRows, err := res.RowsAffected()

	if err != nil {
		log.Println("after Delete pegawai", err.Error())
		return false, errors.New("error setelah Delete")
	}

	if affRows <= 0 {
		log.Println("no record affected")
		return false, errors.New("no record")
	}

	return true, nil
}

func (am *AuthMenu) Show() (user2 []User) {
	rows, e := am.DB.Query(
		`SELECT id_pegawai,
		nama
		FROM pegawai;`)

	if e != nil {
		log.Println(e)
		return
	}

	user2 = make([]User, 0)
	for rows.Next() {
		user := User{}
		rows.Scan(&user.ID, &user.Nama)
		user2 = append(user2, user)
	}
	return user2
}
