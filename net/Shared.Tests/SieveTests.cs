namespace Shared.Tests;

public class SieveTests
{
    // Primes below 100, used to validate IsPrime across a known range.
    private static readonly HashSet<int> PrimesBelow100 = new()
    {
        2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47,
        53, 59, 61, 67, 71, 73, 79, 83, 89, 97
    };

    // --- Constructor ---

    [Theory]
    [InlineData(0)]
    [InlineData(1)]
    [InlineData(-1)]
    [InlineData(int.MinValue)]
    public void Constructor_SizeLessThanTwo_Throws(int size)
    {
        Assert.Throws<ArgumentOutOfRangeException>(() => new Sieve(size));
    }

    [Fact]
    public void Constructor_ValidSize_DoesNotThrow()
    {
        var sieve = new Sieve(100);
        Assert.NotNull(sieve);
    }

    // --- IsPrime: argument validation ---

    [Fact]
    public void IsPrime_NegativeNumber_Throws()
    {
        var sieve = new Sieve(100);
        Assert.Throws<ArgumentOutOfRangeException>(() => sieve.IsPrime(-1));
    }

    [Fact]
    public void IsPrime_NumberEqualToSize_Throws()
    {
        var sieve = new Sieve(100);
        Assert.Throws<ArgumentOutOfRangeException>(() => sieve.IsPrime(100));
    }

    [Fact]
    public void IsPrime_NumberGreaterThanSize_Throws()
    {
        var sieve = new Sieve(100);
        Assert.Throws<ArgumentOutOfRangeException>(() => sieve.IsPrime(101));
    }

    // --- IsPrime: correctness ---

    [Theory]
    [InlineData(0)]
    [InlineData(1)]
    public void IsPrime_ZeroAndOne_AreNotPrime(int num)
    {
        var sieve = new Sieve(100);
        Assert.False(sieve.IsPrime(num));
    }

    [Theory]
    [InlineData(2)]
    [InlineData(3)]
    [InlineData(5)]
    [InlineData(7)]
    [InlineData(97)]
    public void IsPrime_KnownPrimes_ReturnTrue(int num)
    {
        var sieve = new Sieve(100);
        Assert.True(sieve.IsPrime(num));
    }

    [Theory]
    [InlineData(4)]
    [InlineData(6)]
    [InlineData(9)]
    [InlineData(15)]
    [InlineData(25)]
    [InlineData(91)] // 7 * 13, a composite that survives small factor checks
    public void IsPrime_KnownComposites_ReturnFalse(int num)
    {
        var sieve = new Sieve(100);
        Assert.False(sieve.IsPrime(num));
    }

    [Fact]
    public void IsPrime_MatchesKnownPrimesAcrossWholeRange()
    {
        var sieve = new Sieve(100);
        for (int i = 0; i < 100; i++)
        {
            Assert.Equal(PrimesBelow100.Contains(i), sieve.IsPrime(i));
        }
    }

    // --- GetNthPrime: argument validation ---

    [Theory]
    [InlineData(0)]
    [InlineData(-1)]
    public void GetNthPrime_NLessThanOne_Throws(int n)
    {
        var sieve = new Sieve(100);
        Assert.Throws<ArgumentOutOfRangeException>(() => sieve.GetNthPrime(n));
    }

    [Fact]
    public void GetNthPrime_BeyondSieveRange_Throws()
    {
        // Sieve(100) contains 25 primes; asking for the 26th must throw.
        var sieve = new Sieve(100);
        Assert.Throws<ArgumentOutOfRangeException>(() => sieve.GetNthPrime(26));
    }

    // --- GetNthPrime: correctness ---

    [Theory]
    [InlineData(1, 2)]
    [InlineData(2, 3)]
    [InlineData(3, 5)]
    [InlineData(4, 7)]
    [InlineData(5, 11)]
    [InlineData(6, 13)]
    [InlineData(25, 97)]
    public void GetNthPrime_ReturnsExpectedPrime(int n, int expected)
    {
        var sieve = new Sieve(100);
        Assert.Equal(expected, sieve.GetNthPrime(n));
    }

    [Fact]
    public void GetNthPrime_SequenceMatchesIsPrime()
    {
        // The nth prime returned must itself be flagged prime, and the
        // sequence must be strictly increasing.
        var sieve = new Sieve(100);
        int previous = 0;
        for (int n = 1; n <= 25; n++)
        {
            int prime = sieve.GetNthPrime(n);
            Assert.True(sieve.IsPrime(prime));
            Assert.True(prime > previous);
            previous = prime;
        }
    }
}
