// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	hivev1 "github.com/openshift/hive/pkg/apis/hive/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterImageSets implements ClusterImageSetInterface
type FakeClusterImageSets struct {
	Fake *FakeHiveV1
}

var clusterimagesetsResource = schema.GroupVersionResource{Group: "hive.openshift.io", Version: "v1", Resource: "clusterimagesets"}

var clusterimagesetsKind = schema.GroupVersionKind{Group: "hive.openshift.io", Version: "v1", Kind: "ClusterImageSet"}

// Get takes name of the clusterImageSet, and returns the corresponding clusterImageSet object, and an error if there is any.
func (c *FakeClusterImageSets) Get(ctx context.Context, name string, options v1.GetOptions) (result *hivev1.ClusterImageSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterimagesetsResource, name), &hivev1.ClusterImageSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*hivev1.ClusterImageSet), err
}

// List takes label and field selectors, and returns the list of ClusterImageSets that match those selectors.
func (c *FakeClusterImageSets) List(ctx context.Context, opts v1.ListOptions) (result *hivev1.ClusterImageSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterimagesetsResource, clusterimagesetsKind, opts), &hivev1.ClusterImageSetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &hivev1.ClusterImageSetList{ListMeta: obj.(*hivev1.ClusterImageSetList).ListMeta}
	for _, item := range obj.(*hivev1.ClusterImageSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterImageSets.
func (c *FakeClusterImageSets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterimagesetsResource, opts))
}

// Create takes the representation of a clusterImageSet and creates it.  Returns the server's representation of the clusterImageSet, and an error, if there is any.
func (c *FakeClusterImageSets) Create(ctx context.Context, clusterImageSet *hivev1.ClusterImageSet, opts v1.CreateOptions) (result *hivev1.ClusterImageSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterimagesetsResource, clusterImageSet), &hivev1.ClusterImageSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*hivev1.ClusterImageSet), err
}

// Update takes the representation of a clusterImageSet and updates it. Returns the server's representation of the clusterImageSet, and an error, if there is any.
func (c *FakeClusterImageSets) Update(ctx context.Context, clusterImageSet *hivev1.ClusterImageSet, opts v1.UpdateOptions) (result *hivev1.ClusterImageSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterimagesetsResource, clusterImageSet), &hivev1.ClusterImageSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*hivev1.ClusterImageSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterImageSets) UpdateStatus(ctx context.Context, clusterImageSet *hivev1.ClusterImageSet, opts v1.UpdateOptions) (*hivev1.ClusterImageSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusterimagesetsResource, "status", clusterImageSet), &hivev1.ClusterImageSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*hivev1.ClusterImageSet), err
}

// Delete takes name of the clusterImageSet and deletes it. Returns an error if one occurs.
func (c *FakeClusterImageSets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(clusterimagesetsResource, name), &hivev1.ClusterImageSet{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterImageSets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterimagesetsResource, listOpts)

	_, err := c.Fake.Invokes(action, &hivev1.ClusterImageSetList{})
	return err
}

// Patch applies the patch and returns the patched clusterImageSet.
func (c *FakeClusterImageSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *hivev1.ClusterImageSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterimagesetsResource, name, pt, data, subresources...), &hivev1.ClusterImageSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*hivev1.ClusterImageSet), err
}
