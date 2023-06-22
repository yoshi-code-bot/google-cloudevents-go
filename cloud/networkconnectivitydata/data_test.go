// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-googlecetypes. DO NOT EDIT.
// versions:
// 	protoc       				 v3.21.6
// 	protoc-gen-go 				 v1.30.0
// 	protoc-gen-go-googlecetypes  short-sha:74c44f2 (2023-06-15 15:43:30 -0700)
// source: google/events/cloud/networkconnectivity/v1/events.proto

package networkconnectivitydata_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/googleapis/google-cloudevents-go/cloud/networkconnectivitydata"
	"github.com/googleapis/google-cloudevents-go/internal/testhelper"
	"google.golang.org/protobuf/encoding/protojson"
)

// Validate the type can parse test data.
// Goals:
// - "Loose" parsing confirms the expected library experience
// - "Strict" parsing confirms:
//   - no deleted or renamed fields in protos covered in test data
//   - test data does not carry unknown fields
func TestParsingHubEventData(t *testing.T) {
	cases := testhelper.FindTestData(t, "HubEventData", "google/events/cloud/networkconnectivity/v1")

	for name, file := range cases {
		t.Run(name, func(t *testing.T) {
			data, err := os.ReadFile(file)
			if err != nil {
				t.Fatal("os.ReadFile:", err)
			}

			if ext := filepath.Ext(file); ext != ".json" {
				t.Fatalf("test support for %q data not implemented", ext)
			}

			t.Run("loose", func(t *testing.T) {
				out := networkconnectivitydata.HubEventData{}
				pj := protojson.UnmarshalOptions{DiscardUnknown: true}
				if err := pj.Unmarshal(data, &out); err != nil {
					t.Fatalf("protojson.Unmarshal: could not parse %q\n----%s\n----", file, data)
				}
			})

			t.Run("strict", func(t *testing.T) {
				out := networkconnectivitydata.HubEventData{}
				if err := protojson.Unmarshal(data, &out); err != nil {
					t.Fatalf("protojson.Unmarshal: could not parse %q\n----%s\n----", file, data)
				}
			})

		})
	}
}

// Validate the type can parse test data.
// Goals:
// - "Loose" parsing confirms the expected library experience
// - "Strict" parsing confirms:
//   - no deleted or renamed fields in protos covered in test data
//   - test data does not carry unknown fields
func TestParsingServiceClassEventData(t *testing.T) {
	cases := testhelper.FindTestData(t, "ServiceClassEventData", "google/events/cloud/networkconnectivity/v1")

	for name, file := range cases {
		t.Run(name, func(t *testing.T) {
			data, err := os.ReadFile(file)
			if err != nil {
				t.Fatal("os.ReadFile:", err)
			}

			if ext := filepath.Ext(file); ext != ".json" {
				t.Fatalf("test support for %q data not implemented", ext)
			}

			t.Run("loose", func(t *testing.T) {
				out := networkconnectivitydata.ServiceClassEventData{}
				pj := protojson.UnmarshalOptions{DiscardUnknown: true}
				if err := pj.Unmarshal(data, &out); err != nil {
					t.Fatalf("protojson.Unmarshal: could not parse %q\n----%s\n----", file, data)
				}
			})

			t.Run("strict", func(t *testing.T) {
				out := networkconnectivitydata.ServiceClassEventData{}
				if err := protojson.Unmarshal(data, &out); err != nil {
					t.Fatalf("protojson.Unmarshal: could not parse %q\n----%s\n----", file, data)
				}
			})

		})
	}
}

// Validate the type can parse test data.
// Goals:
// - "Loose" parsing confirms the expected library experience
// - "Strict" parsing confirms:
//   - no deleted or renamed fields in protos covered in test data
//   - test data does not carry unknown fields
func TestParsingServiceConnectionMapEventData(t *testing.T) {
	cases := testhelper.FindTestData(t, "ServiceConnectionMapEventData", "google/events/cloud/networkconnectivity/v1")

	for name, file := range cases {
		t.Run(name, func(t *testing.T) {
			data, err := os.ReadFile(file)
			if err != nil {
				t.Fatal("os.ReadFile:", err)
			}

			if ext := filepath.Ext(file); ext != ".json" {
				t.Fatalf("test support for %q data not implemented", ext)
			}

			t.Run("loose", func(t *testing.T) {
				out := networkconnectivitydata.ServiceConnectionMapEventData{}
				pj := protojson.UnmarshalOptions{DiscardUnknown: true}
				if err := pj.Unmarshal(data, &out); err != nil {
					t.Fatalf("protojson.Unmarshal: could not parse %q\n----%s\n----", file, data)
				}
			})

			t.Run("strict", func(t *testing.T) {
				out := networkconnectivitydata.ServiceConnectionMapEventData{}
				if err := protojson.Unmarshal(data, &out); err != nil {
					t.Fatalf("protojson.Unmarshal: could not parse %q\n----%s\n----", file, data)
				}
			})

		})
	}
}

// Validate the type can parse test data.
// Goals:
// - "Loose" parsing confirms the expected library experience
// - "Strict" parsing confirms:
//   - no deleted or renamed fields in protos covered in test data
//   - test data does not carry unknown fields
func TestParsingServiceConnectionPolicyEventData(t *testing.T) {
	cases := testhelper.FindTestData(t, "ServiceConnectionPolicyEventData", "google/events/cloud/networkconnectivity/v1")

	for name, file := range cases {
		t.Run(name, func(t *testing.T) {
			data, err := os.ReadFile(file)
			if err != nil {
				t.Fatal("os.ReadFile:", err)
			}

			if ext := filepath.Ext(file); ext != ".json" {
				t.Fatalf("test support for %q data not implemented", ext)
			}

			t.Run("loose", func(t *testing.T) {
				out := networkconnectivitydata.ServiceConnectionPolicyEventData{}
				pj := protojson.UnmarshalOptions{DiscardUnknown: true}
				if err := pj.Unmarshal(data, &out); err != nil {
					t.Fatalf("protojson.Unmarshal: could not parse %q\n----%s\n----", file, data)
				}
			})

			t.Run("strict", func(t *testing.T) {
				out := networkconnectivitydata.ServiceConnectionPolicyEventData{}
				if err := protojson.Unmarshal(data, &out); err != nil {
					t.Fatalf("protojson.Unmarshal: could not parse %q\n----%s\n----", file, data)
				}
			})

		})
	}
}

// Validate the type can parse test data.
// Goals:
// - "Loose" parsing confirms the expected library experience
// - "Strict" parsing confirms:
//   - no deleted or renamed fields in protos covered in test data
//   - test data does not carry unknown fields
func TestParsingServiceConnectionTokenEventData(t *testing.T) {
	cases := testhelper.FindTestData(t, "ServiceConnectionTokenEventData", "google/events/cloud/networkconnectivity/v1")

	for name, file := range cases {
		t.Run(name, func(t *testing.T) {
			data, err := os.ReadFile(file)
			if err != nil {
				t.Fatal("os.ReadFile:", err)
			}

			if ext := filepath.Ext(file); ext != ".json" {
				t.Fatalf("test support for %q data not implemented", ext)
			}

			t.Run("loose", func(t *testing.T) {
				out := networkconnectivitydata.ServiceConnectionTokenEventData{}
				pj := protojson.UnmarshalOptions{DiscardUnknown: true}
				if err := pj.Unmarshal(data, &out); err != nil {
					t.Fatalf("protojson.Unmarshal: could not parse %q\n----%s\n----", file, data)
				}
			})

			t.Run("strict", func(t *testing.T) {
				out := networkconnectivitydata.ServiceConnectionTokenEventData{}
				if err := protojson.Unmarshal(data, &out); err != nil {
					t.Fatalf("protojson.Unmarshal: could not parse %q\n----%s\n----", file, data)
				}
			})

		})
	}
}

// Validate the type can parse test data.
// Goals:
// - "Loose" parsing confirms the expected library experience
// - "Strict" parsing confirms:
//   - no deleted or renamed fields in protos covered in test data
//   - test data does not carry unknown fields
func TestParsingSpokeEventData(t *testing.T) {
	cases := testhelper.FindTestData(t, "SpokeEventData", "google/events/cloud/networkconnectivity/v1")

	for name, file := range cases {
		t.Run(name, func(t *testing.T) {
			data, err := os.ReadFile(file)
			if err != nil {
				t.Fatal("os.ReadFile:", err)
			}

			if ext := filepath.Ext(file); ext != ".json" {
				t.Fatalf("test support for %q data not implemented", ext)
			}

			t.Run("loose", func(t *testing.T) {
				out := networkconnectivitydata.SpokeEventData{}
				pj := protojson.UnmarshalOptions{DiscardUnknown: true}
				if err := pj.Unmarshal(data, &out); err != nil {
					t.Fatalf("protojson.Unmarshal: could not parse %q\n----%s\n----", file, data)
				}
			})

			t.Run("strict", func(t *testing.T) {
				out := networkconnectivitydata.SpokeEventData{}
				if err := protojson.Unmarshal(data, &out); err != nil {
					t.Fatalf("protojson.Unmarshal: could not parse %q\n----%s\n----", file, data)
				}
			})

		})
	}
}