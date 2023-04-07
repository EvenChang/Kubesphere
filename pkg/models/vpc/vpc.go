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

package vpc

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"

	vpcv1 "kubesphere.io/api/vpc/v1"

	"kubesphere.io/kubesphere/pkg/api"
	"kubesphere.io/kubesphere/pkg/apiserver/query"
	kubesphere "kubesphere.io/kubesphere/pkg/client/clientset/versioned"
	"kubesphere.io/kubesphere/pkg/informers"
	"kubesphere.io/kubesphere/pkg/models/events"
	resourcesv1alpha3 "kubesphere.io/kubesphere/pkg/models/resources/v1alpha3/resource"
	eventsclient "kubesphere.io/kubesphere/pkg/simple/client/events"
)

const orphanFinalizer = "orphan.finalizers.kubesphere.io"

type Interface interface {
	GetVpcNetwork(vpcnetwork string) (*vpcv1.VPCNetwork, error)
	ListVpcNetwork(query *query.Query) (*api.ListResult, error)
	CreateVpcNetwork(vpcnetwork *vpcv1.VPCNetwork) (*vpcv1.VPCNetwork, error)
}

type vpcOperator struct {
	ksclient       kubesphere.Interface
	resourceGetter *resourcesv1alpha3.ResourceGetter
	events         events.Interface
}

func New(informers informers.InformerFactory, ksclient kubesphere.Interface, evtsClient eventsclient.Client) Interface {
	return &vpcOperator{
		resourceGetter: resourcesv1alpha3.NewResourceGetter(informers, nil),
		ksclient:       ksclient,
		events:         events.NewEventsOperator(evtsClient),
	}
}

func (t *vpcOperator) ListVpcNetwork(queryParam *query.Query) (*api.ListResult, error) {

	result, err := t.resourceGetter.List(vpcv1.ResourcePluralVpcNetworks, "", queryParam)
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	return result, nil
}

func (t *vpcOperator) GetVpcNetwork(vpcnetwork string) (*vpcv1.VPCNetwork, error) {
	obj, err := t.resourceGetter.Get(vpcv1.ResourcePluralVpcNetworks, "", vpcnetwork)
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	return obj.(*vpcv1.VPCNetwork), nil
}

func (t *vpcOperator) CreateVpcNetwork(vpcnetwork *vpcv1.VPCNetwork) (*vpcv1.VPCNetwork, error) {
	return t.ksclient.K8sV1().VPCNetworks().Create(context.Background(), vpcnetwork, metav1.CreateOptions{})
}
