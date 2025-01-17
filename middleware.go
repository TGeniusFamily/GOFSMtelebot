package telebot

// MiddlewareFunc represents a middleware processing function,
// which get called before the endpoint group or specific handler.
type MiddlewareFunc func(HandlerFunc) HandlerFunc

// Group is a separated group of handlers, united by the general middleware.
type Group struct {
	b          *Bot
	middleware []MiddlewareFunc
}

// Use adds middleware to the chain.
func (g *Group) Use(middleware ...MiddlewareFunc) {
	g.middleware = append(g.middleware, middleware...)
}

// Handle adds endpoint handler to the bot, combining group's middleware
// For remove state set it at 0
// with the optional given middleware.
func (g *Group) Handle(endpoint interface{}, h HandlerFunc, s State, m ...MiddlewareFunc) {
	g.b.Handle(endpoint, h, s, append(g.middleware, m...)...)
}
