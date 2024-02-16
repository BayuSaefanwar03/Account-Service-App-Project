package main

import (
	"fmt"
	"project1/config"
	"project1/data"
	"project1/topup"
	"project1/transfer"
	"project1/users"
)

var database = config.InitMysql()

func menuAccountService(user data.Users) {
	var input int
	for input != 99 {
		fmt.Println("Saldo anda :", user.Saldo)
		fmt.Println("Pilih menu")
		fmt.Println("1. View Account")
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
		switch input {
		case 1:
			fmt.Println("ID: ", user.ID)
			fmt.Println("Nama: ", user.Nama)
			fmt.Println("No Hp: ", user.HP)
			fmt.Println("Alamat: ", user.Alamat)
			fmt.Println("Created account: ", user.CreatedAt)
			fmt.Println("Updated account: ", user.UpdatedAt)
		case 2:
		case 3:
		case 4:
			nominal, success, err := Topup(user)
			if err != nil {
				fmt.Println("Terjadi kesalahan :", err)
			} else if !success {
				fmt.Println("Anda tidak berhasil topup")
			} else {
				fmt.Println("Selamat anda berhasil topup")
				user.Saldo += nominal
			}
		case 5:
			err := Send(user)
			if err != nil {
				fmt.Println("Terjadi Error:", err.Error())
			} else {
				fmt.Println("Selamat anda berhasil trasfer :)")
			}
		case 6:
			result := topup.HistoryTopup(database, user)
			fmt.Printf("%14s|%10s|%s\n", "Penerima", "nominal", "waktu")
			for i := 0; i < len(result); i++ {
				fmt.Printf("%14s|%10d|%s\n", result[i].HP, result[i].Nominal, result[i].CreatedAt)
			}
		case 7:
			result := transfer.HistoryTransfer(database, user)
			fmt.Printf("%14s|%10s\n", "Penerima", "nominal")
			for i := 0; i < len(result); i++ {
				fmt.Printf("%14s|%10d\n", result[i].HP_Penerima, result[i].Nominal)
			}
		case 8:
		}
	}
	fmt.Println("Terima kasih telah bertransaksi,")
}

func main() {
	config.Migrate(database)
	var input int
	if result := test(); result {
		input = 99
	}
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
				fmt.Println("Terjadi Error:", err.Error())
			} else if !success {
				fmt.Println("No HP/Password salah!")
			} else {
				fmt.Println("Selamat datang,", user.Nama)
				menuAccountService(user)
			}
		}
	}

}

func login() (data.Users, bool, error) {
	var hp string
	var password string
	fmt.Print("Masukkan HP : ")
	fmt.Scanln(&hp)
	fmt.Print("Masukkan Password : ")
	fmt.Scanln(&password)
	return users.Login(database, hp, password)
}

func Register() (bool, error) {
	var newUser data.Users
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

func Updata_Account() error {

	return nil
}

func Topup(user data.Users) (int, bool, error) {
	var nominal int
	fmt.Print("masukan nominal : ")
	fmt.Scanln(&nominal)
	success, err := topup.Newtopup(database, user, nominal)
	return nominal, success, err
}

func Send(user data.Users) error {
	var hp string
	var nominal int
	fmt.Print("Masukkan HP Tujuan : ")
	fmt.Scanln(&hp)
	fmt.Print("Masukkan Nominal : ")
	fmt.Scanln(&nominal)
	return transfer.Send(database, user, hp, nominal)
}

func test() bool {
	transfer.HistoryTransfer(database, data.Users{HP: "081"})
	return false
}
