package upload

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"os"
)

/*
1. 获取本地文件
2. 指定上传目录
3. 上传服务器临时文件夹
4. 上传七牛云等平台
5. 删除服务器临时文件
*/

func UploadImgToCloud(ctx context.Context, file *ghttp.UploadFile) (url string, err error) {
	//定义上传目录
	dirPath := "/tmp/"
	saveFileName, err := file.Save(dirPath, true)
	if err != nil {
		return "", err
	}

	//定义本地文件路径
	localFile := dirPath + saveFileName
	//获取七牛云配置参数
	bucket := g.Cfg().MustGet(ctx, "qiniu.bucket").String()
	accessKey := g.Cfg().MustGet(ctx, "qiniu.accessKey").String()
	secretKey := g.Cfg().MustGet(ctx, "qiniu.secretKey").String()
	//对接sdk
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	// 生成token
	mac := qbox.NewMac(accessKey, secretKey)
	// 上传token
	token := putPolicy.UploadToken(mac)
	// 七牛云上传配置
	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuanan
	cfg.UseHTTPS = true
	cfg.UseCdnDomains = false //CDN图片加速
	// 构建表单上的对象
	formUploader := storage.NewFormUploader(&cfg)
	//上传结果结构体
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	key := saveFileName
	//七牛云表单上传
	err = formUploader.PutFile(ctx, &ret, token, key, localFile, &putExtra)
	if err != nil {
		return "", err
	}
	fmt.Println(ret.Key, ret.Hash, ret.PersistentID)
	// 删除本地临时文件
	err = os.RemoveAll(localFile)
	if err != nil {
		return "", err
	}
	// 返回数据
	url = g.Cfg().MustGet(ctx, "qiniu.url").String() + ret.Key
	return url, nil
}
