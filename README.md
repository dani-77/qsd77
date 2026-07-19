# qsd77

qsd77 is cli tool built with go for quickshell-d77.

## How to

```
git clone https://github.com/dani-77/qsd77.git ~/qsd77

cd ~/qsd77

go build

sudo install qsd77 /usr/bin/
```

Autocompletions for bash, fish and zsh can be generated:

```
qsd77 completion bash
```

## Usage

qsd77 launcher -> calls the launcher menu

qsd77 locker -> lock the screen

qsd77 session -> calls the session menu

qsd77 wallpaper -> calls wallpaper selection

### Choosing a shell/config

All commands accept a `-c`/`--config` flag to target a specific quickshell
config (defaults to `quickshell-d77`):

```
qsd77 launcher -c utumno
```

### Raw IPC passthrough

For any IPC target/action not wrapped by a named command, use `ipc call`
directly, same as `qs ipc call ...`:

```
qsd77 ipc call launcher toggle -c utumno
qsd77 ipc call session open somearg -c quickshell-d77
```

Enjoy
