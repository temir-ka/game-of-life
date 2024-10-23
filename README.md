# game-of-life

Overview

This project implements a simulation of Conway's Game of Life, a cellular automaton devised by mathematician John Horton Conway. The game takes place on a 2D grid of cells, where each cell is either alive or dead. The grid evolves at each tick based on simple rules, simulating life, death, and reproduction.

Features

- Grid Representation:
        
        Live cells are represented by ×.
        Dead cells are represented by ·.

- Evolution Rules:

      Any live cell with fewer than two live neighbors dies (underpopulation).
      Any live cell with two or three live neighbors survives.
      Any live cell with more than three live neighbors dies (overpopulation).
      Any dead cell with exactly three live neighbors becomes a live cell (reproduction).

- Input Format: The program accepts a grid of characters where # represents a live cell and . represents an empty cell.
- Edge Cases: Handles minimum grid size (3x3) and larger grids. Terminates if the input is invalid.

- Command-Line Flags:

      --help: Displays information about the program and its capabilities.
      --verbose: Displays additional information about the grid, including the map size, number of ticks, speed, number of live cells, and more.
      --delay-ms=X: Sets the animation delay between ticks in milliseconds (default: 2500 ms).
      --file=path: Load the initial grid from a file.
      --edges-portal: Enables portal edges where cells that exit the grid reappear on the opposite side.
      --random=WxH: Generates a random grid of the specified width and height.
      --fullscreen: Adjusts the grid to fit the terminal size with empty cells.
      --footprints: Shows traces of visited cells using ∘.
      --colored: Adds color to live cells and footprints.

Development Process

This project was implemented following the rules of Conway’s Game of Life, ensuring correct handling of edge cases such as invalid inputs and conflicting flags. It supports customizable grid size, animation speed, and additional features like random grid generation and color.

Technologies Used

    Go: Used for developing the simulation logic.
    Command-Line Interface (CLI): For user interaction with flags and input handling.

How to Run

    1) Clone the repository
    2) Navigate to the project folder and run the simulation
    3) Optionally, pass flags like --verbose or --delay-ms=X for additional customization
