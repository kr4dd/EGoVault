package SubSystems

import (
	"EGoVault/db"
	"fmt"
	"os"
)

func CliMenu() {
	fmt.Fprintf(os.Stdout, "Welcome to EGoVault:\n\n")

	if len(os.Args[1:]) <= 0 {
		helpMenu()
	} else {
		checkParameters(os.Args[1:])
	}

}

func checkParameters(parameters []string) {
	badParams := "[!] Invalid parameters!! \nUse --help"

	switch parameters[0] {
	case "--seal":
		if len(os.Args[2:]) != 2 {
			fmt.Println(badParams)
			os.Exit(1)
		}
		db.SealMsg(parameters[1], parameters[2], nil)

	case "--unseal":
		if len(os.Args[2:]) != 1 {
			fmt.Println(badParams)
			os.Exit(1)
		}
		db.UnsealMsgShow(parameters[1])
	case "--append":
		if len(os.Args[2:]) != 2 {
			fmt.Println(badParams)
			os.Exit(1)
		}
		db.AppendMsg(parameters[1], parameters[2])
	default:
		helpMenu()
	}

}

func helpMenu() {
	fmt.Print("\n[ARGUMENTS]\n")
	fmt.Println("--help\t\tShow helps menu")
	fmt.Println("--seal\t\tSpecify message and path where you can save the encrypted message")
	fmt.Println("--unseal\tSpecify path of the seal file")
	fmt.Println("--append\tAppend new data to a existent file")

	fmt.Print("\n[EXAMPLES]\n")
	seal := "--seal \"your message using double quotes\" <filePathDestination>"
	unseal := "--unseal <filePathDestination>"
	append := "--append \"your NEW message using double quotes\" <existentFile>"

	fmt.Println("ego run app " + seal)
	fmt.Println("ego run app " + unseal)
	fmt.Println("ego run app " + append)
}
