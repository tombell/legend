# Legend

API server for track information from [rekordbox][rekordbox].

[rekordbox]: https://rekordbox.com/en/

This project couldn't have been built without inspiration from
[PRACT-OBS][pract-obs], [unbox][unbox], and [supbox][supbox].

[pract-obs]: https://github.com/LePopal/PRACT-OBS
[unbox]: https://github.com/erikrichardlarson/unbox
[supbox]: https://github.com/gabek/supbox

Please note that since I built this project for myself, it is built with a
number of assumptions and opinions in mind.

- It assumes you're running macOS (this is my main platform I use rekordbox on)
- It assumes you're only using 2 channels, and the first track is played on the
  first deck

Due to limitations with rekordbox and how we're getting the track information,
the track is only registered as playing/played after one minute.

## Installation

Coming soon.

## Usage

You must have rekordbox running on the system that you plan on running `legend`
on.

Run `legend`, this will start the API server and rekordbox monitor.

Visit http://localhost:8888/public/index.html in your browser to check that it
is serving the overlay

In OBS add a **Browser** source and set the URL to
http://localhost:8888/public/index.html (you can change `localhost` to the IP
address of the system `legend` is running on if you use OBS on another system)
