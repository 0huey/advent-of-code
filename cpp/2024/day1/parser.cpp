#include "parser.hpp"
#include <string>
#include <cctype>

puzzle_input parse(std::ifstream& file)
{
    puzzle_input data;
    std::string line;

    while (std::getline(file, line))
    {
        int num1, num2;
        size_t pos;

        if (line.size() == 0 || !isdigit(line[0]))
        {
            break;
        }

        num1 = std::stoi(line, &pos);

        while (pos < line.size() && !isdigit(line[pos]))
        {
            pos++;
        }

        line = line.substr(pos);

        num2 = std::stoi(line);

        data.list1.push_back(num1);
        data.list2.push_back(num2);
    }

    return data;
}
