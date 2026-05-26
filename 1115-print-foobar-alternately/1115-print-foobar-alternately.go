type FooBar struct {
	n      int
	first  chan struct{}
	second chan struct{}
}

func NewFooBar(n int) *FooBar {
	return &FooBar{n: n, first: make(chan struct{}), second: make(chan struct{})}
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		if i > 0 {
			<-fb.second
		}
		// printFoo() outputs "foo". Do not change or remove this line.
		printFoo()
		fb.first <- struct{}{}
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.first
		// printBar() outputs "bar". Do not change or remove this line.
		printBar()
		if i < fb.n-1 {
			fb.second <- struct{}{}
		}

	}
}