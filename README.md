# i3-glance

i3-glance is a lightweight daemon for the i3 window manager that automatically renames workspaces based on the applications running in each workspace.

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

Or add to your i3 config for automatic startup:

```sh
exec --no-startup-id i3-glance --config ~/.config/i3/app-icons.json
```

See the included `app-icons.json` file for a sample config.
