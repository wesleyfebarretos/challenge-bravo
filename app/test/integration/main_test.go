package integration

import (
	"log"
	"os"
	"path"
	"runtime"
	"testing"
	"time"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../..")
	os.Chdir(dir)
	log.Printf("Setting test root folder to: %v\n", dir)
}

func TestMain(m *testing.M) {
	start := time.Now()

	beforeAll()

	log.Printf("setup took %s seconds\n", time.Since(start))

	exitVal := m.Run()

	afterAll()

	os.Exit(exitVal)
}
