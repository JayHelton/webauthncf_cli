package testing

type Result struct {
	message string
	passed  bool
}

type ConformanceTest struct {
	name        string
	description string
}

func (c *ConformanceTest) Run(i interface{}) Result {
	return Result{}
}

type Pipeline struct {
	tests []*ConformanceTest
}

func (p *Pipeline) Register(t *ConformanceTest) {
	p.tests = append(p.tests, t)
}

func (p *Pipeline) Run(data interface{}) []Result {
	results := []Result{}
	for _, test := range p.tests {
		results = append(results, test.Run(data))
	}
	return results
}

func WithFIDOTests() *Pipeline {
	return &Pipeline{}
}
