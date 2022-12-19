package main

import (
	"bufio"
	"fmt"
	"os"
	"todo/barang"
	"todo/config"
	"todo/user"
)

func main() {
	var inputMenu int = 1
	var cfg = config.ReadConfig()
	var conn = config.ConnectSQL(*cfg)
	var authMenu = user.AuthMenu{DB: conn}
	var barangMenu = barang.BarangMenu{DB: conn}
	for inputMenu != 0 {
		fmt.Println(" ")
		fmt.Println("Selamat Datang di Aplikasi TOKOKU")
		fmt.Println(" ")
		fmt.Println("MENU")
		fmt.Println("===================================")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("0. Exit")
		fmt.Println("===================================")
		fmt.Println("Silakan masukkan pilihan anda : ")
		fmt.Scanln(&inputMenu)
		if inputMenu == 1 {
			var newUser user.User
			in := bufio.NewReader(os.Stdin)
			fmt.Print("Masukkan nama : ")
			name, _ := in.ReadString('\n')
			newUser.Nama = name
			// fmt.Scanln(&newUser.Nama)
			fmt.Print("Masukkan password : ")
			pass, _ := in.ReadString('\n')
			newUser.Password = pass
			// fmt.Scanln(&newUser.Password)
			res, err := authMenu.Register(newUser)
			if err != nil {
				fmt.Println(err.Error())
			}
			if res {
				fmt.Println("Sukses mendaftarkan data")
			} else {
				fmt.Println("Gagal mendaftarn data")
			}
		} else if inputMenu == 2 {
			var inputNama, inputPassword string
			// in := bufio.NewReader(os.Stdin)
			fmt.Println(" ")
			fmt.Println("MENU LOGIN PEGAWAI")
			fmt.Println("=======================")
			fmt.Print("Masukkan nama : ")
			// name, _ := in.ReadString('\n')
			// inputNama = name
			fmt.Scanln(&inputNama)
			fmt.Print("Masukkan password : ")
			// pass, _ := in.ReadString('\n')
			// inputPassword = pass
			fmt.Scanln(&inputPassword)
			res, err := authMenu.Login(inputNama, inputPassword)
			if err != nil {
				fmt.Println(err.Error())
			}

			if res.ID > 0 {
				isLogin := true
				loginMenu := 0
				for isLogin {
					fmt.Println(" ")
					fmt.Println("INVENTORY TOKO")
					fmt.Println("=======================")
					fmt.Println("1. Tambah Barang")
					fmt.Println("9. Logout")
					fmt.Print("Masukkan pilihan anda : ")
					fmt.Scanln(&loginMenu)
					if loginMenu == 1 {
						inputBarang := barang.Barang{}
						fmt.Print("Masukkan Nama Barang : ")
						fmt.Scanln(&inputBarang.Nama)
						fmt.Print("Masukkan Deskripsi Barang : ")
						fmt.Scanln(&inputBarang.Deskripsi)
						fmt.Print("Masukkan Jumlah (Stok) : ")
						fmt.Scanln(&inputBarang.Stok)
						barRes, err := barangMenu.Insert(inputBarang)
						if err != nil {
							fmt.Println(err.Error())
						}
						inputBarang.ID = barRes
						fmt.Println(inputBarang)
					} else if loginMenu == 9 {
						isLogin = false
					}
				}
			}
		}
	}
}
