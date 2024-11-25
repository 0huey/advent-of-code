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

    for (cube_draws game : games)
    {
        size_t red = *std::max_element(game.red.begin(), game.red.end());
        size_t blue = *std::max_element(game.blue.begin(), game.blue.end());
        size_t green = *std::max_element(game.green.begin(), game.green.end());

        answer += red * green * blue;
    }

    std::cout << answer << std::endl;

	return 0;
}
