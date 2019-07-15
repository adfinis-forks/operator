/*
Copyright 2019 The Kube Vault Authors.

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

package v1alpha1

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubevault.dev/operator/apis/engine/v1alpha1"
	scheme "kubevault.dev/operator/client/clientset/versioned/scheme"
)

// GCPAccessKeyRequestsGetter has a method to return a GCPAccessKeyRequestInterface.
// A group's client should implement this interface.
type GCPAccessKeyRequestsGetter interface {
	GCPAccessKeyRequests(namespace string) GCPAccessKeyRequestInterface
}

// GCPAccessKeyRequestInterface has methods to work with GCPAccessKeyRequest resources.
type GCPAccessKeyRequestInterface interface {
	Create(*v1alpha1.GCPAccessKeyRequest) (*v1alpha1.GCPAccessKeyRequest, error)
	Update(*v1alpha1.GCPAccessKeyRequest) (*v1alpha1.GCPAccessKeyRequest, error)
	UpdateStatus(*v1alpha1.GCPAccessKeyRequest) (*v1alpha1.GCPAccessKeyRequest, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GCPAccessKeyRequest, error)
	List(opts v1.ListOptions) (*v1alpha1.GCPAccessKeyRequestList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GCPAccessKeyRequest, err error)
	GCPAccessKeyRequestExpansion
}

// gCPAccessKeyRequests implements GCPAccessKeyRequestInterface
type gCPAccessKeyRequests struct {
	client rest.Interface
	ns     string
}

// newGCPAccessKeyRequests returns a GCPAccessKeyRequests
func newGCPAccessKeyRequests(c *EngineV1alpha1Client, namespace string) *gCPAccessKeyRequests {
	return &gCPAccessKeyRequests{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the gCPAccessKeyRequest, and returns the corresponding gCPAccessKeyRequest object, and an error if there is any.
func (c *gCPAccessKeyRequests) Get(name string, options v1.GetOptions) (result *v1alpha1.GCPAccessKeyRequest, err error) {
	result = &v1alpha1.GCPAccessKeyRequest{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gcpaccesskeyrequests").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GCPAccessKeyRequests that match those selectors.
func (c *gCPAccessKeyRequests) List(opts v1.ListOptions) (result *v1alpha1.GCPAccessKeyRequestList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GCPAccessKeyRequestList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gcpaccesskeyrequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested gCPAccessKeyRequests.
func (c *gCPAccessKeyRequests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("gcpaccesskeyrequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a gCPAccessKeyRequest and creates it.  Returns the server's representation of the gCPAccessKeyRequest, and an error, if there is any.
func (c *gCPAccessKeyRequests) Create(gCPAccessKeyRequest *v1alpha1.GCPAccessKeyRequest) (result *v1alpha1.GCPAccessKeyRequest, err error) {
	result = &v1alpha1.GCPAccessKeyRequest{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("gcpaccesskeyrequests").
		Body(gCPAccessKeyRequest).
		Do().
		Into(result)
	return
}

// Update takes the representation of a gCPAccessKeyRequest and updates it. Returns the server's representation of the gCPAccessKeyRequest, and an error, if there is any.
func (c *gCPAccessKeyRequests) Update(gCPAccessKeyRequest *v1alpha1.GCPAccessKeyRequest) (result *v1alpha1.GCPAccessKeyRequest, err error) {
	result = &v1alpha1.GCPAccessKeyRequest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gcpaccesskeyrequests").
		Name(gCPAccessKeyRequest.Name).
		Body(gCPAccessKeyRequest).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *gCPAccessKeyRequests) UpdateStatus(gCPAccessKeyRequest *v1alpha1.GCPAccessKeyRequest) (result *v1alpha1.GCPAccessKeyRequest, err error) {
	result = &v1alpha1.GCPAccessKeyRequest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gcpaccesskeyrequests").
		Name(gCPAccessKeyRequest.Name).
		SubResource("status").
		Body(gCPAccessKeyRequest).
		Do().
		Into(result)
	return
}

// Delete takes name of the gCPAccessKeyRequest and deletes it. Returns an error if one occurs.
func (c *gCPAccessKeyRequests) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gcpaccesskeyrequests").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *gCPAccessKeyRequests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gcpaccesskeyrequests").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched gCPAccessKeyRequest.
func (c *gCPAccessKeyRequests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GCPAccessKeyRequest, err error) {
	result = &v1alpha1.GCPAccessKeyRequest{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("gcpaccesskeyrequests").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
