package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

	imageParam, _, err := gpujpeg.ReadImageInfo(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer imageParam.Free()

	fmt.Println("image width", imageParam.Width(), " height", imageParam.Height())
	start := time.Now()

	decoder, err := gpujpeg.CreateDecoder()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer decoder.Free()

	//decoder.SetOutput(gpujpeg.GPUJPEG_RGB, gpujpeg.GPUJPEG_444_U8_P012)

	result, err := decoder.Decode(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(time.Since(start))
	ioutil.WriteFile("result.rgb", result, os.ModePerm)
}
