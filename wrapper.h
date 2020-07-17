#pragma once
#ifndef WRAPPER_H
#define WRAPPER_H

#include <stdio.h>
#include <stdlib.h>
#include <libgpujpeg/gpujpeg.h>
#include <libgpujpeg/gpujpeg_common.h>
#include <libgpujpeg/gpujpeg_encoder.h>
#include <libgpujpeg/gpujpeg_decoder.h>
#include <libgpujpeg/gpujpeg_type.h>
#include <libgpujpeg/gpujpeg_version.h>

typedef struct gpujpeg_decoder *gogpujpeg_decoder;
typedef struct gpujpeg_encoder *gogpujpeg_encoder;

typedef struct gpujpeg_parameters *gogpujpeg_parameters;
typedef struct gpujpeg_image_parameters *gogpujpeg_image_parameters;

typedef struct gpujpeg_decoder_output gpujpeg_decoder_output;
typedef struct gpujpeg_encoder_input gpujpeg_encoder_input;

gogpujpeg_decoder create_decoder();
gogpujpeg_encoder create_encoder();

gogpujpeg_parameters malloc_gpujpeg_parameters();
void free_gpujpeg_parameters(gogpujpeg_parameters p);

gogpujpeg_image_parameters malloc_gpujpeg_image_parameters();
void free_gpujpeg_image_parameters(gogpujpeg_image_parameters p);

uint8_t * gpuencode(struct gpujpeg_encoder *encoder,gogpujpeg_parameters param, gogpujpeg_image_parameters param_image,uint8_t* data ,int *image_compressed_size);
#endif /* WRAPPER_H */
