ulong sumOfSquares = 0;
ulong sum = 0;
for (uint num = 1; num <= 100; num++)
{
    sum += num;
    sumOfSquares += num * num;
}
Console.WriteLine(sum * sum - sumOfSquares);
