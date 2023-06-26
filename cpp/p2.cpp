#include <iostream>

int main(int argc, char const *argv[])
{
    int x = 1, y = 2, n = 2;

    while (y < 4000000)
    {
        x += y;
        if (x % 2 == 0)
        {
            n += x;
        }

        y += x;
        if (y % 2 == 0)
        {
            n += y;
        }
    }

    std::cout << n << std::endl;

    return 0;
}