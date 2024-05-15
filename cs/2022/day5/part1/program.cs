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

        int num_stacks = 0;
        int line_num = 0;

        for (; line_num < lines.Length; line_num++) {
            string[] temp = lines[line_num].Trim().Split();

            if (temp.Length == 0 || temp[0] != "1") {
                continue;
            }

            string last_elm = temp[temp.Length - 1];
            num_stacks = int.Parse(last_elm);
            break;
        }

        Stack<char>[] stacks = new Stack<char>[num_stacks];
        for (int i = 0; i < num_stacks; i++) {
            stacks[i] = new();
        }

        int inst_line_num = line_num + 1;
        line_num--;

        for (; line_num >= 0; line_num--) {
            for (int s = 0; s < num_stacks; s++) {

                int i = s * 4 + 1;

                if (i >= lines[line_num].Length) {
                    continue;
                }

                char c = lines[line_num][i];

                if (char.IsAsciiLetter(c)) {
                    stacks[s].Push(c);
                }
            }
        }

        for (; inst_line_num < lines.Length; inst_line_num++) {
            string[] instruction = lines[inst_line_num].Trim().Split();

            if (instruction.Length != 6) {
                continue;
            }

            int num  = int.Parse(instruction[1]);

            //adjust from 1 based index
            int src  = int.Parse(instruction[3]) - 1;
            int dest = int.Parse(instruction[5]) - 1;

            while (num-- > 0) {
                stacks[dest].Push( stacks[src].Pop() );
            }
        }

        foreach(Stack<char> stack in stacks) {
            Console.Write(stack.Peek());
        }
    }
}
