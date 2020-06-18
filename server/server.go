package server

import "rura/bb-server/extcon"

func Server(context *extcon.ExtContext, stop chan int) {

	for true {
		select {
		case <-context.Done():
			return
		}
	}
}
