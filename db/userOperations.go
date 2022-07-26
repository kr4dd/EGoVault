package db

import (
	"encoding/json"
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

const (
	DB_PATH = "/workspaces/EGoVault/db/user.json"
)

type CatalogUser struct {
	Username, Password string
}

//Seal message into a file
func SealMsg(msg, filePathDestination string, additionalKey []byte) {
	fmt.Println("[+] Trying to seal data...")

	var masterKey, sealData []byte

	if additionalKey == nil {
		//Read master key
		masterKey, err = ReadMasterKey()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed reading masterKey while sealing: %v\n", err)
		}

	} else {
		masterKey = additionalKey
	}

	//Cipher msg with a master key
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

	mK, err := ReadMasterKey()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed reading masterKey while unsealing: %v\n", err)
	}
	r, err := readBytesFromUnsealedFile(mK, filePathDestination)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed reading bytes from unsealed file: %v\n", err)
	}
	fmt.Println("unsealed data: \n" + string(r))
}

func getUnsealMsg(masterKey []byte, filePathDestination string) (string, error) {
	r, err := readBytesFromUnsealedFile(masterKey, filePathDestination)
	if err != nil {
		return "", fmt.Errorf("error getting unseal msg: %v", err)
	}

	return string(r), nil
}

func readBytesFromUnsealedFile(masterKey []byte, filePathDestination string) ([]byte, error) {
	//Read bytes from file
	var sealedData, unsealData []byte
	if sealedData, err = ioutil.ReadFile(filePathDestination); err != nil {
		return nil, fmt.Errorf("reading unsealed file failed: %v", err)

	}

	if unsealData, err = ecrypto.Unseal(sealedData, masterKey); err != nil {
		return nil, fmt.Errorf("reading unsealed file failed: %v", err)
	}

	return unsealData, nil
}

//Add info to a previously file
func AppendMsg(msg, filePathDestination string) {
	fmt.Println("[+] Trying to append data...")

	//Read master key for unseal file
	masterKey, err := ReadMasterKey()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed reading masterKey while append: %v\n", err)
	}

	newMsg, err := getUnsealMsg(masterKey, filePathDestination)
	if err != nil {
		log.Fatal(err)
	}

	SealMsg(newMsg+"\n"+msg, filePathDestination, masterKey)
	fmt.Println("[+] Data was appended and Seal ...")

}

func ReadMasterKey() ([]byte, error) {
	fmt.Println("\nInsert your masterKey: ")

	maskedPassword, err := gopass.GetPasswdMasked()
	if err != nil {
		return nil, fmt.Errorf("reading masterKey failed: %v", err)
	}

	return []byte(maskedPassword), nil
}

func ReadUserDB(user, pass string) bool {
	var data CatalogUser
	validUser := false

	UnCipherDBData()

	json.Unmarshal(readDBContent(), &data)

	if (data.Username == user) && (data.Password == pass) {
		validUser = true
	}

	CipherDBData()

	return validUser

}

func readDBContent() []byte {
	f, err := os.Open(DB_PATH)
	if err != nil {
		fmt.Println(err)
	}

	fBytes, _ := ioutil.ReadAll(f)
	defer f.Close()

	return fBytes

}

func CreateAndCipherUserDB(user string, password []byte) {
	var data CatalogUser
	data.Username, data.Password = user, string(password)

	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile(DB_PATH, file, 0644)

	CipherDBData()
}
