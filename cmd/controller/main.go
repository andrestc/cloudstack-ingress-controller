// Copyright 2017 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/tsuru/cloudstack-ingress-controller/controller"
	"k8s.io/apiserver/pkg/server/healthz"
)

func main() {
	dc := &controller.CloudstackController{
		HealthzChecker: healthz.PingHealthz,
	}
	dc.Start()
}
