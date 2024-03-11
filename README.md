# My Connect Four Solver
## Connect Four 
Golang program that runs a simple game of Connect Four on the command line where player 1 (human) plays against a connect four solver also written in Go.
## The Solver
The solver uses the MiniMax algorithm with alpha beta pruning, as well as an optimized move order in order to speed up the algorithm. 
## How to Run the program
The file that runs the CLI game is main/main.go. Simply open up the code in your terminal, navigate to the main/ directory and run (make sure you have go installed)
```
go run main.go
```
## Where to Next
I'm going to keep updating this project to make it run as fast as possible--I want to add transposition tables, use the bitboard approach, add go routines to run the MiniMax for
each of the 7 possible insert columns, among other things to keep making the solver as accurate as possible. I'm also interested in (possibly) eventually creating a simple front
end for it and having the user be able to play the AI, be able to build a position and evaluate it, have the AI play itself, etc.
