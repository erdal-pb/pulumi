// *** WARNING: this file was generated by tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codegentest

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Check codegen of functions with a List parameter.
func FuncWithListParam(ctx *pulumi.Context, args *FuncWithListParamArgs, opts ...pulumi.InvokeOption) (*FuncWithListParamResult, error) {
	var rv FuncWithListParamResult
	err := ctx.Invoke("madeup-package:codegentest:funcWithListParam", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type FuncWithListParamArgs struct {
	A []string `pulumi:"a"`
	B *string `pulumi:"b"`
}


type FuncWithListParamResult struct {
	R string `pulumi:"r"`
}


func FuncWithListParamOutput(ctx *pulumi.Context, args FuncWithListParamOutputArgs, opts ...pulumi.InvokeOption) FuncWithListParamResultTypeOutput {
        return pulumi.ToOutputWithContext(context.Background(), args).
                ApplyT(func(v interface{}) (FuncWithListParamResult, error) {
             		args := v.(FuncWithListParamArgs)
                        r, err := FuncWithListParam(ctx, &args, opts...)
                        return *r, err
                }).(FuncWithListParamResultTypeOutput)
}

type FuncWithListParamOutputArgs struct {
	A pulumi.StringArrayInput `pulumi:"a"`
	B pulumi.StringPtrInput `pulumi:"b"`
}

func (FuncWithListParamOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FuncWithListParamArgs)(nil)).Elem()
}

type FuncWithListParamResultTypeOutput struct { *pulumi.OutputState }

func (FuncWithListParamResultTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FuncWithListParamResult)(nil)).Elem()
}

func (o FuncWithListParamResultTypeOutput) ToFuncWithListParamResultTypeOutput() FuncWithListParamResultTypeOutput {
	return o
}

func (o FuncWithListParamResultTypeOutput) ToFuncWithListParamResultTypeOutputWithContext(ctx context.Context) FuncWithListParamResultTypeOutput {
	return o
}

func (o FuncWithListParamResultTypeOutput) R() pulumi.StringOutput {
	return o.ApplyT(func (v FuncWithListParamResult) string { return v.R }).(pulumi.StringOutput)
}


func init() {
        pulumi.RegisterOutputType(FuncWithListParamResultTypeOutput{})
}
