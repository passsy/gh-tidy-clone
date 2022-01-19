# Tidy clone

Github cli extension (`gh extension`) to clone repos into `~/Projects/github/{org}/{repo}` on the local filesystem

## Install

```bash
gh extension install passsy/gh-tidy-clone
```

## Usage

```bash
gh tidy-clone https://github.com/passsy/deep_pick.git

# Clones into ~/Projects/github/passsy/deep_pick
```

## Development

```bash
go build
gh extension install .
```