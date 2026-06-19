namespace Shared;

// Sieve of Eratosthenes
public class Sieve
{

    // if _flags[i] is true then i is prime
    private readonly bool[] _flags;
    private readonly int _size;
    public Sieve(int size)
    {
        _size = size;
        // defaults false
        _flags = new bool[size];
        InitFlags();
        RunSieve();
    }

    public bool IsPrime(int num)
    {
        if (num < 0 || num >= _size)
        {
            throw new IndexOutOfRangeException();
        }
        return _flags[num];
    }

    private void InitFlags()
    {
        _flags[2] = true;
        // all odd numbers might be prime
        for (int i = 3; i < _size; i += 2)
        {
            _flags[i] = true;
        }
    }

    private void RunSieve()
    {
        int candidate = 3;
        // search upto sqrt(size)
        while (candidate * candidate < _size)
        {
            // starting with 9 (3*3) because evens already false, odds before 9 are prime
            // advance by candidate*2 to skip even multiples of candidate (already false)
            int advance = candidate * 2;
            for (int notPrime = candidate * candidate; notPrime < _size; notPrime += advance)
            {
                _flags[notPrime] = false;
            }
            // find the next candidate
            // next item not already crossed out
            for (candidate += 2; candidate < _size && !_flags[candidate]; candidate += 2) { }
        }
    }
}
