package integration

func afterAll() {
	for _, container := range runningContainers {
		container.Terminate()
	}
}
