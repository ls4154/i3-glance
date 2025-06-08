# i3-glance

i3-glance is a lightweight daemon for the i3 window manager that automatically renames workspaces based on the applications running in each workspace.


Example screenshot with fontawesome icons:

![example](example.png)


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
./i3-glance --config config.json &
```

Or add to your i3 config for automatic startup:

```sh
exec --no-startup-id i3-glance --config ~/path/to/config.json
```

See the included example json config files for how to set up application icons and workspace names.
