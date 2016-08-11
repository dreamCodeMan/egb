package egb

import (
	"encoding/base64"
	"io/ioutil"
	"image"
	"bytes"
	"os"
	"image/jpeg"
)

//ImageBase64ToFile convert base64 string to a image file.
func ImageBase64ToFile(datasource, filename string) error {
	filebytes, err := base64.StdEncoding.DecodeString(datasource)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, filebytes, 0666)
	if err != nil {
		return err
	}
	return nil
}

//ImageBase64ToImg convert base64 string to image.Image.
func ImageBase64ToImg(datasource string) (image.Image, error) {
	filebytes, err := base64.StdEncoding.DecodeString(datasource)
	if err != nil {
		return nil, err
	}
	filebuffer := bytes.NewBuffer(filebytes)
	image, _, err := image.Decode(filebuffer)
	if err != nil {
		return nil, err
	}
	return image, nil
}

//ImageImgToFile convert image to file.
func ImageImgToFile(image image.Image, filename string) error {
	f, err := os.Create(filename)       //创建文件
	if err != nil {
		return err
	}
	defer f.Close()                     //关闭文件
	return jpeg.Encode(f, image, nil)   //写入文件
}

//ImageImgToBase64 convert image.Image to base64 string.
func ImageImgToBase64(image image.Image) string {
	emptyBuff := bytes.NewBuffer(nil)                  //开辟一个新的空buff
	jpeg.Encode(emptyBuff, image, nil)                //img写入到buff
	dist := make([]byte, 50000)                        //开辟存储空间
	base64.StdEncoding.Encode(dist, emptyBuff.Bytes()) //buff转成base64
	return string(dist)
}

//ImageFileToBase64
func ImageFileToBase64(filename string) string {
	ff, _ := ioutil.ReadFile(filename)               //快速读文件
	bufstore := make([]byte, 50000)                       //数据缓存
	base64.StdEncoding.Encode(bufstore, ff)               // 文件转base64
	return string(bufstore)
}




