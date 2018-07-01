package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCatfactsApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CatfactsApi Suite")
}
