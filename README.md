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


## Installation

Download the binary from [GitHub Releases](https://github.com/uu64/gi/releases/latest) and drop it in your `$PATH`.

If you use Homebrew, please execute the following command.

```
$ brew install uu64/tap/gi
```


## Configuration

The configuration file is written in yaml format.
The location of the file is as follows.

- `$HOME/.config/gi/config.yml`

The following is an example of a configuration file.

```yaml
repos:
  - owner: uu64
    name: gitignore
    branch: main
  - owner: github
    name: gitignore
    branch: master
auth:
  token: 0123456789abc
cli:
  pagesize: 30
```

### repos

Set the list of repositories where gitignore templates are stored.
If the length of the list is greater than 1, the prompt to select the repository is shown.

Default is as follows.

```yaml
repos:
  - owner: github
    name: gitignore
    branch: master
```

### auth.token

Set the value of a personal API token of GitHub.

Default is empty.

`gi` uses the GitHub API v3 which has a rate limit.
If you encounter a rate limit error, please refer to the following URL to get a token and set it.

[Creating a personal access token](https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/creating-a-personal-access-token)

### cli.pagesize

Set the maximum number of lines to display for the prompt to select gitignore templates.

Default is 20.


## Demo

Default configuration:

![demo1](docs/img/demo1.gif)

Added your own repository:

![demo2](docs/img/demo2.gif)


## License

[MIT](LICENSE)
