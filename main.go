package main

import (
	"fmt"
	"project1/config"
	"project1/users"
)

var database = config.InitMysql()

func menuLibrary(user string) {
	fmt.Println("Selamat Datang,", user)
	var input int
	for input != 99 {
		fmt.Println("Pilih menu")
		fmt.Println("1. Lihat Daftar Buku")
		fmt.Println("2. Penjam Barang")
		fmt.Println("3. eefefefes")
		fmt.Println("99. Logout")
		fmt.Print("Masukkan pilihan:")
		fmt.Scanln(&input)
	}
	fmt.Println("Selamat Tinggal,", user)
}

func main() {
	config.Migrate(database)
	var input int
	for input != 99 {
		fmt.Println("Pilih menu")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("99. Exit")
		fmt.Print("Masukkan pilihan:")
		fmt.Scanln(&input)
		switch input {
		case 1:
			user := login()
			if user == "" {
				fmt.Println("Login Gagal!!")
			} else {
				menuLibrary(user)
			}

		case 2:
			success, err := Register()
			if err != nil {
				fmt.Println("terjadi kesalahan(tidak bisa mendaftarkan pengguna)", err.Error())
			}

			if success {
				fmt.Println("selamat anda telah terdaftar")
			}
		}
	}

}

func login() string {
	var hp string
	var password string
	var loggedIn users.Users
	fmt.Print("Masukkan HP : ")
	fmt.Scanln(&hp)
	fmt.Print("Masukkan Password : ")
	fmt.Scanln(&password)
	loggedIn, err := users.Login(database, hp, password)
	if err == nil {
		return loggedIn.Nama
	} else {
		return ""
	}
}

func Register() (bool, error) {
	var newUser users.Users
	fmt.Print("Masukkan nama     : ")
	fmt.Scanln(&newUser.Nama)
	fmt.Print("Masukkan nomor HP : ")
	fmt.Scanln(&newUser.HP)
	fmt.Print("Masukkan password : ")
	fmt.Scanln(&newUser.Password)
	fmt.Print("Masukkan alamat   : ")
	fmt.Scanln(&newUser.Alamat)
	return users.Register(database, newUser)
}
