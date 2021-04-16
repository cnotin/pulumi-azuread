# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['GroupArgs', 'Group']

@pulumi.input_type
class GroupArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owners: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 prevent_duplicate_names: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Group resource.
        :param pulumi.Input[str] description: The description for the Group.  Changing this forces a new resource to be created.
        :param pulumi.Input[str] display_name: The display name for the Group. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: A set of members who should be present in this Group. Supported Object types are Users, Groups or Service Principals.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] owners: A set of owners who own this Group. Supported Object types are Users or Service Principals.
        :param pulumi.Input[bool] prevent_duplicate_names: If `true`, will return an error when an existing Group is found with the same name. Defaults to `false`.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if members is not None:
            pulumi.set(__self__, "members", members)
        if name is not None:
            warnings.warn("""This property has been renamed to `display_name` and will be removed in v2.0 of this provider.""", DeprecationWarning)
            pulumi.log.warn("""name is deprecated: This property has been renamed to `display_name` and will be removed in v2.0 of this provider.""")
        if name is not None:
            pulumi.set(__self__, "name", name)
        if owners is not None:
            pulumi.set(__self__, "owners", owners)
        if prevent_duplicate_names is not None:
            pulumi.set(__self__, "prevent_duplicate_names", prevent_duplicate_names)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description for the Group.  Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name for the Group. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A set of members who should be present in this Group. Supported Object types are Users, Groups or Service Principals.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def owners(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A set of owners who own this Group. Supported Object types are Users or Service Principals.
        """
        return pulumi.get(self, "owners")

    @owners.setter
    def owners(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "owners", value)

    @property
    @pulumi.getter(name="preventDuplicateNames")
    def prevent_duplicate_names(self) -> Optional[pulumi.Input[bool]]:
        """
        If `true`, will return an error when an existing Group is found with the same name. Defaults to `false`.
        """
        return pulumi.get(self, "prevent_duplicate_names")

    @prevent_duplicate_names.setter
    def prevent_duplicate_names(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "prevent_duplicate_names", value)


@pulumi.input_type
class _GroupState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 mail_enabled: Optional[pulumi.Input[bool]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 object_id: Optional[pulumi.Input[str]] = None,
                 owners: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 prevent_duplicate_names: Optional[pulumi.Input[bool]] = None,
                 security_enabled: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering Group resources.
        :param pulumi.Input[str] description: The description for the Group.  Changing this forces a new resource to be created.
        :param pulumi.Input[str] display_name: The display name for the Group. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] mail_enabled: Whether the group is mail-enabled.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: A set of members who should be present in this Group. Supported Object types are Users, Groups or Service Principals.
        :param pulumi.Input[str] object_id: The Object ID of the Group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] owners: A set of owners who own this Group. Supported Object types are Users or Service Principals.
        :param pulumi.Input[bool] prevent_duplicate_names: If `true`, will return an error when an existing Group is found with the same name. Defaults to `false`.
        :param pulumi.Input[bool] security_enabled: Whether the group is a security group.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if mail_enabled is not None:
            pulumi.set(__self__, "mail_enabled", mail_enabled)
        if members is not None:
            pulumi.set(__self__, "members", members)
        if name is not None:
            warnings.warn("""This property has been renamed to `display_name` and will be removed in v2.0 of this provider.""", DeprecationWarning)
            pulumi.log.warn("""name is deprecated: This property has been renamed to `display_name` and will be removed in v2.0 of this provider.""")
        if name is not None:
            pulumi.set(__self__, "name", name)
        if object_id is not None:
            pulumi.set(__self__, "object_id", object_id)
        if owners is not None:
            pulumi.set(__self__, "owners", owners)
        if prevent_duplicate_names is not None:
            pulumi.set(__self__, "prevent_duplicate_names", prevent_duplicate_names)
        if security_enabled is not None:
            pulumi.set(__self__, "security_enabled", security_enabled)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description for the Group.  Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name for the Group. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="mailEnabled")
    def mail_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the group is mail-enabled.
        """
        return pulumi.get(self, "mail_enabled")

    @mail_enabled.setter
    def mail_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "mail_enabled", value)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A set of members who should be present in this Group. Supported Object types are Users, Groups or Service Principals.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="objectId")
    def object_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Object ID of the Group.
        """
        return pulumi.get(self, "object_id")

    @object_id.setter
    def object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "object_id", value)

    @property
    @pulumi.getter
    def owners(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A set of owners who own this Group. Supported Object types are Users or Service Principals.
        """
        return pulumi.get(self, "owners")

    @owners.setter
    def owners(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "owners", value)

    @property
    @pulumi.getter(name="preventDuplicateNames")
    def prevent_duplicate_names(self) -> Optional[pulumi.Input[bool]]:
        """
        If `true`, will return an error when an existing Group is found with the same name. Defaults to `false`.
        """
        return pulumi.get(self, "prevent_duplicate_names")

    @prevent_duplicate_names.setter
    def prevent_duplicate_names(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "prevent_duplicate_names", value)

    @property
    @pulumi.getter(name="securityEnabled")
    def security_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the group is a security group.
        """
        return pulumi.get(self, "security_enabled")

    @security_enabled.setter
    def security_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "security_enabled", value)


class Group(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owners: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 prevent_duplicate_names: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Manages a Group within Azure Active Directory.

        > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Read and write all groups` within the `Windows Azure Active Directory` API. In addition it must also have either the `Company Administrator` or `User Account Administrator` Azure Active Directory roles assigned in order to be able to delete groups. You can assign one of the required Azure Active Directory Roles with the **AzureAD PowerShell Module**, which is available for Windows PowerShell or in the Azure Cloud Shell. Please refer to [this documentation](https://docs.microsoft.com/en-us/powershell/module/azuread/add-azureaddirectoryrolemember) for more details.

        ## Example Usage

        *Basic example*

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.Group("example", display_name="A-AD-Group")
        ```

        *A group with members*

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example_user = azuread.User("exampleUser",
            display_name="J Doe",
            password="notSecure123",
            user_principal_name="jdoe@hashicorp.com")
        example_group = azuread.Group("exampleGroup",
            display_name="MyGroup",
            members=[example_user.object_id])
        ```

        ## Import

        Azure Active Directory Groups can be imported using the `object id`, e.g.

        ```sh
         $ pulumi import azuread:index/group:Group my_group 00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description for the Group.  Changing this forces a new resource to be created.
        :param pulumi.Input[str] display_name: The display name for the Group. Changing this forces a new resource to be created.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: A set of members who should be present in this Group. Supported Object types are Users, Groups or Service Principals.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] owners: A set of owners who own this Group. Supported Object types are Users or Service Principals.
        :param pulumi.Input[bool] prevent_duplicate_names: If `true`, will return an error when an existing Group is found with the same name. Defaults to `false`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[GroupArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a Group within Azure Active Directory.

        > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Read and write all groups` within the `Windows Azure Active Directory` API. In addition it must also have either the `Company Administrator` or `User Account Administrator` Azure Active Directory roles assigned in order to be able to delete groups. You can assign one of the required Azure Active Directory Roles with the **AzureAD PowerShell Module**, which is available for Windows PowerShell or in the Azure Cloud Shell. Please refer to [this documentation](https://docs.microsoft.com/en-us/powershell/module/azuread/add-azureaddirectoryrolemember) for more details.

        ## Example Usage

        *Basic example*

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.Group("example", display_name="A-AD-Group")
        ```

        *A group with members*

        ```python
        import pulumi
        import pulumi_azuread as azuread

        example_user = azuread.User("exampleUser",
            display_name="J Doe",
            password="notSecure123",
            user_principal_name="jdoe@hashicorp.com")
        example_group = azuread.Group("exampleGroup",
            display_name="MyGroup",
            members=[example_user.object_id])
        ```

        ## Import

        Azure Active Directory Groups can be imported using the `object id`, e.g.

        ```sh
         $ pulumi import azuread:index/group:Group my_group 00000000-0000-0000-0000-000000000000
        ```

        :param str resource_name: The name of the resource.
        :param GroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owners: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 prevent_duplicate_names: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GroupArgs.__new__(GroupArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["members"] = members
            if name is not None and not opts.urn:
                warnings.warn("""This property has been renamed to `display_name` and will be removed in v2.0 of this provider.""", DeprecationWarning)
                pulumi.log.warn("""name is deprecated: This property has been renamed to `display_name` and will be removed in v2.0 of this provider.""")
            __props__.__dict__["name"] = name
            __props__.__dict__["owners"] = owners
            __props__.__dict__["prevent_duplicate_names"] = prevent_duplicate_names
            __props__.__dict__["mail_enabled"] = None
            __props__.__dict__["object_id"] = None
            __props__.__dict__["security_enabled"] = None
        super(Group, __self__).__init__(
            'azuread:index/group:Group',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            mail_enabled: Optional[pulumi.Input[bool]] = None,
            members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            object_id: Optional[pulumi.Input[str]] = None,
            owners: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            prevent_duplicate_names: Optional[pulumi.Input[bool]] = None,
            security_enabled: Optional[pulumi.Input[bool]] = None) -> 'Group':
        """
        Get an existing Group resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description for the Group.  Changing this forces a new resource to be created.
        :param pulumi.Input[str] display_name: The display name for the Group. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] mail_enabled: Whether the group is mail-enabled.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: A set of members who should be present in this Group. Supported Object types are Users, Groups or Service Principals.
        :param pulumi.Input[str] object_id: The Object ID of the Group.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] owners: A set of owners who own this Group. Supported Object types are Users or Service Principals.
        :param pulumi.Input[bool] prevent_duplicate_names: If `true`, will return an error when an existing Group is found with the same name. Defaults to `false`.
        :param pulumi.Input[bool] security_enabled: Whether the group is a security group.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GroupState.__new__(_GroupState)

        __props__.__dict__["description"] = description
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["mail_enabled"] = mail_enabled
        __props__.__dict__["members"] = members
        __props__.__dict__["name"] = name
        __props__.__dict__["object_id"] = object_id
        __props__.__dict__["owners"] = owners
        __props__.__dict__["prevent_duplicate_names"] = prevent_duplicate_names
        __props__.__dict__["security_enabled"] = security_enabled
        return Group(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description for the Group.  Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The display name for the Group. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="mailEnabled")
    def mail_enabled(self) -> pulumi.Output[bool]:
        """
        Whether the group is mail-enabled.
        """
        return pulumi.get(self, "mail_enabled")

    @property
    @pulumi.getter
    def members(self) -> pulumi.Output[Sequence[str]]:
        """
        A set of members who should be present in this Group. Supported Object types are Users, Groups or Service Principals.
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="objectId")
    def object_id(self) -> pulumi.Output[str]:
        """
        The Object ID of the Group.
        """
        return pulumi.get(self, "object_id")

    @property
    @pulumi.getter
    def owners(self) -> pulumi.Output[Sequence[str]]:
        """
        A set of owners who own this Group. Supported Object types are Users or Service Principals.
        """
        return pulumi.get(self, "owners")

    @property
    @pulumi.getter(name="preventDuplicateNames")
    def prevent_duplicate_names(self) -> pulumi.Output[Optional[bool]]:
        """
        If `true`, will return an error when an existing Group is found with the same name. Defaults to `false`.
        """
        return pulumi.get(self, "prevent_duplicate_names")

    @property
    @pulumi.getter(name="securityEnabled")
    def security_enabled(self) -> pulumi.Output[bool]:
        """
        Whether the group is a security group.
        """
        return pulumi.get(self, "security_enabled")

