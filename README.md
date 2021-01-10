# gi

[![Actions Status: release](https://github.com/uu64/gi/workflows/release/badge.svg)](https://github.com/uu64/gi/actions?query=workflow%3A"release")
[![Actions Status: test](https://github.com/uu64/gi/workflows/test/badge.svg)](https://github.com/uu64/gi/actions?query=workflow%3A"test")

A simple interactive CLI tool to create a gitignore.


## Features

- Copy the gitignore template from the remote repository.
    - By default, you can copy from [https://github.com/github/gitignore](https://github.com/github/gitignore).
    - By changing the [configuration](#Configuration), you can also copy from other repository.

- Up to five gitignore templates can be selected.
    - If multiple templates are selected, they will be merged and output as a single gitignore file.


## Demo

![demo](docs/img/scrennshot.gif)


## Installation

Download the binary from [GitHub Releases](https://github.com/uu64/gi/releases/latest) and drop it in your `$PATH`.

If you use Homebrew, please execute the following command.

```
$ brew install uu64/tap/gi
```


## Configuration

TBD


## License

[MIT](LICENSE)
