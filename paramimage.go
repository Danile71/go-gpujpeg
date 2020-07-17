package gpujpeg

/*
#cgo LDFLAGS: -lgpujpeg
#include <libgpujpeg/gpujpeg.h>
#include <libgpujpeg/gpujpeg_common.h>
#include <libgpujpeg/gpujpeg_encoder.h>
#include <libgpujpeg/gpujpeg_decoder.h>
#include <libgpujpeg/gpujpeg_type.h>
#include <libgpujpeg/gpujpeg_version.h>
#include "wrapper.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

type PixelFormat int

const (
	GPUJPEG_PIXFMT_NONE PixelFormat = -1

	/// 8bit unsigned samples, 1 component
	GPUJPEG_U8

	/// 8bit unsigned samples, 3 components, 4:4:4 sampling,
	/// sample order: comp#0 comp#1 comp#2, interleaved
	GPUJPEG_444_U8_P012

	/// 8bit unsigned samples, 3 components, 4:4:4, planar
	GPUJPEG_444_U8_P0P1P2

	/// 8bit unsigned samples, 3 components, 4:2:2,
	/// order of samples: comp#1 comp#0 comp#2 comp#0, interleaved
	GPUJPEG_422_U8_P1020

	/// 8bit unsigned samples, planar, 3 components, 4:2:2, planar
	GPUJPEG_422_U8_P0P1P2

	/// 8bit unsigned samples, planar, 3 components, 4:2:0, planar
	GPUJPEG_420_U8_P0P1P2

	/// 8bit unsigned samples, 3 components, each pixel padded to 32bits
	/// with zero byte, 4:4:4 sampling, interleaved
	GPUJPEG_444_U8_P012Z

	/// 8bit unsigned samples, 3 components, each pixel padded to 32bits
	/// with all-one bits, 4:4:4 sampling, interleaved
	GPUJPEG_444_U8_P012A
)

func (d PixelFormat) String() string {
	return [...]string{"None", "U8", "444_U8_P012", "444_U8_P0P1P2", "422_U8_P1020", "422_U8_P0P1P2", "420_U8_P0P1P2", "444_U8_P012Z", "444_U8_P012A"}[d]
}

type ParamImage struct {
	param C.gpujpeg_image_parameters
}

func (p *ParamImage) Width() int {
	return int(p.param.width)
}

func (p *ParamImage) Height() int {
	return int(p.param.height)
}

func (p *ParamImage) CompCount() int {
	return int(p.param.comp_count)
}

func (p *ParamImage) ColorSpace() int {
	return int(p.param.color_space)
}

func (p *ParamImage) PixelFormat() int {
	return int(p.param.pixel_format)
}

func (p *ParamImage) SetWidth(w int) {
	p.param.width = C.int(w)
}

func (p *ParamImage) SetHeight(h int) {
	p.param.height = C.int(h)
}

func (p *ParamImage) SetCompCount(c int) {
	p.param.comp_count = C.int(c)
}

func (p *ParamImage) SetColorSpace(c uint32) {
	p.param.color_space = uint32(c)
}

func (p *ParamImage) SetPixelFormat(pf int) {
	p.param.pixel_format = int32(pf)
}

func SetImageParam() *ParamImage {
	p := &ParamImage{}
	C.gpujpeg_image_set_default_parameters(&p.param)
	return p
}

func ReadImageInfo(image []byte) (*ParamImage, int, error) {
	p := &ParamImage{}

	segment_count := C.int(0)

	if C.gpujpeg_decoder_get_image_info((*C.uchar)(unsafe.Pointer(&image[0])), C.int(len(image)), &p.param, &segment_count) != C.int(0) {
		return nil, 0, errors.New("Can't decode image info")
	}

	return p, int(segment_count), nil
}

func SetCustomParam(w, h, c int) *ParamImage {
	p := &ParamImage{}
	C.gpujpeg_image_set_default_parameters(&p.param)
	p.param.width = C.int(w)
	p.param.height = C.int(h)
	p.param.comp_count = C.int(c)
	return p
}
