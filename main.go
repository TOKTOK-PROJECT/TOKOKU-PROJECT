package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"todo/barang"
	"todo/config"
	"todo/item"
	"todo/konsumen"
	"todo/transaksi"
	"todo/user"
)

func screenClear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func colorText(message string, color string) string {
	var colorCode string
	switch color {
	case "red":
		colorCode = "31"
	case "green":
		colorCode = "32"
	case "yellow":
		colorCode = "33"
	case "blue":
		colorCode = "34"
	case "magenta":
		colorCode = "35"
	case "cyan":
		colorCode = "36"
	default:
		colorCode = "0" // reset to default color
	}
	return fmt.Sprintf("\033[%sm%s\033[0m", colorCode, message)

	// fmt.Println(colorText("Hello, World!", "red"))
}

func main() {
	var cfg = config.ReadConfig()
	var conn = config.ConnectSQL(*cfg)
	var authMenu = user.AuthMenu{DB: conn}
	var barangMenu = barang.BarangMenu{DB: conn}
	var konsumMenu = konsumen.KonsumMenu{DB: conn}
	var transaksiMenu = transaksi.TransaksiMenu{DB: conn}
	var itemMenu = item.ItemMenu{DB: conn}

	var inputMenu int = 1

	for inputMenu != 0 {
		fmt.Println(colorText("-- Selamat Datang di Aplikasi TOKOKU --", "cyan"))
		fmt.Println(colorText("\n================= MENU =================", "green"))
		fmt.Println(colorText("1. Login", "yellow"))
		fmt.Println(colorText("0. Exit", "yellow"))
		fmt.Println(colorText("========================================", "green"))
		fmt.Print(colorText("Silakan masukkan pilihan anda : ", "green"))
		fmt.Scanln(&inputMenu)
		screenClear()
		if inputMenu == 1 {
			var inputNama, inputPassword string
			in := bufio.NewReader(os.Stdin)
			fmt.Println(colorText("\n------- MENU LOGIN -------", "green"))
			fmt.Println(colorText("==========================", "green"))
			fmt.Print(colorText("Masukkan nama : ", "yellow"))
			name, _ := in.ReadString('\n')
			name = name[:len(name)-2]
			inputNama = name
			fmt.Print(colorText("Masukkan password : ", "yellow"))
			pass, _ := in.ReadString('\n')
			pass = pass[:len(pass)-2]
			inputPassword = pass
			res, err := authMenu.Login(inputNama, inputPassword)
			if err != nil {
				fmt.Println(err.Error())
			}
			screenClear()
			if res.ID == 1 && inputNama == "admin" && inputPassword == "admin" {
				isLogin := true
				fmt.Println(colorText("\n--- Login sebagai Admin ---", "cyan"))
				for isLogin {
					loginMenu := 0
					fmt.Println(colorText("\n------- MENU ADMIN -------", "green"))
					fmt.Println(colorText("==========================", "green"))
					fmt.Println(colorText("1. Tambah Pegawai", "yellow"))
					fmt.Println(colorText("2. Hapus Pegawai", "yellow"))
					fmt.Println(colorText("3. Hapus Barang", "yellow"))
					fmt.Println(colorText("4. Hapus Pelanggan", "yellow"))
					fmt.Println(colorText("5. Hapus Transaksi", "yellow"))
					fmt.Println(colorText("9. Logout", "yellow"))
					fmt.Println(colorText("==========================", "green"))
					fmt.Print(colorText("Silakan masukkan pilihan : ", "green"))
					fmt.Scanln(&loginMenu)
					screenClear()
					switch loginMenu {
					case 1:
						fmt.Println(colorText("\n--- Halaman Tambah Pegawai ---", "green"))
						fmt.Println(colorText("==============================", "green"))
						var newUser user.User
						in := bufio.NewReader(os.Stdin)
						fmt.Print(colorText("Masukkan nama : ", "yellow"))
						name, _ := in.ReadString('\n')
						name = name[:len(name)-2]
						newUser.Nama = name
						fmt.Print(colorText("Masukkan password : ", "yellow"))
						pass, _ := in.ReadString('\n')
						pass = pass[:len(pass)-2]
						newUser.Password = pass
						res, err := authMenu.Register(newUser)
						if err != nil {
							fmt.Println(err.Error())
						}
						if res {
							fmt.Println(colorText("Sukses mendaftarkan data", "yellow"))
						} else {
							fmt.Println(colorText("Gagal mendaftarn data", "red"))
						}

						fmt.Println(colorText("\n--- DAFTAR PEGAWAI ---", "green"))
						fmt.Println(colorText("======================", "green"))
						fmt.Println(authMenu.Show())
					case 2:
						var deleteUser user.User
						fmt.Println(colorText("\n--- Halaman Hapus Pegawai ---", "green"))
						fmt.Println(colorText("=============================", "green"))
						fmt.Print(colorText("masukkan ID pegawai yang ingin dihapus : ", "yellow"))
						fmt.Scanln(&deleteUser.ID)
						res, err := authMenu.Delete(deleteUser)
						if err != nil {
							fmt.Println(err.Error())
						}
						if res {
							fmt.Println(colorText("Sukses menghapus Pegawai", "yellow"))
						} else {
							fmt.Println(colorText("Gagal menghapus Pegawai", "red"))
						}

						fmt.Println(colorText("\n--- DAFTAR PEGAWAI ---", "green"))
						fmt.Println(colorText("======================", "green"))
						fmt.Println(authMenu.Show())
					case 3:
						var deleteBarang barang.Barang
						fmt.Println(colorText("\n--- Halaman Hapus Barang ---", "green"))
						fmt.Println(colorText("=============================", "green"))
						fmt.Print(colorText("masukkan ID barang yang ingin dihapus : ", "yellow"))
						fmt.Scanln(&deleteBarang.ID)
						res, err := barangMenu.Delete(deleteBarang)
						if err != nil {
							fmt.Println(err.Error())
						}
						if res {
							fmt.Println(colorText("Sukses menghapus Barang", "yellow"))
						} else {
							fmt.Println(colorText("Gagal menghapus Barang", "red"))
						}

						fmt.Println(colorText("\n--- DAFTAR BARANG ---", "green"))
						fmt.Println(colorText("=====================", "green"))
						fmt.Println(barangMenu.Show())
					case 4:
						var deleteKonsumen konsumen.Konsumen
						fmt.Println(colorText("\n--- Halaman Hapus Pelanggan ---", "green"))
						fmt.Println(colorText("===============================", "green"))
						fmt.Println(colorText("masukkan nomor HP pelanggan yang ingin dihapus :", "yellow"))
						fmt.Scanln(&deleteKonsumen.HP)
						res, err := konsumMenu.DeleteKonsumen(deleteKonsumen)
						if err != nil {
							fmt.Println(err.Error())
						}
						if res {
							fmt.Println(colorText("Sukses menghapus data pelanggan", "yellow"))
						} else {
							fmt.Println(colorText("Gagal menghapus data pelanggan", "red"))
						}

						fmt.Println(colorText("\n--- DAFTAR PELANGGAN ---", "green"))
						fmt.Println(colorText("========================", "green"))
						fmt.Println(konsumMenu.Show())
					case 5:
						fmt.Println(colorText("\n--- DAFTAR TRANSAKSI ---", "green"))
						fmt.Println(colorText("========================", "green"))
						fmt.Println(transaksiMenu.Show())
						fmt.Println(colorText("========================", "green"))

						var deleteTransaksi transaksi.Transaksi
						fmt.Println(colorText("\n--- Halaman Hapus Transaksi ---", "green"))
						fmt.Println(colorText("===============================", "green"))
						fmt.Print(colorText("masukkan nomor nota transaksi yang ingin dihapus : ", "yellow"))
						fmt.Scanln(&deleteTransaksi.NoNota)
						res, err := transaksiMenu.DeleteTransaksi(deleteTransaksi)
						if err != nil {
							fmt.Println(err.Error())
						}
						if res {
							fmt.Println(colorText("Sukses menghapus data transaksi", "yellow"))
						} else {
							fmt.Println(colorText("Gagal menghapus data transaksi", "red"))
						}

						fmt.Println(colorText("\n--- DAFTAR TRANSAKSI ---", "green"))
						fmt.Println(colorText("========================", "green"))
						fmt.Println(transaksiMenu.Show())
					case 9:
						isLogin = false
						screenClear()
					}
				}
			}

			if res.ID > 0 && inputNama != "admin" {
				isLogin := true
				fmt.Println(colorText("\n--- Login Sebagai Pegawai ---", "cyan"))
				for isLogin {
					loginMenu := 0
					fmt.Println(colorText("------ MENU PEGAWAI ------", "green"))
					fmt.Println(colorText("==========================", "green"))
					fmt.Println(colorText("1. Tambah Barang", "yellow"))
					fmt.Println(colorText("2. Lihat Barang", "yellow"))
					fmt.Println(colorText("3. Edit Deskripsi", "yellow"))
					fmt.Println(colorText("4. Update Stok Barang", "yellow"))
					fmt.Println(colorText("5. Transaksi dan Pelanggan", "yellow"))
					fmt.Println(colorText("9. Logout", "yellow"))
					fmt.Println(colorText("==========================", "green"))
					fmt.Print(colorText("Masukkan pilihan anda : ", "green"))
					fmt.Scanln(&loginMenu)
					screenClear()
					switch loginMenu {
					case 1:
						inputBarang := barang.Barang{}
						fmt.Println(colorText("\n--- Halaman Tambah Barang ---", "green"))
						fmt.Println(colorText("=============================", "green"))
						in := bufio.NewReader(os.Stdin)
						fmt.Print(colorText("Masukkan Nama Barang : ", "yellow"))
						name, _ := in.ReadString('\n')
						name = name[:len(name)-2]
						inputBarang.Nama = name
						fmt.Print(colorText("Masukkan Deskripsi Barang : ", "yellow"))
						desc, _ := in.ReadString('\n')
						desc = desc[:len(desc)-2]
						inputBarang.Deskripsi = desc
						fmt.Print(colorText("Masukkan Jumlah (Stok) : ", "yellow"))
						fmt.Scanln(&inputBarang.Stok)
						inputBarang.Owner = res.ID
						barRes, err := barangMenu.Insert(inputBarang)
						if err != nil {
							fmt.Println(err.Error())
						}
						inputBarang.ID = barRes
						fmt.Println(inputBarang)

						fmt.Println(colorText("\n--- DAFTAR BARANG ---", "green"))
						fmt.Println(colorText("=====================", "green"))
						fmt.Println(barangMenu.Show())
					case 2:
						fmt.Println(colorText("\n--- DAFTAR BARANG ---", "green"))
						fmt.Println(colorText("=====================", "green"))
						fmt.Println(barangMenu.Show())
					case 3:
						var editBarang barang.Barang
						in := bufio.NewReader(os.Stdin)
						fmt.Println(colorText("\n--- Halaman Edit Deskripsi Barang ---", "green"))
						fmt.Println(colorText("=====================================", "green"))
						fmt.Print(colorText("masukkan ID barang yang deskripsinya akan diedit : ", "yellow"))
						fmt.Scanln(&editBarang.ID)
						fmt.Println(colorText("masukkan Deskripsi terbaru : ", "yellow"))
						desc, _ := in.ReadString('\n')
						desc = desc[:len(desc)-2]
						editBarang.Deskripsi = desc
						res, err := barangMenu.Edit(editBarang)
						if err != nil {
							fmt.Println(err.Error())
						}
						if res {
							fmt.Println(colorText("Sukses Mengedit Deskripsi Barang", "yellow"))
						} else {
							fmt.Println(colorText("Gagal Mengedit Deskripsi Barang", "red"))
						}
						fmt.Println(editBarang)
					case 4:
						fmt.Println(colorText("\n--- Halaman Update Stok Barang ---", "green"))
						fmt.Println(colorText("==================================", "green"))
						var updateStok barang.Barang
						fmt.Print(colorText("masukkan ID barang yang akan diedit : ", "yellow"))
						fmt.Scanln(&updateStok.ID)
						fmt.Print(colorText("masukkan jumlah stok terbaru : ", "yellow"))
						fmt.Scanln(&updateStok.Stok)

						res, err := barangMenu.UpdateStok(updateStok)
						if err != nil {
							fmt.Println(err.Error())
						}
						if res {
							fmt.Println(colorText("Sukses update stok barang", "yellow"))
						} else {
							fmt.Println(colorText("Gagal update stok barang", "red"))
						}
						fmt.Println(updateStok)
					case 5:
						isTransaksi := true
						for isTransaksi {
							var choice int
							fmt.Println(colorText("\n--- Halaman Transaksi ---", "green"))
							fmt.Println(colorText("=========================", "green"))
							fmt.Println(colorText("1. Tambah Pelanggan", "yellow"))
							fmt.Println(colorText("2. Lihat Daftar Pelanggan", "yellow"))
							fmt.Println(colorText("3. Transaksi", "yellow"))
							fmt.Println(colorText("4. Cetak Nota", "yellow"))
							fmt.Println(colorText("9. Logout", "yellow"))
							fmt.Println(colorText("=========================", "green"))
							fmt.Print(colorText("silakan masukkan pilihan anda : ", "green"))
							fmt.Scanln(&choice)
							screenClear()
							switch choice {
							case 1:
								var newKonsumen konsumen.Konsumen
								in := bufio.NewReader(os.Stdin)
								fmt.Println(colorText("\n--- Halaman Tambah Pelanggan ---", "green"))
								fmt.Println(colorText("================================", "green"))
								fmt.Print(colorText("Masukkan nama : ", "yellow"))
								name, _ := in.ReadString('\n')
								name = name[:len(name)-2]
								newKonsumen.Nama = name
								fmt.Print(colorText("Masukkan nomor telepon : ", "yellow"))
								hp, _ := in.ReadString('\n')
								hp = hp[:len(hp)-2]
								newKonsumen.HP = hp
								newKonsumen.IdPegawai = res.ID
								res, err := konsumMenu.RegistKonsumen(newKonsumen)
								if err != nil {
									fmt.Println(err.Error())
								}
								if res {
									fmt.Println(colorText("Sukses mendaftarkan pelanggan", "yellow"))
								} else {
									fmt.Println(colorText("Gagal mendaftarkan pelanggan", "red"))
								}
								screenClear()
							case 2:
								fmt.Println(colorText("\n--- DAFTAR PELANGGAN ---", "green"))
								fmt.Println(colorText("========================", "green"))
								fmt.Println(konsumMenu.Show())
							case 3:
								var newTransaksi transaksi.Transaksi
								in := bufio.NewReader(os.Stdin)
								fmt.Println(colorText("\n--- Halaman Buat Transaksi ---", "green"))
								fmt.Println(colorText("===============================", "green"))
								newTransaksi.IdPegawai = res.ID
								fmt.Print(colorText("Masukkan nomor HP pelanggan : ", "yellow"))
								hp, _ := in.ReadString('\n')
								hp = hp[:len(hp)-2]
								newTransaksi.HpKonsumen = hp
								result, err := transaksiMenu.AddTransaksi(newTransaksi)
								if err != nil {
									fmt.Println(err.Error())
								}
								if result {
									fmt.Println(colorText("Transaksi Berhasil", "yellow"))
								} else {
									fmt.Println(colorText("Transaksi Gagal", "red"))
								}
								fmt.Println(newTransaksi)
								fmt.Println(colorText("\n--- DAFTAR TRANSAKSI ---", "green"))
								fmt.Println(colorText("========================", "green"))
								fmt.Println(transaksiMenu.Show())

								tambahItem := true
								for tambahItem {
									input := 0
									fmt.Println(colorText("1. Tambah Belanjaan", "yellow"))
									fmt.Println(colorText("9. Back", "yellow"))
									fmt.Println(colorText("=========================", "green"))
									fmt.Print(colorText("silakan masukkan pilihan anda : ", "green"))
									fmt.Scanln(&input)
									switch input {
									case 1:
										var newItem item.Item
										fmt.Println(colorText("================================", "green"))
										fmt.Print(colorText("Masukkan nomor nota : ", "yellow"))
										fmt.Scanln(&newItem.NoNota)
										// newItem.NoNota = newTransaksi.NoNota
										fmt.Print(colorText("Masukkan ID barang : ", "yellow"))
										fmt.Scanln(&newItem.IdBarang)
										fmt.Print(colorText("Masukkan jumlah barang : ", "yellow"))
										fmt.Scanln(&newItem.Kuantitas)
										hasil, err := itemMenu.Insert(newItem)
										if err != nil {
											fmt.Println(err.Error())
										}
										if hasil {
											fmt.Println(colorText("Belanjaan ditambahkan", "yellow"))
										} else {
											fmt.Println(colorText("Gagal menambahkan belanjaan", "red"))
										}
									case 9:
										tambahItem = false
									}
								}
							case 4:
								var newCetak transaksi.Transaksi

								fmt.Print(colorText("Masukkan nomor Nota : ", "yellow"))
								fmt.Scanln(&newCetak.NoNota)
								fmt.Println(colorText("\n======== Tokoku ========", "green"))
								fmt.Println(colorText("--- Cetak Nota ---", "green"))
								fmt.Println(" ")
								hasil, err := transaksiMenu.Cetak(newCetak)
								if err != nil {
									fmt.Println(err.Error())
								}
								fmt.Println(hasil)
							case 9:
								isTransaksi = false
							}
						}
					case 9:
						isLogin = false
					}
				}
			}
		}
	}
}
