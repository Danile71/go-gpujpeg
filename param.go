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

type Param struct {
	param C.gpujpeg_parameters
}

func SetParam() *Param {
	p := &Param{}
	C.gpujpeg_set_default_parameters(&p.param)
	// p.param.restart_interval = 1
	// p.param.interleaved = 1
	// p.param.color_space_internal = 82
	return p
}

func (p *Param) RestartInterval() int {
	return int(p.param.restart_interval)
}

func (p *Param) Interleaved() int {
	return int(p.param.interleaved)
}

func (p *Param) ColorSpaceInternal() int {
	return int(p.param.color_space_internal)
}

func (p *Param) SetRestartInterval(i int) {
	p.param.restart_interval = C.int(i)
}

func (p *Param) SetInterleaved(i int) {
	p.param.interleaved = C.int(i)
}

func (p *Param) SetColorSpaceInternal(c uint32) {
	p.param.color_space_internal = c
}
