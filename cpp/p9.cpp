#include <iostream>

int findPythagoreanTriplet(int sum)
{
    for (size_t a = 1; a < sum / 3; a++)
    {
        for (size_t b = 0; b < sum / 2; b++)
        {
            int c = sum - a - b;
            if (a * a + b * b == c * c)
            {
                return a * b * c;
            }
        }
    }
}

int main(int argc, char const *argv[])
{
    int ans = findPythagoreanTriplet(1000);
    std::cout << ans << std::endl;
    return 0;
}
