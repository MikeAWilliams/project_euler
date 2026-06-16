using System;

static List<long> PrimeFactors(long number)
{
    var result = new List<long>();
    for (long divisor = 2; number > 1; divisor++)
    {
        while (number % divisor == 0)
        {
            result.Add(divisor);
            number /= divisor;
        }
    }
    return result;
}

Console.WriteLine($"prime factors of 2 are {string.Join(", ", PrimeFactors(2))}");
Console.WriteLine($"prime factors of 3 are {string.Join(", ", PrimeFactors(3))}");
Console.WriteLine($"prime factors of 4 are {string.Join(", ", PrimeFactors(4))}");
Console.WriteLine($"prime factors of 5 are {string.Join(", ", PrimeFactors(5))}");
Console.WriteLine($"prime factors of 9 are {string.Join(", ", PrimeFactors(9))}");
Console.WriteLine($"prime factors of 13195 are {string.Join(", ", PrimeFactors(13195))}");
Console.WriteLine($"prime factors of 600851475143 are {string.Join(", ", PrimeFactors(600851475143))}");
