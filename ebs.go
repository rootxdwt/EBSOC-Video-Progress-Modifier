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
	var progress int
	var token, userid, videoid string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	text2 := strings.Split(text, "|")
	userid = text2[0]
	token = text2[1]
	videoid = text2[2]
	progress, _ = strconv.Atoi(strings.TrimSuffix(text2[3], "\r\n"))

	for i := 0; i < progress+1; i++ {
		formatted := []byte(userid + "|" + videoid + "|" + strconv.Itoa(i))

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
		req, _ := http.NewRequest("POST", "https://kyg4.ebsoc.co.kr/lecture/api/v1/student/learning/"+videoid+"/progress", s)
		req.Header.Add("X-AUTH-TOKEN", token)
		req.Header.Add("Content-Type", "application/json;charset=UTF-8")
		resp, _ := client.Do(req)
		f, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(f))
	}
}
