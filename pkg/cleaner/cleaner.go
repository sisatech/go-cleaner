package cleaner

// Cleaner TODO
type Cleaner struct {
	errChan       chan error
	onFailActions []func()
	// timeToLive    time.Duration
}

// New TODO
func New() *Cleaner {

	out := new(Cleaner)

	out.errChan = make(chan error)
	out.onFailActions = make([]func(), 0)
	// out.timeToLive = ttl

	go out.wait()
	return out
}

// OnFail TODO
func (c *Cleaner) OnFail(f func()) {
	c.onFailActions = append(c.onFailActions, f)
}

func (c *Cleaner) wait() {

	select {
	case e := <-c.errChan:
		if e != nil {
			for i := len(c.onFailActions) - 1; i >= 0; i-- {
				c.onFailActions[i]()
			}
		}
	}

}

// Resolve TODO
func (c *Cleaner) Resolve(err error) {
	c.errChan <- err
}
