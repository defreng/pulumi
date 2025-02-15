// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"regress-go-8664/example/internal"
)

var _ = internal.GetEnvOrDefault

type ConditionalAccessPolicyConditions struct {
	ClientAppTypes []string `pulumi:"clientAppTypes"`
}

// ConditionalAccessPolicyConditionsInput is an input type that accepts ConditionalAccessPolicyConditionsArgs and ConditionalAccessPolicyConditionsOutput values.
// You can construct a concrete instance of `ConditionalAccessPolicyConditionsInput` via:
//
//	ConditionalAccessPolicyConditionsArgs{...}
type ConditionalAccessPolicyConditionsInput interface {
	pulumi.Input

	ToConditionalAccessPolicyConditionsOutput() ConditionalAccessPolicyConditionsOutput
	ToConditionalAccessPolicyConditionsOutputWithContext(context.Context) ConditionalAccessPolicyConditionsOutput
}

type ConditionalAccessPolicyConditionsArgs struct {
	ClientAppTypes pulumi.StringArrayInput `pulumi:"clientAppTypes"`
}

func (ConditionalAccessPolicyConditionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConditionalAccessPolicyConditions)(nil)).Elem()
}

func (i ConditionalAccessPolicyConditionsArgs) ToConditionalAccessPolicyConditionsOutput() ConditionalAccessPolicyConditionsOutput {
	return i.ToConditionalAccessPolicyConditionsOutputWithContext(context.Background())
}

func (i ConditionalAccessPolicyConditionsArgs) ToConditionalAccessPolicyConditionsOutputWithContext(ctx context.Context) ConditionalAccessPolicyConditionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionalAccessPolicyConditionsOutput)
}

func (i ConditionalAccessPolicyConditionsArgs) ToOutput(ctx context.Context) pulumix.Output[ConditionalAccessPolicyConditions] {
	return pulumix.Output[ConditionalAccessPolicyConditions]{
		OutputState: i.ToConditionalAccessPolicyConditionsOutputWithContext(ctx).OutputState,
	}
}

func (i ConditionalAccessPolicyConditionsArgs) ToConditionalAccessPolicyConditionsPtrOutput() ConditionalAccessPolicyConditionsPtrOutput {
	return i.ToConditionalAccessPolicyConditionsPtrOutputWithContext(context.Background())
}

func (i ConditionalAccessPolicyConditionsArgs) ToConditionalAccessPolicyConditionsPtrOutputWithContext(ctx context.Context) ConditionalAccessPolicyConditionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionalAccessPolicyConditionsOutput).ToConditionalAccessPolicyConditionsPtrOutputWithContext(ctx)
}

// ConditionalAccessPolicyConditionsPtrInput is an input type that accepts ConditionalAccessPolicyConditionsArgs, ConditionalAccessPolicyConditionsPtr and ConditionalAccessPolicyConditionsPtrOutput values.
// You can construct a concrete instance of `ConditionalAccessPolicyConditionsPtrInput` via:
//
//	        ConditionalAccessPolicyConditionsArgs{...}
//
//	or:
//
//	        nil
type ConditionalAccessPolicyConditionsPtrInput interface {
	pulumi.Input

	ToConditionalAccessPolicyConditionsPtrOutput() ConditionalAccessPolicyConditionsPtrOutput
	ToConditionalAccessPolicyConditionsPtrOutputWithContext(context.Context) ConditionalAccessPolicyConditionsPtrOutput
}

type conditionalAccessPolicyConditionsPtrType ConditionalAccessPolicyConditionsArgs

func ConditionalAccessPolicyConditionsPtr(v *ConditionalAccessPolicyConditionsArgs) ConditionalAccessPolicyConditionsPtrInput {
	return (*conditionalAccessPolicyConditionsPtrType)(v)
}

func (*conditionalAccessPolicyConditionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConditionalAccessPolicyConditions)(nil)).Elem()
}

func (i *conditionalAccessPolicyConditionsPtrType) ToConditionalAccessPolicyConditionsPtrOutput() ConditionalAccessPolicyConditionsPtrOutput {
	return i.ToConditionalAccessPolicyConditionsPtrOutputWithContext(context.Background())
}

func (i *conditionalAccessPolicyConditionsPtrType) ToConditionalAccessPolicyConditionsPtrOutputWithContext(ctx context.Context) ConditionalAccessPolicyConditionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionalAccessPolicyConditionsPtrOutput)
}

func (i *conditionalAccessPolicyConditionsPtrType) ToOutput(ctx context.Context) pulumix.Output[*ConditionalAccessPolicyConditions] {
	return pulumix.Output[*ConditionalAccessPolicyConditions]{
		OutputState: i.ToConditionalAccessPolicyConditionsPtrOutputWithContext(ctx).OutputState,
	}
}

type ConditionalAccessPolicyConditionsOutput struct{ *pulumi.OutputState }

func (ConditionalAccessPolicyConditionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConditionalAccessPolicyConditions)(nil)).Elem()
}

func (o ConditionalAccessPolicyConditionsOutput) ToConditionalAccessPolicyConditionsOutput() ConditionalAccessPolicyConditionsOutput {
	return o
}

func (o ConditionalAccessPolicyConditionsOutput) ToConditionalAccessPolicyConditionsOutputWithContext(ctx context.Context) ConditionalAccessPolicyConditionsOutput {
	return o
}

func (o ConditionalAccessPolicyConditionsOutput) ToConditionalAccessPolicyConditionsPtrOutput() ConditionalAccessPolicyConditionsPtrOutput {
	return o.ToConditionalAccessPolicyConditionsPtrOutputWithContext(context.Background())
}

func (o ConditionalAccessPolicyConditionsOutput) ToConditionalAccessPolicyConditionsPtrOutputWithContext(ctx context.Context) ConditionalAccessPolicyConditionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConditionalAccessPolicyConditions) *ConditionalAccessPolicyConditions {
		return &v
	}).(ConditionalAccessPolicyConditionsPtrOutput)
}

func (o ConditionalAccessPolicyConditionsOutput) ToOutput(ctx context.Context) pulumix.Output[ConditionalAccessPolicyConditions] {
	return pulumix.Output[ConditionalAccessPolicyConditions]{
		OutputState: o.OutputState,
	}
}

func (o ConditionalAccessPolicyConditionsOutput) ClientAppTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ConditionalAccessPolicyConditions) []string { return v.ClientAppTypes }).(pulumi.StringArrayOutput)
}

type ConditionalAccessPolicyConditionsPtrOutput struct{ *pulumi.OutputState }

func (ConditionalAccessPolicyConditionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConditionalAccessPolicyConditions)(nil)).Elem()
}

func (o ConditionalAccessPolicyConditionsPtrOutput) ToConditionalAccessPolicyConditionsPtrOutput() ConditionalAccessPolicyConditionsPtrOutput {
	return o
}

func (o ConditionalAccessPolicyConditionsPtrOutput) ToConditionalAccessPolicyConditionsPtrOutputWithContext(ctx context.Context) ConditionalAccessPolicyConditionsPtrOutput {
	return o
}

func (o ConditionalAccessPolicyConditionsPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*ConditionalAccessPolicyConditions] {
	return pulumix.Output[*ConditionalAccessPolicyConditions]{
		OutputState: o.OutputState,
	}
}

func (o ConditionalAccessPolicyConditionsPtrOutput) Elem() ConditionalAccessPolicyConditionsOutput {
	return o.ApplyT(func(v *ConditionalAccessPolicyConditions) ConditionalAccessPolicyConditions {
		if v != nil {
			return *v
		}
		var ret ConditionalAccessPolicyConditions
		return ret
	}).(ConditionalAccessPolicyConditionsOutput)
}

func (o ConditionalAccessPolicyConditionsPtrOutput) ClientAppTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ConditionalAccessPolicyConditions) []string {
		if v == nil {
			return nil
		}
		return v.ClientAppTypes
	}).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConditionalAccessPolicyConditionsInput)(nil)).Elem(), ConditionalAccessPolicyConditionsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConditionalAccessPolicyConditionsPtrInput)(nil)).Elem(), ConditionalAccessPolicyConditionsArgs{})
	pulumi.RegisterOutputType(ConditionalAccessPolicyConditionsOutput{})
	pulumi.RegisterOutputType(ConditionalAccessPolicyConditionsPtrOutput{})
}
