package registry

import "todo/server/interface/handler"

type HandlerRegistry struct {
	ItemHandler handler.ItemHandler
}

func NewRegistry(i handler.ItemHandler) HandlerRegistry {
	return HandlerRegistry{
		ItemHandler: i,
	}
}
