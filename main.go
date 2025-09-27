package main

import (
	"database/sql"
	"fmt"
	"goPasswordGenerator/model"
	"goPasswordGenerator/service"
	"goPasswordGenerator/store"
	"goPasswordGenerator/util"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func clearScreen() {
	cmd := exec.Command("clear") // Default for Linux/macOS
	if os.Getenv("GOOS") == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	menu := `
==========================
   PASSWORD GENERATOR MENU
==========================
1. Generate new password
2. Save password to CSV
3. View saved passwords
4. View specific password
5. Exit
==========================
`

	db, err := sql.Open("sqlite3", "./password.db")
	if err != nil {
		log.Fatal("Error opening database:", err.Error())
	}
	defer db.Close()

	q := `CREATE TABLE if not exists passwords(
		id integer primary key autoincrement,
		label text not null,
		password text not null,
		created_at text not null
	)`

	if _, err := db.Exec(q); err != nil {
		log.Fatal(err.Error())
	}

	store := store.New(db)
	services := service.New(store)

	for {
		var selection int
		fmt.Println(menu)
		fmt.Print("Enter your selection: ")
		fmt.Scan(&selection)
		clearScreen()

		switch selection {
		case 1:
			var useNumbers, useSymbols, label string
			var length int

			fmt.Print("What length do you want for your password: ")
			fmt.Scan(&length)

			fmt.Print("Do you want to use numbers in your password? (y/n): ")
			fmt.Scan(&useNumbers)

			fmt.Print("Do you want to use symbols in your password? (y/n): ")
			fmt.Scan(&useSymbols)

			includeNumbers := useNumbers == "y"
			includeSymbols := useSymbols == "y"

			generatedPass, _ := util.GenerateRandomPassword(length, includeNumbers, includeSymbols)

			fmt.Print("What label do you want to use to save your password? ")
			fmt.Scan(&label)

			newPass := model.Password{Label: label, Password: generatedPass, CreatedAt: time.Now()}
			if _, err := services.CreatePassword(&newPass); err != nil {
				log.Fatal(err.Error())
			}
		case 2:
			var label, password string

			fmt.Print("What label do you want to use to save your password? ")
			fmt.Scan(&label)

			fmt.Print("Enter your password: ")
			fmt.Scan(&password)

			newPass := model.Password{Label: label, Password: password, CreatedAt: time.Now()}
			if _, err := services.CreatePassword(&newPass); err != nil {
				log.Fatal(err.Error())
			}
		case 3:
			passwords, err := services.GetAllPasswords()
			if err != nil {
				log.Fatal(err.Error())
			}

			template := `{{.ID}}) {{.Label}} - {{.Password}}`
			stringTemplate, err := util.PrintPasswordsTemplate(passwords, template)
			if err != nil {
				log.Fatal(err.Error())
			}

			fmt.Println(stringTemplate)
		case 4:
			var selection string
			passwords, err := services.GetAllPasswords()
			if err != nil {
				log.Fatal(err.Error())
			}

			template := `{{.ID}}) {{.Label}}`
			stringTemplate, err := util.PrintPasswordsTemplate(passwords, template)
			if err != nil {
				log.Fatal(err.Error())
			}

			fmt.Println(stringTemplate)
			fmt.Print("Select which password you want to view: ")
			fmt.Scan(&selection)

			selectionConverted, err := strconv.Atoi(selection)
			if err != nil {
				log.Fatal(err.Error())
			}

			password, err := services.GetPasswordById(selectionConverted)
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Printf("\n%s - %s\n", password.Label, password.Password)
		case 5:
			return
		}
	}
}