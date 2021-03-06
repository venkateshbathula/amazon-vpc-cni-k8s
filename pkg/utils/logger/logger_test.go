// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package logger

import (
	"os"
	"testing"

	log "github.com/cihub/seelog"
	"github.com/stretchr/testify/assert"
)

func TestGetLogFileLocationReturnsOverriddenPath(t *testing.T) {
	path := "/tmp/foo"
	_ = os.Setenv(envLogFilePath, path)
	defer os.Unsetenv(envLogFilePath)

	assert.Equal(t, path, GetLogFileLocation("/tmp/bar"))
}

func TestGetLogFileLocationReturnsDefaultPath(t *testing.T) {
	path := "/tmp/foo"
	assert.Equal(t, path, GetLogFileLocation(path))
}

func TestLogLevelReturnsOverriddenLevel(t *testing.T) {
	_ = os.Setenv(envLogLevel, "INFO")
	defer os.Unsetenv(envLogLevel)

	var expectedLogLevel log.LogLevel
	expectedLogLevel = log.InfoLvl
	assert.Equal(t, expectedLogLevel.String(), getLogLevel())
}

func TestLogLevelReturnsDefaultLevelWhenEnvNotSet(t *testing.T) {
	var expectedLogLevel log.LogLevel
	expectedLogLevel = log.DebugLvl
	assert.Equal(t, expectedLogLevel.String(), getLogLevel())
}

func TestLogLevelReturnsDefaultLevelWhenEnvSetToInvalidValue(t *testing.T) {
	_ = os.Setenv(envLogLevel, "EVERYTHING")
	defer os.Unsetenv(envLogLevel)

	var expectedLogLevel log.LogLevel
	expectedLogLevel = log.DebugLvl
	assert.Equal(t, expectedLogLevel.String(), getLogLevel())
}

func TestLogOutputReturnsFileWhenValueNotStdout(t *testing.T) {
	path := "/tmp/foo"

	var expectedOutput = `<rollingfile filename="/tmp/foo" type="date" datepattern="2006-01-02-15" archivetype="none" maxrolls="24" />`
	assert.Equal(t, expectedOutput, getLogOutput(path))
}

func TestLogOutputReturnsConsole(t *testing.T) {
	path := "stdout"

	var expectedOutput = `<console />`
	assert.Equal(t, expectedOutput, getLogOutput(path))
}
