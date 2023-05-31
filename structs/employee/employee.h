#define EMPLOYEE_API __declspec(dllexport)

typedef struct {
    int id;
    char* name;
    float salary;
} Employee;

EMPLOYEE_API void get_employee(Employee* emp);
EMPLOYEE_API void set_employee(Employee* emp, int id, const char* name, float salary);
EMPLOYEE_API void free_employee(Employee* emp);
