#include <fstream>
#include <iostream>
#include <unordered_map>
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

    std::unordered_map<int, int> list2_counts;

    for (int num : data.list2)
    {
        if (list2_counts.contains(num))
        {
            list2_counts[num]++;
        }
        else
        {
            list2_counts[num] = 1;
        }
    }

    int64_t similarity = 0;

    for (int num : data.list1)
    {
        if (list2_counts.contains(num))
        {
            similarity += num * list2_counts[num];
        }
    }

    std::cout << similarity << std::endl;

    return 0;
}
