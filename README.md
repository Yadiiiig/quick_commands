# Quick Commands
A collection of commands to use in your terminal/command prompt.

## Installation

##### Building all commmands at once
```
go run build.go <OS>
```
OS - windows, linux or mac.
##### Enabling commands globally
Add the bin folder to your environment variables. Now you'll be able to use the command everywhere.


## Available Commands

* searchfile `file-/folder name` - Finds all the corresponding paths of files and folder with that specific name. 
* timer `minutes` - Starts a countdown timer for specified time.
* so `-l/-limit int` `search term` - Stackoverflow search, limit flag is optional. From default it returns 7 links. 
* lock - Locks the PC, windows only.
* internetspeed - checks download, upload, ping, ip, isp and the location of the server via headless chrome. This means chrome must be installed!

## Known issues:
Building certain commands on Mac will need: xcode-select (`xcode-select --install`).
This will download the Command Line Tools.

The command internetspeed works via headless chrome. This means chrome must be installed if you run this command.

## To-Do Commands

* Eval
* Weather
* YouTube/Reddit/TikTok Download
* ~~Lock~~
* Matrix
* Cron
* Google
* ~~Stackoverflow search~~
* Useful links saver
* ~~Help command~~ *Needs to be rewritten*
* Repo downloader
* ...

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
