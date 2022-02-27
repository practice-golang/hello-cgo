#ifdef __cplusplus
extern "C" {
#endif

__declspec(dllexport) int add(int a, int b) {
	return a+b;
}

#ifdef __cplusplus
};
#endif