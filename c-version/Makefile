all: simulate analyze

CC = gcc
override CFLAGS += -g -Wno-everything -pthread -lm

# SRCS = $(shell find . -name '.ccls-cache' -type d -prune -o -type f -name '*.c' -print)

# DEBUG Targets
simulate: simulate.c
	$(CC) $(CFLAGS) -O0 simulate.c -o "$@"

analyze: simulate.c analyze.c
	$(CC) $(CFLAGS) -O0 -D_ANALYZE_CPP_ analyze.c -o "$@"

clean:
	rm -f simulate analyze