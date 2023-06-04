#include <stdbool.h>

#if defined(_WIN32) && !defined(__MINGW32__)
#define COMPANY_API __declspec(dllexport)
#else
#define COMPANY_API
#endif

#ifdef __cplusplus
extern "C" {

typedef struct {
    int id;
    char* name;
    float salary;
} employee_data;

typedef struct {
    std::vector<employee_data> data;
    bool sorted;
} employee_list;

struct company {
    employee_list emp_list;
};

COMPANY_API void* init_company();

COMPANY_API void generate_employee_list(void* cmp, int count);
COMPANY_API void free_company(void* cmp);

COMPANY_API void print_employee_list(void* cmp);
}
#endif
