CC = g++
OUT = -o $@
CC_OPTS = -Wall -O2 -std=c++20

ifeq ($(MAKECMDGOALS), debug)
DEBUG = -D_DEBUG -g
endif

COMPILE = $(CC) $(CC_OPTS) $(DEBUG) $(OUT)

all: solve solve-p2

solve: solve.cpp
	$(COMPILE) $^

solve-p2: solve-p2.cpp
	$(COMPILE) $^

debug: all
