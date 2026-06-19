const int size = 5;
var digits = new int[size, size];
digits[2, 2] = 1;
int value = 2;
int x = 3;
int y = 2;
var dir = Direction.Up;
for (int rowSize = 3; rowSize <= size; rowSize += 2)
{
    var steps = rowSize + 2 * (rowSize - 1) + rowSize - 2;
    Console.WriteLine(steps);
    for (int step = 0; step < steps; step++)
    {
        Console.WriteLine($"{x},{y}={value}, dir={dir}");
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
    PrintGrid(digits, size);
    x++;
    y--;
}

PrintGrid(digits, size);

static void PrintGrid(int[,] grid, int n)
{
    for (int j = 0; j < n; j++)
    {
        for (int i = 0; i < n; i++)
        {
            Console.Write(grid[j, i]);
        }
        Console.Write("\n");
    }
}

enum Direction
{
    Right,
    Down,
    Left,
    Up
}
