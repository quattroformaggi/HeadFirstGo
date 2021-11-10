using System;
using System.Threading;

namespace _1_4
{
    class Program
    {
        static void Main(string[] args)
        {
            /*Разработайте программу, в случайных местах на
             * консоли с частотой в 250 мс выводящую строку «C#». сдано2710*/
            ConsoleKeyInfo info;
            Random random = new Random();
            while (true)
            {
                Console.SetCursorPosition(random.Next(1, Console.WindowHeight), random.Next(1, Console.WindowWidth));
                Console.WriteLine("C#");
                if (Console.KeyAvailable)
                {
                    info = Console.ReadKey(true);
                    if (info.Key == ConsoleKey.E)
                        break;
                }
                Thread.Sleep(250);
            }
            Console.Write("Нажмите любую клавишу, чтобы завершить...");
            Console.ReadLine();
        }
    }
}
