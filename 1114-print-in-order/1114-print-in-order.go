type Foo struct {
    s chan struct{}
    t chan struct{}
}

func NewFoo() *Foo {
	return &Foo{
        s:make(chan struct{}),
        t:make(chan struct{}),
	}
}

func (f *Foo) First(printFirst func()) {
	// Do not change this line
	printFirst()
    f.s<-struct{}{}
}

func (f *Foo) Second(printSecond func()) {
    <-f.s
	/// Do not change this line
	printSecond()
    f.t<-struct{}{}
}

func (f *Foo) Third(printThird func()) {
    <-f.t
	// Do not change this line
	printThird()
}