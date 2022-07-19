package db

//Sign user.json with uniqueid of EGoVault executable

import (
	b64 "encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/edgelesssys/ego/ecrypto"
	"github.com/edgelesssys/ego/enclave"
)

func CipherDBData() {
	enclaveUniqueKey, enclaveUniqueKeyInfo, err := enclave.GetUniqueSealKey()
	if err != nil {
		log.Fatal("cannot get unique seal key")
	}

	cipherPlainPasswords(enclaveUniqueKey, enclaveUniqueKeyInfo)

}

func cipherPlainPasswords(enclaveUniqueKey, enclaveUniqueKeyInfo []byte) {
	var data CatalogUser
	json.Unmarshal(readDBContent(), &data)

	p := data.Password
	ep, err := ecrypto.Encrypt([]byte(p), enclaveUniqueKey, enclaveUniqueKeyInfo)
	if err != nil {
		log.Fatal("error while cipher json passwords")
	}

	epBase64 := b64.StdEncoding.EncodeToString(ep)

	data.Password = epBase64

	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile(DB_PATH, file, 0644)
}

func UnCipherDBData() {
	enclaveUniqueKey, enclaveUniqueKeyInfo, err := enclave.GetUniqueSealKey()
	if err != nil {
		log.Fatal("cannot get unique seal key")
	}

	UnCipherPasswords(enclaveUniqueKey, enclaveUniqueKeyInfo)
}

func UnCipherPasswords(enclaveUniqueKey, enclaveUniqueKeyInfo []byte) {
	var data CatalogUser
	json.Unmarshal(readDBContent(), &data)

	epBase64 := data.Password

	ep, err := b64.StdEncoding.DecodeString(epBase64)
	if err != nil {
		log.Fatal("base 64 cannot be decoded")
	}

	p, err := ecrypto.Decrypt(ep, enclaveUniqueKey, enclaveUniqueKeyInfo)
	if err != nil {
		log.Fatal("error while uncipher json passwords")
	}

	data.Password = string(p)

	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile(DB_PATH, file, 0644)
}
