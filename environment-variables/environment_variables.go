package main

import (
	"os"
	"strings"
	"fmt"
	"sync"
)

type FunctionalMap map[string]interface{}

func (c FunctionalMap) get(key string) interface{} {
	return c[key]
}

func (c FunctionalMap) put(key string, val interface{}) {
	c[key] = val
}

func (c FunctionalMap) threadSafetyPut(key string, val interface{}) {
	mutex := &sync.Mutex{}
	mutex.Lock()
	c[key] = val
	mutex.Unlock()
}

func main() {
	environ := os.Environ()
	environments := FunctionalMap{}
	for _, environment := range environ {
		splitEnvironment := strings.Split(environment, "=")
		environments.threadSafetyPut(splitEnvironment[0], splitEnvironment[1])
	}

	for k, v := range environments {
		fmt.Println("Key:", k, "and val:", v)
	}

}
