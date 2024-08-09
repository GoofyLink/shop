package file

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"shop-v2/internal/consts"
	"shop-v2/internal/dao"
	"shop-v2/internal/model"
	"shop-v2/internal/model/entity"
	"shop-v2/internal/service"
	"time"
)

type sFile struct {
}

// 注册服务  给server层生成对应服务
func init() {
	service.RegisterFile(New())
}

func New() *sFile {
	return &sFile{}
}

// 文件上传
/*
 1.定义图片上传位置
 2.校验： 上传位置是否正确安全性校验 1分钟只能上传10次
 3.定义年月日
 4.入库
 5.返回数据
*/
func (s *sFile) UploadFile(ctx context.Context, in model.FileUpdateInput) (out *model.FileUploadOutPut, err error) {
	// 定义图片上传位置
	uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
	if uploadPath == "" {
		// 配置文件
		return nil, gerror.New("读取配置文件失败，上传路径不对")
	}

	if in.Name != "" {
		in.File.Filename = in.Name
	}

	// 校验 1分钟只能上传10次
	count, err := dao.FileInfo.Ctx(ctx).
		Where(dao.FileInfo.Columns().UserId, gconv.Int(ctx.Value(consts.CtxAdminId))).
		WhereGTE(dao.FileInfo.Columns().CreatedAt, gtime.Now().Add(time.Minute)).Count()
	if err != nil {
		return nil, err
	}
	// 避免在代码中写死常量
	if count >= consts.FileMaxUploadCountMinute {
		return nil, gerror.New("上传频繁，1分钟内只能上传十次")
	}

	// 定义年月日
	dataDirName := gtime.Now().Format("Ymd")
	fileName, err := in.File.Save(gfile.Join(uploadPath, dataDirName), in.RandomName)
	if err != nil {
		return nil, err
	}

	// 入库
	data := entity.FileInfo{
		Name:   fileName,
		Src:    gfile.Join(uploadPath, dataDirName, fileName),
		Url:    "/upload/" + dataDirName + "/" + fileName, // 跟上面代码效果一样
		UserId: gconv.Int(ctx.Value(consts.CtxAdminId)),
	}
	id, err := dao.FileInfo.Ctx(ctx).Data(data).OmitEmpty().InsertAndGetId()
	if err != nil {
		return nil, err
	}

	return &model.FileUploadOutPut{
		Id:   uint(id),
		Name: data.Name,
		Src:  data.Src,
		Url:  data.Url,
	}, nil
}
