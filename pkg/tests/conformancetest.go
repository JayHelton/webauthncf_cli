package tests

type Result struct {
	Message         string
	TestDescription string
	TestName        string
	Passed          bool
}

type ConformanceTest struct {
	Name        string
	Description string
	Run         func(interface{}) Result
}
type ConformanceTestSuite struct {
	Name        string
	Description string
	Run         func(interface{}) []Result
}
