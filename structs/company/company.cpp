#include <cstdio>
#include <vector>
#include <algorithm>

#include "company.h"

void* init_company() {
    company* c = new company();
    c->emp_list = nullptr;

    return c;
}

void generate_employee_list_vector(void* cmp, int count) {
    company* c = (company*)cmp;

    std::vector<employee_data> employees = std::vector<employee_data>();
    employees.reserve(count);

    for (int i = 0; i < count; i++) {
        employee_data emp;

        emp.id = i;

        int name_length = 5 + (rand() % 5);
        emp.name = new char[name_length + 1];
        for (int j = 0; j < name_length; j++) {
            emp.name[j] = 'a' + (rand() % 26);
        }

        emp.salary = 5000.0f + (rand() % 5000);
        employees.emplace_back(emp);

        printf("  %d: %s, $%.2f\n", emp.id, emp.name, emp.salary);
    }

    c->emp_list = new employee_list{employees.data(), employees.size(), false};
}

void free_company_vector(void* cmp) {
    company* c = (company*)cmp;
    delete (employee_list*)c->emp_list;
    delete c;
}

void generate_employee_list_cmalloc(void* cmp, int count) {
    company* c = (company*)cmp;

    employee_list* emp_list = (employee_list*)malloc(sizeof(employee_list));
    emp_list->data = (employee_data*)malloc(count * sizeof(employee_data));
    emp_list->size = count;
    emp_list->sorted = false;

    for (int i = 0; i < count; i++) {
        employee_data* emp = &(emp_list->data[i]);

        emp->id = i;

        int name_length = 5 + (rand() % 5);
        emp->name = (char*)malloc((name_length + 1) * sizeof(char));
        for (int j = 0; j < name_length; j++) {
            emp->name[j] = 'a' + (rand() % 26);
        }
        emp->name[name_length] = '\0';

        emp->salary = 5000.0f + (rand() % 5000);

        printf("  %d: %s, $%.2f\n", emp->id, emp->name, emp->salary);
    }

    c->emp_list = emp_list;
}

void free_company_cmalloc(void* cmp) {
    company* c = (company*)cmp;
    employee_list* emp = (employee_list*)c->emp_list;

    for (size_t i = 0; i < emp->size; i++) {
        employee_data* data = &(emp->data[i]);
        free(data->name);
    }

    free(emp->data);
    free(emp);
    free(c);
}

void generate_employee_list(void* cmp, int count) {
    // generate_employee_list_cmalloc(cmp, count);
    generate_employee_list_vector(cmp, count);
}

void free_company(void* cmp) {
    // free_company_cmalloc(cmp);
    free_company_vector(cmp);
}

void print_employee_list(void* cmp) {
    company* c = (company*)cmp;
    employee_list* emp = (employee_list*)c->emp_list;

    printf("#### Employee list:\n");
    for (size_t i = 0; i < emp->size; i++) {
        printf("  %d: %s, $%.2f\n", emp->data[i].id, emp->data[i].name, emp->data[i].salary);
    }
}

#ifdef NOT_LIBRARY
int main() {
    void* cmp = init_company();

    generate_employee_list(cmp, 5);
    print_employee_list(cmp);

    free_company(cmp);
}
#endif