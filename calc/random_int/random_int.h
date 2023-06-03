#ifdef __cplusplus
extern "C" {
#endif

#define CALC_API __declspec(dllexport)

CALC_API int* random_int_pointer_array(int length);
CALC_API void random_int_by_pointer_set(int* result, int* length);

#ifdef __cplusplus
}
#endif
