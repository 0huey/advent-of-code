#include <fstream>
#include <iostream>
#include <vector>
#include <string>
#include "parser.hpp"

ChalMap::ChalMap(std::ifstream& file) {
    std::string line;

    int line_len = 0;
    int y = 0;

    while (std::getline(file, line)) {
        std::vector<map_value> map_line;

        if (y == 0) {
            line_len = line.size();
        }
        else if (line.size() != line_len) {
            std::cout << "mismatched line len at line " << y << "\n";
            std::exit(1);
        }

        for (int x = 0; x < line.size(); x++) {
            switch (line[x]) {
                case '.':
                    map_line.push_back(BLANK);
                    break;
                case '#':
                    map_line.push_back(OBSTRUCTION);
                    break;
                case '^':
                    map_line.push_back(START);
                    start = point(x, y);
                    break;
                default:
                    std::cout << "unknown char at" << x << " " << y << "\n";
                    std::exit(1);
            }
        }
        y++;
        map.push_back(map_line);
    }
}

bool ChalMap::PointOnMap(point p) {
    return p.x >= 0 && p.y >= 0 && p.y < map.size() && p.x < map[0].size();
}

map_value ChalMap::GetValueAtPoint(point p) {
    if (!PointOnMap(p)) {
        return (map_value)-1;
    }
    return map[p.y][p.x];
}

map_value ChalMap::SetValueAtPoint(point p, map_value v) {
    if (!PointOnMap(p)) {
        return (map_value)-1;
    }
    map_value old = map[p.y][p.x];
    map[p.y][p.x] = v;
    return old;
}

point::point() {
    x = -1;
    y = -1;
}

point::point(int tempx, int tempy) {
    x = tempx;
    y = tempy;
}

point point::operator+(const point& other) {
    return point(x + other.x, y + other.y);
}

bool point::operator==(const point& other) {
    return x == other.x && y == other.y;
}
