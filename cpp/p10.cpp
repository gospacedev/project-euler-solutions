#include <iostream>

bool isPrime(int num)
{
    if (num <= 1)
    {
        return false;
    }

    for (int i = 2; i * i <= num; ++i)
    {
        if (num % i == 0)
        {
            return false;
        }
    }

    return true;
}

int main(int argc, char const *argv[])
{
    long long ans = 0;
    for (size_t i = 0; i < 2000000; i++)
    {
        if (isPrime(i))
        {
            ans += i;
        }
    }

    std::cout << ans << std::endl;

    return 0;
}
