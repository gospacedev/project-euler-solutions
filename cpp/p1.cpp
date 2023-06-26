#include <iostream>

/*
This generates the even-valued Fibonacci sequence terms below four million and adds them all up
*/


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
