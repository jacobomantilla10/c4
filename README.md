# My Connect Four Solver
## Connect Four 
Golang program that runs a simple game of Connect Four on the command line where player 1 (human) plays against a connect four solver also written in Go.
## The Solver
The solver uses the MiniMax algorithm with alpha beta pruning, optimized move order, a transposition, and an opening database.
## How to Run the program
There are two ways to run the app right now, CLI and Web.

For web: navigate to ./cmd/web/ and run
```
go run .
```
For the CLI: navigate to ./cmd/cli/ and run the same command.
## Where to Next
Right now I'm focused on polishing up the Web application so that I can release this project. In the future I'd also like to rethink/refactor some details in
the implementation of the engine.
