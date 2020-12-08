// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/openshift/hive/pkg/apis/hive/v1"
	scheme "github.com/openshift/hive/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SelectorSyncSetsGetter has a method to return a SelectorSyncSetInterface.
// A group's client should implement this interface.
type SelectorSyncSetsGetter interface {
	SelectorSyncSets() SelectorSyncSetInterface
}

// SelectorSyncSetInterface has methods to work with SelectorSyncSet resources.
type SelectorSyncSetInterface interface {
	Create(ctx context.Context, selectorSyncSet *v1.SelectorSyncSet, opts metav1.CreateOptions) (*v1.SelectorSyncSet, error)
	Update(ctx context.Context, selectorSyncSet *v1.SelectorSyncSet, opts metav1.UpdateOptions) (*v1.SelectorSyncSet, error)
	UpdateStatus(ctx context.Context, selectorSyncSet *v1.SelectorSyncSet, opts metav1.UpdateOptions) (*v1.SelectorSyncSet, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.SelectorSyncSet, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.SelectorSyncSetList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.SelectorSyncSet, err error)
	SelectorSyncSetExpansion
}

// selectorSyncSets implements SelectorSyncSetInterface
type selectorSyncSets struct {
	client rest.Interface
}

// newSelectorSyncSets returns a SelectorSyncSets
func newSelectorSyncSets(c *HiveV1Client) *selectorSyncSets {
	return &selectorSyncSets{
		client: c.RESTClient(),
	}
}

// Get takes name of the selectorSyncSet, and returns the corresponding selectorSyncSet object, and an error if there is any.
func (c *selectorSyncSets) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.SelectorSyncSet, err error) {
	result = &v1.SelectorSyncSet{}
	err = c.client.Get().
		Resource("selectorsyncsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SelectorSyncSets that match those selectors.
func (c *selectorSyncSets) List(ctx context.Context, opts metav1.ListOptions) (result *v1.SelectorSyncSetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.SelectorSyncSetList{}
	err = c.client.Get().
		Resource("selectorsyncsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested selectorSyncSets.
func (c *selectorSyncSets) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("selectorsyncsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a selectorSyncSet and creates it.  Returns the server's representation of the selectorSyncSet, and an error, if there is any.
func (c *selectorSyncSets) Create(ctx context.Context, selectorSyncSet *v1.SelectorSyncSet, opts metav1.CreateOptions) (result *v1.SelectorSyncSet, err error) {
	result = &v1.SelectorSyncSet{}
	err = c.client.Post().
		Resource("selectorsyncsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(selectorSyncSet).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a selectorSyncSet and updates it. Returns the server's representation of the selectorSyncSet, and an error, if there is any.
func (c *selectorSyncSets) Update(ctx context.Context, selectorSyncSet *v1.SelectorSyncSet, opts metav1.UpdateOptions) (result *v1.SelectorSyncSet, err error) {
	result = &v1.SelectorSyncSet{}
	err = c.client.Put().
		Resource("selectorsyncsets").
		Name(selectorSyncSet.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(selectorSyncSet).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *selectorSyncSets) UpdateStatus(ctx context.Context, selectorSyncSet *v1.SelectorSyncSet, opts metav1.UpdateOptions) (result *v1.SelectorSyncSet, err error) {
	result = &v1.SelectorSyncSet{}
	err = c.client.Put().
		Resource("selectorsyncsets").
		Name(selectorSyncSet.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(selectorSyncSet).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the selectorSyncSet and deletes it. Returns an error if one occurs.
func (c *selectorSyncSets) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("selectorsyncsets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *selectorSyncSets) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("selectorsyncsets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched selectorSyncSet.
func (c *selectorSyncSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.SelectorSyncSet, err error) {
	result = &v1.SelectorSyncSet{}
	err = c.client.Patch(pt).
		Resource("selectorsyncsets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
