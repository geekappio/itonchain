package seaweedfs

import (
	"net/http"
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"fmt"
	"io"
	"errors"
	"encoding/json"
)

type UploadResponse struct {
	Fid string `json:"fid"`
	FileName string `json:"fileName"`
	FileUrl string `json:"fileUrl"`
	Size int64 	`json:"size"`
}

type UpdateResponse struct {
	Name string `json:"name"`
	Size int64 	`json:"size"`
}

type DeleteResponse struct {
	Size int64 	`json:"size"`
}

func UploadFile(filename string, url string) (*UploadResponse, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return UploadFileContent(filename, content, url)
}

func UploadFileContent(fileName string, content []byte, url string) (*UploadResponse, error) {
	model := new(UploadResponse)
	err := uploadFile("POST", url, fileName, content, model)
	return model, err
}

func UpdateFile(filename string, url string) (*UpdateResponse, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return UpdateFileContent(filename, content, url)
}

func UpdateFileContent(fileName string, content []byte, url string) (*UpdateResponse, error) {
	model := new(UpdateResponse)
	err := uploadFile("PUT", url, fileName, content, model)
	return model, err
}

func DeleteFile(url string) (*DeleteResponse, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	model := new(DeleteResponse)
	err = buildModel(resp, model)
	return model, err
}

func uploadFile(method string, url string, fileName string, content []byte, model interface{}) (error) {
	bodyBuf := bytes.NewBufferString("")
	bodyWriter := multipart.NewWriter(bodyBuf)
	_, err := bodyWriter.CreateFormFile("userfile", fileName)
	if err != nil {
		return err
	}
	boundary := bodyWriter.Boundary()
	closeBuf := bytes.NewBufferString(fmt.Sprintf("\r\n--%s--\r\n", boundary))
	requestReader := io.MultiReader(bodyBuf, bytes.NewReader(content), closeBuf)
	req, err := http.NewRequest(method, url, requestReader)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "multipart/form-data; boundary="+boundary)
	req.ContentLength = int64(bodyBuf.Len()) + int64(len(content)) + int64(closeBuf.Len())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	return buildModel(resp, model)
}

func buildModel(resp *http.Response, model interface{}) error {
	body, _ := ioutil.ReadAll(resp.Body)

	// FIXME 不知道哪里有毛病会返回201但是却通讯正常
	if 2 == resp.StatusCode / 100 {
		return json.Unmarshal(body, model)
	} else {
		return errors.New(fmt.Sprintf("网络通讯异常：StatusCode=%v, Body=%v", resp.StatusCode, string(body)))
	}
}
