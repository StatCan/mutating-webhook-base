package main

import (
	"encoding/json"

	"github.com/statcan/mutating-webhook-base/pkg/mutatingwebhook"
	"k8s.io/api/admission/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Based on:
// https://medium.com/ovni/writing-a-very-basic-kubernetes-mutating-admission-webhook-398dbbcb63ec
// https://github.com/alex-leonhardt/k8s-mutate-webhook

func main() {
	mutatingwebhook.ListenAndMutate(
		&customMutator{},
		mutatingwebhook.MutatingWebhookConfigs{},
	)
}

// Define the variables that you will need.
var (
// define the variables
)

// Initialize the variables.
func init() {
	// initialize the variables via arguments
	// flag.StringVar(&variable, "argument", "default-value", "Argument description.")
}

// Define the variables you need in the struct
type customMutator struct {
	// add variables you may need in your mutating code
}

// This is the function that will be called to mutate
func (cm *customMutator) Mutate(request v1beta1.AdmissionRequest) (v1beta1.AdmissionResponse, error) {
	response := v1beta1.AdmissionResponse{}

	// Default response
	response.Allowed = true
	response.UID = request.UID

	// Decode the object you are trying to mutate.
	// Here's an example with a Pod:

	// Decode the pod object
	var err error
	//pod := v1.Pod{}
	// if err := json.Unmarshal(request.Object.Raw, &pod); err != nil {
	// 	return response, fmt.Errorf("unable to decode Pod %w", err)
	// }

	// Add the logic you wish to implement and create patches:
	patches := []map[string]interface{}{
		{
			"op":    "add",
			"path":  "/spec/tolerations/-",
			"value": "",
		},
	}

	// If there are any patches, they will be appended to the
	if len(patches) > 0 {

		patchType := v1beta1.PatchTypeJSONPatch
		response.PatchType = &patchType

		response.AuditAnnotations = map[string]string{
			// Add annotations to clearly denote that actions have
			// been performed on objects
		}

		response.Patch, err = json.Marshal(patches)
		if err != nil {
			return response, err
		}

		response.Result = &metav1.Status{
			Status: metav1.StatusSuccess,
		}
	}

	return response, nil
}
