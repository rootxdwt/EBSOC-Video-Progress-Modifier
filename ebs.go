package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	
	var userid int
	fmt.Println("input user id: ")
	fmt.Scanln(&userid)

	fmt.Println("input video id: ")
	var videoid int
	fmt.Scanln(&videoid)

	fmt.Println("input progress: ")
	var progress int
	fmt.Scanln(&progress)

	fmt.Println("input token: ")
	var token string
	fmt.Scanln(&token)

	for i := 0; i < progress+1; i++ {
		formatted := []byte(strconv.Itoa(userid) + "|" + strconv.Itoa(videoid) + "|" + strconv.Itoa(i))

		for (len(formatted) % 16) != 0 {
			formatted = append(formatted, 0x00)
		}

		block, _ := aes.NewCipher([]byte("l40jsfljasln32uf"))
		md := cipher.NewCBCEncrypter(block, []byte("asjfknal3bafjl23"))
		store := make([]byte, len(formatted))
		
		md.CryptBlocks(store, formatted)
		b := base64.URLEncoding.EncodeToString(store)

		q := []byte(`{"encriptedProgressRate":"` + b + `"}`)
		s := bytes.NewBuffer(q)

		client := &http.Client{}
		req, _ := http.NewRequest("POST", "https://kyg4.ebsoc.co.kr/lecture/api/v1/student/learning/"+strconv.Itoa(videoid)+"/progress", s)
		req.Header.Add("X-AUTH-TOKEN", token)
		req.Header.Add("Content-Type", "application/json;charset=UTF-8")
		resp, _ := client.Do(req)
		f, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(f))
	}
}
