#include <iostream>
#include <fstream>
#include <vector>
#include <algorithm>
#include "parser.hpp"

int main(int argc, char* argv[])
{
	if (argc < 2)
	{
		return 1;
	}

	std::ifstream chal_file(argv[1]);

	if (!chal_file.is_open())
	{
		return 2;
	}

    std::vector<cube_draws> games = parse(chal_file);

    size_t answer = 0;

    const size_t max_red = 12;
    const size_t max_green = 13;
    const size_t max_blue = 14;

    for (cube_draws game : games)
    {
        size_t red = *std::max_element(game.red.begin(), game.red.end());
        size_t blue = *std::max_element(game.blue.begin(), game.blue.end());
        size_t green = *std::max_element(game.green.begin(), game.green.end());

        if (red <= max_red && blue <= max_blue && green <= max_green)
        {
            answer += game.game_num;
        }
    }

    std::cout << answer << std::endl;

	return 0;
}
