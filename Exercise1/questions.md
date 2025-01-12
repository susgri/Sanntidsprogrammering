Exercise 1 - Theory questions
-----------------------------

### Concepts

What is the difference between *concurrency* and *parallelism*?
> Concurrency is doing multiple tasks at the same time by switching between them. The tasks do not specifically run at the same time. Parallelism however, is when the tasks physically run at the same time using for example different cores.  

What is the difference between a *race condition* and a *data race*? 
> Race condition is when the results depend on the ordering of the tasks, while data race is when more than one thread is accesing the same memory at the same time causing upredictable behavior. 
 
*Very* roughly - what does a *scheduler* do, and how does it do it?
> A scheduler decides what thread to run next. We have cooperative scheduling, where once a task start running it will do so until it gives away the control, and preemtive scheduling, where tasks can forcibly be ended at any time.  


### Engineering

Why would we use multiple threads? What kinds of problems do threads solve?
> Threads allows the system to compute multiple task at the same time and is therefore useful to perform tasks in parallell. Threads solves therefore problems where programs need to do multiple tasks simultaniously. 

Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they, and why would we rather use them over threads?
> Fibers are similar to threads, but they use cooperative scheduling. Therefore the fibers will run until they yield, meaning they will only start and stop execution in "well defined" spaces, which will not cause loss of data. Threads however often use preemtive scheduling, which means they can be interrupted anytime. 

Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
> Concurrent programs makes the programmer's life easier by making it possoble to do multiple tasks at once. However it introduces the possibilities of problems when threads wants to use the same resources as well as bugs appearing only for specific timing instances, making it harder to debug. 
 
What do you think is best - *shared variables* or *message passing*?
> It seems like message passing is the way to go as this will not cause inconsistencies and will prevent different threads to try to modify the same variable at the same time. 