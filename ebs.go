//important variables

//userid := 8794402                                                                                                                                                                                                                                                                //input the value of "memberSeq" in your cookie(int)
//videoid := 374                                                                                                                                                                                                                                                                   //insert videoid here (4 digits int)
//progress := 100                                                                                                                                                                                                                                                                  // any number you want(int between 0-100)
//token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJPbmxpbmUgQ2xhc3MiLCJtZW1iZXJOYW1lIjoi7Jik7J2A7LSdIiwibWVtYmVyU2Nob29sQ29kZSI6Ik0wNjI2MyIsImV4cCI6MTYxNDk0NDk0NywiaWF0IjoxNjE0Njg1NzQ3LCJtZW1iZXJJZCI6InN0YXJ3YXJzOTM5NSJ9.kgnKNWwu2aDzxu4XSpknKxVPkoAwAfQG58qlHWXzvhw" //input the value of "access" in your cookie(string)

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
	//important variables
	var userid, videoid, progress int
	var token string
	fmt.Println("input key: ")
	fmt.Scanf("%s|%s|%s|%s", &userid, &token, &videoid, &progress)

	for i := 0; i < progress+1; i++ {
		formatted := []byte(strconv.Itoa(userid) + "|" + strconv.Itoa(videoid) + "|" + strconv.Itoa(i))

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
		req, _ := http.NewRequest("POST", "https://kyg4.ebsoc.co.kr/lecture/api/v1/student/learning/"+strconv.Itoa(videoid)+"/progress", s)
		req.Header.Add("X-AUTH-TOKEN", token)
		req.Header.Add("Content-Type", "application/json;charset=UTF-8")
		resp, _ := client.Do(req)
		f, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(f))
	}
}
