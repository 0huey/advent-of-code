#include <fstream>
#include <iostream>
#include <vector>
#include <algorithm>
#include "parser.hpp"

struct redirect {
    point pos;
    point dir;

    bool operator==(const redirect& other) {
        return pos == other.pos && dir == other.dir;
    }
};

void RotateDir(point& dir);

int main(int argc, char* argv[])
{
    if (argc < 2)
    {
        return 1;
    }

    std::ifstream chal_file(argv[1]);

    if (!chal_file.is_open()) {
        return 2;
    }

    ChalMap map(chal_file);
    chal_file.close();

    point guard = map.start;
    point dir(0, -1);
    point next;
    int loops = 0;
    std::vector<point> positions;

    while (map.PointOnMap(guard)) {
        if (std::find(positions.begin(), positions.end(), guard) == positions.end()) {
            positions.push_back(guard);
        }
        next = guard + dir;
        if (!map.PointOnMap(next)) {
            break;
        }

        if (map.GetValueAtPoint(next) == OBSTRUCTION) {
            RotateDir(dir);
        }
        else {
            guard = next;
        }
    }


    for (point swap : positions) {
        if (swap == map.start) {
            continue;
        }
        map_value backup = map.SetValueAtPoint(swap, OBSTRUCTION);

        std::vector<redirect> redirects;
        guard = map.start;
        dir = point(0, -1);

        while (map.PointOnMap(guard)) {
            next = guard + dir;

            if (!map.PointOnMap(next)) {
                break;
            }

            if (map.GetValueAtPoint(next) == OBSTRUCTION) {
                redirect redir = redirect{.pos = guard, .dir = dir};
                if (std::find(redirects.begin(), redirects.end(), redir) != redirects.end()) {
                    loops++;
                    break;
                }
                redirects.push_back(redir);
                RotateDir(dir);
            }
            else {
                guard = next;
            }
        }
        map.SetValueAtPoint(swap, backup);
    }

    std::cout << "part1 posisions: " << positions.size() << "; part2 loops: "<< loops << std::endl;

    return 0;
}

void RotateDir(point& dir) {
    if (dir.y == -1) {
        dir = point(1, 0);
    } else if (dir.y == 1 ) {
        dir = point(-1, 0);
    } else if (dir.x == -1) {
        dir = point(0, -1);
    } else if (dir.x == 1) {
        dir = point(0, 1);
    }
}
