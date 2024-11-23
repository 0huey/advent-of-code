#include <iostream>
#include <fstream>
#include <string>

const std::string digits = "0123456789";

int main(int argc, char* argv[])
{
	if (argc < 2)
	{
		return 1;
	}

	std::ifstream chal_file (argv[1]);

	if (!chal_file.is_open())
	{
		return 2;
	}

	int result = 0;

	std::string line;

	while (std::getline(chal_file, line))
	{
		char num[3];
		num[0] = line[ line.find_first_of(digits, 0) ];
		num[1] = line[ line.find_last_of(digits, line.npos) ];
		num[2] = '\0';

		result += std::stoi( std::string(num) );
	}

	std::cout << result << "\n";

	return 0;
}
