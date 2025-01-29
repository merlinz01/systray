//go:build windows

package main

import (
	"fmt"
	"unsafe"

	"github.com/merlinz01/systray"
	"golang.org/x/sys/windows"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIconFromFilePath("example/app.ico")
	systray.SetTooltip("Unicode works here 棒棒嗒")
	systray.SetOpenOnLeftClick(false)
	systray.SetOpenOnRightClick(true)

	systray.AddMenuItem("Quit").SetCallback(systray.Quit)

	systray.AddSeparator()

	trayOpenedCount := 0

	mOpenedCount := systray.AddMenuItem("<this text will be replaced>")
	mOpenedCount.Disable()

	mChange := systray.AddMenuItem("Click to change me")
	mChange.SetCallback(func() {
		mChange.SetTitle("I've Changed")
	})

	mChecked := systray.AddMenuItem("Checked")
	mChecked.Check()
	mChecked.SetCallback(func() {
		if mChecked.Checked() {
			mChecked.Uncheck()
			mChecked.SetTitle("Unchecked")
		} else {
			mChecked.Check()
			mChecked.SetTitle("Checked")
		}
	})

	mEnabled := systray.AddMenuItem("Click to disable me")
	mEnabled.SetCallback(func() {
		mEnabled.SetTitle("Disabled")
		mEnabled.Disable()
	})

	systray.AddMenuItem("I do nothing").SetIconFromFilePath("example/app.ico")

	subMenuTop := systray.AddMenuItem("This is a submenu")

	subMenuMiddle := subMenuTop.AddSubMenuItem("This is a submenu of the submenu")

	subMenuMiddle.AddSubMenuItem("Panic!").SetCallback(func() {
		panic("panic button pressed")
	})

	subMenuMiddle.AddSubMenuItem("This is a submenu of the submenu of the submenu")

	systray.AddSeparator()

	mToggle := systray.AddMenuItem("Hide/show some menu items")
	shown := true
	toggle := func() {
		if shown {
			mEnabled.Hide()
			mChange.Hide()
			mChecked.Hide()
			shown = false
		} else {
			mEnabled.Show()
			mChange.Show()
			mChecked.Show()
			shown = true
		}
	}
	mToggle.SetCallback(toggle)

	systray.AddMenuItem("Reset all items").SetCallback(func() {
		systray.ResetMenu()
		systray.AddMenuItem("Quit").SetCallback(systray.Quit)
		trayOpenedCount = -1
	})

	systray.OnTrayOpened(func() {
		if trayOpenedCount == -1 {
			return
		}
		trayOpenedCount++
		mOpenedCount.SetTitle(fmt.Sprintf("The menu has been opened %d time(s)", trayOpenedCount))
	})
}

func onExit() {
	showMessageBox("Goodbye", "onExit called")
}

func showMessageBox(title string, message string) {
	hUser32 := windows.NewLazySystemDLL("user32.dll")
	hMessageBox := hUser32.NewProc("MessageBoxW")
	pTitle, err := windows.UTF16PtrFromString(title)
	if err != nil {
		panic(err)
	}
	pMessage, err := windows.UTF16PtrFromString(message)
	if err != nil {
		panic(err)
	}
	hMessageBox.Call(0, uintptr(unsafe.Pointer(pMessage)), uintptr(unsafe.Pointer(pTitle)), 0x40)
}
