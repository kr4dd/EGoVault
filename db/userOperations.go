package db

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/edgelesssys/ego/ecrypto"
	"github.com/howeyc/gopass"
)

var (
	err error
)

//Seal message into a file
func SealMsg(msg string, filePathDestination string, additionalKey []byte) {
	fmt.Println("[+] Trying to seal data...")
	var masterKey []byte
	if additionalKey == nil {
		//Read master key
		masterKey = readMasterKey()
	} else {
		masterKey = additionalKey
	}

	//Cipher msg with a master key
	var sealData []byte
	if sealData, err = ecrypto.SealWithProductKey([]byte(msg), masterKey); err != nil {
		log.Fatal(err)
	}

	//Write data that was seal with master key previously, into a .txt file
	f, err := os.OpenFile(filePathDestination, os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	if _, err := f.WriteString(string(sealData)); err != nil {
		log.Println(err)
	}

}

//Unseal message from a file
func UnsealMsgShow(filePathDestination string) {
	fmt.Println("[+] Trying to unseal data...")

	fmt.Println("Unsealed data: \n" + string(readBytesFromUnsealedFile(readMasterKey(), filePathDestination)))
}

func getUnsealMsg(masterKey []byte, filePathDestination string) string {
	return string(readBytesFromUnsealedFile(masterKey, filePathDestination))
}

func readBytesFromUnsealedFile(masterKey []byte, filePathDestination string) []byte {
	var err error

	//Read bytes from file
	var sealedData []byte
	if sealedData, err = ioutil.ReadFile(filePathDestination); err != nil {
		log.Fatal(err)
	}

	var unsealData []byte
	if unsealData, err = ecrypto.Unseal(sealedData, masterKey); err != nil {
		log.Fatal(err)
	}

	return unsealData
}

//Add info to a previously file
func AppendMsg(msg string, filePathDestination string) {
	fmt.Println("[+] Trying to append data...")

	//Read master key for unseal file
	masterKey := readMasterKey()

	newMsg := getUnsealMsg(masterKey, filePathDestination) + "\n" + msg

	SealMsg(newMsg, filePathDestination, masterKey)
	fmt.Println("[+] Data was appended and Seal ...")

}

func readMasterKey() []byte {
	fmt.Println("\nInsert your masterKey: ")

	maskedPassword, err := gopass.GetPasswdMasked()
	if err != nil {
		panic("Input masterKey failed")
	}

	return []byte(maskedPassword)
}
