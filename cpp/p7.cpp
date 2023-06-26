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
    int isPrimeCounter = 0;
    int num = 0;
    int ans = 0;

    while (isPrimeCounter <= 10001)
    {
        num++;

        if (isPrime(num))
        {
            isPrimeCounter++;
        }
        
        if (isPrimeCounter == 10001)
        {
            ans = num;
            break;
        }
    }

    std::cout << ans << std::endl;
    
    return 0;
}
