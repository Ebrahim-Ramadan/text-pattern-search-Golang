# CLI-based Search Tool in Golang. v1.0

This is a simple command-line utility I wrote in Go that searches for a specific text pattern within files in a given directory and its subs. It displays the filenames, line numbers, and lines containing the search pattern, too (except err ofcourse)

## Table of Contents

- [Usage](#usage)
- [Getting Started](#getting-started)
- [Python Comparison](#Python-Comparison)
- [Process Time](#Process-Time)
- [Contributing](#contributing)
- [Contact](#Contact)


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

## Python Comparison
- **Performance: Go is a compiled language, while Python is an interpreted language. This difference in execution model often gives Go an advantage in terms of raw performance, especially for CPU-bound tasks like text searching.**
- **Concurrency: Go has built-in support for concurrency with goroutines and channels. This can lead to efficient parallel processing of tasks, which is beneficial in scenarios like searching through multiple files concurrently.**
- **Memory Management: Go uses a garbage collector that is optimized for low-latency applications. This can lead to more predictable memory usage and less impact on performance compared to Python's memory management.**
- **Static Typing: Go is statically typed, which allows the compiler to perform optimizations based on the known data types. Python's dynamic typing can introduce some runtime overhead.** 
- **Optimized Libraries: Go's standard library includes optimized routines for common tasks like text processing, which can contribute to better performance.**

## Process Time
some search per line outputs as it took 535.8µs, 533µs, 506.9µs, 525.9µs, 67.9µs, 999.2µs, 531.5µs
<br>
the values are pretty similar to each other but the way they vary someitmes is still confusing to me, that's why for more accurate timing measurements, I might consider running my program in a controlled environment, such as on a dedicated machine with minimal background processes. Also, I could perform multiple runs and calculate the average processing time to get a clearer picture of the actual processing time for each line. More to come, and please any contributions are very needed indeed.
### Contributing
Contributions are welcome! If you find any issues or have ideas to improve the utility, feel free to open an issue or submit a pull request. Please follow the standard GitHub practices for contributions.
### Contact
you can always reach me at ramadanebrahim791@gmail.com
<br/>
have a look at [my portfolio](https://ebrahim-ramadan.vercel.app/)










