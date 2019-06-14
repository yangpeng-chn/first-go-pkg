package test

type Example struct {
	Text string
}

func (exa Example) Hello() string {
	return exa.Text
}
