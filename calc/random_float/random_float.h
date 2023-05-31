#ifdef __cplusplus
extern "C" {
#endif

#define CALC_API __declspec(dllexport)

CALC_API float* random_float_pointer_array(int length);
CALC_API void random_float_by_pointer_set(float* result, int* length);

#ifdef __cplusplus
}
#endif
