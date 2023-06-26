#include <iostream>

int main(int argc, char const *argv[])
{
    int n;

    for (size_t i = 0; i < 1000; i++)
    {
        if (i % 3 == 0 || i % 5 == 0)
        {
            n += i;
        }
    }

    std::cout << n << std::endl;

    return 0;
}
