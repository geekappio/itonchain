package seaweedfs

import (
	"net/http"
	"io"
	"bytes"
	"mime/multipart"
	"fmt"
	"os"
	"log"
)

func UploadFile(filename string, targetUrl string) (*http.Response, error) {
	bodyBuf := bytes.NewBufferString("")
	bodyWriter := multipart.NewWriter(bodyBuf)

	// use the body_writer to write the Part headers to the buffer
	_, err := bodyWriter.CreateFormFile("userfile", filename)
	if err != nil {
		log.Print("error writing to buffer")
		return nil, err
	}

	// the file data will be the second part of the body
	fh, err := os.Open(filename)
	if err != nil {
		log.Print("error opening file")
		return nil, err
	}
	// need to know the boundary to properly close the part myself.
	boundary := bodyWriter.Boundary()
	//close_string := fmt.Sprintf("\r\n--%s--\r\n", boundary)
	closeBuf := bytes.NewBufferString(fmt.Sprintf("\r\n--%s--\r\n", boundary))

	// use multi-reader to defer the reading of the file data until
	// writing to the socket buffer.
	requestReader := io.MultiReader(bodyBuf, fh, closeBuf)
	fi, err := fh.Stat()
	if err != nil {
		log.Printf("Error Stating file: %s", filename)
		return nil, err
	}
	req, err := http.NewRequest("PUT", targetUrl, requestReader)
	if err != nil {
		return nil, err
	}

	// Set headers for multipart, and Content Length
	req.Header.Add("Content-Type", "multipart/form-data; boundary="+boundary)
	req.ContentLength = fi.Size() + int64(bodyBuf.Len()) + int64(closeBuf.Len())

	return http.DefaultClient.Do(req)
}

func UpdateFile(filename string, targetUrl string) (*http.Response, error) {
	bodyBuf := bytes.NewBufferString("")
	bodyWriter := multipart.NewWriter(bodyBuf)

	// use the body_writer to write the Part headers to the buffer
	_, err := bodyWriter.CreateFormFile("userfile", filename)
	if err != nil {
		log.Print("error writing to buffer")
		return nil, err
	}

	// the file data will be the second part of the body
	fh, err := os.Open(filename)
	if err != nil {
		log.Print("error opening file")
		return nil, err
	}
	// need to know the boundary to properly close the part myself.
	boundary := bodyWriter.Boundary()
	//close_string := fmt.Sprintf("\r\n--%s--\r\n", boundary)
	closeBuf := bytes.NewBufferString(fmt.Sprintf("\r\n--%s--\r\n", boundary))

	// use multi-reader to defer the reading of the file data until
	// writing to the socket buffer.
	requestReader := io.MultiReader(bodyBuf, fh, closeBuf)
	fi, err := fh.Stat()
	if err != nil {
		log.Printf("Error Stating file: %s", filename)
		return nil, err
	}
	req, err := http.NewRequest("POST", targetUrl, requestReader)
	if err != nil {
		return nil, err
	}

	// Set headers for multipart, and Content Length
	req.Header.Add("Content-Type", "multipart/form-data; boundary="+boundary)
	req.ContentLength = fi.Size() + int64(bodyBuf.Len()) + int64(closeBuf.Len())

	return http.DefaultClient.Do(req)
}

func DeleteFile(url string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(req)
}