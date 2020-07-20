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

type ColorSpace uint

const (
	GPUJPEG_NONE                ColorSpace = 0
	GPUJPEG_RGB                            = 1
	GPUJPEG_YCBCR_BT601                    = 2
	GPUJPEG_YCBCR_BT601_256LVLS            = 3
	GPUJPEG_YCBCR_JPEG                     = GPUJPEG_YCBCR_BT601_256LVLS
	GPUJPEG_YCBCR_BT709                    = 4
	GPUJPEG_YCBCR                          = GPUJPEG_YCBCR_BT709
	GPUJPEG_YUV                            = 5
)

func (d ColorSpace) String() string {
	return [...]string{"None", "RGB", "YCBCR_BT601", "YCBCR_BT601_256LVLS/GPUJPEG_YCBCR_JPEG", "GPUJPEG_YCBCR_BT709/GPUJPEG_YCBCR", "GPUJPEG_YUV"}[d]
}

type PixelFormat int

const (
	GPUJPEG_PIXFMT_NONE PixelFormat = -1

	/// 8bit unsigned samples, 1 component
	GPUJPEG_U8 = 1

	/// 8bit unsigned samples, 3 components, 4:4:4 sampling,
	/// sample order: comp#0 comp#1 comp#2, interleaved
	GPUJPEG_444_U8_P012 = 2

	/// 8bit unsigned samples, 3 components, 4:4:4, planar
	GPUJPEG_444_U8_P0P1P2 = 3

	/// 8bit unsigned samples, 3 components, 4:2:2,
	/// order of samples: comp#1 comp#0 comp#2 comp#0, interleaved
	GPUJPEG_422_U8_P1020 = 4

	/// 8bit unsigned samples, planar, 3 components, 4:2:2, planar
	GPUJPEG_422_U8_P0P1P2 = 5

	/// 8bit unsigned samples, planar, 3 components, 4:2:0, planar
	GPUJPEG_420_U8_P0P1P2 = 6

	/// 8bit unsigned samples, 3 components, each pixel padded to 32bits
	/// with zero byte, 4:4:4 sampling, interleaved
	GPUJPEG_444_U8_P012Z = 7

	/// 8bit unsigned samples, 3 components, each pixel padded to 32bits
	/// with all-one bits, 4:4:4 sampling, interleaved
	GPUJPEG_444_U8_P012A = 8
)

func (d PixelFormat) String() string {
	return [...]string{"None", "U8", "444_U8_P012", "444_U8_P0P1P2", "422_U8_P1020", "422_U8_P0P1P2", "420_U8_P0P1P2", "444_U8_P012Z", "444_U8_P012A"}[d]
}

type ParamImage struct {
	param C.gogpujpeg_image_parameters
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

func (p *ParamImage) ColorSpace() ColorSpace {
	return ColorSpace(p.param.color_space)
}

func (p *ParamImage) PixelFormat() PixelFormat {
	return PixelFormat(p.param.pixel_format)
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

func (p *ParamImage) SetColorSpace(c ColorSpace) {
	p.param.color_space = uint32(c)
}

func (p *ParamImage) SetPixelFormat(pf PixelFormat) {
	p.param.pixel_format = int32(pf)
}

func SetImageParam() (p *ParamImage) {
	p = &ParamImage{param: C.malloc_gpujpeg_image_parameters()}
	C.gpujpeg_image_set_default_parameters(p.param)
	return
}

func ReadImageInfo(image []byte) (*ParamImage, int, error) {
	p := &ParamImage{param: C.malloc_gpujpeg_image_parameters()}

	segment_count := C.int(0)

	if C.gpujpeg_decoder_get_image_info((*C.uchar)(unsafe.Pointer(&image[0])), C.int(len(image)), p.param, &segment_count) != C.int(0) {
		return nil, 0, errors.New("Can't decode image info")
	}

	return p, int(segment_count), nil
}

func (p *ParamImage) Free() {
	C.free_gpujpeg_image_parameters(p.param)
	p.param = nil
	p = nil
}
