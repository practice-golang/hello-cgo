#include <cstdio>
#include <vector>
#include "random_float.h"

float* random_float_pointer_array(int length) {
    std::vector<float> vec(length);

    for (int i = 0; i < length; i++) {
        vec[i] = (float)rand() / RAND_MAX;

        printf("C: %f\n", vec[i]);
    }

    float* result = vec.data();

    return result;
}

void random_float_by_pointer_set(float* result, int* length_p) {
    int* length = length_p;
    // *length = 10;

    std::vector<float> vec(*length);

    for (int i = 0; i < *length; i++) {
        vec[i] = (float)rand() / RAND_MAX;

        printf("C: %f\n", vec[i]);
    }

    for (int i = 0; i < *length; i++) {
        result[i] = vec[i];
    }
}
