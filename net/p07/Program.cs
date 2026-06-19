Shared.Sieve sieve = new Shared.Sieve(1000000);

int i = 10001;
Console.WriteLine($"{i}th prime is {sieve.GetNthPrime(i)}");
