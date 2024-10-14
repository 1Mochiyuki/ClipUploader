# README

## About

This is the official Wails Svelte template.

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on <http://localhost:34115>. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.

## Roadmap

- [:white_check_mark:] Create perfences file (if doesnt exist) with sane defaults
- [:white_check_mark] Save settings to file
- [] Add auth to file hosts
- [:white_check_mark] Actually upload to a file host
- [ ] Add retrieving upload history
- [ ] Add setting login information per host
- [ ] Copy all links to clipboard
- [ ] Notification for when files are uploaded
- [ ] Toggle notification sound
- [ ] Deploy to Vercel :)
