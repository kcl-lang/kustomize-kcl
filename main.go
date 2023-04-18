package main

import (
	"fmt"
	"os"

	"github.com/KusionStack/krm-kcl/pkg/edit"
	"sigs.k8s.io/kustomize/kyaml/kio"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

func main() {
	rw := &kio.ByteReadWriter{Reader: os.Stdin, Writer: os.Stdout, KeepReaderAnnotations: true}
	p := kio.Pipeline{
		Inputs:  []kio.Reader{rw},             // read the inputs into a slice
		Filters: []kio.Filter{filter{rw: rw}}, // run the filter against the inputs
		Outputs: []kio.Writer{rw},             // copy the inputs to the output
	}
	if err := p.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

// API defines the input API schema as a struct
type API struct {
	Spec struct {
		// Source is a required field for providing a KCL script inline.
		Source string `json:"source" yaml:"source"`
		// Params are the parameters in key-value pairs format.
		Params map[string]interface{} `json:"params,omitempty" yaml:"params,omitempty"`
	} `json:"spec" yaml:"spec"`
}

// filter implements kio.Filter
type filter struct {
	rw *kio.ByteReadWriter
}

// Filter checks each input and ensures that all containers have cpu and memory
// reservations set, otherwise it returns an error.
func (f filter) Filter(in []*yaml.RNode) ([]*yaml.RNode, error) {
	api := f.parseAPI()
	st := &edit.SimpleTransformer{
		Name:           "kcl-function-run",
		Source:         api.Spec.Source,
		FunctionConfig: f.rw.FunctionConfig,
	}
	return st.Transform(in)
}

// parseAPI parses the functionConfig into an API struct.
func (f *filter) parseAPI() API {
	// Parse the input function config.
	var api API
	if err := yaml.Unmarshal([]byte(f.rw.FunctionConfig.MustString()), &api); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	return api
}
