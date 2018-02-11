# Changelog

## 2.10.5-7 (2018/02/11)

* Move ia32/x64 to win32/win64 for arch def
* Add portapp.json file
* Uncheck run app in setup

## 2.10.5-6 (2018/02/09)

* Ability to pass custom args to the portable process

## 2.10.5-5 (2018/02/08)

* Upgrade to Rocket.Chat 2.10.5

## 2.10.3-4 (2018/02/03)

* Upgrade to Rocket.Chat 2.10.3

## 2.10.2-3 (2018/01/26)

* Upgrade to Rocket.Chat 2.10.2

## 2.10.1-2 (2018/01/13)

* Rebuild electron.asar to push data directly in `data` subfolder
* No need to override USERPROFILE environment variable anymore

> :warning: **UPGRADE NOTES**
> * Move everything in `data\AppData\Roaming\Rocket.Chat+\*` to `data` folder and remove folder `data\AppData`.

## 2.10.1-1 (2017/11/24)

* Initial version based on Rocket.Chat 2.10.1
