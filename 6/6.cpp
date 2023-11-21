#include <iostream>
#include <cmath>
#include <iomanip>
#include <algorithm>
#include <fenv.h>
#include <cstring>

using namespace std;

class Interval {
public:
    double L, R;

    Interval(double left, double right) : L(left), R(right) {}

    // Перегрузка оператора сложения для интервалов
    Interval operator+(const Interval& other) const {
        Interval result(L,R);
        fesetround(FE_DOWNWARD);
        result.L = this->L + other.L;
        fesetround(FE_UPWARD);
        result.R = this->R + other.R;
        return result;
    }

    // Перегрузка оператора вычитания для интервалов
    Interval operator-(const Interval& other) const {
        Interval result(L,R);
        fesetround(FE_DOWNWARD);
        result.L = this->L - other.R;
        fesetround(FE_UPWARD);
        result.R = this->R - other.L;
        return result;
    }

    // Перегрузка оператора умножения для интервалов
    Interval operator*(const Interval& other) const {
        Interval result(L,R);
        double a = this->L;
        double b = this->R;
        double c = other.L;
        double d = other.R;
        fesetround(FE_DOWNWARD);
        result.L = min(min(a * c, a * d), min(b * c, b * d));
        fesetround(FE_UPWARD);
        result.R = max(max(a * c, a * d), max(b * c, b * d));
        return result;
    }

    // Перегрузка оператора деления для интервалов
    Interval operator/(const Interval& other) const {
        if (other.L == 0 || other.R == 0) {
            // Деление на интервал, содержащий 0, результат - интервал от минус бесконечности до плюс бесконечности
            return Interval(-INFINITY, INFINITY);
        }
        fesetround(FE_DOWNWARD);
        double min_el = min(min(this->L / other.L, this->L / other.R), min(this->R / other.L, this->R / other.R));
        fesetround(FE_UPWARD);
        double max_el = max(max(this->L / other.L, this->L / other.R), max(this->R / other.L, this->R / other.R));

        return Interval(min_el, max_el);
    }
};

void testOperation(const char* operation, Interval& a, const Interval& b) {
    printf("\n\t\tТестирование операции %s\n", operation);
    printf("    N\t\t\tleft\t\t\tright\t\t\tmed\t\t  wid\n");
    int count = 1;
    for (int i = 1; i <= 1000000000; ++i) {
        if (std::strcmp(operation, "сложения") == 0)
            a = a + b;
        else if (std::strcmp(operation, "вычитания") == 0)
            a = a - b;
        else if (std::strcmp(operation, "умножения") == 0)
            a = a * b;
        else if (std::strcmp(operation, "деления") == 0)
            a = a / b;

        if (i == 200000000 * count) {
            printf("%d\t%.15f\t%.15f\t%.15f\t%.2e\n", i, a.L, a.R, (a.L + a.R) / 2, (a.L + a.R) / 2 - a.L);
            count++;
        }
    }
    printf("\n\n");
}

int main() {
    Interval a(0.0, 0.0);
    Interval b(0.000000001, 0.000000001);

    testOperation("сложения", a, b);

    a.L = 1.0;
    a.R = 1.0;
    b.L = 0.000000001;
    b.R = 0.000000001;
    testOperation("вычитания", a, b);

    a.L = 1.0;
    a.R = 1.0;
    b.L = 1.000000001;
    b.R = 1.000000001;
    testOperation("умножения", a, b);

    a.L = 1.0;
    a.R = 1.0;
    b.L = 1.000000001;
    b.R = 1.000000001;
    testOperation("деления", a, b);

    return 0;
}
