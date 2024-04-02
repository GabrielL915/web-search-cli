# web-search-cli
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/GabrielL915/web-search-cli?label=version)
## Description
  A simple CLI tool to search the web using different search engines and browsers.
  currently working on windows

## Usage

need go install

```bash
go build -o ws-cli ./cmd
```
then

```bash
./ws-cli --query "example" --engine google --browser firefox
```

## Options

- `--query`: The search query.
- `--engine`: The search engine to use (google or duckduckgo).
- `--browser`: The browser to use (chrome, firefox.).
