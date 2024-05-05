using System;
using System.IO;
using System.Collections.Generic;

public static class Solution {
    public static void Main(string[] args) {
        if (args.Length < 1) {
            return;
        }

        string data = File.ReadAllText(args[0]);

        var calories = new List<int>();

        string[] elves = data.Split("\n\n");

        foreach (string elf in elves) {
            string[] str_calories = elf.Split("\n");

            int elf_calories = 0;

            foreach (string cal in str_calories) {
                if (string.IsNullOrWhiteSpace(cal)) {
                    continue;
                }

                elf_calories += int.Parse(cal);
            }

            calories.Add(elf_calories);
        }

        calories.Sort();
        calories.Reverse();

        int top_three = 0;
        int count = 0;

        foreach (int cal in calories) {
            top_three += cal;
            count++;
            if (count >= 3) {
                break;
            }
        }
        Console.WriteLine(top_three);
    }
}
