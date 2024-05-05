using System;
using System.IO;

public static class Solution {
    public static void Main(string[] args) {
        if (args.Length < 1) {
            return;
        }

        string data = File.ReadAllText(args[0]);

        int highest_calories = 0;

        string[] elves = data.Split("\n\n");

        foreach (string elf in elves) {
            string[] calories = elf.Split("\n");

            int elf_calories = 0;

            foreach (string cal in calories) {
                if (string.IsNullOrWhiteSpace(cal)) {
                    continue;
                }

                elf_calories += int.Parse(cal);
            }

            if (elf_calories > highest_calories) {
                highest_calories = elf_calories;
            }
        }

        Console.WriteLine(highest_calories);
    }
}
