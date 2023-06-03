#include <cstdio>
#include <vector>
#include "random_int.h"

int* random_int_pointer_array(int length) {
    std::vector<int> vec(length);

    for (int i = 0; i < length; i++) {
        vec[i] = (int)rand();

        printf("C: %d\n", vec[i]);
    }

    int* result = vec.data();

    return result;
}

void random_int_by_pointer_set(int* result, int* length_p) {
    int* length = length_p;
    // *length = 10;

    std::vector<int> vec(*length);

    for (int i = 0; i < *length; i++) {
        vec[i] = (int)rand();

        printf("C: %d\n", vec[i]);
    }

    for (int i = 0; i < *length; i++) {
        result[i] = vec[i];
    }
}
