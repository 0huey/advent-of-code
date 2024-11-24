#include <iostream>
#include <fstream>
#include <string>
#include <array>

const std::string digits = "123456789";

const std::array<std::string, 9> words = {
	"one",
    "two",
    "three",
    "four",
    "five",
    "six",
    "seven",
    "eight",
    "nine"
};

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

	int result = 0;

	std::string line;

	while (std::getline(chal_file, line))
	{
		char line_val[3];

		size_t digit_pos = line.find_first_of(digits, 0);
		size_t word_pos  = std::string::npos;
		size_t word_index = 0;

        for (size_t i = 0; i < words.size(); i++)
        {
			size_t pos = line.find(words[i]);
			if (pos < word_pos)
			{
				word_pos = pos;
				word_index = i;
			}
        }

		if (digit_pos < word_pos)
		{
			line_val[0] = line[digit_pos];
		}
		else
		{
			line_val[0] = word_index + '1';
		}


		digit_pos = line.find_last_of(digits, line.npos);
		word_pos = 0;

		for (size_t i = 0; i < words.size(); i++)
        {
			size_t pos = line.rfind(words[i]);
			if (pos != std::string::npos && pos > word_pos)
			{
				word_pos = pos;
				word_index = i;
			}
        }

		if (digit_pos == std::string::npos || digit_pos < word_pos)
		{
			line_val[1] = word_index + '1';
		}
		else
		{
			line_val[1] = line[digit_pos];
		}

		line_val[2] = '\0';
		result += std::stoi(std::string(line_val));
	}

	std::cout << result << "\n";

	return 0;
}
