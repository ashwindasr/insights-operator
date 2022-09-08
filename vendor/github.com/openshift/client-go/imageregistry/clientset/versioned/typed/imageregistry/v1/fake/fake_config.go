// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	imageregistryv1 "github.com/openshift/api/imageregistry/v1"
	applyconfigurationsimageregistryv1 "github.com/openshift/client-go/imageregistry/applyconfigurations/imageregistry/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeConfigs implements ConfigInterface
type FakeConfigs struct {
	Fake *FakeImageregistryV1
}

var configsResource = schema.GroupVersionResource{Group: "imageregistry.operator.openshift.io", Version: "v1", Resource: "configs"}

var configsKind = schema.GroupVersionKind{Group: "imageregistry.operator.openshift.io", Version: "v1", Kind: "Config"}

// Get takes name of the config, and returns the corresponding config object, and an error if there is any.
func (c *FakeConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *imageregistryv1.Config, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(configsResource, name), &imageregistryv1.Config{})
	if obj == nil {
		return nil, err
	}
	return obj.(*imageregistryv1.Config), err
}

// List takes label and field selectors, and returns the list of Configs that match those selectors.
func (c *FakeConfigs) List(ctx context.Context, opts v1.ListOptions) (result *imageregistryv1.ConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(configsResource, configsKind, opts), &imageregistryv1.ConfigList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &imageregistryv1.ConfigList{ListMeta: obj.(*imageregistryv1.ConfigList).ListMeta}
	for _, item := range obj.(*imageregistryv1.ConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested configs.
func (c *FakeConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(configsResource, opts))
}

// Create takes the representation of a config and creates it.  Returns the server's representation of the config, and an error, if there is any.
func (c *FakeConfigs) Create(ctx context.Context, config *imageregistryv1.Config, opts v1.CreateOptions) (result *imageregistryv1.Config, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(configsResource, config), &imageregistryv1.Config{})
	if obj == nil {
		return nil, err
	}
	return obj.(*imageregistryv1.Config), err
}

// Update takes the representation of a config and updates it. Returns the server's representation of the config, and an error, if there is any.
func (c *FakeConfigs) Update(ctx context.Context, config *imageregistryv1.Config, opts v1.UpdateOptions) (result *imageregistryv1.Config, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(configsResource, config), &imageregistryv1.Config{})
	if obj == nil {
		return nil, err
	}
	return obj.(*imageregistryv1.Config), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeConfigs) UpdateStatus(ctx context.Context, config *imageregistryv1.Config, opts v1.UpdateOptions) (*imageregistryv1.Config, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(configsResource, "status", config), &imageregistryv1.Config{})
	if obj == nil {
		return nil, err
	}
	return obj.(*imageregistryv1.Config), err
}

// Delete takes name of the config and deletes it. Returns an error if one occurs.
func (c *FakeConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(configsResource, name, opts), &imageregistryv1.Config{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(configsResource, listOpts)

	_, err := c.Fake.Invokes(action, &imageregistryv1.ConfigList{})
	return err
}

// Patch applies the patch and returns the patched config.
func (c *FakeConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *imageregistryv1.Config, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(configsResource, name, pt, data, subresources...), &imageregistryv1.Config{})
	if obj == nil {
		return nil, err
	}
	return obj.(*imageregistryv1.Config), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied config.
func (c *FakeConfigs) Apply(ctx context.Context, config *applyconfigurationsimageregistryv1.ConfigApplyConfiguration, opts v1.ApplyOptions) (result *imageregistryv1.Config, err error) {
	if config == nil {
		return nil, fmt.Errorf("config provided to Apply must not be nil")
	}
	data, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}
	name := config.Name
	if name == nil {
		return nil, fmt.Errorf("config.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(configsResource, *name, types.ApplyPatchType, data), &imageregistryv1.Config{})
	if obj == nil {
		return nil, err
	}
	return obj.(*imageregistryv1.Config), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeConfigs) ApplyStatus(ctx context.Context, config *applyconfigurationsimageregistryv1.ConfigApplyConfiguration, opts v1.ApplyOptions) (result *imageregistryv1.Config, err error) {
	if config == nil {
		return nil, fmt.Errorf("config provided to Apply must not be nil")
	}
	data, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}
	name := config.Name
	if name == nil {
		return nil, fmt.Errorf("config.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(configsResource, *name, types.ApplyPatchType, data, "status"), &imageregistryv1.Config{})
	if obj == nil {
		return nil, err
	}
	return obj.(*imageregistryv1.Config), err
}
