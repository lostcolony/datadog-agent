// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2018 Datadog, Inc.

// +build !kubeapiserver

package custommetrics

// GetStatus returns the status info of the Custom Metrics Server.
func GetStatus() map[string]interface{} {
	status := make(map[string]interface{})
	status["Error"] = "The Custom Metrics Server is not enabled"
	return status
}
