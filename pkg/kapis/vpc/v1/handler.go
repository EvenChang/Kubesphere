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
	"kubesphere.io/kubesphere/pkg/api"
	"kubesphere.io/kubesphere/pkg/apiserver/query"
	"kubesphere.io/kubesphere/pkg/client/informers/externalversions"
	vpclister "kubesphere.io/kubesphere/pkg/client/listers/vpc/v1"
)

type handler struct {
	vpcLister vpclister.VPCNetworkLister
}

func newHandler(ksInformers externalversions.SharedInformerFactory) *handler {
	return &handler{
		vpcLister: ksInformers.K8s().V1().VPCNetworks().Lister(),
	}
}

func (h *handler) getVpcNetwork(request *restful.Request, response *restful.Response) {

	vpcnetworks, err := h.vpcLister.List(query.New().Selector())

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

type vpcResponse struct {
	CIDR         string `json:"cidr"`
	SubnetLength int    `json:"subnetLength"`
}
