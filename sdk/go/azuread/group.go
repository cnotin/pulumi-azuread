// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Group within Azure Active Directory.
//
// > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Read and write all groups` within the `Windows Azure Active Directory` API. In addition it must also have either the `Company Administrator` or `User Account Administrator` Azure Active Directory roles assigned in order to be able to delete groups. You can assign one of the required Azure Active Directory Roles with the **AzureAD PowerShell Module**, which is available for Windows PowerShell or in the Azure Cloud Shell. Please refer to [this documentation](https://docs.microsoft.com/en-us/powershell/module/azuread/add-azureaddirectoryrolemember) for more details.
//
// ## Example Usage
//
// *Basic example*
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azuread/sdk/v4/go/azuread"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := azuread.NewGroup(ctx, "example", &azuread.GroupArgs{
// 			DisplayName: pulumi.String("A-AD-Group"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// *A group with members*
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azuread/sdk/v4/go/azuread"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleUser, err := azuread.NewUser(ctx, "exampleUser", &azuread.UserArgs{
// 			DisplayName:       pulumi.String("J Doe"),
// 			Password:          pulumi.String("notSecure123"),
// 			UserPrincipalName: pulumi.String("jdoe@hashicorp.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = azuread.NewGroup(ctx, "exampleGroup", &azuread.GroupArgs{
// 			DisplayName: pulumi.String("MyGroup"),
// 			Members: pulumi.StringArray{
// 				exampleUser.ObjectId,
// 			},
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
// Azure Active Directory Groups can be imported using the `object id`, e.g.
//
// ```sh
//  $ pulumi import azuread:index/group:Group my_group 00000000-0000-0000-0000-000000000000
// ```
type Group struct {
	pulumi.CustomResourceState

	// The description for the Group.  Changing this forces a new resource to be created.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name for the Group. Changing this forces a new resource to be created.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Whether the group is mail-enabled.
	MailEnabled pulumi.BoolOutput `pulumi:"mailEnabled"`
	// A set of members who should be present in this Group. Supported Object types are Users, Groups or Service Principals.
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// Deprecated: This property has been renamed to `display_name` and will be removed in v2.0 of this provider.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Object ID of the Group.
	ObjectId pulumi.StringOutput `pulumi:"objectId"`
	// A set of owners who own this Group. Supported Object types are Users or Service Principals.
	Owners pulumi.StringArrayOutput `pulumi:"owners"`
	// If `true`, will return an error when an existing Group is found with the same name. Defaults to `false`.
	PreventDuplicateNames pulumi.BoolPtrOutput `pulumi:"preventDuplicateNames"`
	// Whether the group is a security group.
	SecurityEnabled pulumi.BoolOutput `pulumi:"securityEnabled"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		args = &GroupArgs{}
	}

	var resource Group
	err := ctx.RegisterResource("azuread:index/group:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("azuread:index/group:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	// The description for the Group.  Changing this forces a new resource to be created.
	Description *string `pulumi:"description"`
	// The display name for the Group. Changing this forces a new resource to be created.
	DisplayName *string `pulumi:"displayName"`
	// Whether the group is mail-enabled.
	MailEnabled *bool `pulumi:"mailEnabled"`
	// A set of members who should be present in this Group. Supported Object types are Users, Groups or Service Principals.
	Members []string `pulumi:"members"`
	// Deprecated: This property has been renamed to `display_name` and will be removed in v2.0 of this provider.
	Name *string `pulumi:"name"`
	// The Object ID of the Group.
	ObjectId *string `pulumi:"objectId"`
	// A set of owners who own this Group. Supported Object types are Users or Service Principals.
	Owners []string `pulumi:"owners"`
	// If `true`, will return an error when an existing Group is found with the same name. Defaults to `false`.
	PreventDuplicateNames *bool `pulumi:"preventDuplicateNames"`
	// Whether the group is a security group.
	SecurityEnabled *bool `pulumi:"securityEnabled"`
}

type GroupState struct {
	// The description for the Group.  Changing this forces a new resource to be created.
	Description pulumi.StringPtrInput
	// The display name for the Group. Changing this forces a new resource to be created.
	DisplayName pulumi.StringPtrInput
	// Whether the group is mail-enabled.
	MailEnabled pulumi.BoolPtrInput
	// A set of members who should be present in this Group. Supported Object types are Users, Groups or Service Principals.
	Members pulumi.StringArrayInput
	// Deprecated: This property has been renamed to `display_name` and will be removed in v2.0 of this provider.
	Name pulumi.StringPtrInput
	// The Object ID of the Group.
	ObjectId pulumi.StringPtrInput
	// A set of owners who own this Group. Supported Object types are Users or Service Principals.
	Owners pulumi.StringArrayInput
	// If `true`, will return an error when an existing Group is found with the same name. Defaults to `false`.
	PreventDuplicateNames pulumi.BoolPtrInput
	// Whether the group is a security group.
	SecurityEnabled pulumi.BoolPtrInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// The description for the Group.  Changing this forces a new resource to be created.
	Description *string `pulumi:"description"`
	// The display name for the Group. Changing this forces a new resource to be created.
	DisplayName *string `pulumi:"displayName"`
	// A set of members who should be present in this Group. Supported Object types are Users, Groups or Service Principals.
	Members []string `pulumi:"members"`
	// Deprecated: This property has been renamed to `display_name` and will be removed in v2.0 of this provider.
	Name *string `pulumi:"name"`
	// A set of owners who own this Group. Supported Object types are Users or Service Principals.
	Owners []string `pulumi:"owners"`
	// If `true`, will return an error when an existing Group is found with the same name. Defaults to `false`.
	PreventDuplicateNames *bool `pulumi:"preventDuplicateNames"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// The description for the Group.  Changing this forces a new resource to be created.
	Description pulumi.StringPtrInput
	// The display name for the Group. Changing this forces a new resource to be created.
	DisplayName pulumi.StringPtrInput
	// A set of members who should be present in this Group. Supported Object types are Users, Groups or Service Principals.
	Members pulumi.StringArrayInput
	// Deprecated: This property has been renamed to `display_name` and will be removed in v2.0 of this provider.
	Name pulumi.StringPtrInput
	// A set of owners who own this Group. Supported Object types are Users or Service Principals.
	Owners pulumi.StringArrayInput
	// If `true`, will return an error when an existing Group is found with the same name. Defaults to `false`.
	PreventDuplicateNames pulumi.BoolPtrInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

type GroupInput interface {
	pulumi.Input

	ToGroupOutput() GroupOutput
	ToGroupOutputWithContext(ctx context.Context) GroupOutput
}

func (*Group) ElementType() reflect.Type {
	return reflect.TypeOf((*Group)(nil))
}

func (i *Group) ToGroupOutput() GroupOutput {
	return i.ToGroupOutputWithContext(context.Background())
}

func (i *Group) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutput)
}

func (i *Group) ToGroupPtrOutput() GroupPtrOutput {
	return i.ToGroupPtrOutputWithContext(context.Background())
}

func (i *Group) ToGroupPtrOutputWithContext(ctx context.Context) GroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPtrOutput)
}

type GroupPtrInput interface {
	pulumi.Input

	ToGroupPtrOutput() GroupPtrOutput
	ToGroupPtrOutputWithContext(ctx context.Context) GroupPtrOutput
}

type groupPtrType GroupArgs

func (*groupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil))
}

func (i *groupPtrType) ToGroupPtrOutput() GroupPtrOutput {
	return i.ToGroupPtrOutputWithContext(context.Background())
}

func (i *groupPtrType) ToGroupPtrOutputWithContext(ctx context.Context) GroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPtrOutput)
}

// GroupArrayInput is an input type that accepts GroupArray and GroupArrayOutput values.
// You can construct a concrete instance of `GroupArrayInput` via:
//
//          GroupArray{ GroupArgs{...} }
type GroupArrayInput interface {
	pulumi.Input

	ToGroupArrayOutput() GroupArrayOutput
	ToGroupArrayOutputWithContext(context.Context) GroupArrayOutput
}

type GroupArray []GroupInput

func (GroupArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Group)(nil))
}

func (i GroupArray) ToGroupArrayOutput() GroupArrayOutput {
	return i.ToGroupArrayOutputWithContext(context.Background())
}

func (i GroupArray) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupArrayOutput)
}

// GroupMapInput is an input type that accepts GroupMap and GroupMapOutput values.
// You can construct a concrete instance of `GroupMapInput` via:
//
//          GroupMap{ "key": GroupArgs{...} }
type GroupMapInput interface {
	pulumi.Input

	ToGroupMapOutput() GroupMapOutput
	ToGroupMapOutputWithContext(context.Context) GroupMapOutput
}

type GroupMap map[string]GroupInput

func (GroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Group)(nil))
}

func (i GroupMap) ToGroupMapOutput() GroupMapOutput {
	return i.ToGroupMapOutputWithContext(context.Background())
}

func (i GroupMap) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupMapOutput)
}

type GroupOutput struct {
	*pulumi.OutputState
}

func (GroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Group)(nil))
}

func (o GroupOutput) ToGroupOutput() GroupOutput {
	return o
}

func (o GroupOutput) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return o
}

func (o GroupOutput) ToGroupPtrOutput() GroupPtrOutput {
	return o.ToGroupPtrOutputWithContext(context.Background())
}

func (o GroupOutput) ToGroupPtrOutputWithContext(ctx context.Context) GroupPtrOutput {
	return o.ApplyT(func(v Group) *Group {
		return &v
	}).(GroupPtrOutput)
}

type GroupPtrOutput struct {
	*pulumi.OutputState
}

func (GroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil))
}

func (o GroupPtrOutput) ToGroupPtrOutput() GroupPtrOutput {
	return o
}

func (o GroupPtrOutput) ToGroupPtrOutputWithContext(ctx context.Context) GroupPtrOutput {
	return o
}

type GroupArrayOutput struct{ *pulumi.OutputState }

func (GroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Group)(nil))
}

func (o GroupArrayOutput) ToGroupArrayOutput() GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) ToGroupArrayOutputWithContext(ctx context.Context) GroupArrayOutput {
	return o
}

func (o GroupArrayOutput) Index(i pulumi.IntInput) GroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Group {
		return vs[0].([]Group)[vs[1].(int)]
	}).(GroupOutput)
}

type GroupMapOutput struct{ *pulumi.OutputState }

func (GroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Group)(nil))
}

func (o GroupMapOutput) ToGroupMapOutput() GroupMapOutput {
	return o
}

func (o GroupMapOutput) ToGroupMapOutputWithContext(ctx context.Context) GroupMapOutput {
	return o
}

func (o GroupMapOutput) MapIndex(k pulumi.StringInput) GroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Group {
		return vs[0].(map[string]Group)[vs[1].(string)]
	}).(GroupOutput)
}

func init() {
	pulumi.RegisterOutputType(GroupOutput{})
	pulumi.RegisterOutputType(GroupPtrOutput{})
	pulumi.RegisterOutputType(GroupArrayOutput{})
	pulumi.RegisterOutputType(GroupMapOutput{})
}
