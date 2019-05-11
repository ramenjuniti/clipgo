<img src="https://i.gyazo.com/532cc01852ed06b797535fba2c5f195c.jpg" width=128>

# Clip Go

![Chrome Web Store](https://img.shields.io/chrome-web-store/v/khkfegmjjbijinlbmffohmiofhpnjlja.svg)
![](https://img.shields.io/chrome-web-store/users/khkfegmjjbijinlbmffohmiofhpnjlja.svg)

This is a chrome extension for running a selected golang code in [The Go Playground](https://play.golang.org/).

[![Image from Gyazo](https://i.gyazo.com/b5a55cee8b61acb4388ed0aa767a55d4.gif)](https://gyazo.com/b5a55cee8b61acb4388ed0aa767a55d4)

## Install

### For Users

[![](https://developer.chrome.com/webstore/images/ChromeWebStore_BadgeWBorder_v2_340x96.png)](https://chrome.google.com/webstore/detail/clip-go/khkfegmjjbijinlbmffohmiofhpnjlja)

### For Developers

1. Setup for macOS

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

   [![Image from Gyazo](https://i.gyazo.com/80b67452913a6147aa89cd05c6c78f4a.png)](https://gyazo.com/80b67452913a6147aa89cd05c6c78f4a)

5. Click the LOAD UNPACKED button and select the `build` directory.

   [![Image from Gyazo](https://i.gyazo.com/837cf2b32fbe485cb1b360aa31e052c3.png)](https://gyazo.com/837cf2b32fbe485cb1b360aa31e052c3)

## Test

```
make test
```

## LICENSE

This software is released under the MIT License, see LICENSE.
