#include <fstream>
#include <iostream>
//#include <filesystem>
#include <sstream>

const std::string mul_str  = "mul(";
const std::string do_str   = "do()";
const std::string dont_str = "don't()";
const std::string digits   = "0123456789";

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
/*
    std::filesystem::path path = argv[1];
    std::uintmax_t file_size = std::filesystem::file_size(path);
    char* buffer = new char[file_size+1];
    chal_file.read(buffer, file_size);
    buffer[file_size] = '\0';
    std::string data = std::string(buffer);
    delete[] buffer;
    chal_file.close();
*/

    std::stringstream buffer;
    buffer << chal_file.rdbuf();
    std::string data = buffer.str();

    bool enabled = true;
    uint64_t result = 0;
    uint64_t num1, num2;
    size_t pos = 0;

    while (pos < data.size())
    {
        if (!enabled)
        {
            pos = data.find(do_str, pos);

            if (pos == std::string::npos)
                break;

            enabled = true;
            pos += do_str.size();
        }

        size_t mul_pos = data.find(mul_str, pos);
        size_t dont_pos = data.find(dont_str, pos);

        if (mul_pos == std::string::npos)
            break;

        if (mul_pos > dont_pos)
        {
            enabled = false;
            pos = dont_pos + 1;
            continue;
        }

        pos = mul_pos + mul_str.size();

        if (pos > data.size() || !isdigit(data[pos]))
            continue;

        size_t num_end = data.find_first_not_of(digits, pos);
        std::string number_slice = data.substr(pos, num_end - pos);
        num1 = std::stoll(number_slice);
        pos = num_end;

        if (pos + 1 >= data.size() || data[pos] != ',' || !isdigit(data[pos+1]))
            continue;

        pos++;
        num_end = data.find_first_not_of(digits, pos);
        number_slice = data.substr(pos, num_end - pos);
        num2 = std::stoll(number_slice);
        pos = num_end;

        if (pos > data.size() || data[pos] != ')')
            continue;

        result += num1 * num2;
    }

    std::cout << result << std::endl;

    return 0;
}
