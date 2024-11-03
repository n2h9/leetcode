INCLUDE_PATH ?= ../../../include/c

solution: solution.o solution_test.o
	gcc solution.o solution_test.o -o solution

solution.o: solution.c
	gcc -Wall -I $(INCLUDE_PATH) -c solution.c

solution_test.o: solution_test.c
	gcc -Wall -I $(INCLUDE_PATH) -c solution_test.c

run: solution
	./solution

clean:
	rm -f *.o solution
