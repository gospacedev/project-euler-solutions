#include <iostream>
#include <numeric>

int main(int argc, char const *argv[])
{
    int smallestMultiple = 1;
    for (size_t i = 2; i < 20; i++)
    {
        smallestMultiple = std::lcm(smallestMultiple, i);
    }

    std::cout << smallestMultiple << std::endl;

    return 0;
}