#include <fstream>
#include <algorithm>
#include <iostream>
#include "parser.hpp"

bool safety_check(std::vector<int64_t> report, bool problem=false, size_t problem_i=0)
{
    if (problem && problem_i > report.size() - 1)
    {
        return false;
    }
    else if (problem)
    {
        report.erase(report.begin() + problem_i);
    }

    int64_t dir = 1;

    for (size_t i = 0; i < report.size() - 1; i++)
    {
        int64_t diff = report[i] - report[i+1];

        if (i == 0 && diff < 0)
        {
            dir = -1;
        }

        if ( (dir == 1 && diff < 0) || (dir == -1 && diff > 0) || (abs(diff) > 3) || (diff == 0) )
        {
            if (problem)
            {
                return false;
            }
            return (safety_check(report, true, i) || safety_check(report, true, i-1) || safety_check(report, true, i+1));
        }
    }

    return true;
}

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

    std::vector<std::vector<int64_t>> data = parse(chal_file);

    chal_file.close();

    size_t safe_reports = 0;

    for (std::vector<int64_t>& report : data)
    {
        if (safety_check(report))
        {
            safe_reports++;
        }
    }

    std::cout << safe_reports << std::endl;

    return 0;
}
