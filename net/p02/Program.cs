using System;

static bool IsEven(int n)
{
    return n % 2 == 0;
}

static int SumEvenFib(int maxFib)
{
    int result = 0;
    int last = 1;
    int current = 1;
    while (current < maxFib)
    {
        if (IsEven(current))
        {
            result += current;
        }
        int tmp = last;
        last = current;
        current = tmp + last;
    }
    return result;
}

Console.WriteLine(SumEvenFib(4000000));
