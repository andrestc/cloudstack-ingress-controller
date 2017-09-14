// Copyright 2017 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package controller

import (
	"log"
	"strings"

	"k8s.io/apiserver/pkg/server/healthz"

	"github.com/spf13/pflag"
	api "k8s.io/api/core/v1"
	extensions "k8s.io/api/extensions/v1beta1"
	"k8s.io/ingress/core/pkg/ingress"
	k8sctl "k8s.io/ingress/core/pkg/ingress/controller"
	"k8s.io/ingress/core/pkg/ingress/defaults"
	"k8s.io/ingress/core/pkg/ingress/store"
)

var (
	defIngressClass = "cloudstack"
)

type CloudstackController struct {
	healthz.HealthzChecker

	nodeLister store.NodeLister
	svcLister  store.ServiceLister
}

func (c *CloudstackController) Start() {
	ic := k8sctl.NewIngressController(c)
	defer func() {
		log.Printf("Shutting down ingress controller...")
		ic.Stop()
	}()
	ic.Start()
}

func (c *CloudstackController) SetConfig(cfgMap *api.ConfigMap) {
	log.Printf("SetConfig: %+v", cfgMap)
}

func (c *CloudstackController) OnUpdate(updatePayload ingress.Configuration) error {
	var services []string
	for _, b := range updatePayload.Backends {
		services = append(services, b.Name)
	}
	log.Printf("[OnUpdate] Services: %s", strings.Join(services, ", "))
	return nil
}

func (c *CloudstackController) BackendDefaults() defaults.Backend {
	return defaults.Backend{}
}

func (c *CloudstackController) Name() string {
	return "Cloudstack Controller"
}

func (c *CloudstackController) Info() *ingress.BackendInfo {
	return &ingress.BackendInfo{
		Name:       "cloudstack",
		Release:    "0.0.1",
		Repository: "git://github.com/tsuru/cloudstack-ingress-controller",
	}
}

func (c *CloudstackController) ConfigureFlags(p *pflag.FlagSet) {
}

func (c *CloudstackController) OverrideFlags(flags *pflag.FlagSet) {
	ic, _ := flags.GetString("ingress-class")

	if ic == "" {
		ic = defIngressClass
	}

	flags.Set("ingress-class", ic)
}

func (c *CloudstackController) SetListers(lister ingress.StoreLister) {
	c.nodeLister = lister.Node
	c.svcLister = lister.Service
}

func (c *CloudstackController) DefaultIngressClass() string {
	return "cloudstack"
}

// UpdateIngressStatus updates IP of the ingress based on the Cloudstack LoadBalancer
func (c *CloudstackController) UpdateIngressStatus(ing *extensions.Ingress) []api.LoadBalancerIngress {
	// get the ingress cloudstack VIP and add its IP to loadBalancerIngress

	// data, err := json.Marshal(ing)
	// if err != nil {
	// 	log.Printf("err json UpdateIngressStatus %v", err)
	// }
	// log.Printf("UpdateIngressStatus: %s", data)
	return nil
}
