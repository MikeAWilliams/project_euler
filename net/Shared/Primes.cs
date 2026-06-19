namespace Shared;

// Sieve of Eratosthenes
public class Sieve
{

    // if _flags[i] is true then i is prime
    private readonly bool[] _flags;
    private readonly int _count;
    public Sieve(int count)
    {
        _count = count;
        // defaults false
        _flags = new bool[count];
        InitFlags();
        RunSieve();
    }

    public bool IsPrime(int num)
    {
        if (num > _count)
        {
            throw new IndexOutOfRangeException();
        }
        return _flags[num];
    }

    private void InitFlags()
    {
        _flags[2] = true;
        // all odd numbers might be prime
        for (int i = 3; i < _count; i += 2)
        {
            _flags[i] = true;
        }
    }

    private void RunSieve()
    {
        int candidate = 3;
        // search upto sqrt(count)
        while (candidate * candidate < _count)
        {
            // starting with 9 (3*3) because evens already false, odds before 9 are prime
            // advance by candidate*2
            for (int notPrime = candidate * candidate; notPrime < _count; notPrime += candidate * 2)
            {
                _flags[notPrime] = false;
            }
            // find the next candidate
            for (candidate += 2; !_flags[candidate]; candidate += 2) { }
        }
    }
}
