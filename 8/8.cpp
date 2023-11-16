#include <iostream>
#include <cmath>
#include <algorithm>
#include <fenv.h>

class Interval {
public:
    double L, R;

    Interval(double left, double right) : L(left), R(right) {}
    
    bool operator<(const Interval& other) const {
        return R < other.L;
    }

    bool operator>(const Interval& other) const {
        return L > other.R;
    }

    bool operator==(const Interval& other) const {
        return (L == other.L) && (R == other.R);
    }

    // Оператор сложения
    Interval operator+(const Interval& other) const {
        Interval result(L + other.L, R + other.R);
        return result;
    }

       // Оператор нестандартного вычитания
    Interval operator-(const Interval& other) const {
        double tmp[] = {L - other.L, R - other.R};
        std::sort(tmp, tmp + 2);
        Interval result(tmp[0], tmp[1]);
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
        if (*this * other > Interval(0, 0)) {
            double tmp[] = {L / other.L, R / other.R};
            std::sort(tmp, tmp + 2);
            Interval result(tmp[0], tmp[1]);
            return result;
        } else if (*this > Interval(0, 0) && other > Interval(0, 0)) {
            double tmp[] = {L / other.R, R / other.L};
            std::sort(tmp, tmp + 2);
            Interval result(tmp[0], tmp[1]);
            return result;
        } else if (L <= 0 && R >= 0 && other.L > 0 && other.R > 0) {
            double tmp = 1 / other.L;
            return Interval(tmp * L, tmp * R);
        } else if (L <= 0 && R >= 0 && other.L < 0 && other.R < 0) {
            double tmp = 1 / other.R;
            return Interval(tmp * L, tmp * R);
        } else {
            throw std::runtime_error("Данные вычисления не предусмотрены библиотекой");
        }
    }
};

int main() {
    Interval a(0.0, 0.0);
    Interval b(0.000000001, 0.000000001);

    printf("\nТестирование операции сложения\n");
    int count = 1;
    for (int i = 1; i <= 1000000000; ++i) {
        a = a + b;

        if (i == 200000000 * count) {
            printf("[%d, %.15f, %.15f, %.15f, %.e]\n", i, a.L, a.R, (a.L + a.R) / 2, (a.L + a.R) / 2 - a.L);
            count++;
        }
    }
    
    a.L = 1.0;
    a.R = 1.0;
    b.L = 0.000000001;
    b.R = 0.000000001;
    printf("\nТестирование операции вычитания\n");
    count = 1;
    for (int i = 1; i <= 1000000000; ++i) {
        a = a - b;

        if (i == 200000000 * count) {
            printf("[%d, %.15f, %.15f, %.15f, %.e]\n", i, a.L, a.R, (a.L + a.R) / 2, (a.L + a.R) / 2 - a.L);
            count++;
        }
    }
    
    a.L = 1.0;
    a.R = 1.0;
    b.L = 1.000000001;
    b.R = 1.000000001;
    printf("\nТестирование операции умножения\n");
    count = 1;
    for (int i = 1; i <= 1000000000; ++i) {
        a = a * b;

        if (i == 200000000 * count) {
            printf("[%d, %.15f, %.15f, %.15f, %.e]\n", i, a.L, a.R, (a.L + a.R) / 2, (a.L + a.R) / 2 - a.L);
            count++;
        }
    }
    
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
