# Poker Hand Evaluator

This project implements a 5-card poker hand evaluator in Go (Golang). It serves as a personal exercise to explore different approaches to solving this common programming challenge.

## Features

- Evaluates 5-card poker hands
- Implements two distinct approaches:
  1. A naive, readable implementation
  2. A high-performance bit manipulation

## Project Structure

- `cards`: Package containing the naive implementation
- `main.go`: Contains the bit manipulation implementation

## Implementations

### 1. Naive Approach (`cards` package)

This implementation prioritizes readability and ease of understanding. It's ideal for learning purposes and for those new to poker hand evaluation algorithms.

### 2. Bit Manipulation Approach (`main.go`)

This approach uses bitwise operations to represent and evaluate poker hands. It offers significantly improved performance compared to the naive approach, making it suitable for high-volume evaluations.

## Usage

To use this evaluator, clone the repository and run:

```bash
go run main.go