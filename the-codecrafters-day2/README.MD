# Base Converter CLI

A simple Go command-line tool to convert numbers between **Hexadecimal**, **Binary**, and **Decimal**. This project was built using only the Go standard library (`fmt` and `strconv`).

## Features
- **Hex to Decimal**: Converts hex strings (like `1E`) to base-10.
- **Binary to Decimal**: Converts binary strings (like `1010`) to base-10.
- **Decimal to All**: Converts decimal numbers to both Binary and Hex.
- **Validation**: Automatically detects if a number is invalid for its base (e.g., `1G` is not hex).
- **Negative Support**: Supports negative decimal inputs.

## How to Run

1.  **Open your terminal** and navigate to the project folder:
    ```bash
    cd thecodecrafterthon-day2/
    ```

2.  **Run the program** directly using Go:
    ```bash
    go run main.go
    ```

3.  **Enter your commands** using the format `<number> <base>`:
    ```text
    > 1E hex
    > 1010 bin
    > 255 dec
    ```

4.  **To Exit**, simply type:
    ```text
    > quit
    ```

## Project Requirements Met
- Supports `hex`, `bin`, and `dec` inputs.
- Uppercase Hex output (using `%X`).
- Handles "garbage" input without crashing.
- Supports negative numbers for decimal.
- Zero external dependencies.