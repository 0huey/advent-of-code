#include <fstream>
#include <algorithm>
#include <iostream>
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

    puzzle_input data = parse(chal_file);

    if (data.list1.size() != data.list2.size())
    {
        std::cout << "mismatched list sizes\n";
        return 3;
    }

    std::sort(data.list1.begin(), data.list1.end());
    std::sort(data.list2.begin(), data.list2.end());

    int sum_diff = 0;

    for (int i = 0; i < data.list1.size(); i++)
    {
        sum_diff += abs(data.list1[i] - data.list2[i]);
    }

    std::cout << sum_diff << std::endl;

    return 0;
}
