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

type Decoder struct {
	decoder *C.gpujpeg_decoder
}

func CreateDecoder() (*Decoder, error) {
	d := &Decoder{decoder: C.create_decoder()}
	if d.decoder == nil {
		return nil, errors.New("Can't create decoder")
	}
	return d, nil
}

func (d *Decoder) Init(param *Param, paramImage *ParamImage) {
	C.gpujpeg_decoder_init(d.decoder, &param.param, &paramImage.param)
}

func (d *Decoder) SetOutput() {
	C.gpujpeg_decoder_set_output_format(d.decoder, C.GPUJPEG_RGB,
		C.GPUJPEG_444_U8_P012)
	// or eg. GPUJPEG_YCBCR_JPEG and GPUJPEG_422_U8_P1020
}

func (d *Decoder) Decode(image []byte) (data []byte, err error) {
	var decoder_output C.gpujpeg_decoder_output
	C.gpujpeg_decoder_output_set_default(&decoder_output)

	if C.gpujpeg_decoder_decode(d.decoder, (*C.uchar)(unsafe.Pointer(&image[0])), C.int(len(image)), &decoder_output) != C.int(0) {
		return nil, errors.New("Can't decode")
	}

	data = make([]byte, int(decoder_output.data_size))

	copy(data, *(*[]byte)(unsafe.Pointer(&decoder_output.data)))

	return
}

func (d *Decoder) Free() {
	C.gpujpeg_decoder_destroy(d.decoder)
}
