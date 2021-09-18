package test_cases

type Result struct {
	Message          string
	TestDescription  string
	TestName         string
	SuiteName        string
	SuiteDescription string
	Passed           bool
}

type ConformanceTest struct {
	Name        string
	Description string
	Run         func(r *Result, data AttestationRequestAndResponse) *Result
}

func (c *ConformanceTest) Execute(data AttestationRequestAndResponse) *Result {
	r := Result{
		TestName:        c.Name,
		TestDescription: c.Description,
		Passed:          true,
	}
	return c.Run(&r, data)
}

type ConformanceTestSuite struct {
	Name        string
	Description string
	Run         func(c *ConformanceTestSuite, data AttestationRequestAndResponse) []*Result
}

func (c *ConformanceTestSuite) Execute(data AttestationRequestAndResponse) []*Result {
	return c.Run(c, data)
}
