CC = g++
OUT = -o $@
CC_OPTS = -Wall -O2 -std=c++20

all: solve solve-p2

solve: solve.cpp parser.o
	$(CC) $(CC_OPTS) $(OUT) $^

solve-p2: solve-p2.cpp parser.o
	$(CC) $(CC_OPTS) $(OUT) $^

parser.o: parser.cpp
	$(CC) $(CC_OPTS) $(OUT) -c $^
