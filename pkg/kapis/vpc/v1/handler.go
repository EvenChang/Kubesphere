/*
Copyright 2020 KubeSphere Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	"github.com/emicklei/go-restful"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/klog"
	vpcv1 "kubesphere.io/api/vpc/v1"
	"kubesphere.io/kubesphere/pkg/api"
	"kubesphere.io/kubesphere/pkg/apiserver/query"
	kubesphere "kubesphere.io/kubesphere/pkg/client/clientset/versioned"
	vpclister "kubesphere.io/kubesphere/pkg/client/listers/vpc/v1"
	"kubesphere.io/kubesphere/pkg/informers"
	"kubesphere.io/kubesphere/pkg/models/vpc"
	"kubesphere.io/kubesphere/pkg/simple/client/events"
)

type handler struct {
	vpc       vpc.Interface
	vpcLister vpclister.VPCNetworkLister
}

func newHandler(factory informers.InformerFactory, ksclient kubesphere.Interface, evtsClient events.Client) *handler {
	return &handler{
		vpc: vpc.New(factory, ksclient, evtsClient),
	}
}

func (h *handler) getVpcNetwork(request *restful.Request, response *restful.Response) {

	vpcnetworks, err := h.vpc.ListVpcNetwork(query.New())

	if err != nil {
		if errors.IsNotFound(err) {
			api.HandleNotFound(response, request, err)
			return
		} else {
			api.HandleInternalError(response, request, err)
			return
		}
	}

	response.WriteAsJson(vpcnetworks)
}

func (h *handler) createVpcNetwork(request *restful.Request, response *restful.Response) {

	var vpcnetwork vpcv1.VPCNetwork

	err := request.ReadEntity(&vpcnetwork)

	if err != nil {
		klog.Error(err)
		api.HandleBadRequest(response, request, err)
		return
	}

	created, err := h.vpc.CreateVpcNetwork(&vpcnetwork)

	if err != nil {
		klog.Error(err)
		if errors.IsNotFound(err) {
			api.HandleNotFound(response, request, err)
			return
		}
		if errors.IsForbidden(err) {
			api.HandleForbidden(response, request, err)
			return
		}
		api.HandleBadRequest(response, request, err)
		return
	}

	response.WriteEntity(created)
}

type vpcResponse struct {
	CIDR         string `json:"cidr"`
	SubnetLength int    `json:"subnetLength"`
}
