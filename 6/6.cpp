#include <iostream>
#include <cmath>
#include <algorithm>
#include <fenv.h>

class Interval {
public:
    double L, R;

    Interval(double left, double right) : L(left), R(right) {}

    // Оператор сложения
    Interval operator+(const Interval& other) const {
        Interval result(L + other.L, R + other.R);
        return result;
    }

    // Оператор вычитания
    Interval operator-(const Interval& other) const {
        Interval result(L - other.R, R - other.L);
        return result;
    }

    // Оператор умножения
    Interval operator*(const Interval& other) const {
        double products[] = {L * other.L, L * other.R, R * other.L, R * other.R};
        std::sort(products, products + 4);
        Interval result(products[0], products[3]);
        return result;
    }

    // Оператор деления
    Interval operator/(const Interval& other) const {
        double ratios[] = {L / other.L, L / other.R, R / other.L, R / other.R};
        std::sort(ratios, ratios + 4);
        Interval result(ratios[0], ratios[3]);
        return result;
    }
};

int main() {
    fesetround(FE_DOWNWARD);
    Interval a(0.0, 0.0);
    Interval b(0.000000001, 0.000000001);

    // printf("\nТестирование операции сложения\n");
    int count = 1;
    // for (int i = 1; i <= 1000000000; ++i) {
    //     a = a + b;

    //     if (i == 200000000 * count) {
    //         printf("[%d, %.15f, %.15f, %.15f, %.e]\n", i, a.L, a.R, (a.L + a.R) / 2, (a.L + a.R) / 2 - a.L);
    //         count++;
    //     }
    // }
    
    // a.L = 1.0;
    // a.R = 1.0;
    // b.L = 0.000000001;
    // b.R = 0.000000001;
    // printf("\nТестирование операции вычитания\n");
    // count = 1;
    // for (int i = 1; i <= 1000000000; ++i) {
    //     a = a - b;

    //     if (i == 200000000 * count) {
    //         printf("[%d, %.15f, %.15f, %.15f, %.e]\n", i, a.L, a.R, (a.L + a.R) / 2, (a.L + a.R) / 2 - a.L);
    //         count++;
    //     }
    // }
    
    // a.L = 1.0;
    // a.R = 1.0;
    // b.L = 1.000000001;
    // b.R = 1.000000001;
    // printf("\nТестирование операции умножения\n");
    // count = 1;
    // for (int i = 1; i <= 1000000000; ++i) {
    //     a = a * b;

    //     if (i == 200000000 * count) {
    //         printf("[%d, %.15f, %.15f, %.15f, %.e]\n", i, a.L, a.R, (a.L + a.R) / 2, (a.L + a.R) / 2 - a.L);
    //         count++;
    //     }
    // }
    
    a.L = 1.0;
    a.R = 1.0;
    b.L = 1.000000001;
    b.R = 1.000000001;
    printf("Тестирование операции деления\n");
    count = 1;
    for (int i = 1; i <= 1000000000; ++i) {
        a = a / b;

        if (i == 200000000 * count) {
            printf("[%d, %.15f, %.15f, %.15f, %.e]\n", i, a.L, a.R, (a.L + a.R) / 2, (a.L + a.R) / 2 - a.L);
            count++;
        }
    }

    return 0;
}
