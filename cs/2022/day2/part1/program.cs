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

        List<Strategy> strats = new();

        foreach (string line in lines) {
            if (string.IsNullOrWhiteSpace(line)) {
                continue;
            }
            strats.Add(new Strategy(line));
        }

        int result = 0;

        foreach (Strategy strat in strats) {
            result += strat.chal_score;
        }

        Console.WriteLine(result);
    }
}

public class Strategy {
    public enum RPS_choice : int {
        rock,
        paper,
        scissors
    };

    public enum RPS_result : int {
        win = 6,
        draw = 3,
        loss = 0
    };

    private RPS_choice[] RPS_choice_map = {
        RPS_choice.scissors,    //beat by rock
        RPS_choice.rock,        //paper
        RPS_choice.paper        //scissors
    };

    public RPS_choice opponent;
    public RPS_choice response;
    public RPS_result outcome;

    public int chal_score;

    public Strategy(string line) {
        string[] s = line.Split(" ");

        if (s.Length != 2) {
            throw new ArgumentException("bad line" + line);
        }

        switch (s[0]) {
            case "A":
                opponent = RPS_choice.rock;
                break;
            case "B":
                opponent = RPS_choice.paper;
                break;
            case "C":
                opponent = RPS_choice.scissors;
                break;
            default:
                throw new ArgumentException("bad line" + line);
        }

        switch (s[1]) {
            case "X":
                response = RPS_choice.rock;
                break;
            case "Y":
                response = RPS_choice.paper;
                break;
            case "Z":
                response = RPS_choice.scissors;
                break;
            default:
                throw new ArgumentException("bad line" + line);
        }

        if (opponent == response) {
            outcome = RPS_result.draw;
        }
        else if (RPS_choice_map[(int)response] == opponent) {
            outcome = RPS_result.win;
        }
        else {
            outcome = RPS_result.loss;
        }

        chal_score = (int)outcome + (int)response + 1;
    }
}
