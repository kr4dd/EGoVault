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

var (
	err error
)

func RequireCredentials() bool {
	user, err := readUser()
	if err != nil {
		log.Fatal(err)
	}

	pass, err := readPasswd()
	if err != nil {
		log.Fatal(err)
	}

	//TODO: limpiar caracteres extranhos y inputs raros para User
	u, p := strings.TrimSuffix(strings.Trim(user, " "), "\n"), pass

	return db.ReadUsersDB(u, string(p))

}

func readPasswd() ([]byte, error) {
	fmt.Println("\nInsert your password: ")

	maskedPassword, err := gopass.GetPasswdMasked()
	if err != nil {
		return nil, fmt.Errorf("Reading password failed: %v\n", err)
	}

	return []byte(maskedPassword), nil
}

func readUser() (string, error) {
	fmt.Println("\nInsert your user: ")

	user, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("Reading user failed: %v\n", err)
	}

	return user, nil
}
