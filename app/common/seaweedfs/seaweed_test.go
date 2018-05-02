package seaweedfs

import (
	"testing"
	"fmt"
	"encoding/json"
)

func TestUploadFile(t *testing.T) {
	targetUrl := "http://35.194.190.28:9333/submit"
	fileName := "D:\\tmp\\test.xml"
	resp, err := UploadFile(fileName, targetUrl)
	if err != nil {
		fmt.Println(err)
	}
	bytes, _ := json.Marshal(resp)
	println(string(bytes))
}

func TestUpdateFile(t *testing.T) {
	targetUrl := "http://35.194.190.28:9082/3,499b5562cb"
	fileName := "D:\\tmp\\OPML.xml"
	resp, err := UpdateFile(fileName, targetUrl)
	if err != nil {
		fmt.Println(err)
	}
	bytes, _ := json.Marshal(resp)
	println(string(bytes))
}

func TestDeleteFile(t *testing.T) {
	targetUrl := "http://35.194.190.28:9082/3,499b5562cb"
	resp, err :=DeleteFile(targetUrl)
	if err != nil {
		fmt.Println(err)
	}
	bytes, _ := json.Marshal(resp)
	println(string(bytes))
}

