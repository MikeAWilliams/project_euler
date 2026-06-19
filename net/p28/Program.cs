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
        // rowSize - 2 reflects that 2 is the middle of this grid needs adjustment
        if (x == rowSize && y == rowSize)
        {
            dir = Direction.Left;
        }
        else if (x == rowSize - 2 && y == rowSize)
        {
            dir = Direction.Down;
        }
        else if (x == rowSize - 2 && y == rowSize - 2)
        {
            dir = Direction.Right;
        }
        else if (x == rowSize && y == rowSize - 2)
        {
            dir = Direction.Up;
        }
    }
    PrintGrid(digits, size);
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
