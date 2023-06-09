/*
Copyright 2020 The KubeSphere Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1 "kubesphere.io/api/vpc/v1"
	scheme "kubesphere.io/kubesphere/pkg/client/clientset/versioned/scheme"
)

// VPCNetworksGetter has a method to return a VPCNetworkInterface.
// A group's client should implement this interface.
type VPCNetworksGetter interface {
	VPCNetworks() VPCNetworkInterface
}

// VPCNetworkInterface has methods to work with VPCNetwork resources.
type VPCNetworkInterface interface {
	Create(ctx context.Context, vPCNetwork *v1.VPCNetwork, opts metav1.CreateOptions) (*v1.VPCNetwork, error)
	Update(ctx context.Context, vPCNetwork *v1.VPCNetwork, opts metav1.UpdateOptions) (*v1.VPCNetwork, error)
	UpdateStatus(ctx context.Context, vPCNetwork *v1.VPCNetwork, opts metav1.UpdateOptions) (*v1.VPCNetwork, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.VPCNetwork, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.VPCNetworkList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.VPCNetwork, err error)
	VPCNetworkExpansion
}

// vPCNetworks implements VPCNetworkInterface
type vPCNetworks struct {
	client rest.Interface
}

// newVPCNetworks returns a VPCNetworks
func newVPCNetworks(c *K8sV1Client) *vPCNetworks {
	return &vPCNetworks{
		client: c.RESTClient(),
	}
}

// Get takes name of the vPCNetwork, and returns the corresponding vPCNetwork object, and an error if there is any.
func (c *vPCNetworks) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.VPCNetwork, err error) {
	result = &v1.VPCNetwork{}
	err = c.client.Get().
		Resource("vpcnetworks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VPCNetworks that match those selectors.
func (c *vPCNetworks) List(ctx context.Context, opts metav1.ListOptions) (result *v1.VPCNetworkList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.VPCNetworkList{}
	err = c.client.Get().
		Resource("vpcnetworks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested vPCNetworks.
func (c *vPCNetworks) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("vpcnetworks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a vPCNetwork and creates it.  Returns the server's representation of the vPCNetwork, and an error, if there is any.
func (c *vPCNetworks) Create(ctx context.Context, vPCNetwork *v1.VPCNetwork, opts metav1.CreateOptions) (result *v1.VPCNetwork, err error) {
	result = &v1.VPCNetwork{}
	err = c.client.Post().
		Resource("vpcnetworks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vPCNetwork).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a vPCNetwork and updates it. Returns the server's representation of the vPCNetwork, and an error, if there is any.
func (c *vPCNetworks) Update(ctx context.Context, vPCNetwork *v1.VPCNetwork, opts metav1.UpdateOptions) (result *v1.VPCNetwork, err error) {
	result = &v1.VPCNetwork{}
	err = c.client.Put().
		Resource("vpcnetworks").
		Name(vPCNetwork.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vPCNetwork).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *vPCNetworks) UpdateStatus(ctx context.Context, vPCNetwork *v1.VPCNetwork, opts metav1.UpdateOptions) (result *v1.VPCNetwork, err error) {
	result = &v1.VPCNetwork{}
	err = c.client.Put().
		Resource("vpcnetworks").
		Name(vPCNetwork.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vPCNetwork).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the vPCNetwork and deletes it. Returns an error if one occurs.
func (c *vPCNetworks) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("vpcnetworks").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *vPCNetworks) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("vpcnetworks").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched vPCNetwork.
func (c *vPCNetworks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.VPCNetwork, err error) {
	result = &v1.VPCNetwork{}
	err = c.client.Patch(pt).
		Resource("vpcnetworks").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
