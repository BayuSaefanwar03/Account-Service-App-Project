package main

import (
	"Account-Service/src/data"
	"Account-Service/src/database"
	"Account-Service/src/handle"
	"Account-Service/src/topup"
	"Account-Service/src/transfer"
	"Account-Service/src/users"
	"fmt"
	"time"
)

var db = database.InitMysql()

func main() {
	database.Migrate(db)
	var input int

	for input != 99 {
		fmt.Println("\n-----LOGIN-----")
		fmt.Println(" [1] Login")
		fmt.Println(" [2] Register")
		fmt.Println(" [3] Forget Password")
		fmt.Println("[99] Exit")
		fmt.Print("==> ")
		handle.ScanInt(&input)
		fmt.Println("")
		switch input {
		case 1:
			if user, success := Login(); success {
				accountService(user)
			}
		case 2:
			Register()
		case 3:
		case 99:
		default:
			fmt.Println("[Error] Wrong Input")
		}
	}
	fmt.Println("Thnk u for coming & Goodbye :)")
}

func accountService(user *data.Users) {
	fmt.Println("\nWelcome,", user.Nama)
	var input int
	for input != 99 {
		fmt.Printf("\nSaldo : Rp.%s\n", handle.ToRP(user.Saldo))
		fmt.Println("-----HOME-----")
		fmt.Println(" [1] View Profile")
		fmt.Println(" [2] Update Profile")
		fmt.Println(" [3] Delete Account")
		fmt.Println(" [4] Topup")
		fmt.Println(" [5] Transfer")
		fmt.Println(" [6] History Topup")
		fmt.Println(" [7] History Trasfer")
		fmt.Println(" [8] View Profile Other User")
		fmt.Println("[99] Logout")
		fmt.Print("==> ")
		handle.ScanInt(&input)
		fmt.Println("")
		switch input {
		case 1:
			View_Profile(*user, true)
		case 2:
			Update_Profile(user)
		case 3:
			if do := Delete_Account(*user); do {
				input = 99
			}
		case 4:
			Topup(user)
		case 5:
			Transfer(user)
		case 6:
			Topup_History(*user)
		case 7:
			Transfer_History(*user)
		case 8:
			View_Other_Profile()
		case 99:
		default:
			fmt.Println("[Error] Wrong Input")
		}
	}
	fmt.Println("Goodbye", user.Nama)
}

func Login() (*data.Users, bool) {
	var hp, password string
	fmt.Println("|--[LOGIN]")
	fmt.Print("|--> HP : ")
	handle.Scan(&hp)
	fmt.Print("|--> Password : ")
	handle.Scan(&password)

	user, err := users.Login(db, hp, password)
	if err != nil {
		fmt.Println("[Error]", err)
		return &data.Users{}, false
	} else if user.Nama == "" {
		fmt.Println("[Error] Invalid HP or Password")
		return &data.Users{}, false
	} else {
		return &user, true
	}
}

func Register() {
	var user data.Users
	fmt.Println("|--[Register]")
	fmt.Print("|--> Name : ")
	handle.Scan(&user.Nama)
	fmt.Print("|--> No HP : ")
	handle.Scan(&user.HP)
	fmt.Print("|--> Password : ")
	handle.Scan(&user.Password)
	fmt.Print("|--> Address : ")
	handle.Scan(&user.Alamat)
	user.CreatedAt = time.Time{}
	user.UpdatedAt = time.Time{}

	success, err := users.Register(db, user)
	if err != nil {
		fmt.Println("[Error]", err)
	} else if !success {
		fmt.Println("[Error] The cellphone number is registered")
	} else {
		fmt.Println("Account successfully registered :)")
	}
}

func View_Profile(user data.Users, access bool) {
	if access {
		fmt.Println("|--[View Profile]")
	} else {
		fmt.Printf("|--[See %s's Profile]\n", user.Nama)
	}
	fmt.Println("|--> Name         :", user.Nama)
	fmt.Println("|--> Phone Number :", user.HP)
	fmt.Println("|--> Address      :", user.Alamat)
	if access {
		fmt.Println("|--> Create At    :", user.CreatedAt.Format("01/02/2006 03:04:05"))
		fmt.Println("|--> Last Update  :", user.UpdatedAt.Format("01/02/2006 03:04:05"))
	}
}

func Update_Profile(user *data.Users) {
	var input int
	var user_old = *user
	fmt.Println("|--[Edit Profile]")
	fmt.Println("|--[1] Name")
	fmt.Println("|--[2] Phone Number")
	fmt.Println("|--[3] Address")
	fmt.Println("|--[4] Password")
	fmt.Println("Press [Enter] to exit")
	fmt.Print("==> ")
	handle.ScanInt(&input)
	fmt.Println("")
	switch input {
	case 1:
		fmt.Print("Name --> ")
		handle.Scan(&user.Nama)
	case 2:
		fmt.Print("HP --> ")
		handle.Scan(&user.HP)
	case 3:
		fmt.Print("Address --> ")
		handle.Scan(&user.Alamat)
	case 4:
		fmt.Print("Password --> ")
		handle.Scan(&user.Password)
	case 0:
		return
	}
	if user.Nama == "" || user.HP == "" || user.Alamat == "" {
		fmt.Println("[Error] Empty is prohibited")
		*user = user_old
	} else {
		user.UpdatedAt = time.Time{}
		success, err := users.Update(db, user_old, *user)
		if err != nil {
			fmt.Println("[Error]", err)
		} else if !success {
			fmt.Println("[Error] Failed to change profile")
		} else {
			fmt.Println("Successfully changed profile")
		}
	}
}

func Delete_Account(user data.Users) bool {
	var input int
	fmt.Println("|--[Delete Profile]")
	fmt.Println("| Are you sure? don't leave me :(")
	fmt.Println("|--[1] Yes :(")
	fmt.Println("|--[2] No")
	fmt.Print("==> ")
	handle.ScanInt(&input)
	fmt.Println("")
	switch input {
	case 1:
		success, err := users.Delete(db, user)
		if err != nil {
			fmt.Println("[Error]", err)
		} else if !success {
			fmt.Println("[Error] Failed to delete account")
		} else {
			fmt.Println(user.Nama, "Nooooo...... ˃⌓ <")
			return true
		}
	case 2:
		fmt.Println("Hooray... :)")
	case 0:
		return false
	}
	return false
}

func Topup(user *data.Users) {
	var nominal int
	fmt.Println("|--[Topup saldo]")
	fmt.Println("|")
	fmt.Println("| press [enter] to cancel")
	fmt.Println("|--Nominal :")
	fmt.Print("|--> Rp.")
	handle.ScanInt(&nominal)
	fmt.Println("")
	if nominal < 0 {
		fmt.Println("[Error] Nominal cannot be less than zero")
		return
	} else if nominal == 0 {
		fmt.Println("Cencel")
		return
	}

	success, err := topup.Topup(db, user, nominal)
	if err != nil {
		fmt.Println("[Error]", err)
	} else if !success {
		fmt.Println("[Error] Failed to add saldo")
	} else {
		fmt.Println("Successfully added")
	}
}

func Topup_History(user data.Users) {
	data := topup.HistoryTopup(db, user)

	if len(data) == 0 {
		fmt.Println("No topup history")
	}

	for i := 0; i < len(data); i++ {
		if i == 0 {
			fmt.Printf("|%-2s|%-10s|%-19s|\n", "ID", "Nominal", "Date")
		}
		fmt.Printf("|%-2d|%-10d|%-19s|\n", data[i].TopupID, data[i].Nominal, data[i].CreatedAt.Format("01/02/2006 03:04:05"))
	}
}

func Transfer(user *data.Users) {
	var to string
	var nominal int
	fmt.Println("|--[Transfer]")
	fmt.Print("|--> To : ")
	handle.Scan(&to)
	fmt.Print("|--> Rp.")
	handle.ScanInt(&nominal)
	fmt.Println("")

	if to == user.HP {
		fmt.Println("[Error] Can't send to yourself")
		return
	} else if nominal <= 0 {
		fmt.Println("[Error] Nominal cannot be less than one")
		return
	} else if nominal > user.Saldo {
		fmt.Println("[Error] Nominal > Your Saldo")
		return
	}
	receiver, _ := users.View(db, to)
	if receiver.Nama == "" {
		fmt.Println("[Error] The recipient's cellphone number was not found")
		return
	}

	err := transfer.Send(db, user, receiver, nominal)
	if err != nil {
		fmt.Println("[Error]", err)
	} else {
		fmt.Println("Successfully saldo transfer")
	}
}

func Transfer_History(user data.Users) {
	data := transfer.HistoryTransfer(db, user)

	if len(data) == 0 {
		fmt.Println("No topup history")
	}

	for i := 0; i < len(data); i++ {
		if i == 0 {
			fmt.Printf("|%-2s|%-6s|%-8s|%-7s|%-19s|\n", "ID", "Sender", "Receiver", "Nominal", "Date")
		}
		fmt.Printf("|%-2d|%-6s|%-8s|%-7d|%-19s|\n", data[i].TransferID, data[i].HP_Pengirim, data[i].HP_Penerima, data[i].Nominal, data[i].CreatedAt.Format("01/02/2006 03:04:05"))
	}
}

func View_Other_Profile() {
	var hp string
	fmt.Println("|--[View Profile Other User]")
	fmt.Print("|--> No Hp : ")
	handle.Scan(&hp)
	fmt.Println("")

	receiver, _ := users.View(db, hp)
	if receiver.Nama == "" {
		fmt.Println("[Error] The recipient's cellphone number was not found")
		return
	}

	View_Profile(receiver, false)
}
