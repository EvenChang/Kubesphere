/*
Copyright 2019 The KubeSphere Authors.

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
	"net/http"

	"github.com/emicklei/go-restful"
	"k8s.io/apimachinery/pkg/runtime/schema"
	v1 "kubesphere.io/api/vpc/v1"
	"kubesphere.io/kubesphere/pkg/api"
	"kubesphere.io/kubesphere/pkg/apiserver/runtime"
	"kubesphere.io/kubesphere/pkg/client/informers/externalversions"
)

const (
	GroupName = "k8s.ovn.org"
)

var GroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1"}

func Resource(resource string) schema.GroupResource {
	return GroupVersion.WithResource(resource).GroupResource()
}

func AddToContainer(container *restful.Container, ksInformers externalversions.SharedInformerFactory) error {
	webservice := runtime.NewWebService(GroupVersion)
	handler := newHandler(ksInformers)

	webservice.Route(webservice.GET("/vpc").
		Reads("").
		To(handler.getVpcNetwork).
		Returns(http.StatusOK, api.StatusOK, v1.VPCNetworkSpec{})).
		Doc("Api for vpcnetowkrs")

	container.Add(webservice)

	return nil
}
