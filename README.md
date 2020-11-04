# gomd 📘

[![Build Status](https://travis-ci.org/wojciechkepka/gomd.svg?branch=master)](https://travis-ci.org/wojciechkepka/gomd)

Quickly display formatted markdown files locally in browser.  

![working-app](https://raw.githubusercontent.com/wojciechkepka/gomd/master/assets/gomd.gif)

## About
`gomd` sets up an http server rendering markdown files in selected flavour and theme.  

## Features
- **Simple**
  - No setup required. `gomd` comes as a single binary with all the batteries included, no need for static assets etc.
- **Monitoring files**
  - `gomd` will monitor the directory for any changes and update files whenever a file is modified or added.
- **Hot reloading**
  - On file update `gomd` will trigger a reload of the tab in the browser.
- **Code Highlight**
  - Blocks of code in most common programming languages will automatically be color highlighted.
- **Selectable themes**
  - Choose from available code themes:
    - `solarized`, `monokai`, `paraiso` available in dark and light versions
    - `dracula`, `github`, `vs`, `xcode`

## Installing
 - **AUR**
   - Available in package `gomd-git`.
 - **Prebuilt binaries**
   - Available [here](https://github.com/wojciechkepka/gomd/releases)
   - macOS, windows, linux
 - **From source**
   - Requires
     - go
   - `git clone https://github.com/wojciechkepka/gomd`
   - `cd ./gomd`
   - Linux: `go build && sudo cp gomd /usr/bin/`
   - macOS: `go build && sudo cp gomd /usr/local/bin`

## Configuration
By default `gomd` will look for files in current working directory `.` and bind to `127.0.0.1:5001`.  

To change default port and address use `--bind-port` and `--bind-addr`.  
For example:  
 - `gomd --bind-port 1337 --bind-addr 192.168.0.1`

To view a different directory use:  
 - `gomd --dir /some/different/directory`

You can view the files in dark and light mode by using the switch available on rendered file view.

## Contributing
If you wish to contribute thanks! There are some things to keep in mind while editing the source.  
When making any changes to sass files the whole project has to be built with make for it to 
generate appropriate go files containing styles.  

 - Build requirements
   - go
   - sassc or sass
   - make
   - pkger (`go get github.com/markbates/pkger/cmd/pkger`)
   - clean-css-cli (`npm install clean-css-cli -g`)
 - To run tests, recompile styles and build the project run `make`
 - To build executables for multiple platforms run `make buildall`
 - To run all tests execute `make tests`
 - To recompile styles run `make styles`

## License
[MIT](https://github.com/wojciechkepka/gomd/blob/master/LICENSE)
