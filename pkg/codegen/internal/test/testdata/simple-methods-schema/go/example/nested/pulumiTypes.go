// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package nested

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Baz struct {
	Hello *string `pulumi:"hello"`
	World *string `pulumi:"world"`
}

// BazInput is an input type that accepts BazArgs and BazOutput values.
// You can construct a concrete instance of `BazInput` via:
//
//          BazArgs{...}
type BazInput interface {
	pulumi.Input

	ToBazOutput() BazOutput
	ToBazOutputWithContext(context.Context) BazOutput
}

type BazArgs struct {
	Hello pulumi.StringPtrInput `pulumi:"hello"`
	World pulumi.StringPtrInput `pulumi:"world"`
}

func (BazArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Baz)(nil)).Elem()
}

func (i BazArgs) ToBazOutput() BazOutput {
	return i.ToBazOutputWithContext(context.Background())
}

func (i BazArgs) ToBazOutputWithContext(ctx context.Context) BazOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BazOutput)
}

type BazOutput struct{ *pulumi.OutputState }

func (BazOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Baz)(nil)).Elem()
}

func (o BazOutput) ToBazOutput() BazOutput {
	return o
}

func (o BazOutput) ToBazOutputWithContext(ctx context.Context) BazOutput {
	return o
}

func (o BazOutput) Hello() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Baz) *string { return v.Hello }).(pulumi.StringPtrOutput)
}

func (o BazOutput) World() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Baz) *string { return v.World }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(BazOutput{})
}