// Copyright (c) 2017 Intel Corporation
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

package main

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	logLevel := flag.String("log", "warn",
		"log messages above specified level; one of debug, warn, error, fatal or panic")

	flag.Parse()

	if err := SetLoggingLevel(*logLevel); err != nil {
		fmt.Fprint(os.Stderr, err)
	}

	if err := ksmTestPrepare(); err != nil {
		ksmTestCleanup()
		fmt.Fprint(os.Stderr, err)
	}

	exit := m.Run()

	ksmTestCleanup()

	os.Exit(exit)
}
