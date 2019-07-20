package server

func inLineIntRange(min, max int) int {
	return inLineRand().Intn(max-min) + min
}
