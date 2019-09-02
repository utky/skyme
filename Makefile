.PHONY: test build

.SUFFIXES:
.SUFFIXES: .c .o
CC=gcc
CFLAGS=-I. -Wall
OBJS=main.o
APP=skyme

test: build
	./run_test.sh ./$(APP)

build: $(APP)

clean:
	rm $(OBJS)
	rm $(APP)

.c.o:
	$(CC) -c $(CFLAGS) $<

$(APP): $(OBJS)
	$(CC) -o $(APP) $^
