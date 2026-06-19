const int size = 2000000;
var sieve = new Shared.Sieve(size);
ulong sum = 0;
for (uint i = 2; i < size; i++)
{
    if (sieve.IsPrime((int)(i))) sum += i;
}
Console.WriteLine(sum);
