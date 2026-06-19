for (int a = 1; a < 1000; a++)
{
    for (int b = a + 1; b < 1000; b++)
    {
        int c = 1000 - a - b;
        if (c < b || c < a) break;
        if (a * a + b * b == c * c)
        {
            Console.WriteLine($"a={a}, b={b}, c={c}");
            Console.WriteLine($"a*a={a * a}, b*b={b * b}, c*c={c * c}");
            Console.WriteLine($"a*b*c={a * b * c}");
        }
    }
}
