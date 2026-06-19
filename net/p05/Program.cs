static bool TestForRemainderInRange(long num, int max)
{
    for (int div = 2; div <= max; div++)
    {
        if (num % div > 0)
        {
            return false;
        }

    }
    return true;
}

static long FindSmallestNoRemainder(int maxDivisor, long maxCandidate)
{
    for (long candidate = maxDivisor; candidate < maxCandidate; candidate++)
    {
        if (TestForRemainderInRange(candidate, maxDivisor))
        {
            return candidate;
        }
    }
    return 0;
}

Console.WriteLine(TestForRemainderInRange(2520, 10));
Console.WriteLine(FindSmallestNoRemainder(10, 10000000000));
Console.WriteLine(FindSmallestNoRemainder(20, 10000000000));
