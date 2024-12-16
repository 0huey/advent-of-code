#include "parser.hpp"
#include <string>
#include <cctype>

std::vector<std::vector<int64_t>> parse(std::ifstream& file)
{
    std::vector<std::vector<int64_t>> data;
    std::string line;

    while (std::getline(file, line))
    {
        //puzzle calls each line in input a "report"
        std::vector<int64_t> report;
        size_t pos;

        if (line.size() == 0 || !isdigit(line[0]))
        {
            break;
        }

        while (line.size() > 0)
        {
            report.push_back(std::stoll(line, &pos));

            while (pos < line.size() && !isdigit(line[pos]))
            {
                pos++;
            }

            line = line.substr(pos);
        }
        data.push_back(report);
    }

    return data;
}
