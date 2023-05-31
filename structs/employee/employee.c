// #include <stdio.h>
// #include <stdlib.h>
#include <string.h>

#include "employee.h"

void get_employee(Employee* emp) {
    emp->id = 1;
    emp->name = strdup("John Doe");
    emp->salary = 5000.0f;
}

void set_employee(Employee* emp, int id, const char* name, float salary) {
    emp->id = id;
    emp->name = strdup(name);
    emp->salary = salary;
}

void free_employee(Employee* emp) {
    free(emp->name);
}
