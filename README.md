# systray

systray is a Windows-specific Go library to place an icon and menu in the notification area.

## Usage

This package comes with an example application that shows how to use the library.
See [example/main.go](example/main.go) for more details.

To run the example app, run the following commands:

```sh
git clone https://github.com/merlinz01/systray
cd systray

go run ./example
# or, to hide the console window when not running in a terminal
go run -ldflags "-H=windowsgui" ./example
```

## License

This project is licensed under the Apache License, Version 2.0 - see the [LICENSE](LICENSE) file for details.

## Credits

This project is based on the [fyne.io/systray](https://github.com/fyne-io/systray) module,
with non-Windows code removed and some API changes.
