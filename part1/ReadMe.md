# Calibration Value Calculator

This program reads a file containing lines of text and calculates a "calibration value" for each line. The calibration value is formed by combining the first and last digits on each line.

## Features

- Extracts the first and last digits from each line.
- Combines them into a two-digit number.
- Sums the calibration values for all lines.
- Handles cases where there is only one digit by using it twice.

## Usage

1. Place your input data in a file named `doc.txt`.
2. Run the program to calculate the sum of calibration values:
   ```bash
   go run main.go
