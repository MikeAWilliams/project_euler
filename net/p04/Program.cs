static bool IsPalindrome(int num)
{
    string strNum = num.ToString();
    string revNum = new string(strNum.Reverse().ToArray());
    return revNum == strNum;
}

static int BruteForceSearchForPalindromeProduct(int max)
{
    int result = 0;
    for (int i = max; i > 0; i--)
    {
        for (int j = max; j > 0; j--)
        {
            var product = i * j;
            if (product > result && IsPalindrome(product))
            {
                result = product;
            }
        }
    }
    return result;
}

Console.WriteLine(IsPalindrome(9999));
Console.WriteLine(IsPalindrome(9899));
Console.WriteLine(IsPalindrome(9889));
Console.WriteLine(BruteForceSearchForPalindromeProduct(100));
Console.WriteLine(BruteForceSearchForPalindromeProduct(1000));
