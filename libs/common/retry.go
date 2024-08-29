package common

import (
	"fmt"
	"reflect"
	"time"
)

var retryCount int = 3

// Retry (UNTESTED) is a wrapper which retries a function a default specified amount of times (typically 3). Not type safe.
func Retry(fn interface{}, args ...interface{}) ([]interface{}, error) {
	fnValue := reflect.ValueOf(fn)

	if fnValue.Kind() != reflect.Func {
		return nil, fmt.Errorf("provided interface is not a function")
	}

	var result []interface{}
	var err error

	for i := 0; i < retryCount; i++ {
		argValues := make([]reflect.Value, len(args))
		for j, arg := range args {
			argValues[j] = reflect.ValueOf(arg)
		}

		results := fnValue.Call(argValues)

		if len(results) > 1 && results[len(results)-1].Type().Implements(reflect.TypeOf((*error)(nil)).Elem()) {
			if err, ok := results[len(results)-1].Interface().(error); ok && err != nil {
				// If an error is returned, wait before retrying
				if i < retryCount-1 {
					time.Sleep(1 * time.Second) // TODO: Exponential backoff?
					continue
				}
				return nil, err
			}
		}

		result = make([]interface{}, len(results))
		for j, r := range results {
			result[j] = r.Interface()
		}
		return result, nil
	}

	return nil, fmt.Errorf("function failed after %d retries: %v ", retryCount, err)
}
