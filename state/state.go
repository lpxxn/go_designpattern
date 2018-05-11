package state

import "fmt"

type State interface {
	DoAction(c *Context)
}

type Run struct { }

func (r *Run) DoAction(c *Context) {
	c.SetContext(new(Stop))
	fmt.Println("running......")
}


type Stop struct { }

func (s *Stop) DoAction(c *Context) {
	c.SetContext(new(Resume))
	fmt.Println("stopping...")
}

type Resume struct { }

func (r *Resume) DoAction(c *Context) {
	c.SetContext(new(Run))
	fmt.Println("Resuming....")
}


type Context struct {
	state State
}


func NewContext() *Context{
	c := new(Context)
	c.state = new(Resume)
	return c
}

func (c *Context) SetContext(s State) {
	c.state = s
}


func (c *Context) DoAction() {
	c.state.DoAction(c)
}