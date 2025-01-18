#include <iostream>

int collatzSequenceLength(int num)
{
    int length = 1;
    while (num != 1)
    {
        if (num % 2 == 0)
        {
            num = num / 2;
        }
        else
        {
            num = 3 * num + 1;
        }
        length++;
    }

    return length;
}

int main(int argc, char const *argv[])
{
    int currentLength, longestLength = 0, startNum = 0;
    for (size_t i = 0; i < 100000; i++)
    {
        currentLength = collatzSequenceLength(i);
        if (currentLength >= longestLength)
        {
            currentLength = longestLength;
            startNum = i;
        }
    }

    std::cout << startNum << std::endl;

    return 0;
}
