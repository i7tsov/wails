package dispatcher

import (
	"errors"
	"github.com/wailsapp/wails/v2/internal/frontend"
	"strconv"
	"strings"
)

func (d *Dispatcher) mustAtoI(input string) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		d.log.Error("cannot convert %s to integer!", input)
	}
	return result
}

func (d *Dispatcher) processWindowMessage(message string, sender frontend.Frontend) (string, error) {
	if len(message) < 2 {
		return "", errors.New("Invalid Event Message: " + message)
	}

	switch message[1] {
	case 'c':
		go sender.WindowCenter()
	case 'T':
		title := message[2:]
		go sender.WindowSetTitle(title)
	case 'F':
		go sender.WindowFullscreen()
	case 'f':
		go sender.WindowUnFullscreen()
	case 's':
		parts := strings.Split(message[3:], ":")
		w := d.mustAtoI(parts[0])
		h := d.mustAtoI(parts[1])
		go sender.WindowSetSize(w, h)
	case 'p':
		parts := strings.Split(message[3:], ":")
		x := d.mustAtoI(parts[0])
		y := d.mustAtoI(parts[1])
		go sender.WindowSetPos(x, y)
	case 'H':
		go sender.WindowHide()
	case 'S':
		go sender.WindowShow()
	case 'M':
		go sender.WindowMaximise()
	case 'U':
		go sender.WindowUnmaximise()
	case 'm':
		go sender.WindowMinimise()
	case 'u':
		go sender.WindowUnminimise()
	case 'Z':
		parts := strings.Split(message[3:], ":")
		w := d.mustAtoI(parts[0])
		h := d.mustAtoI(parts[1])
		go sender.WindowSetMaxSize(w, h)
	case 'z':
		parts := strings.Split(message[3:], ":")
		w := d.mustAtoI(parts[0])
		h := d.mustAtoI(parts[1])
		go sender.WindowSetMinSize(w, h)
	case 'C':
		sender.Quit()
	default:
		d.log.Error("unknown Window message: %s", message)
	}

	return "", nil
}