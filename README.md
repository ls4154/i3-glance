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
./i3-glance --config config.toml &
```

See the included example toml config files for how to set up application icons and workspace names.

## i3 Config Example

Add to your `~/.config/i3/config`:

```
exec --no-startup-id i3-glance --config ~/.config/i3-glance/config.toml

bar {
    # Use FontAwesome for icons (check your font name with fc-list)
    font pango:DejaVu Sans Mono, Font Awesome 7 Free 10
    status_command i3status
}
```
