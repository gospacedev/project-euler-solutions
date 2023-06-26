#include <iostream>

bool isPal(int n)
{
    std::string s = std::to_string(n);

    for (size_t i = 0, j = s.size() - 1; i < j; i++, j--)
    {
        if (s[i] != s[j])
        {
            return false;
        }
    }
    return true;
}

int main()
{
    int max = 999;
    int min = 100;
    int product;
    int max_product = 0;

    for (int i = min; i <= max; i++)
    {
        for (int j = min; j <= max; j++)
        {
            product = i * j;
            if (product > max_product && isPal(product))
            {
                max_product = product;
            }
        }
    }

    std::cout << max_product;

    return 0;
}