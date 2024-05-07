using System;
using System.IO;
using System.Collections.Generic;

public static class Solution {
    public static void Main(string[] args) {
        if (args.Length < 1) {
            return;
        }

        string data = File.ReadAllText(args[0]);
        string[] lines = data.Split("\n");

        HashSet<char> compartment1 = new HashSet<char>();
        HashSet<char> compartment2 = new HashSet<char>();

        int sum_of_types = 0;

        foreach (string line in lines) {

            if (string.IsNullOrWhiteSpace(line)) {
                continue;
            }
            if (line.Length % 2 != 0) {
                throw new ArgumentException("odd length string");
            }

            compartment1.Clear();
            compartment2.Clear();

            int midpoint = line.Length / 2;

            for (int i = 0; i < midpoint; i++) {
                compartment1.Add(line[i]);
            }

            for (int i = midpoint; i < line.Length; i++) {
                compartment2.Add(line[i]);
            }

            compartment1.IntersectWith(compartment2);

            if (compartment1.Count == 0) {
                throw new ArgumentException("no intersection");
            }

            char intersect = '\0';

            foreach (char c in compartment1) {
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
        }

        Console.WriteLine(sum_of_types);
    }
}
