package seaweedfs

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/geekappio/itonchain/app/config"
	"github.com/linxGnu/goseaweedfs"
	"github.com/linxGnu/goseaweedfs/model"
)

var SeaWeedFS *goseaweedfs.Seaweed

// InitSeaWeedFS 根据配置文件的配置，初始化SeaWeedFS
func InitSeaWeedFS() {
	// check master url
	sfsConfig := config.App.SeaWeedFS
	var filers []string
	if sfsConfig.SwFSFilerUrls != "" {
		filers = strings.Split(sfsConfig.SwFSFilerUrls, ";")
	}
	SeaWeedFS = goseaweedfs.NewSeaweed(sfsConfig.SwFSSchema,
		sfsConfig.SwFSMasterUrl, filers, sfsConfig.ChunkSize*1024*1024,
		sfsConfig.Duration*time.Minute)
}

// SubmitRourceContent 提交文件内容到seaweedfs
func SubmitRourceContent(fileName string, content []byte, args url.Values) (result *model.SubmitResult, err error) {
	filePart := model.NewFilePartFromReader(bytes.NewBuffer(content), fileName, int64(len(content)))
	return SeaWeedFS.SubmitFilePart(filePart, args)
}

// UploadResourceContent 上传文件内容到seaweedsfs
func UploadResourceContent(fileName string, content []byte, args url.Values) (cm *model.ChunkManifest, fileID string, err error) {
	filePart := model.NewFilePartFromReader(bytes.NewBuffer(content), fileName, int64(len(content)))
	return SeaWeedFS.UploadFilePart(filePart)
}

// UpdateResourceContent 更新文件内容
func UpdateResourceContent(fileName string, fid string, content []byte, args url.Values, delFirst bool) (fileID string, err error) {
	filePart := model.NewFilePartFromReader(bytes.NewBuffer(content), fileName, int64(len(content)))
	filePart.FileID = fid
	return SeaWeedFS.ReplaceFilePart(filePart, delFirst)
}

// Read resource content from SeaWeedFS
func DownloadResourceContent(fid string) ([]byte, error){
	fullurl, err :=SeaWeedFS.LookupFileID(fid, nil, true)
	if err != nil {
		return nil, err
	}

	resp, respErr := http.Get(fullurl)
	if respErr != nil {
		return nil, respErr
	}

	return ioutil.ReadAll(resp.Body)
}