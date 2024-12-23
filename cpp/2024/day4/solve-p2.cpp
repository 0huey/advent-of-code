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
bool is_x_mas(const std::vector<std::string>& data, int row, int col);
bool in_matrix_bounds(const std::vector<std::string>& data, int row, int col);

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

    for (int row = 0; row < (int)data.size(); row++)
    {
        for (int col = 0; col < (int)data[row].size(); col++)
        {
            num_xmas += is_x_mas(data, row, col);
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

const std::string MS_str = "MS";

bool is_x_mas(const std::vector<std::string>& data, int row, int col)
{
/*
 M   M
   A
 S   S
*/
    if (data[row][col] != 'A')
    {
        return false;
    }

    if (DEBUG)
    {
        std::cout << data[row][col] << " R" << row << " C" << col << std::endl;
    }

    int diag_r_1 = row - 1;
    int diag_c_1 = col - 1;
    int diag_r_2 = row + 1;
    int diag_c_2 = col + 1;

    if (!in_matrix_bounds(data, diag_r_1, diag_c_1) || !in_matrix_bounds(data, diag_r_2, diag_c_2))
    {
        return false;
    }
    if (MS_str.find(data[diag_r_1][diag_c_1]) == std::string::npos || MS_str.find(data[diag_r_2][diag_c_2]) == std::string::npos)
    {
        return false;
    }
    if (data[diag_r_1][diag_c_1] == data[diag_r_2][diag_c_2])
    {
        return false;
    }

    diag_r_1 = row + 1;
    diag_c_1 = col - 1;
    diag_r_2 = row - 1;
    diag_c_2 = col + 1;

    if (!in_matrix_bounds(data, diag_r_1, diag_c_1) || !in_matrix_bounds(data, diag_r_2, diag_c_2))
    {
        return false;
    }
    if (MS_str.find(data[diag_r_1][diag_c_1]) == std::string::npos || MS_str.find(data[diag_r_2][diag_c_2]) == std::string::npos)
    {
        return false;
    }
    if (data[diag_r_1][diag_c_1] == data[diag_r_2][diag_c_2])
    {
        return false;
    }

    return true;
}

bool in_matrix_bounds(const std::vector<std::string>& data, int row, int col)
{
    if (row >= 0 && row < (int)data.size() && col >= 0 && col < (int)data[row].size())
    {
        return true;
    }
    return false;
}
