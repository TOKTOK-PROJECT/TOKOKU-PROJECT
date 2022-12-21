package main

import (
	"bufio"
	"fmt"
	"os"
	"todo/barang"
	"todo/config"
	"todo/konsumen"
	"todo/user"
)

func main() {
	var inputMenu int = 1
	var cfg = config.ReadConfig()
	var conn = config.ConnectSQL(*cfg)
	var authMenu = user.AuthMenu{DB: conn}
	var barangMenu = barang.BarangMenu{DB: conn}
	var konsumMenu = konsumen.KonsumMenu{DB: conn}

	for inputMenu != 0 {
		fmt.Println("-- Selamat Datang di Aplikasi TOKOKU --")
		fmt.Println("\n================= MENU =================")
		fmt.Println("1. Login")
		fmt.Println("0. Exit")
		fmt.Println("========================================")
		fmt.Println("Silakan masukkan pilihan anda : ")
		fmt.Scanln(&inputMenu)
		if inputMenu == 1 {
			var inputNama, inputPassword string
			in := bufio.NewReader(os.Stdin)
			fmt.Println("\n------- MENU LOGIN -------")
			fmt.Println("==========================")
			fmt.Print("Masukkan nama : ")
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

			if res.ID == 1 && inputNama == "admin" && inputPassword == "admin" {
				isLogin := true
				fmt.Println("\n--- Login sebagai Admin ---")
				for isLogin {
					loginMenu := 0
					fmt.Println("\n------- MENU ADMIN -------")
					fmt.Println("==========================")
					fmt.Println("1. Tambah Pegawai")
					fmt.Println("2. Hapus Pegawai")
					fmt.Println("3. Hapus Barang")
					fmt.Println("4. Hapus Pelanggan")
					fmt.Println("5. Hapus Transaksi")
					fmt.Println("9. Logout")
					fmt.Println("==========================")
					fmt.Println("Silakan masukkan pilihan:")
					fmt.Scanln(&loginMenu)
					switch loginMenu {
					case 1:
						fmt.Println("\n--- Halaman Tambah Pegawai ---")
						fmt.Println("==============================")
						var newUser user.User
						in := bufio.NewReader(os.Stdin)
						fmt.Print("Masukkan nama : ")
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

						fmt.Println("\n--- DAFTAR PEGAWAI ---")
						fmt.Println("======================")
						fmt.Println(authMenu.Show())
					case 2:
						var deleteUser user.User
						fmt.Println("\n--- Halaman Hapus Pegawai ---")
						fmt.Println("=============================")
						fmt.Println("masukkan ID pegawai yang ingin dihapus :")
						fmt.Scanln(&deleteUser.ID)
						res, err := authMenu.Delete(deleteUser)
						if err != nil {
							fmt.Println(err.Error())
						}
						if res {
							fmt.Println("Sukses menghapus Pegawai")
						} else {
							fmt.Println("Gagal menghapus Pegawai")
						}

						fmt.Println("\n--- DAFTAR PEGAWAI ---")
						fmt.Println("======================")
						fmt.Println(authMenu.Show())
					case 3:
						var deleteBarang barang.Barang
						fmt.Println("\n--- Halaman Hapus Barang ---")
						fmt.Println("=============================")
						fmt.Println("masukkan ID barang yang ingin dihapus :")
						fmt.Scanln(&deleteBarang.ID)
						res, err := barangMenu.Delete(deleteBarang)
						if err != nil {
							fmt.Println(err.Error())
						}
						if res {
							fmt.Println("Sukses menghapus Barang")
						} else {
							fmt.Println("Gagal menghapus Barang")
						}

						fmt.Println("\n--- DAFTAR BARANG ---")
						fmt.Println("=====================")
						fmt.Println(barangMenu.Show())
					case 4:
						var deleteKonsumen konsumen.Konsumen
						fmt.Println("\n--- Halaman Hapus Pelanggan ---")
						fmt.Println("===============================")
						fmt.Println("masukkan nomor HP pelanggan yang ingin dihapus :")
						fmt.Scanln(&deleteKonsumen.HP)
						res, err := konsumMenu.DeleteKonsumen(deleteKonsumen)
						if err != nil {
							fmt.Println(err.Error())
						}
						if res {
							fmt.Println("Sukses menghapus data pelanggan")
						} else {
							fmt.Println("Gagal menghapus data pelanggan")
						}

						fmt.Println("\n--- DAFTAR PELANGGAN ---")
						fmt.Println("========================")
						fmt.Println(konsumMenu.Show())
					case 5:

					case 9:
						isLogin = false
					}
				}

			}

			if res.ID > 0 && inputNama != "admin" {
				isLogin := true
				fmt.Println("\n--- Login Sebagai Pegawai ---")
				for isLogin {
					loginMenu := 0
					fmt.Println("------ MENU PEGAWAI ------")
					fmt.Println("==========================")
					fmt.Println("1. Tambah Barang")
					fmt.Println("2. Lihat Barang")
					fmt.Println("3. Edit Deskripsi")
					fmt.Println("4. Update Stok Barang")
					fmt.Println("5. Tambah Pelanggan")
					fmt.Println("6. Transaksi")
					fmt.Println("9. Logout")
					fmt.Print("Masukkan pilihan anda : ")
					fmt.Scanln(&loginMenu)
					switch loginMenu {
					case 1:
						inputBarang := barang.Barang{}
						fmt.Println("\n--- Halaman Tambah Barang ---")
						fmt.Println("=============================")
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
					case 2:
						fmt.Println("\n--- DAFTAR BARANG ---")
						fmt.Println("=====================")
						fmt.Println(barangMenu.Show())
					case 3:
						var editBarang barang.Barang
						in := bufio.NewReader(os.Stdin)
						fmt.Println("\n--- Halaman Edit Deskripsi Barang ---")
						fmt.Println("=====================================")
						fmt.Println("masukkan ID barang yang deskripsinya akan diedit :")
						fmt.Scanln(&editBarang.ID)
						fmt.Println("masukkan Deskripsi terbaru :")
						desc, _ := in.ReadString('\n')
						desc = desc[:len(desc)-2]
						editBarang.Deskripsi = desc
						res, err := barangMenu.Edit(editBarang)
						if err != nil {
							fmt.Println(err.Error())
						}
						if res {
							fmt.Println("Sukses Mengedit Deskripsi Barang")
						} else {
							fmt.Println("Gagal Mengedit Deskripsi Barang")
						}
						fmt.Println(editBarang)
					case 4:
						fmt.Println("\n--- Halaman Update Stok Barang ---")
						fmt.Println("==================================")
						var updateStok barang.Barang
						fmt.Println("masukkan ID barang yang akan diedit :")
						fmt.Scanln(&updateStok.ID)
						fmt.Println("masukkan jumlah stok terbaru")
						fmt.Scanln(&updateStok.Stok)

						res, err := barangMenu.UpdateStok(updateStok)
						if err != nil {
							fmt.Println(err.Error())
						}
						if res {
							fmt.Println("Sukses update stok barang")
						} else {
							fmt.Println("Gagal update stok barang")
						}
						fmt.Println(updateStok)
					case 5:
						var newKonsumen konsumen.Konsumen
						in := bufio.NewReader(os.Stdin)
						fmt.Println("\n--- Halaman Tambah Pelanggan ---")
						fmt.Println("================================")
						fmt.Print("Masukkan nama : ")
						name, _ := in.ReadString('\n')
						name = name[:len(name)-2]
						newKonsumen.Nama = name
						fmt.Print("Masukkan nomor telepon : ")
						hp, _ := in.ReadString('\n')
						hp = hp[:len(hp)-2]
						newKonsumen.HP = hp
						// fmt.Print("Masukkan ID Pegawai : ")
						// idpeg, _ := in.ReadString('\n')
						// idpeg = idpeg[:len(idpeg)-2]
						newKonsumen.IdPegawai = res.ID
						res, err := konsumMenu.RegistKonsumen(newKonsumen)
						if err != nil {
							fmt.Println(err.Error())
						}
						if res {
							fmt.Println("Sukses mendaftarkan pelanggan")
						} else {
							fmt.Println("Gagal mendaftarn pelanggan")
						}

					case 6:

					case 9:
						isLogin = false
					}
				}
			}
		}
	}
}
