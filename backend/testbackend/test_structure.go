package testbackend

import (
	"testing"

	"github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/backend"
	"github.com/owulveryck/onnx-go/internal/pb-onnx"
	"github.com/stretchr/testify/assert"
	"gorgonia.org/tensor"
)

// TestCase describes an integration test
type TestCase struct {
	Title           string
	Model           *pb.ModelProto
	Inputs          []tensor.Tensor
	ExpectedOutputs []tensor.Tensor
}

// RunTest Returns a function to be executed against the ComputationBackend.
// The return function should be executed via a call to testing.Run(...)
func (tc *TestCase) RunTest(b backend.ComputationBackend) func(t *testing.T) {
	return func(t *testing.T) {
		m := onnx.NewModel(b)
		err := m.DecodeProto(tc.Model)
		if err != nil {
			if _, ok := err.(*onnx.ErrNotImplemented); ok {
				t.Skip(err)
			}
			t.Fatal(err)
		}
		for i := range tc.Inputs {
			err := m.SetInput(i, tc.Inputs[i])
			if err != nil {
				t.Fatal(err)
			}
		}

		err = b.Run()
		if err != nil {
			if _, ok := err.(*onnx.ErrNotImplemented); ok {
				t.Skip(err)
			}
			t.Fatal(err)
		}
		outputs := m.GetOutputs()
		if len(outputs) != len(tc.ExpectedOutputs) {
			t.Fatalf("expected %v output, got %v", len(tc.ExpectedOutputs), len(outputs))
		}
		for i := range outputs {
			assert.Equal(t, outputs[i], tc.ExpectedOutputs[i], "the two tensors should be equal.")
		}

	}

}
