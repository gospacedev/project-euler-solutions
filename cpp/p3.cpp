#include <cmath>
#include <iostream>

void factorize(long long n)
{
    for (long long i = 2; i <= sqrt(n); i++)
    {
        while (n % i == 0)
        {
            std::cout << i << " ";
            n /= i;
        }
    }

    if (n > 1)
    {
        std::cout << n << std::endl;
    }
}

int main(int argc, char const *argv[])
{
    long long n = 600851475143;

    factorize(n);

    return 0;
}
