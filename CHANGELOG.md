# Changelog

## 2.17.5-15 (2020/02/07)

* Rocket.Chat 2.17.5

## 2.17.3-14 (2020/02/03)

* Rocket.Chat 2.17.3

## 2.17.2-13 (2019/12/23)

* Rocket.Chat 2.17.2
* Add cleanup config
* Portapps 1.31.0

## 2.17.1-12 (2019/12/11)

* Rocket.Chat 2.17.1

## 2.17.0-11 (2019/12/03)

* Rocket.Chat 2.17.0

## 2.16.2-10 (2019/11/10)

* Rocket.Chat 2.16.2
* Portapps 1.30.1

## 2.16.1-9 (2019/11/06)

* Rocket.Chat 2.16.1
* Portapps 1.30.0

## 2.16.0-8 (2019/10/15)

* Rocket.Chat 2.16.0
* Portapps 1.28.0

## 2.15.5-7 (2019/08/10)

* Rocket.Chat 2.15.5 (rebuilt)

## 2.15.5-6 (2019/08/09)

* Rocket.Chat 2.15.5
* Portapps 1.25.0

## 2.15.3-5 (2019/04/30)

* Rocket.Chat 2.15.3
* Portapps 1.22.2

## 2.15.2-4 (2019/04/18)

* Rocket.Chat 2.15.2
* Portapps 1.22.1

## 2.15.1-3 (2019/03/15)

* Rocket.Chat 2.15.1
* Portapps 1.20.2

## 2.15.0-2 (2019/02/25)

* Rocket.Chat 2.15.0
* Switch to Travis CI

## 2.14.7-23 (2019/01/18)

* Rocket.Chat 2.14.7

## 2.14.5-22 (2018/11/04)

* Rocket.Chat 2.14.5

## 2.14.4-21 (2018/11/22)

* Rocket.Chat 2.14.4

## 2.14.3-20 (2018/11/14)

* Rocket.Chat 2.14.3

## 2.14.1-19 (2018/10/26)

* Rocket.Chat 2.14.1

## 2.14.0-18 (2018/10/11)

* Rocket.Chat 2.14.0

## 2.13.3-17 (2018/10/05)

* Fix `update.json` cannot be written
* Write update settings properly

## 2.13.3-16 (2018/10/03)

* Fix Windows desktop notifications not working (Issue #3)

## 2.13.3-15 (2018/09/19)

* Rocket.Chat 2.13.3

## 2.13.2-14 (2018/09/17)

* electron userData path overrided by app (Issue #2)
* `Rocket.Chat+.exe` renamed `Rocket.Chat.exe` (Issue #1)

## 2.13.2-13 (2018/09/11)

* Rocket.Chat 2.13.2
* Go 1.11
* Use [go mod](https://golang.org/cmd/go/#hdr-Module_maintenance) instead of `dep`

## 2.13.1-12 (2018/08/31)

* Rocket.Chat 2.13.1
* win32 releases restored

## 2.13.0-11 (2018/08/28)

* Rocket.Chat 2.13.0
* No more win32 releases

## 2.12.1-10 (2018/08/15)

* Rocket.Chat 2.12.1

## 2.12.0-9 (2018/08/05)

* Rocket.Chat 2.12.0

## 2.11.0-8 (2018/06/11)

* Rocket.Chat 2.11.0

## 2.10.5-7 (2018/02/11)

* Move ia32/x64 to win32/win64 for arch def
* Add portapp.json file
* Uncheck run app in setup

## 2.10.5-6 (2018/02/09)

* Ability to pass custom args to the portable process

## 2.10.5-5 (2018/02/08)

* Rocket.Chat 2.10.5

## 2.10.3-4 (2018/02/03)

* Rocket.Chat 2.10.3

## 2.10.2-3 (2018/01/26)

* Rocket.Chat 2.10.2

## 2.10.1-2 (2018/01/13)

* Rebuild electron.asar to push data directly in `data` subfolder
* No need to override USERPROFILE environment variable anymore

> :warning: **UPGRADE NOTES**
> * Move everything in `data\AppData\Roaming\Rocket.Chat+\*` to `data` folder and remove folder `data\AppData`.

## 2.10.1-1 (2017/11/24)

* Initial version based on Rocket.Chat 2.10.1
