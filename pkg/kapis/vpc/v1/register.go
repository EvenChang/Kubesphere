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
	restfulspec "github.com/emicklei/go-restful-openapi"
	"k8s.io/apimachinery/pkg/runtime/schema"
	v1 "kubesphere.io/api/vpc/v1"
	vpcv1 "kubesphere.io/api/vpc/v1"
	"kubesphere.io/kubesphere/pkg/api"
	"kubesphere.io/kubesphere/pkg/apiserver/runtime"
	kubesphere "kubesphere.io/kubesphere/pkg/client/clientset/versioned"
	"kubesphere.io/kubesphere/pkg/constants"
	"kubesphere.io/kubesphere/pkg/informers"
	"kubesphere.io/kubesphere/pkg/server/errors"
)

const (
	GroupName = "k8s.ovn.org"
)

var GroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1"}

func Resource(resource string) schema.GroupResource {
	return GroupVersion.WithResource(resource).GroupResource()
}

func AddToContainer(container *restful.Container, factory informers.InformerFactory, ksclient kubesphere.Interface) error {
	webservice := runtime.NewWebService(GroupVersion)
	handler := newHandler(factory, ksclient)

	webservice.Route(webservice.GET("/vpcnetworks").
		To(handler.listVpcNetwork).
		Doc("List all vpcnetowkrs resources").
		Returns(http.StatusOK, api.StatusOK, api.ListResult{}).
		Metadata(restfulspec.KeyOpenAPITags, []string{constants.VpcNetworkTag}))

	webservice.Route(webservice.GET("/vpcnetwork/{vpcnetwork}").
		To(handler.getVpcNetwork).
		Doc("Get vpcnetowkrs resources").
		Returns(http.StatusOK, api.StatusOK, api.ListResult{}).
		Metadata(restfulspec.KeyOpenAPITags, []string{constants.VpcNetworkTag}))

	webservice.Route(webservice.POST("/vpcnetwork").
		To(handler.createVpcNetwork).
		Reads(vpcv1.VPCNetwork{}).
		Doc("Create vpcnetwork").
		Returns(http.StatusOK, api.StatusOK, v1.VPCNetwork{}).
		Metadata(restfulspec.KeyOpenAPITags, []string{constants.VpcNetworkTag}))

	webservice.Route(webservice.DELETE("/vpcnetwork/{vpcnetwork}").
		To(handler.deleteVpcNetwork).
		Param(webservice.PathParameter("vpcnetwork", "vpcnetwork name")).
		Doc("Delete vpcnetwork").
		Returns(http.StatusOK, api.StatusOK, errors.None).
		Metadata(restfulspec.KeyOpenAPITags, []string{constants.VpcNetworkTag}))

	container.Add(webservice)

	return nil
}
