package auth

import (
	"fmt"
	"log"
	"strings"
	"os"
	"bufio"
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
	
	u, p := strings.ToLower(strings.Trim(user, " ")), pass
	
	//TODO: check user credentials from a server folder
	fmt.Println(u, p)

	return false 
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