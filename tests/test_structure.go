package tests

import (
	"testing"

	"github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/internal/pb-onnx"
	"gorgonia.org/tensor"
)

// ComputationBackend is a backend that can run the graph
type ComputationBackend interface {
	onnx.Backend
	Run() error
}

// TestCase describes an integration test
type TestCase struct {
	Title           string
	Model           *pb.ModelProto
	Inputs          []tensor.Tensor
	ExpectedOutputs []tensor.Tensor
}

// TestOperator ...
func TestOperator(testCase TestCase) func(t *testing.T) {
	return func(t *testing.T) {

	}
}
