// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build system
// +build system

package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"testing"

	"golang.org/x/oauth2"
)

func TestSystem(t *testing.T) {
	ctx := context.Background()
	baseURL := os.Getenv("BASE_URL")
	idToken := os.Getenv("ID_TOKEN")
	if baseURL == "" || idToken == "" {
		t.Fatal("both BASE_URL and ID_TOKEN must be set")
	}
	tok := oauth2.Token{
		AccessToken: idToken,
	}
	ts := oauth2.StaticTokenSource(&tok)
	client := oauth2.NewClient(ctx, ts)
	resp, err := client.Get(baseURL)
	if err != nil {
		t.Fatalf("unable to complete request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("resp.StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	gotBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("unable to read response body: %v", err)
	}
	wantBody := "Hello World!\n"
	if string(gotBody) != wantBody {
		t.Fatalf("resp.Body = %q, want %q", gotBody, wantBody)
	}
}
