#include <fstream>
#include <vector>

struct cube_draws {
    size_t game_num;
    std::vector<size_t> red;
    std::vector<size_t> green;
    std::vector<size_t> blue;
};

std::vector<cube_draws> parse(std::ifstream& file);
