// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ttnmage

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// EndToEnd namespace
type EndToEnd mg.Namespace

var (
	testDatabaseName = "ttn_lorawan_test"
	databaseURI      = "postgresql://root@localhost:26257/" + testDatabaseName + "?sslmode=disable"
)

func getTestURL() string {
	nodeEnv := os.Getenv("NODE_ENV")
	port := "1885"
	if nodeEnv == "development" {
		port = "8080"
	}
	return "http://localhost:" + port
}

func (endToEnd EndToEnd) execFromNodeBin(cmd string, verbose bool, appendArgs ...string) (func(args ...string) error, error) {
	if _, err := os.Stat(nodeBin(cmd)); os.IsNotExist(err) {
		mg.Deps(Js.DevDeps)
	}

	out := os.Stdout
	if !mg.Verbose() && !verbose {
		out = nil
	}
	return func(args ...string) error {
		_, err := sh.Exec(nil, out, os.Stderr, nodeBin(cmd), append(args, appendArgs...)...)
		return err
	}, nil
}

func (endToEnd EndToEnd) cypress() (func(args ...string) error, error) {
	return endToEnd.execFromNodeBin("cypress", true, "--config", getTestURL())
}

// Prepare prepares the server for running end to end tests.
func (endToEnd EndToEnd) Prepare() {
	os.Setenv("TTN_LW_IS_DATABASE_URI", databaseURI)
	mg.Deps(Js.BuildIfNecessary)
	mg.SerialDeps(Dev.DBErase, Dev.InitStack, EndToEnd.DBDump)
}

// DBDump performs a database dump to the .cache folder.
func (EndToEnd) DBDump() error {
	if mg.Verbose() {
		fmt.Printf("Execute database dump\n")
	}
	return execDockerComposeWithOutput(filepath.Join(".cache", "sqldump.sql"), "exec", "cockroach", "./cockroach", "dump", testDatabaseName, "--insecure")
}

// DBRestore restores the database using a previously generated dump.
func (EndToEnd) DBRestore() error {
	if mg.Verbose() {
		fmt.Printf("Restore database from dump")
	}
	n, err := node()
	if err != nil {
		return err
	}
	return n("./.mage/restore-db-dump.js", "--db", testDatabaseName)
}

// Cypress runs the Cypress end-to-end tests.
func (endToEnd EndToEnd) Cypress() error {
	if mg.Verbose() {
		fmt.Println("Running Cypress E2E Tests")
	}
	mg.Deps(endToEnd.WaitUntilReady)
	cypress, err := endToEnd.cypress()
	if err != nil {
		return err
	}
	return cypress("run", "--config-file", "./config/cypress.json")
}

// CypressOpen runs the Cypress end-to-end tests.
func (endToEnd EndToEnd) CypressOpen() error {
	if mg.Verbose() {
		fmt.Println("Running Cypress in interactive mode")
	}
	cypress, err := endToEnd.cypress()
	if err != nil {
		return err
	}

	return cypress("open", "--config-file", "./config/cypress.json")
}

// StartTestStack starts TTS in end-to-end test configuration.
func (endToEnd EndToEnd) StartTestStack() error {
	if mg.Verbose() {
		fmt.Println("Starting stack in end-to-end test configuration")
	}
	os.Setenv("TTN_LW_IS_DATABASE_URI", databaseURI)
	mg.Deps(Js.BuildIfNecessary)
	execGo("run", "./cmd/ttn-lw-stack", "start")
	return nil
}

// WaitUntilReady waits until the web endpoints become available.
func (endToEnd EndToEnd) WaitUntilReady() error {
	if mg.Verbose() {
		fmt.Println("Waiting for the stack to be ready...")
	}
	for i := 0; i < 100; i++ {
		resp, _ := http.Get(getTestURL() + "/oauth")
		if resp != nil && resp.StatusCode == 200 {
			return nil
		}
		time.Sleep(1000 * time.Millisecond)
	}
	return errors.New("Could not connect to server")
}

// RunFrontend starts the frontend end-to-end tests from scratch. Meant for CI
// usage only.
func (endToEnd EndToEnd) RunFrontend() {
	if mg.Verbose() {
		fmt.Println("Running end-to-end frontend based tests")
	}
	mg.Deps(EndToEnd.Cypress)
}

// Run starts the end-to-end tests from scratch.
func (endToEnd EndToEnd) Run() {
	if mg.Verbose() {
		fmt.Println("Running end-to-end tests")
	}
	mg.Deps(EndToEnd.RunFrontend)
}
