package core

import "github.com/VladyslavLukyanenko/GopherAlert/core"

type Monitor interface {
	Monitor(core.Monitor)

}
