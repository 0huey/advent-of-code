#include <vector>
#include <fstream>

struct puzzle_input
{
    std::vector<int> list1;
    std::vector<int> list2;
};

puzzle_input parse(std::ifstream& file);
