#include <cstdio>
#include <vector>
#include <algorithm>

#include "company.h"

void* init_company() {
    company* c = new company();

    return c;
}

void generate_employee_list_vector(void* cmp, int count) {
    company* c = (company*)cmp;

    c->emp_list.data.reserve(count);

    for (int i = 0; i < count; i++) {
        employee_data emp;

        emp.id = i;

        int name_length = 5 + (rand() % 5);
        emp.name = new char[name_length + 1];
        for (int j = 0; j < name_length; j++) {
            emp.name[j] = 'a' + (rand() % 26);
        }

        emp.salary = 5000.0f + (rand() % 5000);
        c->emp_list.data.emplace_back(emp);

        printf("  %d: %s, $%.2f\n", emp.id, emp.name, emp.salary);
    }
}

void free_company_vector(void* cmp) {
    company* c = (company*)cmp;
    delete c;
}

void generate_employee_list(void* cmp, int count) {
    generate_employee_list_vector(cmp, count);
}

void free_company(void* cmp) {
    free_company_vector(cmp);
}

void print_employee_list(void* cmp) {
    company* c = (company*)cmp;
    for (auto const& emp : c->emp_list.data) {
        printf("  %d: %s, $%.2f\n", emp.id, emp.name, emp.salary);
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