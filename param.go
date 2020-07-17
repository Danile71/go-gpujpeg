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
	param C.gogpujpeg_parameters
}

func SetParam() *Param {
	p := &Param{param: C.malloc_gpujpeg_parameters()}
	C.gpujpeg_set_default_parameters(p.param)
	return p
}

func (p *Param) Free() {
	C.free_gpujpeg_parameters(p.param)
	p.param = nil
	p = nil
}

// get params
func (p *Param) RestartInterval() int {
	return int(p.param.restart_interval)
}

func (p *Param) Interleaved() int {
	return int(p.param.interleaved)
}

func (p *Param) ColorSpaceInternal() int {
	return int(p.param.color_space_internal)
}

func (p *Param) Quality() int {
	return int(p.param.quality)
}

func (p *Param) SamplingFactor(i int) (byte, byte) {
	return byte(p.param.sampling_factor[i].vertical), byte(p.param.sampling_factor[i].horizontal)
}

//set params
func (p *Param) SetRestartInterval(i int) {
	p.param.restart_interval = C.int(i)
}

func (p *Param) SetInterleaved(i int) {
	p.param.interleaved = C.int(i)
}

func (p *Param) SetColorSpaceInternal(c uint32) {
	p.param.color_space_internal = c
}

func (p *Param) SetQuality(q int) {
	p.param.quality = C.int(q)
}

func (p *Param) SetSamplingFactor(i int, v, h byte) {
	p.param.sampling_factor[i].vertical = C.uchar(v)
	p.param.sampling_factor[i].horizontal = C.uchar(h)
}
