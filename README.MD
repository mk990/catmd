# catmd

**catmd** is a command-line tool written in Go that allows you to convert text files into GitHub-flavored Markdown with code blocks. It supports various programming languages, making it easy to create Markdown documents with well-formatted code snippets.

## Features

- Convert plain text files into GitHub-flavored Markdown.
- Support for code highlighting in various programming languages, including Go, Python, Bash, PHP, Rust, HTML, and more.
- Accepts both command-line arguments and filenames from stdin for flexible usage.
- Simple and easy to use.

## Getting Started

### Prerequisites

Before using **catmd**, make sure you have Go installed on your system.

### Installation

```bash
go install -v github.com/mk990/catmd
```

### Usage

```bash
catmd file1.txt file2.txt ...
```

OR

```bash
ls -1 file* | catmd
```

The program will process the provided files, convert their content to Markdown, and print the result to the console.

Supported Languages

catmd supports code highlighting in various programming languages, including:

- Go
- Python
- Bash
- PHP
- Rust
- HTML
- ...

### License

This project is licensed under the GPL-v3 License - see the LICENSE file for details.

### Acknowledgments

- GitHub for providing the Markdown format and code highlighting.

Happy coding and enjoy creating beautifully formatted Markdown documents with catmd!
