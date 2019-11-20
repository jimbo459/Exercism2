#include "armstrong_numbers.h"
#include <math.h>
#include <stdio.h>

int is_armstrong_number(int num) {
    if (num == 0) {
        return 1;
    }

    int numOfDigits = log10(num)+1;
    int sum = 0;
    int tempNum = num;

    if (tempNum >= 10) {
        while (tempNum != 0) {
            sum += pow((tempNum % 10), numOfDigits);
            tempNum /= 10;
        }
    } else {
        sum = pow(tempNum, 1);
    }

    if (num == sum) {
        return 1;
    } else {
        return 0;
    }
}

