# Laravel Flaky Test Helper

A CLI tool to repeatedly run Laravel tests until they fail. Useful for debugging intermittent test failures in your application.

## Installation

### Download
Download and extract the latest [release](https://github.com/Kazuto/laravel-flaky-helper/releases).

Move the script to a directory in your `$PATH` for local usage:

```sh
mv flaky /usr/local/bin
```

### Build from Source
Clone the repository and navigate into the project directory:

```sh
git clone https://github.com/kazuto/laravel-flaky-helper
cd laravel-flaky-helper
```

Build the executable:

```sh
go build -o flaky
```

Move the script to a directory in your `$PATH` for local usage:

```sh
mv flaky /usr/local/bin
```

## Usage

Run a specific test repeatedly until failure:

```sh
flaky ExampleTest
```

### Options

- `--max <count>`: Maximum number of times to run the test before stopping (default: infinite)

Example with a limit and delay:

```sh
flaky --max 10 ExampleTest
```

## Requirements

- PHP 8+
- Laravel 9+
- PHPUnit installed in the project

## How It Works

The tool runs the specified test repeatedly using PHPUnit until it fails or reaches the maximum iteration limit. This helps in identifying flaky tests that pass inconsistently.

## Contributing

Feel free to submit issues or pull requests to improve the tool.

## License

MIT License
