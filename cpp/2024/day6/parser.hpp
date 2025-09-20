#include <fstream>
#include <vector>

class point {
public:
    int x;
    int y;

    point();
    point(int x, int y);

    point operator+(const point& other);
    bool operator==(const point& other);
};

enum map_value {
    BLANK,
    OBSTRUCTION,
    START
};

class ChalMap {
public:
    std::vector<std::vector<map_value>> map;
    point start;

    ChalMap(std::ifstream& file);
    bool PointOnMap(point p);
    map_value GetValueAtPoint(point p);
    map_value SetValueAtPoint(point p, map_value v);
};
