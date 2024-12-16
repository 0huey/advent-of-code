#include <fstream>
#include <iostream>
//#include <filesystem>
#include <sstream>

const std::string mul_str = "mul(";

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

    uint64_t result = 0;
    uint64_t num1, num2;
    size_t pos;

    while ((pos = data.find(mul_str)) != std::string::npos)
    {
        if (pos + mul_str.size() >= data.size())
        {
            break;
        }

        data = data.substr(pos + mul_str.size());

        if (!isdigit(data[0]))
            continue;

        num1 = std::stoull(data, &pos);

        if ( (pos + 1 > data.size()) || (data[pos] != ',') || !isdigit(data[pos+1]) )
            continue;

        data = data.substr(pos + 1);

        num2 = std::stoull(data, &pos);

        if (pos > data.size() || data[pos] != ')')
            continue;

        //std::cout << " multiplying " << num1 << " and " << num2 << std::endl;

        result += num1 * num2;
    }

    std::cout << result << std::endl;

    return 0;
}
