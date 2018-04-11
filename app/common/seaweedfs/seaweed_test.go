package seaweedfs

import (
	"net/http"
	"io/ioutil"
	"testing"
	"fmt"
)

func TestUploadFile(t *testing.T) {
	targetUrl := "http://35.194.190.28:9333/submit"
	fileName := "upload.html"
	resp, err := UploadFile(fileName, targetUrl)
	if err != nil {
		fmt.Println(err)
	}

	if http.StatusOK == resp.StatusCode {
		fmt.Println(resp.Body)
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(body))
	}
}

func TestUpdateFile(t *testing.T) {
	targetUrl := "http://35.194.190.28:9082/6,40caa48662"
	fileName := "upload.html"
	resp, err :=UpdateFile(fileName, targetUrl)
	if err != nil {
		fmt.Println(err)
	}

	if http.StatusOK == resp.StatusCode {
		fmt.Println(resp.Body)
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(body))
	}
}

func TestDeleteFile(t *testing.T) {
	targetUrl := "http://35.194.190.28:9082/6,40caa48662"
	resp, err :=DeleteFile(targetUrl)
	if err != nil {
		fmt.Println(err)
	}

	if http.StatusOK == resp.StatusCode {
		fmt.Println(resp.Body)
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(body))
	}
}

