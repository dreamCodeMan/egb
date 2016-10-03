package egb

import (
	"errors"
	"image"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"math"
	"os"
	"path"

	"code.google.com/p/graphics-go/graphics"
	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
)

//ImageProcessor 图片剪裁参数
type ImageProcessor struct {
	Leftpoint int
	Toppoint  int
	Width     int
	Height    int
}

//Resize 缩放图片
//InterpolationFunction的多种分类:
//NearestNeighbor,Bilinear,Bicubic,MitchellNetravali,Lanczos2,Lanczos3
func Resize(width, height uint, img image.Image, interp resize.InterpolationFunction) image.Image {
	outImagePointer := resize.Resize(width, height, img, interp)
	return outImagePointer
}

//GetImageInfo 获取图片的宽高信息
//传入图片的存储路径 以及 图片的name    返回图片的 宽 高
func GetImageInfo(imagePath string, imageName string) (width int, height int) {
	imageRealPath := imagePath + imageName
	imageFile, err := os.Open(imageRealPath)
	defer imageFile.Close()
	if err != nil {
		panic(err)
	}
	img, _, err := image.Decode(imageFile)
	if err != nil {
		panic(err)
	}
	imgWidth := img.Bounds().Max.X
	imgHeight := img.Bounds().Max.Y
	return imgWidth, imgHeight
}

//ProcessImage 剪裁图片
//传入图片的存储路径 name  以及剪裁参数     返回图片的name
func ProcessImage(imagePath, imageName string, processConfig ImageProcessor) (isSuc bool, getOutImageName string) {
	imageExt := FileGetExt(imageName)
	if imagePath == "" || imageName == "" {
		return false, "图片地址错误"
	}
	if processConfig.Width == 0 && processConfig.Height == 0 {
		return false, "剪裁参数错误"
	}
	if imageExt != "png" && imageExt != "jpg" && imageExt != "jpeg" && imageExt != "gif" {
		return false, "不能识别该后缀图片"
	}
	var cutImageWidth = processConfig.Width
	var cutImageHeight = processConfig.Height
	var cutLeftPoint = processConfig.Leftpoint
	var cutTopPoint = processConfig.Toppoint
	imageRealPath := imagePath + imageName
	imageFile, err := os.Open(imageRealPath)
	if err != nil {
		return false, err.Error()
	}
	img, _, err := image.Decode(imageFile)
	if err != nil {
		return false, err.Error()
	}
	imgWidth := img.Bounds().Max.X
	imgHeight := img.Bounds().Max.Y

	if cutImageWidth > imgWidth && cutImageHeight > imgHeight {
		return true, imageName
	}
	//	resizeWidth, resizeHeight := GetCutImageWidthAndHeight(imgWidth, imgHeight, cutImageWidth, cutImageHeight)
	//等比缩放图片
	//	outImagePointer := Resize(uint(cutImageWidth), uint(cutImageHeight), img, resize.Lanczos3)
	//按规格剪切
	var outImagePointer image.Image
	if cutImageHeight != 0 && cutImageWidth != 0 {
		outImagePointer, _ = cutter.Crop(img, cutter.Config{
			Width:  cutImageWidth,
			Height: cutImageHeight,
			Anchor: image.Point{cutLeftPoint, cutTopPoint},
			Mode:   cutter.TopLeft, // optional, default value
		})
	}
	//判断如果裁剪点坐标存在小于0的情况，需要手动处理
	var jpg draw.Image
	jpg = image.NewRGBA(image.Rect(0, 0, cutImageWidth, cutImageHeight))
	if cutLeftPoint < 0 && cutTopPoint < 0 {
		draw.Draw(jpg, jpg.Bounds(), outImagePointer, outImagePointer.Bounds().Min.Sub(image.Pt(-cutLeftPoint, -cutTopPoint)), draw.Over)
	} else if cutLeftPoint < 0 {
		draw.Draw(jpg, jpg.Bounds(), outImagePointer, outImagePointer.Bounds().Min.Sub(image.Pt(-cutLeftPoint, 0)), draw.Over)
	} else if cutTopPoint < 0 {
		draw.Draw(jpg, jpg.Bounds(), outImagePointer, outImagePointer.Bounds().Min.Sub(image.Pt(0, -cutTopPoint)), draw.Over)
	}
	//输出图片
	outputImage, err := os.Create(imageRealPath)
	if err != nil {
		return false, err.Error()
	}
	if cutLeftPoint >= 0 && cutTopPoint >= 0 {
		switch imageExt {
		case "jpg", "jpeg":
			err = jpeg.Encode(outputImage, outImagePointer, nil)
			if err != nil {
				return false, err.Error()
			}
		case "png":
			png.Encode(outputImage, outImagePointer)
		case "gif":
			gif.Encode(outputImage, outImagePointer, nil)
		}
	} else {
		switch imageExt {
		case "jpg", "jpeg":
			err = jpeg.Encode(outputImage, jpg, nil)
			if err != nil {
				return false, err.Error()
			}
		case "png":
			png.Encode(outputImage, jpg)
		case "gif":
			gif.Encode(outputImage, jpg, nil)
		}
	}

	return true, imageName
}

//GetCutImageWidthAndHeight 通过计算图片缩放比，返回缩放后的宽高
func GetCutImageWidthAndHeight(inputWidth int, inputHeight int, cutWidth int, cutHeight int) (width int, height int) {
	if cutWidth == 0 && cutHeight == 0 {
		return inputWidth, inputHeight
	}
	if inputWidth < cutWidth || inputHeight < cutHeight {
		return inputWidth, inputHeight
	}
	if cutWidth == 0 {
		return int(float32(inputWidth) / (float32(inputHeight) / float32(cutHeight))), cutHeight
	}
	if cutHeight == 0 {
		return cutWidth, int(float32(inputHeight) / (float32(inputWidth) / float32(cutWidth)))
	}
	if float32(inputWidth)/float32(cutWidth) == float32(inputHeight)/float32(cutHeight) {
		return cutWidth, cutHeight
	}
	if float32(inputWidth)/float32(cutWidth) > float32(inputHeight)/float32(cutHeight) {
		width = int(float32(inputWidth) / (float32(inputHeight) / float32(cutHeight)))
		height = cutHeight
	} else {
		width = cutWidth
		height = int(float32(inputHeight) / (float32(inputWidth) / float32(cutWidth)))
	}
	return width, height
}

//CopyImage 拷贝一份图片
func CopyImage(imagePath, imageName, copyName string) (img string) {
	inputImageRealPath := imagePath + imageName
	r, err := os.Open(inputImageRealPath)
	if err != nil {
		panic(err)
	}
	outImageRealPath := imagePath + copyName
	w, err := os.Create(outImageRealPath)
	defer w.Close()
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(w, r)
	if err != nil {
		panic(err)
	}
	return copyName
}

//RotateImage 旋转图片  参数为角度
//旋转后图片大小
//newW = w * cos(angle) + h * sin(angle)
//newH = w * sin(angle) + h * cos(angle)
func RotateImage(imagePath, imageName string, angle float64) (err error) {
	if imagePath == "" || imageName == "" {
		e := errors.New("图片地址错误")
		return e
	}
	angle = (float64(angle) * math.Pi) / 180
	imageExt := FileGetExt(imageName)
	if imageExt != "png" && imageExt != "jpg" && imageExt != "jpeg" && imageExt != "gif" {
		e := errors.New("不能识别该后缀图片")
		return e
	}
	imgRealPath := imagePath + imageName
	src, e := LoadImage(imgRealPath)
	if e != nil {
		return e
	}
	w, h := GetImageInfo(imagePath, imageName)
	newwidth := float64(w)*math.Cos(angle) + float64(h)*math.Sin(angle)
	newheight := float64(w)*math.Sin(angle) + float64(h)*math.Cos(angle)
	dst := image.NewRGBA(image.Rect(0, 0, int(newwidth), int(newheight)))
	e = graphics.Rotate(dst, src, &graphics.RotateOptions{angle})
	if e != nil {
		return e
	}
	e = SaveImage(imgRealPath, dst)
	if e != nil {
		return e
	}
	return nil
}

// LoadImage decodes an image from a file.
func LoadImage(path string) (img image.Image, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	img, _, err = image.Decode(file)
	return
}

//SaveImage 将image.image保存进某个path  以png格式
func SaveImage(path string, img image.Image) (err error) {
	imgfile, err := os.Create(path)
	defer imgfile.Close()
	// 以PNG格式保存文件
	err = png.Encode(imgfile, img)
	if err != nil {
		return
	}
	return
}

//ImageThumbnail 图片缩略
//@params
//@return
func ImageThumbnail(file string, width uint, height uint, to string) {
	// 打开图片并解码
	fileOrigin, _ := os.Open(file)
	switch path.Ext(file) {
	case ".jpeg", ".jpg":
		origin, _ := jpeg.Decode(fileOrigin)
		defer fileOrigin.Close()
		canvas := resize.Thumbnail(width, height, origin, resize.Lanczos3)
		fileOut, err := os.Create(to)
		if err != nil {
			log.Fatal(err)
		}
		defer fileOut.Close()
		jpeg.Encode(fileOut, canvas, &jpeg.Options{Quality: 80})
	case ".png":
		origin, _ := png.Decode(fileOrigin)
		defer fileOrigin.Close()
		canvas := resize.Thumbnail(width, height, origin, resize.Lanczos3)
		fileOut, err := os.Create(to)
		if err != nil {
			log.Fatal(err)
		}
		defer fileOut.Close()
		png.Encode(fileOut, canvas)
	}
}
