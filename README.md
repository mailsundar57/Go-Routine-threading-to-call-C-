# Go-Routine-threading-to-call-C-
The program calls C++ executable from golang. Sometimes, we may need to call precompiled C/C++ program. The main function of C++ is designed to take in argument on run time. 

Go routines are designed to ease the process of concurrency using different threads in  your processor. Here we exploit that to call c++ file from concurrent threads created by go routine. There are two c++ files. a.cpp takes in a character argument and a1.cpp takes an integer argument. These two files are run from the go file concurrently. 



