using System;

static void SayHello(string name)
{
    Console.WriteLine($"Hello, {name}!");
}

static bool IsDivisibleBy(int num, int den)
{
    return num % den == 0;
}

static int SumOfMultiplesThreeAndFive(int maxNum)
{
    int result = 0;
    for (int i = 0; i < maxNum; i++)
    {
        if (IsDivisibleBy(i, 3) || IsDivisibleBy(i, 5))
        {
            result += i;
        }
    }
    return result;
}

Console.WriteLine(SumOfMultiplesThreeAndFive(1000));
