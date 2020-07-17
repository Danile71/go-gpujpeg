package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/Danile71/go-gpujpeg"
)

const file = "test.jpg"

func main() {
	err := gpujpeg.InitDevice(0)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	imageParam, count, err := gpujpeg.ReadImageInfo(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("width", imageParam.Width(), imageParam.Height(), imageParam.PixelFormat(), imageParam.ColorSpace(), imageParam.CompCount(), count)
	start := time.Now()

	decoder, err := gpujpeg.CreateDecoder()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer decoder.Free()

	decoder.SetOutput(gpujpeg.GPUJPEG_YCBCR_JPEG, gpujpeg.GPUJPEG_420_U8_P0P1P2)

	result, err := decoder.Decode(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("result", len(result), time.Since(start))

}
