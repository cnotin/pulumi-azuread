// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azuread/sdk/v5/go/azuread"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := azuread.NewCustomDirectoryRole(ctx, "example", &azuread.CustomDirectoryRoleArgs{
// 			Description: pulumi.String("Allows reading applications and updating groups"),
// 			DisplayName: pulumi.String("My Custom Role"),
// 			Enabled:     pulumi.Bool(true),
// 			Permissions: CustomDirectoryRolePermissionArray{
// 				&CustomDirectoryRolePermissionArgs{
// 					AllowedResourceActions: pulumi.StringArray{
// 						pulumi.String("microsoft.directory/applications/basic/update"),
// 						pulumi.String("microsoft.directory/applications/create"),
// 						pulumi.String("microsoft.directory/applications/standard/read"),
// 					},
// 				},
// 				&CustomDirectoryRolePermissionArgs{
// 					AllowedResourceActions: pulumi.StringArray{
// 						pulumi.String("microsoft.directory/groups/allProperties/read"),
// 						pulumi.String("microsoft.directory/groups/allProperties/read"),
// 						pulumi.String("microsoft.directory/groups/basic/update"),
// 						pulumi.String("microsoft.directory/groups/create"),
// 						pulumi.String("microsoft.directory/groups/delete"),
// 					},
// 				},
// 			},
// 			Version: pulumi.String("1.0"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// This resource does not support importing.
type CustomDirectoryRole struct {
	pulumi.CustomResourceState

	// The description of the custom directory role.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name of the custom directory role.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Indicates whether the role is enabled for assignment.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// The object ID of the custom directory role.
	ObjectId pulumi.StringOutput `pulumi:"objectId"`
	// A collection of `permissions` blocks as documented below.
	Permissions CustomDirectoryRolePermissionArrayOutput `pulumi:"permissions"`
	// Custom template identifier that is typically used if one needs an identifier to be the same across different directories. Changing this forces a new resource to be created.
	TemplateId pulumi.StringOutput `pulumi:"templateId"`
	// - The version of the role definition. This can be any arbitrary string between 1-128 characters.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewCustomDirectoryRole registers a new resource with the given unique name, arguments, and options.
func NewCustomDirectoryRole(ctx *pulumi.Context,
	name string, args *CustomDirectoryRoleArgs, opts ...pulumi.ResourceOption) (*CustomDirectoryRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.Permissions == nil {
		return nil, errors.New("invalid value for required argument 'Permissions'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	var resource CustomDirectoryRole
	err := ctx.RegisterResource("azuread:index/customDirectoryRole:CustomDirectoryRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomDirectoryRole gets an existing CustomDirectoryRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomDirectoryRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomDirectoryRoleState, opts ...pulumi.ResourceOption) (*CustomDirectoryRole, error) {
	var resource CustomDirectoryRole
	err := ctx.ReadResource("azuread:index/customDirectoryRole:CustomDirectoryRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomDirectoryRole resources.
type customDirectoryRoleState struct {
	// The description of the custom directory role.
	Description *string `pulumi:"description"`
	// The display name of the custom directory role.
	DisplayName *string `pulumi:"displayName"`
	// Indicates whether the role is enabled for assignment.
	Enabled *bool `pulumi:"enabled"`
	// The object ID of the custom directory role.
	ObjectId *string `pulumi:"objectId"`
	// A collection of `permissions` blocks as documented below.
	Permissions []CustomDirectoryRolePermission `pulumi:"permissions"`
	// Custom template identifier that is typically used if one needs an identifier to be the same across different directories. Changing this forces a new resource to be created.
	TemplateId *string `pulumi:"templateId"`
	// - The version of the role definition. This can be any arbitrary string between 1-128 characters.
	Version *string `pulumi:"version"`
}

type CustomDirectoryRoleState struct {
	// The description of the custom directory role.
	Description pulumi.StringPtrInput
	// The display name of the custom directory role.
	DisplayName pulumi.StringPtrInput
	// Indicates whether the role is enabled for assignment.
	Enabled pulumi.BoolPtrInput
	// The object ID of the custom directory role.
	ObjectId pulumi.StringPtrInput
	// A collection of `permissions` blocks as documented below.
	Permissions CustomDirectoryRolePermissionArrayInput
	// Custom template identifier that is typically used if one needs an identifier to be the same across different directories. Changing this forces a new resource to be created.
	TemplateId pulumi.StringPtrInput
	// - The version of the role definition. This can be any arbitrary string between 1-128 characters.
	Version pulumi.StringPtrInput
}

func (CustomDirectoryRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*customDirectoryRoleState)(nil)).Elem()
}

type customDirectoryRoleArgs struct {
	// The description of the custom directory role.
	Description *string `pulumi:"description"`
	// The display name of the custom directory role.
	DisplayName string `pulumi:"displayName"`
	// Indicates whether the role is enabled for assignment.
	Enabled bool `pulumi:"enabled"`
	// A collection of `permissions` blocks as documented below.
	Permissions []CustomDirectoryRolePermission `pulumi:"permissions"`
	// Custom template identifier that is typically used if one needs an identifier to be the same across different directories. Changing this forces a new resource to be created.
	TemplateId *string `pulumi:"templateId"`
	// - The version of the role definition. This can be any arbitrary string between 1-128 characters.
	Version string `pulumi:"version"`
}

// The set of arguments for constructing a CustomDirectoryRole resource.
type CustomDirectoryRoleArgs struct {
	// The description of the custom directory role.
	Description pulumi.StringPtrInput
	// The display name of the custom directory role.
	DisplayName pulumi.StringInput
	// Indicates whether the role is enabled for assignment.
	Enabled pulumi.BoolInput
	// A collection of `permissions` blocks as documented below.
	Permissions CustomDirectoryRolePermissionArrayInput
	// Custom template identifier that is typically used if one needs an identifier to be the same across different directories. Changing this forces a new resource to be created.
	TemplateId pulumi.StringPtrInput
	// - The version of the role definition. This can be any arbitrary string between 1-128 characters.
	Version pulumi.StringInput
}

func (CustomDirectoryRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customDirectoryRoleArgs)(nil)).Elem()
}

type CustomDirectoryRoleInput interface {
	pulumi.Input

	ToCustomDirectoryRoleOutput() CustomDirectoryRoleOutput
	ToCustomDirectoryRoleOutputWithContext(ctx context.Context) CustomDirectoryRoleOutput
}

func (*CustomDirectoryRole) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDirectoryRole)(nil)).Elem()
}

func (i *CustomDirectoryRole) ToCustomDirectoryRoleOutput() CustomDirectoryRoleOutput {
	return i.ToCustomDirectoryRoleOutputWithContext(context.Background())
}

func (i *CustomDirectoryRole) ToCustomDirectoryRoleOutputWithContext(ctx context.Context) CustomDirectoryRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDirectoryRoleOutput)
}

// CustomDirectoryRoleArrayInput is an input type that accepts CustomDirectoryRoleArray and CustomDirectoryRoleArrayOutput values.
// You can construct a concrete instance of `CustomDirectoryRoleArrayInput` via:
//
//          CustomDirectoryRoleArray{ CustomDirectoryRoleArgs{...} }
type CustomDirectoryRoleArrayInput interface {
	pulumi.Input

	ToCustomDirectoryRoleArrayOutput() CustomDirectoryRoleArrayOutput
	ToCustomDirectoryRoleArrayOutputWithContext(context.Context) CustomDirectoryRoleArrayOutput
}

type CustomDirectoryRoleArray []CustomDirectoryRoleInput

func (CustomDirectoryRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomDirectoryRole)(nil)).Elem()
}

func (i CustomDirectoryRoleArray) ToCustomDirectoryRoleArrayOutput() CustomDirectoryRoleArrayOutput {
	return i.ToCustomDirectoryRoleArrayOutputWithContext(context.Background())
}

func (i CustomDirectoryRoleArray) ToCustomDirectoryRoleArrayOutputWithContext(ctx context.Context) CustomDirectoryRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDirectoryRoleArrayOutput)
}

// CustomDirectoryRoleMapInput is an input type that accepts CustomDirectoryRoleMap and CustomDirectoryRoleMapOutput values.
// You can construct a concrete instance of `CustomDirectoryRoleMapInput` via:
//
//          CustomDirectoryRoleMap{ "key": CustomDirectoryRoleArgs{...} }
type CustomDirectoryRoleMapInput interface {
	pulumi.Input

	ToCustomDirectoryRoleMapOutput() CustomDirectoryRoleMapOutput
	ToCustomDirectoryRoleMapOutputWithContext(context.Context) CustomDirectoryRoleMapOutput
}

type CustomDirectoryRoleMap map[string]CustomDirectoryRoleInput

func (CustomDirectoryRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomDirectoryRole)(nil)).Elem()
}

func (i CustomDirectoryRoleMap) ToCustomDirectoryRoleMapOutput() CustomDirectoryRoleMapOutput {
	return i.ToCustomDirectoryRoleMapOutputWithContext(context.Background())
}

func (i CustomDirectoryRoleMap) ToCustomDirectoryRoleMapOutputWithContext(ctx context.Context) CustomDirectoryRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDirectoryRoleMapOutput)
}

type CustomDirectoryRoleOutput struct{ *pulumi.OutputState }

func (CustomDirectoryRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDirectoryRole)(nil)).Elem()
}

func (o CustomDirectoryRoleOutput) ToCustomDirectoryRoleOutput() CustomDirectoryRoleOutput {
	return o
}

func (o CustomDirectoryRoleOutput) ToCustomDirectoryRoleOutputWithContext(ctx context.Context) CustomDirectoryRoleOutput {
	return o
}

type CustomDirectoryRoleArrayOutput struct{ *pulumi.OutputState }

func (CustomDirectoryRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomDirectoryRole)(nil)).Elem()
}

func (o CustomDirectoryRoleArrayOutput) ToCustomDirectoryRoleArrayOutput() CustomDirectoryRoleArrayOutput {
	return o
}

func (o CustomDirectoryRoleArrayOutput) ToCustomDirectoryRoleArrayOutputWithContext(ctx context.Context) CustomDirectoryRoleArrayOutput {
	return o
}

func (o CustomDirectoryRoleArrayOutput) Index(i pulumi.IntInput) CustomDirectoryRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CustomDirectoryRole {
		return vs[0].([]*CustomDirectoryRole)[vs[1].(int)]
	}).(CustomDirectoryRoleOutput)
}

type CustomDirectoryRoleMapOutput struct{ *pulumi.OutputState }

func (CustomDirectoryRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomDirectoryRole)(nil)).Elem()
}

func (o CustomDirectoryRoleMapOutput) ToCustomDirectoryRoleMapOutput() CustomDirectoryRoleMapOutput {
	return o
}

func (o CustomDirectoryRoleMapOutput) ToCustomDirectoryRoleMapOutputWithContext(ctx context.Context) CustomDirectoryRoleMapOutput {
	return o
}

func (o CustomDirectoryRoleMapOutput) MapIndex(k pulumi.StringInput) CustomDirectoryRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CustomDirectoryRole {
		return vs[0].(map[string]*CustomDirectoryRole)[vs[1].(string)]
	}).(CustomDirectoryRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomDirectoryRoleInput)(nil)).Elem(), &CustomDirectoryRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomDirectoryRoleArrayInput)(nil)).Elem(), CustomDirectoryRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomDirectoryRoleMapInput)(nil)).Elem(), CustomDirectoryRoleMap{})
	pulumi.RegisterOutputType(CustomDirectoryRoleOutput{})
	pulumi.RegisterOutputType(CustomDirectoryRoleArrayOutput{})
	pulumi.RegisterOutputType(CustomDirectoryRoleMapOutput{})
}
