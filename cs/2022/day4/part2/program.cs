using System;
using System.IO;
using System.Collections.Generic;

public static class Solution {
    public static void Main(string[] args) {
        if (args.Length < 1) {
            return;
        }

        string data = File.ReadAllText(args[0]);
        data = data.Trim();
        string[] lines = data.Split("\n");

        int overlap = 0;

        foreach (string s in lines) {
            string[] pair = s.Trim().Split(",");

            if (pair.Length != 2) {
                throw new ArgumentException("pair doesn't contain 2 entries");
            }

            string[] elf1_str = pair[0].Split("-");
            if (elf1_str.Length != 2) {
                throw new ArgumentException("elf assignment doesn't contain 2 entries");
            }
            int elf1_min = int.Parse(elf1_str[0]);
            int elf1_max = int.Parse(elf1_str[1]);

            string[] elf2_str = pair[1].Split("-");
            if (elf2_str.Length != 2) {
                throw new ArgumentException("elf assignment doesn't contain 2 entries");
            }
            int elf2_min = int.Parse(elf2_str[0]);
            int elf2_max = int.Parse(elf2_str[1]);

            if (elf1_min >= elf2_min && elf1_min <= elf2_max) {
                overlap++;
            }
            else if (elf2_min >= elf1_min && elf2_min <= elf1_max) {
                overlap++;
            }
        }
        Console.WriteLine(overlap);
    }
}
