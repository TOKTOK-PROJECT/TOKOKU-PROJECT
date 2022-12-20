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
		fmt.Println("-- Selamat Datang di Aplikasi TOKOKU --")
		fmt.Println(" ")
		fmt.Println("================= MENU =================")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("0. Exit")
		fmt.Println("========================================")
		fmt.Println("Silakan masukkan pilihan anda : ")
		fmt.Scanln(&inputMenu)
		if inputMenu == 1 {
			var newUser user.User
			in := bufio.NewReader(os.Stdin)
			fmt.Print("Masukkan nama pegawai : ")
			name, _ := in.ReadString('\n')
			name = name[:len(name)-2]
			newUser.Nama = name
			fmt.Print("Masukkan password : ")
			pass, _ := in.ReadString('\n')
			pass = pass[:len(pass)-2]
			newUser.Password = pass
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
			in := bufio.NewReader(os.Stdin)
			fmt.Println(" ")
			fmt.Println("==========MENU LOGIN PEGAWAI============")
			fmt.Println(" ")
			fmt.Print("Masukkan nama anda : ")
			name, _ := in.ReadString('\n')
			name = name[:len(name)-2]
			inputNama = name
			fmt.Print("Masukkan password : ")
			pass, _ := in.ReadString('\n')
			pass = pass[:len(pass)-2]
			inputPassword = pass
			res, err := authMenu.Login(inputNama, inputPassword)
			if err != nil {
				fmt.Println(err.Error())
			}

			if res.ID > 0 {
				isLogin := true
				loginMenu := 0
				for isLogin {
					fmt.Println(" ")
					fmt.Println("============ INVENTORY TOKO ============")
					fmt.Println(" ")
					fmt.Println("1. Tambah Barang")
					fmt.Println("2. Lihat Barang")
					fmt.Println("9. Logout")
					fmt.Print("Masukkan pilihan anda : ")
					fmt.Scanln(&loginMenu)
					if loginMenu == 1 {
						inputBarang := barang.Barang{}
						in := bufio.NewReader(os.Stdin)
						fmt.Print("Masukkan Nama Barang : ")
						name, _ := in.ReadString('\n')
						name = name[:len(name)-2]
						inputBarang.Nama = name
						fmt.Print("Masukkan Deskripsi Barang : ")
						desc, _ := in.ReadString('\n')
						desc = desc[:len(desc)-2]
						inputBarang.Deskripsi = desc
						fmt.Print("Masukkan Jumlah (Stok) : ")
						fmt.Scanln(&inputBarang.Stok)
						inputBarang.Owner = res.ID
						barRes, err := barangMenu.Insert(inputBarang)
						if err != nil {
							fmt.Println(err.Error())
						}
						inputBarang.ID = barRes
						fmt.Println(inputBarang)
					} else if loginMenu == 2 {
						fmt.Println(barangMenu.Show())
					} else if loginMenu == 9 {
						isLogin = false
					}
				}
			}
		}
	}
}
