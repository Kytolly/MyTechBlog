package server

import (
	"mytechblog/utils"
	"context"
    "net/http"
    "net/url"
    "log/slog"
	"github.com/tencentyun/cos-go-sdk-v5"
	"mime/multipart"
	msg "mytechblog/utils/errormsg"
)

var APPID = utils.APPID
var SecretId = utils.SecretId
var SecretKey = utils.SecretKey
var Bucket = utils.Bucket
var CosServer = utils.CosServer

func UploadFile(file multipart.File, header *multipart.FileHeader)(string, int){
    u, _ := url.Parse(CosServer)
    b := &cos.BaseURL{BucketURL: u}
    c := cos.NewClient(b, &http.Client{
        Transport: &cos.AuthorizationTransport{
            SecretID:  SecretId, 
            SecretKey: SecretKey, 
        },
    })

    key := "blog-img/" + header.Filename
    opt := &cos.ObjectPutOptions{
        // ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
        //     ContentType: "text/html",
        // },
        ACLHeaderOptions: &cos.ACLHeaderOptions{
            // 如果不是必要操作，建议上传文件时不要给单个文件设置权限，避免达到限制。若不设置默认继承桶的权限。
            XCosACL: "private",
        },
    }
    _, err := c.Object.Put(context.Background(), key, file, opt)
    if err != nil {
        slog.Info("Failed to Upload Object", "error", err)
		return "", msg.ERROR
    }
	url := CosServer + key 
	return url, msg.SUCCESS
}