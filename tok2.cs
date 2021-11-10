using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace _1_8
{
    class Program
    {
        /* Напишите метод, принимающий параметром массив чисел,
        и который будет возвращать минимальное значение,
        максимальное значение и среднее значение среди положительных
        элементов массива. Можете попробовать сделать это двумя
        способами: через out параметры и при помощи кортежа. сдано2710 */
        static void Main(string[] args)
        {
            Random r = new Random();
            int[] testarr = new int[10];
            for (int i = 0; i < testarr.Length; i++)
            {
                testarr[i] = r.Next(-10000, 10000);
                Console.Write(testarr[i] + "\t");
            }
            MinMaxAvg(testarr);
        }
        public static void MinMaxAvg(int[] array)
        {
            int sum = 0;
            int down_to_min = int.MaxValue;
            int up_to_max = int.MinValue;
            foreach(int num in array)
            {
                if(num > -1)
                {
                    if (num < down_to_min)
                        down_to_min = num;
                    else if (num > up_to_max)
                        up_to_max = num;
                    sum += num;
                }
            }
            Console.WriteLine("\n\nСумма:"+ sum+ "\nМин:"+ down_to_min+ "\nМакс:"+ up_to_max);
            Console.ReadKey();
        }
    }
}
