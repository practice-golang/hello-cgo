#ifdef __cplusplus
extern "C" {
#endif

#include <stdbool.h>

#ifdef _WIN32
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
    employee_data* data;
    size_t size;
    bool sorted;
} employee_list;

struct company {
    void* emp_list;
};

COMPANY_API void* init_company();

COMPANY_API void generate_employee_list(void* cmp, int count);
COMPANY_API void free_company(void* cmp);

COMPANY_API void print_employee_list(void* cmp);

#ifdef __cplusplus
}
#endif