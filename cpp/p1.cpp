#include <iostream>

// This adds all the multiples of 3 or 5 below 1000
int main(int argc, char const *argv[])
{
    int ans;

    for (size_t i = 0; i < 1000; i++)
    {
        if (i % 3 == 0 || i % 5 == 0)
        {
            ans += i;
        }
    }

    std::cout << ans << std::endl;

    return 0;
}
