// Copyright 2017 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"net/http"

	"github.com/pingcap/pd/server"
	"github.com/unrolled/render"
)

type hotStatusHandler struct {
	*server.Handler
	rd *render.Render
}

func newHotStatusHandler(handler *server.Handler, rd *render.Render) *hotStatusHandler {
	return &hotStatusHandler{
		Handler: handler,
		rd:      rd,
	}
}

func (h *hotStatusHandler) GetHotRegions(w http.ResponseWriter, r *http.Request) {
	h.rd.JSON(w, http.StatusOK, h.GetHotWriteRegions())
}

func (h *hotStatusHandler) GetHotStores(w http.ResponseWriter, r *http.Request) {
	h.rd.JSON(w, http.StatusOK, h.GetHotWriteStores())
}
