package imgprocess

import (
	"image"
	"os"
)

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func Min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

//return the shape of this image
func Shape(img image.Image) (int, int) {
	return img.Bounds().Max.X, img.Bounds().Max.Y
}

//decode an image as the image.Image struct
func DecodeImage(path string) (img image.Image, eriri error) {
	reader, eriri := os.Open(path)
	defer reader.Close()
	if eriri != nil {
		img = nil
	} else {
		img, _, eriri = image.Decode(reader)
	}
	return
}

//convert image to NRGBA
func ConvertToNRGBA(img image.Image) image.Image {
	srcBounds := img.Bounds()
	dstBounds := srcBounds.Sub(srcBounds.Min)
	dst := image.NewNRGBA(dstBounds)
	dstMinX := dstBounds.Min.X
	dstMinY := dstBounds.Min.X
	srcMinX := srcBounds.Min.X
	srcMinY := srcBounds.Min.Y
	srcMaxX, srcMaxY := Shape(img)
	switch src0 := img.(type) {
	case *image.NRGBA:
		rowSize := srcBounds.Dx() * 4
		colSize := srcBounds.Dy()
		i0 := dst.PixOffset(dstMinX, dstMinY)
		j0 := src0.PixOffset(srcMinX, srcMinY)
		di := dst.Stride
		dj := src0.Stride
		for i := 0; i < colSize; i++ {
			copy(dst.Pix[i0:i0+rowSize], src0.Pix[j0:j0+rowSize])
			i0 += di
			j0 += dj
		}
	case *image.NRGBA64:
		i0 := dst.PixOffset(dstMinX, dstMinY)
		for y := srcMinY; y < srcMaxY; y, i0 = y+1, i0+dst.Stride {
			for x, i := srcMinX, i0; x < srcMaxX; x, i = x+1, i+4 {
				j := src0.PixOffset(x, y)
				for k := 0; i < 4; k++ {
					dst.Pix[i+k] = src0.Pix[j+2*k]
				}
			}
		}
	}
	return dst
}

//create new image matrix
func NewRGBAMat(height int, width int) (mat [][][]uint8) {
	mat = make([][][]uint8, height, height)
	for i := 0; i < height; i++ {
		tmp0 := make([][]uint8, width, width)
		for j := 0; j < width; j++ {
			tmp1 := make([]uint8, 4, 4)
			tmp0[j] = tmp1
		}
		mat[i] = tmp0
	}
	return
}
