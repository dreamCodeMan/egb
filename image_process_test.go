/**
 * Created by angelina-zf on 16/12/5.
 */
package egb

import "testing"

const (
	inputImagePath = "./testdata/image/1.jpg"
	outputResizeImagePath = "./testdata/image/1_resize.jpg"
	outputThumbImagePath = "./testdata/image/1_thumb.jpg"
)

func TestResizeImage(t *testing.T) {
	err := ResizeImage(inputImagePath, 200, 200, outputResizeImagePath)
	if err != nil {
		t.Fatal("error : " + err.Error())
	}
}

func TestThumbImage(t *testing.T) {
	err := ThumbImage(inputImagePath, 200, 200, outputThumbImagePath)
	if err != nil {
		t.Fatal("error : " + err.Error())
	}
}

func TestGetImageInfo(t *testing.T) {
	w, h, err := GetImageInfo(outputThumbImagePath)
	if err != nil {
		t.Fatal("error : " + err.Error())
	}
	if w != 200 {
		t.Errorf("width should be 200 but get %d", w)
	}
	if h != 112 {
		t.Errorf("height should be 112 but get %d", h)
	}
}
