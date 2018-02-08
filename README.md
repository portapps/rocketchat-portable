<p align="center"><a href="https://github.com/portapps/rocketchat-portable" target="_blank"><img width="100" src="https://github.com/portapps/rocketchat-portable/blob/master/res/papp.png"></a></p>

<p align="center">
  <a href="https://github.com/portapps/rocketchat-portable/releases/latest"><img src="https://img.shields.io/github/release/portapps/rocketchat-portable.svg?style=flat-square" alt="GitHub release"></a>
  <a href="https://github.com/portapps/rocketchat-portable/releases/latest"><img src="https://img.shields.io/github/downloads/portapps/rocketchat-portable/total.svg?style=flat-square" alt="Total downloads"></a>
  <a href="https://ci.appveyor.com/project/crazy-max/rocketchat-portable"><img src="https://img.shields.io/appveyor/ci/crazy-max/rocketchat-portable.svg?style=flat-square" alt="AppVeyor"></a>
  <a href="https://goreportcard.com/report/github.com/portapps/rocketchat-portable"><img src="https://goreportcard.com/badge/github.com/portapps/rocketchat-portable?style=flat-square" alt="Go Report"></a>
  <a href="https://www.codacy.com/app/crazy-max/rocketchat-portable"><img src="https://img.shields.io/codacy/grade/8fa6391e00fc4a7e9eb1b1108d5655a2.svg?style=flat-square" alt="Code Quality"></a>
  <a href="https://beerpay.io/portapps/portapps"><img src="https://img.shields.io/beerpay/portapps/portapps.svg?style=flat-square" alt="Beerpay"></a>
  <a href="https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=WQD7AQGPDEPSG"><img src="https://img.shields.io/badge/donate-paypal-7057ff.svg?style=flat-square" alt="Donate Paypal"></a>
</p>

## About

[Rocket.Chat](https://rocket.chat/) portable app made with 🚀 [Portapps](https://github.com/portapps).<br />
Tested on Windows 7, Windows 8.1 and Windows 10.

## Installation

There are different kinds of artifacts :

* `rocketchat-portable-{ia32,x64}-x.x.x-x-setup.exe` : Full portable release of Rocket.Chat as a setup. **Recommended way**!
* `rocketchat-portable-{ia32,x64}-x.x.x-x.7z` : Full portable release of Rocket.Chat as a 7z archive.
* `rocketchat-portable-{ia32,x64}.exe` : Only the portable binary (must be renamed `rocketchat-portable.exe`)
* `rocketchat-setup-x.x.x.exe` : The original release from the [official website](https://rocket.chat/download).

### Fresh installation

Install `rocketchat-portable-{ia32,x64}-x.x.x-x-setup.exe` where you want then run `rocketchat-portable.exe`.

### App already installed

If **you have already installed Rocket.Chat from the original release**, do the same thing as a fresh installation and move :

* `%APPDATA%\Rocket.Chat+` to `data`.

Run `rocketchat-portable.exe` and then you can [remove](https://support.microsoft.com/en-us/instantanswers/ce7ba88b-4e95-4354-b807-35732db36c4d/repair-or-remove-programs) Rocket.Chat from your computer.

### Upgrade

**For an upgrade**, simply download and install the [latest rocketchat-portable setup](https://github.com/portapps/rocketchat-portable/releases/latest).

## How can i help ?

All kinds of contributions are welcomed :raised_hands:!<br />
The most basic way to show your support is to star :star2: the project, or to raise issues :speech_balloon:<br />
But we're not gonna lie to each other, I'd rather you buy me a beer or two :beers:!

[![Beerpay](https://beerpay.io/portapps/portapps/badge.svg?style=beer-square)](https://beerpay.io/portapps/portapps)
or [![Paypal](https://raw.githubusercontent.com/portapps/portapps/master/res/paypal.png)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=WQD7AQGPDEPSG)

## License

MIT. See `LICENSE` for more details.<br />
Rocket icon credit to [Squid Ink](http://thesquid.ink).
