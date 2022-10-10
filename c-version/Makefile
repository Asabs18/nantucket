all: simulate analyze visualize

CC = gcc
override CFLAGS += -g -Wno-everything -pthread

# DEBUG Targets
simulate: simulate.c
	$(CC) $(CFLAGS) -O0 simulate.c -o "$@"

analyze: simulate.c analyze.c
	$(CC) $(CFLAGS) -O0 -D_ANALYZE_CPP_ analyze.c -o "$@" -lm

viualize: visualize.c
	$(CC) $(CFLAGS) -O0 visualize.c -o "$@"

clean:
	rm -f simulate analyze visualize