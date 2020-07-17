package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/Danile71/go-gpujpeg"
)

const file = "test.rgb"

func main() {
	err := gpujpeg.InitDevice(0)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	param := gpujpeg.SetParam()
	defer param.Free()

	param.SetInterleaved(1)
	param.SetRestartInterval(16)
	param.SetColorSpaceInternal(gpujpeg.GPUJPEG_YCBCR_BT601_256LVLS)
	param.SetQuality(100)

	paramImage := gpujpeg.SetImageParam()
	defer paramImage.Free()

	paramImage.SetCompCount(3)
	paramImage.SetWidth(2442)
	paramImage.SetHeight(1342)
	paramImage.SetColorSpace(gpujpeg.GPUJPEG_RGB)
	paramImage.SetPixelFormat(gpujpeg.GPUJPEG_U8)

	// param.SetSamplingFactor(0, 4, 4)
	// param.SetSamplingFactor(1, 2, 1)
	// param.SetSamplingFactor(2, 2, 1)

	start := time.Now()
	defer func() {
		fmt.Println("result", time.Since(start))
	}()

	encoder, err := gpujpeg.CreateEncoder()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer encoder.Free()

	result, err := encoder.Encode(data, param, paramImage)
	if err != nil {
		fmt.Println(err)
		return
	}
	ioutil.WriteFile("result.jpeg", result, os.ModePerm)

}
