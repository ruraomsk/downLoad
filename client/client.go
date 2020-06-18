package client

import "rura/bb-server/extcon"

func Client(context *extcon.ExtContext, stop chan int) {

	for true {
		select {
		case <-context.Done():
			return
		}
	}
}
