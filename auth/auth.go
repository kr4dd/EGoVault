package auth

import (
	"EGoVault/db"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/howeyc/gopass"
)

func RequireCredentials() bool {
	db.UnCipherDBData()

	fmt.Println("[*] Login:")
	user, err := readUser()
	if err != nil {
		log.Fatal(err)
	}

	pass, err := readPasswd()
	if err != nil {
		log.Fatal(err)
	}

	//TODO: clean strange user inputs
	u, p := strings.TrimSuffix(strings.Trim(user, " "), "\n"), pass

	return db.ReadUserDB(u, string(p))

}

func readPasswd() ([]byte, error) {
	fmt.Println("\nInsert your password: ")

	maskedPassword, err := gopass.GetPasswdMasked()
	if err != nil {
		return nil, fmt.Errorf("reading password failed: %v", err)
	}

	return []byte(maskedPassword), nil
}

func readUser() (string, error) {
	fmt.Println("\nInsert your user: ")

	user, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("reading user failed: %v", err)
	}

	return user, nil
}

func RequireUserCreation() {
	if _, err := os.Stat(db.DB_PATH); err == nil {
		fmt.Println("[!] User was registred previously")
	} else {
		//User doesnt exists, create it
		createUser()
		db.CipherDBData()
	}

}

func createUser() {
	fmt.Println("[*] Create user: ")
	user, err := readUser()
	if err != nil {
		log.Fatal(err)
	}

	pass, err := readPasswd()
	if err != nil {
		log.Fatal(err)
	}

	//TODO: clean strange user inputs
	u, p := strings.TrimSuffix(strings.Trim(user, " "), "\n"), pass

	db.CreateUserDB(u, p)
}
