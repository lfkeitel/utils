package channels

type maxChanGroup struct {
	numOfConnections int
	goChan           chan bool
	maxConnections   int
}

func NewMaxChanGroup(max int) maxChanGroup {
	return maxChanGroup{
		maxConnections: max,
	}
}

func (c *maxChanGroup) Add(delta int) {
	if c.goChan == nil {
		c.goChan = make(chan bool)
	}
	c.numOfConnections += delta
	return
}

func (c *maxChanGroup) Done() {
	c.Add(-1)
	if c.numOfConnections < c.maxConnections {
		c.goChan <- true
	}
	return
}

func (c *maxChanGroup) Wait() {
	if c.numOfConnections < c.maxConnections {
		return
	}
	<-c.goChan
	return
}
