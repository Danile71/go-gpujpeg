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
)

func InitDevice(deviceId int) error {
	if C.gpujpeg_init_device(C.int(deviceId), 0) != C.int(0) {
		return errors.New("Can't init device")
	}
	return nil
}
