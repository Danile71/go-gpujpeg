#include <libgpujpeg/gpujpeg.h>
#include <libgpujpeg/gpujpeg_common.h>
#include <libgpujpeg/gpujpeg_encoder.h>
#include <libgpujpeg/gpujpeg_decoder.h>
#include <libgpujpeg/gpujpeg_type.h>
#include <libgpujpeg/gpujpeg_version.h>
#include "wrapper.h"

gogpujpeg_decoder create_decoder() {
    return gpujpeg_decoder_create(0);
}

gogpujpeg_encoder create_encoder() {
    return gpujpeg_encoder_create(0);
}

uint8_t * encode(gogpujpeg_encoder encoder,gogpujpeg_parameters param, gogpujpeg_image_parameters param_image,uint8_t* data ,int *image_compressed_size) {
    uint8_t* image_compressed = NULL;
	gpujpeg_encoder_input encoder_input;
	gpujpeg_encoder_input_set_image(&encoder_input, data);
    	if (gpujpeg_encoder_encode(encoder, param, param_image, &encoder_input, &image_compressed,
		image_compressed_size) != 0) {
    }
    return image_compressed;
}

gogpujpeg_parameters malloc_gpujpeg_parameters() {
    return malloc(sizeof(struct gpujpeg_parameters));
}

void free_gpujpeg_parameters(gogpujpeg_parameters params) {
    free(params);
}

gogpujpeg_image_parameters malloc_gpujpeg_image_parameters() {
    return malloc(sizeof(struct gpujpeg_image_parameters));
}

void free_gpujpeg_image_parameters(gogpujpeg_image_parameters params) {
    free(params);
}