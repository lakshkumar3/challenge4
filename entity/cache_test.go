package entity

import (
	"fmt"
	"testing"
)

var (
	SaveEquationCachefunc func() error
)

type CacheMock struct{}

func (e CacheMock) SaveEquationCache() error {
	return SaveEquationCachefunc()
}

func TestCache_SaveEquationCache(t *testing.T) {
	cache = &CacheMock{}
	SaveEquationCachefunc = func() error {
		return nil
	}
	err := cache.SaveEquationCache()
	if err != nil {
		t.Fail()
	}

}
func TestNegativeCache_SaveEquationCache(t *testing.T) {
	cache = &CacheMock{}
	SaveEquationCachefunc = func() error {
		return fmt.Errorf("error 404")
	}
	err := cache.SaveEquationCache()
	if err == nil {
		t.Fail()
	}

}
