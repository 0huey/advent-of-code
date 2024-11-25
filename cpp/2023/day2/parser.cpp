#include <string>
#include <cctype>
#include <cstring>
#include "parser.hpp"

const std::string digits = "0123456789";
const std::string colors = "rgb";

std::vector<cube_draws> parse(std::ifstream& file)
{
    std::vector<cube_draws> games;
    cube_draws cubes;
    std::string line;

    while (std::getline(file, line))
    {
        memset(&cubes, 0, sizeof(cubes));

        size_t i = line.find_first_of(digits);

        if (i == std::string::npos)
            break;

        line = line.substr(i);
        cubes.game_num = std::stoull(line, &i);

        while (i < line.size())
        {
            i = line.find_first_of(digits, i);

            if (i == std::string::npos)
                break;

            line = line.substr(i);
            size_t num = std::stoull(line, &i);
            i = line.find_first_of(colors, i);

            if (i == std::string::npos)
                break;

            switch (line[i])
            {
                case 'r':
                    cubes.red.push_back(num);
                    break;
                case 'g':
                    cubes.green.push_back(num);
                    break;
                case 'b':
                    cubes.blue.push_back(num);
                    break;
            }
        }

        games.push_back(cubes);
    }

    return games;
}
