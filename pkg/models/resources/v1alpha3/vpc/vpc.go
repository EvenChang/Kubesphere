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
	"k8s.io/apimachinery/pkg/runtime"

	vpcv1 "kubesphere.io/api/vpc/v1"

	"kubesphere.io/kubesphere/pkg/api"
	"kubesphere.io/kubesphere/pkg/apiserver/query"
	informers "kubesphere.io/kubesphere/pkg/client/informers/externalversions"
	"kubesphere.io/kubesphere/pkg/models/resources/v1alpha3"
)

type vpcGetter struct {
	sharedInformers informers.SharedInformerFactory
}

func New(sharedInformers informers.SharedInformerFactory) v1alpha3.Interface {
	return &vpcGetter{sharedInformers: sharedInformers}
}

func (d *vpcGetter) Get(_, name string) (runtime.Object, error) {
	return d.sharedInformers.K8s().V1().VPCNetworks().Lister().Get(name)
}

func (d *vpcGetter) List(_ string, query *query.Query) (*api.ListResult, error) {

	vpcnetworks, err := d.sharedInformers.K8s().V1().VPCNetworks().Lister().List(query.Selector())
	if err != nil {
		return nil, err
	}

	var result []runtime.Object
	for _, vpcnetwork := range vpcnetworks {
		result = append(result, vpcnetwork)
	}

	return v1alpha3.DefaultList(result, query, d.compare, d.filter), nil
}

func (d *vpcGetter) compare(left runtime.Object, right runtime.Object, field query.Field) bool {

	leftVpcnetwork, ok := left.(*vpcv1.VPCNetwork)
	if !ok {
		return false
	}

	rightVpcnetwork, ok := right.(*vpcv1.VPCNetwork)
	if !ok {
		return false
	}

	return v1alpha3.DefaultObjectMetaCompare(leftVpcnetwork.ObjectMeta, rightVpcnetwork.ObjectMeta, field)
}

func (d *vpcGetter) filter(object runtime.Object, filter query.Filter) bool {
	role, ok := object.(*vpcv1.VPCNetwork)

	if !ok {
		return false
	}

	return v1alpha3.DefaultObjectMetaFilter(role.ObjectMeta, filter)
}
