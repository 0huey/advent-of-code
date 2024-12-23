#include <fstream>
#include <iostream>
#include <string>
#include <vector>

#ifdef _DEBUG
const bool DEBUG = true;
#else
const bool DEBUG = false;
#endif

std::vector<std::string> parse(std::ifstream& file);
size_t count_xmas(const std::vector<std::string>& data, int64_t row, int64_t col);
bool walk_xmas(const std::vector<std::string>& data, int64_t row, int64_t col, int64_t delta_row, int64_t delta_col);
bool in_matrix_bounds(const std::vector<std::string>& data, int64_t row, int64_t col);

int main(int argc, char* argv[])
{
    if (DEBUG)
    {
        std::cout << "DEBUG ON\n";
    }

	if (argc < 2)
    {
		return 1;
    }

	std::ifstream chal_file(argv[1]);

	if (!chal_file.is_open())
    {
		return 2;
    }

    const std::vector<std::string> data = parse(chal_file);

    chal_file.close();

    size_t num_xmas = 0;

    for (int64_t row = 0; row < (int64_t)data.size(); row++)
    {
        for (int64_t col = 0; col < (int64_t)data[row].size(); col++)
        {
            num_xmas += count_xmas(data, row, col);
        }
    }

    std::cout << num_xmas << std::endl;

    return 0;
}

std::vector<std::string> parse(std::ifstream& file)
{
    std::vector<std::string> data;
    std::string line;
    size_t line_length = 0;

    while (std::getline(file, line))
    {
        if (line_length == 0)
        {
            line_length = line.size();
        }

        if (line_length == line.size())
        {
            data.push_back(line);
        }
        else if (line_length != line.size())
        {
            std::cout << "skipped parsing line " << data.size() << std::endl;
        }
    }

    return data;
}

size_t count_xmas(const std::vector<std::string>& data, int64_t row, int64_t col)
{
    if (data[row][col] != 'X')
    {
        return 0;
    }

    if (DEBUG)
    {
        std::cout << data[row][col] << " R" << row << " C" << col << std::endl;
    }

    size_t num_xmas = 0;

    for (int64_t delta_row = -1; delta_row <= 1; delta_row++)
    {
        for (int64_t delta_col = -1; delta_col <= 1; delta_col++)
        {
            if (delta_row == 0 && delta_col == 0)
            {
                continue;
            }

            if (walk_xmas(data, row, col, delta_row, delta_col))
            {
                num_xmas++;
            }
        }
    }

    return num_xmas;
}

const std::string xmas_str = "XMAS";

bool walk_xmas(const std::vector<std::string>& data, int64_t row, int64_t col, int64_t delta_row, int64_t delta_col)
{
    for (int64_t letter = 1; letter < (int64_t)xmas_str.size(); letter++)
    {
        row += delta_row;
        col += delta_col;

        if (!in_matrix_bounds(data, row, col))
        {
            return false;
        }

        if (DEBUG)
        {
            std::cout << data[row][col] << " R" << row << " C" << col << std::endl;
        }

        if (data[row][col] != xmas_str[letter])
        {
            return false;
        }
    }

    if (DEBUG)
    {
        std::cout << "XMAS TRUE\n";
    }

    return true;
}

bool in_matrix_bounds(const std::vector<std::string>& data, int64_t row, int64_t col)
{
    if (row >= 0 && row < (int64_t)data.size() && col >= 0 && col < (int64_t)data[row].size())
    {
        return true;
    }
    return false;
}
