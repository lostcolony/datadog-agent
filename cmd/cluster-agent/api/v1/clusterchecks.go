// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2018 Datadog, Inc.

// +build clusterchecks

package v1

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/DataDog/datadog-agent/cmd/cluster-agent/api/types"
	"github.com/DataDog/datadog-agent/pkg/clusteragent/clusterchecks"
)

var clusterCheckHandler *clusterchecks.Handler

// Install registers v1 API endpoints
func installClusterCheckEndpoints(r *mux.Router, sc types.ServerContext) {
	if sc.ClusterCheckHandler == nil {
		return
	}
	clusterCheckHandler = sc.ClusterCheckHandler

	// TODO
	//r.HandleFunc("/clusterchecks/status/{nodeName}", postCheckStatus).Methods("POST")
	//r.HandleFunc("/clusterchecks/configs/{nodeName}", getCheckConfigs).Methods("GET")
	r.HandleFunc("/clusterchecks", getAllCheckConfigs).Methods("GET")
}

// getAllCheckConfigs is used by the clustercheck config
func getAllCheckConfigs(w http.ResponseWriter, r *http.Request) {
	if clusterCheckHandler == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	configs, err := clusterCheckHandler.GetAllConfigs()

	slcB, err := json.Marshal(configs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(slcB) != 0 {
		w.WriteHeader(http.StatusOK)
		w.Write(slcB)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}
