3: Sharing a variable
---------------------

- When I use runtime.GOMAXPROCS(2) the result is different each time I am running the program. I am guessing this is because it uses two different threads that tries to change the same variable at the same time, causing the result to be somewhat random and not 0. 

- When I changed it to runtime.GOMAXPROCS(1) the results is 0 each time I run the program. From this I am guessing that the function runtime.GOMAXPROCS(x) chooses how many threads that are used. Meaning that when i chang xe to 1 the program will first increment i by 1000000 and then decrement it with the same causing the result to be zero each time. 

4: Sharing a variable, but properly
-----------------------------------

