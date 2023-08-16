#include <iostream>
#include <cmath>
#include <vector>

int triangularNumber(int n)
{
    return n * (n + 1) / 2;
}

int countDivisors(int n)
{
    int count = 0;
    int squareRoot = sqrt(n);

    for (size_t i = 1; i <= squareRoot; i++)
    {
        if (n % i == 0)
        {
            count += 2;
        }
        if (squareRoot * squareRoot == n)
        {
            count--;
        }
    }

    return count;
}

int main(int argc, char const *argv[])
{
    for (size_t i = 1;; i++)
    {
        if (countDivisors(triangularNumber(i)) >= 500)
        {
            std::cout << triangularNumber(i) << std::endl;
            break;
        }
    }

    return 0;
}
