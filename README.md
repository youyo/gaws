# gaws

[![CircleCI](https://circleci.com/gh/youyo/gaws.svg?style=svg)](https://circleci.com/gh/youyo/gaws)

gaws is a command to complement aws-cli.

## Install

### Install with Homebrew on macOS

```
brew tap youyo/gaws
brew install gaws
```

or download binary.

```
wget https://github.com/youyo/gaws/releases/latest/download/gaws_darwin_amd64.zip
```

### Install for Linux

```
wget https://github.com/youyo/gaws/releases/latest/download/gaws_linux_amd64.tar.gz
```

## Enabling shell autocompletion

### Using zsh

Write to your .zshrc file.

```
source <(gaws completion zsh)
```
