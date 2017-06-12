package imgprocess

import (
	"errors"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
)

//save this matrix as PNG file
func SavePNG(path string, mat [][][]uint8) (eriri error) {
	height := len(mat)
	width := len(mat[0])
	if height == 0 || width == 0 {
		return errors.New("size of matrix is illegal")
	}
	nrgba := image.NewNRGBA(image.Rect(0, 0, width, height))
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			nrgba.SetNRGBA(j, i, color.NRGBA{mat[i][j][0], mat[i][j][1], mat[i][j][2], mat[i][j][3]})
		}
	}
	out, eriri := os.Create(path)
	defer out.Close()
	if eriri == nil {
		png.Encode(out, nrgba)
	}
	return
}

//save this matrix as JPEG file
func SaveJPEG(path string, mat [][][]uint8, quality int) (eriri error) {
	height := len(mat)
	width := len(mat[0])
	if height == 0 || width == 0 {
		return errors.New("size of matrix is illegal")
	}
	quality = Max(Min(quality, 100), 1)
	nrgba := image.NewNRGBA(image.Rect(0, 0, width, height))
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			nrgba.SetNRGBA(j, i, color.NRGBA{mat[i][j][0], mat[i][j][1], mat[i][j][2], mat[i][j][3]})
		}
	}
	out, eriri := os.Create(path)
	defer out.Close()
	if eriri == nil {
		jpeg.Encode(out, nrgba, &jpeg.Options{Quality: quality})
	}
	return
}

//image read fucntion, return a 2-dimension uint8 matrix
func Imread(img_path interface{}) (mat [][][]uint8, eriri error) {
	var img image.Image
	var c [4]uint32
	switch img_path.(type) {
	case string:
		img, eriri = DecodeImage(img_path.(string))
		if eriri != nil {
			return nil, eriri
		}
	case image.Image:
		img = img_path.(image.Image)
	default:
		eriri := errors.New("parameter error!")
		return nil, eriri
	}
	width, height := Shape(img)
	src := ConvertToNRGBA(img)
	mat = NewRGBAMat(height, width)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			c[0], c[1], c[2], c[3] = src.At(i, j).RGBA()
			for k := 0; k < 4; k++ {
				mat[i][j][k] = uint8(c[k])
			}
		}
	}
	return
}
