const int size = 2000000;
var sieve = new Shared.Sieve(size);
ulong sum = 0;
for (int i = 2; i < size; i++)
{
    if (sieve.IsPrime(i)) sum += (uint)i;
}
Console.WriteLine(sum);
