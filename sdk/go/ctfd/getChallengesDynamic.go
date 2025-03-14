// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ctfd

import (
	"context"
	"reflect"

	"github.com/ctfer-io/pulumi-ctfd/sdk/v2/go/ctfd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetChallengesDynamic(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetChallengesDynamicResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetChallengesDynamicResult
	err := ctx.Invoke("ctfd:index/getChallengesDynamic:getChallengesDynamic", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getChallengesDynamic.
type GetChallengesDynamicResult struct {
	Challenges []GetChallengesDynamicChallenge `pulumi:"challenges"`
	// The ID of this resource.
	Id string `pulumi:"id"`
}

func GetChallengesDynamicOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetChallengesDynamicResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetChallengesDynamicResultOutput, error) {
		options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
		return ctx.InvokeOutput("ctfd:index/getChallengesDynamic:getChallengesDynamic", nil, GetChallengesDynamicResultOutput{}, options).(GetChallengesDynamicResultOutput), nil
	}).(GetChallengesDynamicResultOutput)
}

// A collection of values returned by getChallengesDynamic.
type GetChallengesDynamicResultOutput struct{ *pulumi.OutputState }

func (GetChallengesDynamicResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetChallengesDynamicResult)(nil)).Elem()
}

func (o GetChallengesDynamicResultOutput) ToGetChallengesDynamicResultOutput() GetChallengesDynamicResultOutput {
	return o
}

func (o GetChallengesDynamicResultOutput) ToGetChallengesDynamicResultOutputWithContext(ctx context.Context) GetChallengesDynamicResultOutput {
	return o
}

func (o GetChallengesDynamicResultOutput) Challenges() GetChallengesDynamicChallengeArrayOutput {
	return o.ApplyT(func(v GetChallengesDynamicResult) []GetChallengesDynamicChallenge { return v.Challenges }).(GetChallengesDynamicChallengeArrayOutput)
}

// The ID of this resource.
func (o GetChallengesDynamicResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetChallengesDynamicResult) string { return v.Id }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetChallengesDynamicResultOutput{})
}
