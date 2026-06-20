static ulong Triangle(ulong n)
{
    return n * (n + 1) / 2;
}

static ulong Pentagon(ulong n)
{
    return n * (3 * n - 1) / 2;
}

static ulong Hexagon(ulong n)
{
    return n * (2 * n - 1);
}


Console.WriteLine(Triangle(285));
Console.WriteLine(Pentagon(165));
Console.WriteLine(Hexagon(143));

Console.WriteLine("Starting to build search lists");
const ulong maxSearch = 100000;
var tr = new List<ulong>();
var pe = new List<ulong>();
var he = new List<ulong>();
for (ulong t = 1; t < maxSearch; t++)
{
    tr.Add(Triangle(t));
    pe.Add(Pentagon(t));
    he.Add(Hexagon(t));
}

Console.WriteLine("Starting to search");
foreach (var (t, tItem) in tr.Index())
{
    var pIndex = pe.BinarySearch(tItem);
    if (pIndex < 0)
    {
        continue;
    }
    var hIndex = he.BinarySearch(tItem);
    if (hIndex < 0)
    {
        continue;
    }
    Console.WriteLine($"t{t + 1},p{pIndex + 1},h{hIndex + 1} {tItem}");
}
Console.WriteLine("done");
