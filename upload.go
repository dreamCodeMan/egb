package egb

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

/**
文件上传类
面向对象的写法
example:
rootpath := "文件的存储路径"("./static/xx/xx/xx")
upload := NewUpload(rootpath)
uploadinfo := upload.UploadFile(this.Ctx.Request, name)
*/

//上传需要初始化的对象
type upload struct {
	Rootpath string //上传文件后存储文件的目录
}

//UploadReturnInfo 执行上传后返回的信息
type UploadReturnInfo struct {
	Err      error  //错误，如果存在，则上传失败
	URL      string //文件地址 rootpath-1(/static/xx/xx/xx) + filename + ext(直接用于显示用,即直接存储在数据库中的)
	Filename string //文件名 filename
	Ext      string //文件后缀 ext
}

var (
	allowSize = 1024 //1G
)

//NewUpload 构造方法
//必须使用这个方法初始化一个Upload对象
func NewUpload(rootpath string) *upload {
	upload := new(upload)
	upload.Rootpath = rootpath
	return upload
}

/*
	检验大小
*/
func checkSize(size int64) bool {
	if size/(1024*1024) > int64(allowSize) {
		return false
	}
	return true
}

//Sizer 获取文件大小
type Sizer interface {
	Size() int64
}

//Stat 获取文件信息的接口
type Stat interface {
	Stat() (os.FileInfo, error)
}

//UploadFile 上传文件
func (u *upload) UploadFile(request *http.Request, name string) UploadReturnInfo {
	file, handle, err := request.FormFile(name)
	defer file.Close()
	if err != nil {
		return UploadReturnInfo{
			Err: err,
		}
	}
	//获取文件名
	filename := handle.Filename
	inputfileprefix := FileGetPrefix(filename)
	outputFileName := inputfileprefix
	//检验大小
	if sizeInterface, ok := file.(Sizer); ok {
		if !checkSize(sizeInterface.Size()) {
			return UploadReturnInfo{
				Err: errors.New("文件大小太大,超过限制"),
			}
		}
	} else if statInterface, ok := file.(Stat); ok {
		fileInfo, _ := statInterface.Stat()
		if !checkSize(fileInfo.Size()) {
			return UploadReturnInfo{
				Err: errors.New("文件大小太大,超过限制"),
			}
		}
	} else {
		return UploadReturnInfo{
			Err: errors.New("无法获取上传文件大小"),
		}
	}
	//获取文件名后缀
	fileExt := FileGetExt(filename)
	fileExt = strings.ToLower(fileExt)
	//使用MD5值作为新的文件名
	/*
		md5h := md5.New()
		_, _ = io.Copy(md5h, inputfile)
		outputFileName = fmt.Sprintf("%x", md5h.Sum(nil))
	*/
	//向文件名后面加上时间戳 保证唯一
	outputFileName = outputFileName + TimeNowUnix()
	//首先创建目录
	os.MkdirAll(StringSubStr(u.Rootpath, 2, len(u.Rootpath)-2), os.ModePerm)
	//拷贝到新文件
	outputfile, err := os.OpenFile(u.Rootpath+outputFileName+"."+fileExt, os.O_WRONLY|os.O_CREATE, 0666)
	defer outputfile.Close()
	if err != nil {
		return UploadReturnInfo{
			Err: err,
		}
	}
	_, copyerr := io.Copy(outputfile, file)
	if copyerr != nil {
		return UploadReturnInfo{
			Err: err,
		}
	}
	return UploadReturnInfo{
		Err:      nil,
		URL:      StringSubStr(u.Rootpath, 1, len(u.Rootpath)-1) + fmt.Sprintf("%s.%s", outputFileName, fileExt),
		Filename: outputFileName,
		Ext:      fileExt,
	}
}
