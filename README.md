# i3-glance

i3-glance is a daemon for the i3 window manager that automatically renames workspaces based on the windows they contain.

## Install

Using go install:

```sh
go install github.com/ls4154/i3-glance@latest
```

Or build manually:

```sh
git clone https://github.com/ls4154/i3-glance.git
cd i3-glance
go build .
```

## Usage

```sh
./i3-glance --config app-icons.json &
```
