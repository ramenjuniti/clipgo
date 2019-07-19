<div align="center">
  <img src="https://i.gyazo.com/eedc49e827385ac25527264bac38d8c4.png" width=256>
</div>

# Clip Go

[![CircleCI](https://circleci.com/gh/ramenjuniti/clipgo.svg?style=svg)](https://circleci.com/gh/ramenjuniti/clipgo)
[![Go Report Card](https://goreportcard.com/badge/github.com/ramenjuniti/clipgo)](https://goreportcard.com/report/github.com/ramenjuniti/clipgo)
![Chrome Web Store](https://img.shields.io/chrome-web-store/v/khkfegmjjbijinlbmffohmiofhpnjlja.svg)
![](https://img.shields.io/chrome-web-store/users/khkfegmjjbijinlbmffohmiofhpnjlja.svg)

This is a chrome extension for running a selected golang code in [The Go Playground](https://play.golang.org/).

A code validation and formatting use [go/format](https://golang.org/pkg/go/format/).

<div align="center">
  <img src="https://i.gyazo.com/072747dfe9190e3e4ecf1e097218a48e.gif">
</div>

## Dependency

- go 1.12

## Install

### For Users

[![](https://developer.chrome.com/webstore/images/ChromeWebStore_BadgeWBorder_v2_340x96.png)](https://chrome.google.com/webstore/detail/clip-go/khkfegmjjbijinlbmffohmiofhpnjlja)

**Reload your browser**

### For Developers

1. Setup (using homebrew)

```
brew install go
brew install yarn
```

2. Clone this repository and build.

```
git clone https://github.com/ramenjuniti/clipgo.git
cd clipgo
make install
make build
```

3. Open the Extension Management page by navigating to `chrome://extensions`.

4. Enable Developer Mode by clicking the toggle switch next to Developer mode.

<div align="center">
  <img src="https://i.gyazo.com/80b67452913a6147aa89cd05c6c78f4a.png">
</div>

5. Click the LOAD UNPACKED button and select the `build` directory.

<div align="center">
  <img src="https://i.gyazo.com/837cf2b32fbe485cb1b360aa31e052c3.png">
</div>

## Test

```
make test
```

## LICENSE

This software is released under the MIT License, see LICENSE.
