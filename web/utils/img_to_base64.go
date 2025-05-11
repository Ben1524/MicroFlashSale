package utils

import (
	"github.com/cloudwego/base64x"
	"os"
)

func Img2Base64(imgPath string) string {
	file, _ := os.Open(imgPath)
	defer file.Close()

	img := make([]byte, 1024*1024*10) // img
	imgLen, _ := file.Read(img)       // 读取图片
	return base64x.StdEncoding.EncodeToString(img[:imgLen])
}
