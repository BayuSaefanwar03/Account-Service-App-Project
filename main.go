package main

import (
	"fmt"
	"project1/config"
	"project1/users"
)

var database = config.InitMysql()

func menuAccountService(user string) {
	fmt.Println("Selamat Datang,", user)
	var input int
	for input != 99 {
		fmt.Println("Pilih menu")
		fmt.Println("1. Account")
		fmt.Println("2. Update Account")
		fmt.Println("3. Delete Account")
		fmt.Println("4. Top-Up")
		fmt.Println("5. Transfer")
		fmt.Println("6. History Top-Up")
		fmt.Println("7. History Transfer")
		fmt.Println("8. All User")
		fmt.Println("99. Logout")
		fmt.Print("Masukkan pilihan:")
		fmt.Scanln(&input)
	}
	fmt.Println("Terima kasih telah bertransaksi,", user)
}

func main() {
	config.Migrate(database)
	var input int
	for input != 99 {
		fmt.Println("Pilih menu")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("99. Exit")
		fmt.Print("Masukkan pilihan:")
		fmt.Scanln(&input)
		switch input {
		case 1:
			success, err := Register()
			if err != nil {
				fmt.Println("terjadi kesalahan(tidak bisa mendaftarkan pengguna)", err.Error())
			}

			if success {
				fmt.Println("selamat anda telah terdaftar")
			}
		case 2:
			user, success, err := login()
			if err != nil {
				fmt.Println("Terjadi error: ", err)
			} else if !success {
				fmt.Println("no telp atau password salah!")
			} else {
				fmt.Println("Selamat Datang", user.Nama)
				menuAccountService(user)
			}

		}
	}

}

func login() (users.Users, bool, error) {
	var hp string
	var password string
	var loggedIn users.Users
	fmt.Print("Masukkan HP : ")
	fmt.Scanln(&hp)
	fmt.Print("Masukkan Password : ")
	fmt.Scanln(&password)
	return users.Login(database, hp, password)
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
