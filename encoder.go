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

type Encoder struct {
	encoder C.gogpujpeg_encoder
}

func CreateEncoder() (*Encoder, error) {
	d := &Encoder{encoder: C.create_encoder()}
	if d.encoder == nil {
		return nil, errors.New("Can't create encoder")
	}
	return d, nil
}

func (e *Encoder) Encode(image []byte, param *Param, paramImage *ParamImage) (data []byte, err error) {
	size := C.int(0)
	result := C.encode(e.encoder, param.param, paramImage.param, (*C.uchar)(unsafe.Pointer(&image[0])), &size)

	if size == C.int(0) {
		return nil, errors.New("Can't encode")
	}

	data = make([]byte, int(size))

	copy(data, *(*[]byte)(unsafe.Pointer(&result)))
	C.gpujpeg_image_destroy(result)
	return
}

func (d *Encoder) Free() {
	C.gpujpeg_encoder_destroy(d.encoder)
}
