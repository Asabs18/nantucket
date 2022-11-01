all: nantucket simulate analyze visualize

CC = gcc
override CFLAGS += -g -Wno-everything -pthread

# DEBUG Targets
nantucket: nantucket.c log.c opt.c args.h log.h opt.h virus.h random.h
	$(CC) $(CFLAGS) -O0 nantucket.c log.c opt.c -o "$@" -lm -pthread -lrt

simulate: simulate.c
	$(CC) $(CFLAGS) -O0 simulate.c -o "$@" -lm

analyze: simulate.c analyze.c
	$(CC) $(CFLAGS) -O0 -D_ANALYZE_CPP_ analyze.c -o "$@" -lm

visualize: visualize.c
	$(CC) $(CFLAGS) -O0 visualize.c -o "$@"

clean:
	rm -f nantucket simulate analyze visualize