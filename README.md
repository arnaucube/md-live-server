# md-live-server
Server that renders markdown files and live updates the page each time that the file is updated.

![screenshot00](https://raw.githubusercontent.com/arnaucube/md-live-server/master/screenshot00.png 'screenshot00')

## Usage
Put the binary file `md-live-server` in your `$PATH`, and then go to the directory where are the markdown files that want to live render, and just use:
```
> md-live-server
```
And then go to the browser at `http://127.0.0.1:8080`

## Features
- [x] server rendering .md files
- [x] live reload when .md file changes
- [x] directory files list on `/` endpoint
- [ ] graphviz
- [ ] colour `<code>` with syntax highlighting
