// +build kubeall kubernetes

// NOTE: we have build tags to differentiate kubernetes tests from non-kubernetes tests. This is done because minikube
// is heavy and can interfere with docker related tests in terratest. To avoid overloading the system, we run the
// kubernetes tests separately from the others.

package k8s

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Test that RunKubectlAndGetOutputE will run kubectl and return the output by running a can-i command call.
func TestRunKubectlAndGetOutputReturnsOutput(t *testing.T) {
	options := NewKubectlOptions("", "")
	output, err := RunKubectlAndGetOutputE(t, options, "auth", "can-i", "get", "pods")
	require.NoError(t, err)
	require.Equal(t, output, "yes")
}
