package models

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/sharabao13/fdfs_client"
)

func FdfsUploadByFileName(fileName string) (groupName, fileId string, err error) {
	logs.Info("================== Start Connection ==================")
	fdfsClient, err := fdfs_client.NewFdfsClient("D:\\iHome\\iHome\\conf\\client.conf")
	if err != nil {
		logs.Info("new fdfsclient error", err.Error())
		return "", "", err
	}
	logs.Info("================== Start Upload ==================")
	uploadResponse, err := fdfsClient.UploadByFilename(fileName)
	if err != nil {
		logs.Info("new fdfsclient error", err.Error())
		return "", "", err
	}
	logs.Info("==================", uploadResponse.GroupName, "==================")
	logs.Info("==================", uploadResponse.RemoteFileId, "==================")
	return uploadResponse.GroupName, uploadResponse.RemoteFileId, nil
}
