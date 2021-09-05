package pkg

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"log"
)

func GetMD5Hash(target string) string {
	data, err := ioutil.ReadFile(target)
	if err != nil {
		log.Fatalf("Open %s Failed with error: ", target, err)
	}

	hasher := md5.New()
	hasher.Write(data)
	return hex.EncodeToString(hasher.Sum(nil))
}
