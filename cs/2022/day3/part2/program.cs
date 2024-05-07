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

        HashSet<char>[] elf_group = new HashSet<char>[3];

        for (int i = 0; i < elf_group.Length; i++) {
            elf_group[i] = new HashSet<char>();
        }

        int sum_of_types = 0;

        if (lines.Length % 3 != 0) {
            throw new ArgumentException("num lines not divisible by 3");
        }

        for (int line = 0; line < lines.Length; line += 3) {
            for (int elf = 0; elf < elf_group.Length; elf++) {
                foreach (char c in lines[line + elf]) {
                    elf_group[elf].Add(c);
                }
            }

            elf_group[0].IntersectWith(elf_group[1]);
            elf_group[0].IntersectWith(elf_group[2]);

            if (elf_group[0].Count == 0) {
                throw new ArgumentException("no intersection");
            }

            char intersect = '\0';

            foreach (char c in elf_group[0]) {
                intersect = c;
                break;
            }

            if (char.IsLower(intersect)) {
                sum_of_types += intersect - 'a' + 1;
            }
            else if (char.IsUpper(intersect)) {
                sum_of_types += intersect - 'A' + 27;
            }
            else {
                throw new ArgumentException("intersection is not a letter");
            }

            for (int i = 0; i < elf_group.Length; i++) {
                elf_group[i].Clear();
            }
        }

        Console.WriteLine(sum_of_types);
    }
}
