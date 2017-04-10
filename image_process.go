package egb

import (
	"errors"
	"image"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"math"
	"os"
	"path"

	"code.google.com/p/graphics-go/graphics"
	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
)

//ImageProcessor 图片剪裁参数
type ImageProcessor struct {
	//距离左边距离
	LeftPoint int
	//距离上边距离
	TopPoint int
	//裁剪后的宽
	Width int
	//剪裁后的高
	Height int
}

//ResizeImage 图片缩放
//@params	inputImagePath width height outputFileName
//@return	error
func ResizeImage(inputImagePath string, width, height uint, outputFileName string) error {
	inputImage, err := LoadImage(inputImagePath)
	if err != nil {
		return err
	}
	outputImage := resize.Resize(width, height, inputImage, resize.Lanczos3)
	err = SaveImage(outputFileName, outputImage)
	if err != nil {
		return err
	}
	return nil
}

//GetImageInfo 获取图片的宽高信息
//@params	inputFilePath
//@return	(int,int,error)
func GetImageInfo(inputFilePath string) (int, int, error) {
	if !FileExists(inputFilePath) {
		return 0, 0, errors.New("图片文件不存在,请检查地址")
	}
	imageFile, err := os.Open(inputFilePath)
	if err != nil {
		return 0, 0, err
	}
	defer imageFile.Close()
	img, _, err := image.Decode(imageFile)
	if err != nil {
		return 0, 0, err
	}
	imgWidth := img.Bounds().Max.X
	imgHeight := img.Bounds().Max.Y
	return imgWidth, imgHeight, nil
}

//ProcessImage 剪裁图片
//传入图片的存储路径 name  以及剪裁参数     返回图片的name
func ProcessImage(inputFilePath string, processConfig ImageProcessor) error {
	if !FileExists(inputFilePath) {
		return errors.New("图片文件不存在,请检查地址")
	}
	if processConfig.Width == 0 && processConfig.Height == 0 {
		return errors.New("剪裁参数错误")
	}
	imageExt := FileGetExt(inputFilePath)
	if imageExt != "png" && imageExt != "jpg" && imageExt != "jpeg" && imageExt != "gif" {
		return errors.New("图片后缀必须是png,jpg,jpeg,gif中的一种")
	}
	var cutImageWidth = processConfig.Width
	var cutImageHeight = processConfig.Height
	var cutLeftPoint = processConfig.LeftPoint
	var cutTopPoint = processConfig.TopPoint
	imageFile, err := os.Open(inputFilePath)
	if err != nil {
		return err
	}
	defer imageFile.Close()
	img, _, err := image.Decode(imageFile)
	if err != nil {
		return err
	}
	imgWidth := img.Bounds().Max.X
	imgHeight := img.Bounds().Max.Y
	if cutImageWidth > imgWidth && cutImageHeight > imgHeight {
		return nil
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
			Anchor: image.Point{X: cutLeftPoint, Y: cutTopPoint},
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
	outputImage, err := os.Create(inputFilePath)
	if err != nil {
		return err
	}
	defer outputImage.Close()
	if cutLeftPoint >= 0 && cutTopPoint >= 0 {
		switch imageExt {
		case "jpg", "jpeg":
			err = jpeg.Encode(outputImage, outImagePointer, nil)
			if err != nil {
				return err
			}
		case "png":
			err = png.Encode(outputImage, outImagePointer)
			if err != nil {
				return err
			}
		case "gif":
			err = gif.Encode(outputImage, outImagePointer, nil)
			if err != nil {
				return err
			}
		}
	} else {
		switch imageExt {
		case "jpg", "jpeg":
			err = jpeg.Encode(outputImage, jpg, nil)
			if err != nil {
				return err
			}
		case "png":
			err = png.Encode(outputImage, jpg)
			if err != nil {
				return err
			}
		case "gif":
			err = gif.Encode(outputImage, jpg, nil)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

//RotateImage 旋转图片  参数为角度(顺时针)
//@params	inputImagePath outputImagePath angle
//@return	error
func RotateImage(inputImagePath, outputImagePath string, angle float64) error {
	if !FileExists(inputImagePath) {
		return errors.New("图片地址错误,请检查")
	}
	angle = (float64(angle) * math.Pi) / 180
	imageExt := FileGetExt(inputImagePath)
	if imageExt != "png" && imageExt != "jpg" && imageExt != "jpeg" && imageExt != "gif" {
		return errors.New("图片后缀必须是png,jpg,jpeg,gif中的一种")
	}
	srcImage, err := LoadImage(inputImagePath)
	if err != nil {
		return err
	}
	w, h, err := GetImageInfo(inputImagePath)
	if err != nil {
		return err
	}
	newWidth := float64(w)*math.Cos(angle) + float64(h)*math.Sin(angle)
	newHeight := float64(w)*math.Sin(angle) + float64(h)*math.Cos(angle)
	dst := image.NewRGBA(image.Rect(0, 0, int(newWidth), int(newHeight)))
	err = graphics.Rotate(dst, srcImage, &graphics.RotateOptions{Angle: angle})
	if err != nil {
		return err
	}
	err = SaveImage(outputImagePath, dst)
	if err != nil {
		return err
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
//@params	path img
//@return	error
func SaveImage(path string, img image.Image) error {
	imgFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer imgFile.Close()
	// 以PNG格式保存文件
	err = png.Encode(imgFile, img)
	if err != nil {
		return err
	}
	return nil
}

//ImageThumbnail 生成缩略图  最大宽高,不会改变图片的比例
//@params	inputFile maxWidth maxHeight outputFile
//@return	error
func ThumbImage(inputFileName string, maxWidth, maxHeight uint, outputFileName string) error {
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		return err
	}
	defer inputFile.Close()
	switch path.Ext(inputFileName) {
	case ".jpg", ".jpeg":
		inputImage, err := jpeg.Decode(inputFile)
		if err != nil {
			return err
		}
		newImage := resize.Thumbnail(maxWidth, maxHeight, inputImage, resize.Lanczos3)
		outputFile, err := os.Create(outputFileName)
		if err != nil {
			return err
		}
		defer outputFile.Close()
		err = jpeg.Encode(outputFile, newImage, &jpeg.Options{Quality: 80})
		if err != nil {
			return err
		}
	case ".png":
		inputImage, err := png.Decode(inputFile)
		if err != nil {
			return err
		}
		newImage := resize.Thumbnail(maxWidth, maxHeight, inputImage, resize.Lanczos3)
		outputFile, err := os.Create(outputFileName)
		if err != nil {
			return err
		}
		defer outputFile.Close()
		err = png.Encode(outputFile, newImage)
		if err != nil {
			return err
		}
	default:
		return errors.New("图片文件后缀错误,请传入png,jpg,jpeg图片")
	}
	return nil
}
