// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"simple-resource-schema/example/internal"
)

type FooResource struct {
	pulumi.ResourceState

	Foo ResourceOutput `pulumi:"foo"`
}

// NewFooResource registers a new resource with the given unique name, arguments, and options.
func NewFooResource(ctx *pulumi.Context,
	name string, args *FooResourceArgs, opts ...pulumi.ResourceOption) (*FooResource, error) {
	if args == nil {
		args = &FooResourceArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FooResource
	err := ctx.RegisterRemoteComponentResource("foo::FooResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type fooResourceArgs struct {
	Foo *Resource `pulumi:"foo"`
}

// The set of arguments for constructing a FooResource resource.
type FooResourceArgs struct {
	Foo ResourceInput
}

func (FooResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fooResourceArgs)(nil)).Elem()
}

type FooResourceInput interface {
	pulumi.Input

	ToFooResourceOutput() FooResourceOutput
	ToFooResourceOutputWithContext(ctx context.Context) FooResourceOutput
}

func (*FooResource) ElementType() reflect.Type {
	return reflect.TypeOf((**FooResource)(nil)).Elem()
}

func (i *FooResource) ToFooResourceOutput() FooResourceOutput {
	return i.ToFooResourceOutputWithContext(context.Background())
}

func (i *FooResource) ToFooResourceOutputWithContext(ctx context.Context) FooResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FooResourceOutput)
}

func (i *FooResource) ToOutput(ctx context.Context) pulumix.Output[*FooResource] {
	return pulumix.Output[*FooResource]{
		OutputState: i.ToFooResourceOutputWithContext(ctx).OutputState,
	}
}

type FooResourceOutput struct{ *pulumi.OutputState }

func (FooResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FooResource)(nil)).Elem()
}

func (o FooResourceOutput) ToFooResourceOutput() FooResourceOutput {
	return o
}

func (o FooResourceOutput) ToFooResourceOutputWithContext(ctx context.Context) FooResourceOutput {
	return o
}

func (o FooResourceOutput) ToOutput(ctx context.Context) pulumix.Output[*FooResource] {
	return pulumix.Output[*FooResource]{
		OutputState: o.OutputState,
	}
}

func (o FooResourceOutput) Foo() ResourceOutput {
	return o.ApplyT(func(v *FooResource) ResourceOutput { return v.Foo }).(ResourceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FooResourceInput)(nil)).Elem(), &FooResource{})
	pulumi.RegisterOutputType(FooResourceOutput{})
}
