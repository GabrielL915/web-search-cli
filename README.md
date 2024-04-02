# web-search-cli

![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/GabrielL915/web-search-cli?label=version)

## Description

web-search-cli is designed to do web searches in the terminal, optimized for use on Windows systems

## Prerequisites

Before you begin, ensure you have Go installed on your system

## Installation

To install the web-search-cli tool, follow these steps:

1. Clone the repository and navigate to the project directory
2. Build the project using the Go compiler:

```bash
go build -o ws-cli ./cmd
```

## Usage

After installation, here's how to use the tool:

```bash
./ws-cli --query "example search" --engine google --browser firefox
```

## Getting Help

For a list of available commands and options, use the `--help` flag:

```bash
./ws-cli --help
```

## Options

Here's a brief overview of the options you can use with web-search-cli:

- `--query`: Specifies the search query
- `--engine`: Selects the search engine (options: `google`, `duckduckgo`)
- `--browser`: Chooses the web browser for displaying results (options: `chrome`, `firefox`)

