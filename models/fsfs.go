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

// 上传文件到fastDFS系统
//func UploadFile(fileName string) string {
//	client, err := fdfs_client.NewClientWithConfig("D:\\iHome\\iHome\\conf\\fastdfs.conf")
//	if err != nil {
//		fmt.Println("打开fast客户端失败", err.Error())
//		return ""
//	}
//	defer client.Destory()
//	fileId, err := client.UploadByFilename(fileName)
//	if err != nil {
//		fmt.Println("上传文件失败", err.Error())
//		return ""
//	}
//	return fileId
//}
//
//// 下载文件
//func DownLoadFile(fileId, tempFile string) {
//	client, err := fdfs_client.NewClientWithConfig("D:\\iHome\\iHome\\conf\\fastdfs.conf")
//	if err != nil {
//		fmt.Println("打开fast客户端失败", err.Error())
//		return
//	}
//	defer client.Destory()
//	if err = client.DownloadToFile(fileId, tempFile, 0, 0); err != nil {
//		fmt.Println("下载文件失败", err.Error())
//		return
//	}
//}
//
//// 删除
//func DeleteFile(fileId string) {
//	client, err := fdfs_client.NewClientWithConfig("D:\\iHome\\iHome\\conf\\fastdfs.conf")
//	if err != nil {
//		fmt.Println("打开fast客户端失败", err.Error())
//		return
//	}
//	defer client.Destory()
//	if err = client.DeleteFile(fileId); err != nil {
//		fmt.Println("删除文件失败", err.Error())
//		return
//	}
//}
//
//// 从配置文件中读取服务器的ip和端口配置
//func FileServerAddr() string {
//	file, err := os.Open("D:\\iHome\\iHome\\conf\\fastdfs.conf")
//	if err != nil {
//		fmt.Println(err)
//		return ""
//	}
//	reader := bufio.NewReader(file)
//	for {
//		line, err := reader.ReadString('\n')
//		line = strings.TrimSpace(line)
//		if err != nil {
//			return ""
//		}
//		line = strings.TrimSuffix(line, "\n")
//		str := strings.SplitN(line, "=", 2)
//		switch str[0] {
//		case "http_server_port":
//			return str[1]
//		}
//	}
//}
