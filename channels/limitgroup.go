package channels

// A LimitGroup is a group used to limit the number of something. An example would be to
// limit the number of goroutines running at the same time
type LimitGroup struct {
	current int
	done    chan bool
	max     int
}

// Limit sets the maximum limit
func (c *LimitGroup) Limit(l int) {
	c.startDoneChannel()
	c.max = l
}

// Add adds delta number to the LimitGroup
func (c *LimitGroup) Add(delta int) {
	c.startDoneChannel()
	c.current += delta
	return
}

// Done removes one from the current LimitGroup
func (c *LimitGroup) Done() {
	c.Add(-1)
	if c.current < c.max {
		select {
		case c.done <- true:
		default:
		}
	}
	return
}

// Wait will block until there are less than the max limit available. This would be placed
// at the end of a loop that's starting goroutines to wait for an available slot
// before starting a new one.
func (c *LimitGroup) Wait() {
	c.startDoneChannel()
	if c.current < c.max {
		return
	}
	<-c.done
	return
}

func (c *LimitGroup) startDoneChannel() {
	if c.done == nil {
		c.done = make(chan bool)
	}
}
