/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package yaml implements fuzzers for yaml deserialization routines in
// Kubernetes. These targets are compatible with the github.com/dvyukov/go-fuzz
// fuzzing framework.
package yaml

import (
	"bytes"

	"gopkg.in/yaml.v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	sigyaml "sigs.k8s.io/yaml"
)

// FuzzDuration is a fuzz target for unmarshaling Duration defined in "k8s.io/apimachinery/pkg/apis/meta/v1".
// This target also checks that the unmarshaled result can be marshaled back to the input.
func FuzzDuration(b []byte) int {
	var unmarshalResult struct {
		D metav1.Duration `json:"d"`
	}
	if err := yaml.Unmarshal(b, &unmarshalResult); err != nil {
		return 0
	}
	marshalResult, err := yaml.Marshal(&unmarshalResult)
	if err != nil {
		panic(err)
	}
	if !bytes.Equal(marshalResult, b) {
		panic("marshalResult != input")
	}
	return 1
}

// FuzzMicroTime is a fuzz target for unmarshaling MicroTime defined in "k8s.io/apimachinery/pkg/apis/meta/v1".
// This target also checks that the unmarshaled result can be marshaled back to the input.
func FuzzMicroTime(b []byte) int {
	var unmarshalResult struct {
		T metav1.MicroTime `json:"t"`
	}
	if err := yaml.Unmarshal(b, &unmarshalResult); err != nil {
		return 0
	}
	marshalResult, err := yaml.Marshal(&unmarshalResult)
	if err != nil {
		panic(err)
	}
	if !bytes.Equal(marshalResult, b) {
		panic("marshalResult != input")
	}
	return 1
}

// FuzzSigYaml is a fuzz target for "sigs.k8s.io/yaml" unmarshaling.
func FuzzSigYaml(b []byte) int {
	t := struct{}{}
	m := map[string]interface{}{}
	var out int
	if err := sigyaml.Unmarshal(b, &m); err == nil {
		out = 1
	}
	if err := sigyaml.Unmarshal(b, &t); err == nil {
		out = 1
	}
	return out
}

// FuzzTime is a fuzz target for unmarshaling Time defined in "k8s.io/apimachinery/pkg/apis/meta/v1".
// This target also checks that the unmarshaled result can be marshaled back to the input.
func FuzzTime(b []byte) int {
	var unmarshalResult struct {
		T metav1.Time `json:"t"`
	}
	if err := yaml.Unmarshal(b, &unmarshalResult); err != nil {
		return 0
	}
	marshalResult, err := yaml.Marshal(&unmarshalResult)
	if err != nil {
		panic(err)
	}
	if !bytes.Equal(marshalResult, b) {
		panic("marshalResult != input")
	}
	return 1
}

// FuzzYamlV2 is a fuzz target for "gopkg.in/yaml.v2" unmarshaling.
func FuzzYamlV2(b []byte) int {
	t := struct{}{}
	m := map[string]interface{}{}
	var out int
	if err := yaml.Unmarshal(b, &m); err == nil {
		out = 1
	}
	if err := yaml.Unmarshal(b, &t); err == nil {
		out = 1
	}
	return out
}
