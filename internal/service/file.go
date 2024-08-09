// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"shop-v2/internal/model"
)

type (
	IFile interface {
		// 文件上传
		/*
		   1.定义图片上传位置
		   2.校验： 上传位置是否正确安全性校验 1分钟只能上传10次
		   3.定义年月日
		   4.入库
		   5.返回数据
		*/
		UploadFile(ctx context.Context, in model.FileUpdateInput) (out *model.FileUploadOutPut, err error)
	}
)

var (
	localFile IFile
)

func File() IFile {
	if localFile == nil {
		panic("implement not found for interface IFile, forgot register?")
	}
	return localFile
}

func RegisterFile(i IFile) {
	localFile = i
}
