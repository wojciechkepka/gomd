# gomd 📘
s

[![Build Status](https://travis-ci.org/wojciechkepka/gomd.svg?branch=master)](https://travis-ci.org/wojciechkepka/gomd)

Quickly display formatted markdown files in your browser.  

![working-app](https://raw.githubusercontent.com/wojciechkepka/gomd/master/assets/gomd.gif)

## About
`gomd` sets up an http server rendering markdown files in selected flavour and theme.  

## Features
- **Simple**
  - No setup required. `gomd` comes with all the batteries included, no need for static assets etc.
- **Monitoring files**
  - `gomd` will monitor the directory for any changes and update the files whenever any file is modified or added.
- **Hot reloading**
  - Whenever a file is updated `gomd` will trigger a reload of tab in browser.
- **Code Highlight**
  - All blocks of code in most common languages will be color highlighted.
- **Selectable themes**
  - Choose from available code themes:
    - `solarized`, `monokai`, `paraiso` available in dark and light versions
    - `dracula`, `github`, `vs`, `xcode`

## Installing
 - **AUR**
   - Available in package `gomd-git`.
 - **Build from source**
   - Requires
     - sassc or sass
     - make
     - clean-css-cli (`npm install clean-css-cli -g`)
     - go
   - `git clone https://github.com/wojciechkepka/gomd`
   - `cd ./gomd`
   - Linux: `make && sudo cp gomd /usr/bin/`
   - macOS: `make && sudo cp gomd /usr/local/bin`

## Configuration
By default when running `gomd` it will look for files in `.` and bind to `127.0.0.1:5001`.

To change default port and address use `--bind-port` and `--bind-addr`.
For example:
    `gomd --bind-port 1337 --bind-addr 192.168.0.1`

To view a different directory use:
    `gomd --dir /some/different/directory`

You can view the files in dark and light mode.


## License
[MIT](https://github.com/wojciechkepka/gomd/blob/master/LICENSE)
