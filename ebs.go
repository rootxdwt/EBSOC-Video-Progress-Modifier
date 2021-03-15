package main

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	//important variables
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text: ")
	text, _ := reader.ReadString('\n')
	text2 := strings.Split(strings.TrimSpace(text), "|")
	progress, _ := strconv.Atoi(text2[3])

	for i := 0; progress >= i; i++ {
		formatted := []byte(text2[0] + "|" + text2[2] + "|" + strconv.Itoa(i))

		for (len(formatted) % 16) != 0 {
			formatted = append(formatted, 0x00)
		}

		block, _ := aes.NewCipher([]byte("l40jsfljasln32uf"))
		md := cipher.NewCBCEncrypter(block, []byte("asjfknal3bafjl23"))
		store := make([]byte, len(formatted))
		//32
		md.CryptBlocks(store, formatted)
		b := base64.URLEncoding.EncodeToString(store)

		q := []byte(`{"encriptedProgressRate":"` + b + `"}`)
		s := bytes.NewBuffer(q)

		client := &http.Client{}
		req, _ := http.NewRequest("POST", "https://"+text2[4]+".ebsoc.co.kr/lecture/api/v1/student/learning/"+text2[2]+"/progress", s)
		req.Header.Add("X-AUTH-TOKEN", text2[1])
		req.Header.Add("Content-Type", "application/json;charset=UTF-8")
		resp, _ := client.Do(req)
		if resp != nil {
			f, _ := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()
			fmt.Println(string(f))
		} else {
			fmt.Println("no response body")
		}

	}
}
