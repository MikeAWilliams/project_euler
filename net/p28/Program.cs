enum Direction
{
    Right,
    Down,
    Left,
    Up
}

class Program
{
    [System.Diagnostics.CodeAnalysis.SuppressMessage("Style", "IDE0051:Remove unused private members", Justification = "Debug helper")]
    static void PrintGrid(int[,] grid, int n)
    {
        // widest value is n*n; pad every cell to that width plus a separating space
        int width = (n * n).ToString().Length;
        for (int j = 0; j < n; j++)
        {
            for (int i = 0; i < n; i++)
            {
                Console.Write(grid[j, i].ToString().PadLeft(width));
                Console.Write(' ');
            }
            Console.Write("\n");
        }
    }

    static void Main()
    {
        const int size = 1001;
        var digits = new int[size, size];
        digits[size / 2, size / 2] = 1;
        int value = 2;
        int x = size / 2 + 1;
        int y = size / 2;
        var dir = Direction.Up;
        for (int rowSize = 3; rowSize <= size; rowSize += 2)
        {
            var steps = rowSize + 2 * (rowSize - 1) + rowSize - 2;
            for (int step = 0; step < steps; step++)
            {
                digits[y, x] = value;
                value += 1;
                switch (dir)
                {
                    case Direction.Right:
                        x++;
                        break;
                    case Direction.Down:
                        y--;
                        break;
                    case Direction.Left:
                        x--;
                        break;
                    case Direction.Up:
                        y++;
                        break;
                }
                // ring corners, measured from the grid center
                int radius = rowSize / 2;
                int low = size / 2 - radius;
                int high = size / 2 + radius;
                if (x == high && y == high)
                {
                    dir = Direction.Left;
                }
                else if (x == low && y == high)
                {
                    dir = Direction.Down;
                }
                else if (x == low && y == low)
                {
                    dir = Direction.Right;
                }
                else if (x == high && y == low)
                {
                    dir = Direction.Up;
                }
            }
            x++;
            y--;
        }

        //PrintGrid(digits, size);
        ulong sum = 0;
        // diag 1
        for (int i = 0; i < size; i++)
        {
            // first diagonal
            sum += (uint)digits[i, i];

            // second diagonal
            if (size - i - 1 != i)
            {
                sum += (uint)digits[size - i - 1, i];
            }
        }
        Console.WriteLine(sum);
    }
}
