#pragma once
#ifndef WRAPPER_H
#define WRAPPER_H

#include <libgpujpeg/gpujpeg.h>
#include <libgpujpeg/gpujpeg_common.h>
#include <libgpujpeg/gpujpeg_encoder.h>
#include <libgpujpeg/gpujpeg_decoder.h>
#include <libgpujpeg/gpujpeg_type.h>
#include <libgpujpeg/gpujpeg_version.h>

typedef struct gpujpeg_decoder gpujpeg_decoder;
typedef struct gpujpeg_parameters gpujpeg_parameters;
typedef struct gpujpeg_image_parameters gpujpeg_image_parameters;
typedef struct gpujpeg_decoder_output gpujpeg_decoder_output;

gpujpeg_decoder *create_decoder();

#endif /* WRAPPER_H */
