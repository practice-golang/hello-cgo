#ifdef __cplusplus
extern "C" {
#endif

#include <stdbool.h>

#if defined(_WIN32) && !defined(__MINGW32__)
#define COMPANY_API __declspec(dllexport)
#else
#define COMPANY_API
#endif

typedef struct {
    int id;
    char* name;
    float salary;
} employee_data;

typedef struct {
#ifdef __cplusplus
    std::vector<employee_data> data;
#else
    employee_data* data;
    size_t size;
#endif
    bool sorted;
} employee_list;

struct company {
#ifdef __cplusplus
    employee_list emp_list;
#else
    void* emp_list;
#endif
};

COMPANY_API void* init_company();

COMPANY_API void generate_employee_list(void* cmp, int count);
COMPANY_API void free_company(void* cmp);

COMPANY_API void print_employee_list(void* cmp);

#ifdef __cplusplus
}
#endif
