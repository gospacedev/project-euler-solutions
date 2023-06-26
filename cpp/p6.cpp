#include <iostream>
#include <cmath>

int main(int argc, char const *argv[])
{
    int sumOfSquares = 0;
    int squareOfSums = 0;

    for (size_t i = 0; i < 100; i++)
    {
        sumOfSquares += pow(i, 2);
        squareOfSums += i;
    }

    squareOfSums = pow(squareOfSums, 2);
    
    int ans = squareOfSums - sumOfSquares;

    std::cout << ans << std::endl;
    
    return 0;
}
