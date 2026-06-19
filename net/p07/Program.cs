Shared.Sieve sieve = new Shared.Sieve(100);

for (int i = 0; i < 30; i++)
{
    Console.WriteLine($"{i} is prime {sieve.IsPrime(i)}");
}
