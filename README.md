# Search Tool in Go

This is a simple command-line utility I wrote in Go that searches for a specific text pattern within files in a given directory and its subs. It displays the filenames, line numbers, and lines containing the search pattern, too (except err ofcourse)

## Table of Contents

- [Usage](#usage)
- [Getting Started](#getting-started)
- [Contributing](#contributing)
- [License](#license)

## Usage

To use the search tool, provide the directory to search in and the text pattern to search for as command-line arguments:

```
go run main.go /path/to/search/directory "pattern"
```
Replace /path/to/search/directory with the directory you want to search in, and "pattern" with the text pattern you're looking for. please make sure of the order as
```
	directory := os.Args[1]
	pattern := os.Args[2]
```
is to xtract the command-line arguments from os.Args array. The first argument is the executable's name itself, so the actual arguments start from index 1.
<br>
Also 
bufio is a temporary storage area used to reduce the number of actual I/O operations. Instead of reading or writing data byte by byte or character by character, bufio allows you to work with larger chunks of data (known as the buffer), improving performance and efficiency.
<br>

## Getting Started
### Clone the Repository:

```
git clone https://github.com/your-username/search-tool.git
cd search-tool
```
<br>
### Run the Tool:

```
go run main.go /path/to/search/directory "-the pattern-"
```
<br>
### Adjust Buffer Size:

If you encounter the "bufio.Scanner: token too long" error like I experienced bearlier, you can adjust the buffer size in the searchInFile function in the main.go file. A larger buffer size can handle longer lines more efficiently. but also remember to find a balance between buffer size and memory usage. (optimum)
<br>

### Contributing
Contributions are welcome! If you find any issues or have ideas to improve the utility, feel free to open an issue or submit a pull request. Please follow the standard GitHub practices for contributions.
