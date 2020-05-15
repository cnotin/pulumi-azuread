# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class Application(pulumi.CustomResource):
    app_roles: pulumi.Output[list]
    """
    A collection of `app_role` blocks as documented below. For more information https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles

      * `allowedMemberTypes` (`list`) - Specifies whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications (that are accessing this application in daemon service scenarios) by setting to `Application`, or to both.
      * `description` (`str`) - Permission help text that appears in the admin app assignment and consent experiences.
      * `display_name` (`str`) - Display name for the permission that appears in the admin consent and app assignment experiences.
      * `id` (`str`) - The unique identifier of the `app_role`.
      * `isEnabled` (`bool`) - Determines if the app role is enabled: Defaults to `true`.
      * `value` (`str`) - Specifies the value of the roles claim that the application should expect in the authentication and access tokens.
    """
    application_id: pulumi.Output[str]
    """
    The Application ID.
    """
    available_to_other_tenants: pulumi.Output[bool]
    """
    Is this Azure AD Application available to other tenants? Defaults to `false`.
    """
    group_membership_claims: pulumi.Output[str]
    """
    Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Defaults to `SecurityGroup`. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
    """
    homepage: pulumi.Output[str]
    """
    The URL to the application's home page. If no homepage is specified this defaults to `https://{name}`.
    """
    identifier_uris: pulumi.Output[list]
    """
    A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
    """
    logout_url: pulumi.Output[str]
    """
    The URL of the logout page.
    """
    name: pulumi.Output[str]
    """
    The display name for the application.
    """
    oauth2_allow_implicit_flow: pulumi.Output[bool]
    """
    Does this Azure AD Application allow OAuth2.0 implicit flow tokens? Defaults to `false`.
    """
    oauth2_permissions: pulumi.Output[list]
    """
    A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by a `oauth2_permission` block as documented below.

      * `adminConsentDescription` (`str`) - The description of the admin consent.
      * `adminConsentDisplayName` (`str`) - The display name of the admin consent.
      * `id` (`str`) - The unique identifier for one of the `OAuth2Permission` or `AppRole` instances that the resource application exposes.
      * `isEnabled` (`bool`) - Determines if the app role is enabled: Defaults to `true`.
      * `type` (`str`) - Type of an application: `webapp/api` or `native`. Defaults to `webapp/api`. For `native` apps type `identifier_uris` property can not not be set.
      * `userConsentDescription` (`str`) - The description of the user consent.
      * `userConsentDisplayName` (`str`) - The display name of the user consent.
      * `value` (`str`) - Specifies the value of the roles claim that the application should expect in the authentication and access tokens.
    """
    object_id: pulumi.Output[str]
    """
    The Application's Object ID.
    """
    owners: pulumi.Output[list]
    """
    A list of Azure AD Object IDs that will be granted ownership of the application. Defaults to the Object ID of the caller creating the application. If a list is specified the caller Object ID will no longer be included unless explicitly added to the list. 
    """
    public_client: pulumi.Output[bool]
    """
    Is this Azure AD Application a public client? Defaults to `false`.
    """
    reply_urls: pulumi.Output[list]
    """
    A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.
    """
    required_resource_accesses: pulumi.Output[list]
    """
    A collection of `required_resource_access` blocks as documented below.

      * `resourceAccesses` (`list`) - A collection of `resource_access` blocks as documented below.
        * `id` (`str`) - The unique identifier for one of the `OAuth2Permission` or `AppRole` instances that the resource application exposes.
        * `type` (`str`) - Specifies whether the id property references an `OAuth2Permission` or an `AppRole`. Possible values are `Scope` or `Role`.

      * `resourceAppId` (`str`) - The unique identifier for the resource that the application requires access to. This should be equal to the appId declared on the target resource application.
    """
    type: pulumi.Output[str]
    """
    Type of an application: `webapp/api` or `native`. Defaults to `webapp/api`. For `native` apps type `identifier_uris` property can not not be set.
    """
    def __init__(__self__, resource_name, opts=None, app_roles=None, available_to_other_tenants=None, group_membership_claims=None, homepage=None, identifier_uris=None, logout_url=None, name=None, oauth2_allow_implicit_flow=None, oauth2_permissions=None, owners=None, public_client=None, reply_urls=None, required_resource_accesses=None, type=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an Application within Azure Active Directory.

        > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write owned by applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_azuread as azuread

        example = azuread.Application("example",
            app_roles=[{
                "allowedMemberTypes": [
                    "User",
                    "Application",
                ],
                "description": "Admins can manage roles and perform all task actions",
                "display_name": "Admin",
                "isEnabled": True,
                "value": "Admin",
            }],
            available_to_other_tenants=False,
            homepage="https://homepage",
            identifier_uris=["https://uri"],
            oauth2_allow_implicit_flow=True,
            owners=["00000004-0000-0000-c000-000000000000"],
            reply_urls=["https://replyurl"],
            required_resource_accesses=[
                {
                    "resourceAccess": [
                        {
                            "id": "...",
                            "type": "Role",
                        },
                        {
                            "id": "...",
                            "type": "Scope",
                        },
                        {
                            "id": "...",
                            "type": "Scope",
                        },
                    ],
                    "resourceAppId": "00000003-0000-0000-c000-000000000000",
                },
                {
                    "resourceAccess": [{
                        "id": "...",
                        "type": "Scope",
                    }],
                    "resourceAppId": "00000002-0000-0000-c000-000000000000",
                },
            ],
            type="webapp/api")
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] app_roles: A collection of `app_role` blocks as documented below. For more information https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles
        :param pulumi.Input[bool] available_to_other_tenants: Is this Azure AD Application available to other tenants? Defaults to `false`.
        :param pulumi.Input[str] group_membership_claims: Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Defaults to `SecurityGroup`. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
        :param pulumi.Input[str] homepage: The URL to the application's home page. If no homepage is specified this defaults to `https://{name}`.
        :param pulumi.Input[list] identifier_uris: A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
        :param pulumi.Input[str] logout_url: The URL of the logout page.
        :param pulumi.Input[str] name: The display name for the application.
        :param pulumi.Input[bool] oauth2_allow_implicit_flow: Does this Azure AD Application allow OAuth2.0 implicit flow tokens? Defaults to `false`.
        :param pulumi.Input[list] oauth2_permissions: A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by a `oauth2_permission` block as documented below.
        :param pulumi.Input[list] owners: A list of Azure AD Object IDs that will be granted ownership of the application. Defaults to the Object ID of the caller creating the application. If a list is specified the caller Object ID will no longer be included unless explicitly added to the list. 
        :param pulumi.Input[bool] public_client: Is this Azure AD Application a public client? Defaults to `false`.
        :param pulumi.Input[list] reply_urls: A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.
        :param pulumi.Input[list] required_resource_accesses: A collection of `required_resource_access` blocks as documented below.
        :param pulumi.Input[str] type: Type of an application: `webapp/api` or `native`. Defaults to `webapp/api`. For `native` apps type `identifier_uris` property can not not be set.

        The **app_roles** object supports the following:

          * `allowedMemberTypes` (`pulumi.Input[list]`) - Specifies whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications (that are accessing this application in daemon service scenarios) by setting to `Application`, or to both.
          * `description` (`pulumi.Input[str]`) - Permission help text that appears in the admin app assignment and consent experiences.
          * `display_name` (`pulumi.Input[str]`) - Display name for the permission that appears in the admin consent and app assignment experiences.
          * `id` (`pulumi.Input[str]`) - The unique identifier of the `app_role`.
          * `isEnabled` (`pulumi.Input[bool]`) - Determines if the app role is enabled: Defaults to `true`.
          * `value` (`pulumi.Input[str]`) - Specifies the value of the roles claim that the application should expect in the authentication and access tokens.

        The **oauth2_permissions** object supports the following:

          * `adminConsentDescription` (`pulumi.Input[str]`) - The description of the admin consent.
          * `adminConsentDisplayName` (`pulumi.Input[str]`) - The display name of the admin consent.
          * `id` (`pulumi.Input[str]`) - The unique identifier for one of the `OAuth2Permission` or `AppRole` instances that the resource application exposes.
          * `isEnabled` (`pulumi.Input[bool]`) - Determines if the app role is enabled: Defaults to `true`.
          * `type` (`pulumi.Input[str]`) - Type of an application: `webapp/api` or `native`. Defaults to `webapp/api`. For `native` apps type `identifier_uris` property can not not be set.
          * `userConsentDescription` (`pulumi.Input[str]`) - The description of the user consent.
          * `userConsentDisplayName` (`pulumi.Input[str]`) - The display name of the user consent.
          * `value` (`pulumi.Input[str]`) - Specifies the value of the roles claim that the application should expect in the authentication and access tokens.

        The **required_resource_accesses** object supports the following:

          * `resourceAccesses` (`pulumi.Input[list]`) - A collection of `resource_access` blocks as documented below.
            * `id` (`pulumi.Input[str]`) - The unique identifier for one of the `OAuth2Permission` or `AppRole` instances that the resource application exposes.
            * `type` (`pulumi.Input[str]`) - Specifies whether the id property references an `OAuth2Permission` or an `AppRole`. Possible values are `Scope` or `Role`.

          * `resourceAppId` (`pulumi.Input[str]`) - The unique identifier for the resource that the application requires access to. This should be equal to the appId declared on the target resource application.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['app_roles'] = app_roles
            __props__['available_to_other_tenants'] = available_to_other_tenants
            __props__['group_membership_claims'] = group_membership_claims
            __props__['homepage'] = homepage
            __props__['identifier_uris'] = identifier_uris
            __props__['logout_url'] = logout_url
            __props__['name'] = name
            __props__['oauth2_allow_implicit_flow'] = oauth2_allow_implicit_flow
            __props__['oauth2_permissions'] = oauth2_permissions
            __props__['owners'] = owners
            __props__['public_client'] = public_client
            __props__['reply_urls'] = reply_urls
            __props__['required_resource_accesses'] = required_resource_accesses
            __props__['type'] = type
            __props__['application_id'] = None
            __props__['object_id'] = None
        super(Application, __self__).__init__(
            'azuread:index/application:Application',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, app_roles=None, application_id=None, available_to_other_tenants=None, group_membership_claims=None, homepage=None, identifier_uris=None, logout_url=None, name=None, oauth2_allow_implicit_flow=None, oauth2_permissions=None, object_id=None, owners=None, public_client=None, reply_urls=None, required_resource_accesses=None, type=None):
        """
        Get an existing Application resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] app_roles: A collection of `app_role` blocks as documented below. For more information https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles
        :param pulumi.Input[str] application_id: The Application ID.
        :param pulumi.Input[bool] available_to_other_tenants: Is this Azure AD Application available to other tenants? Defaults to `false`.
        :param pulumi.Input[str] group_membership_claims: Configures the `groups` claim issued in a user or OAuth 2.0 access token that the app expects. Defaults to `SecurityGroup`. Possible values are `None`, `SecurityGroup`, `DirectoryRole`, `ApplicationGroup` or `All`.
        :param pulumi.Input[str] homepage: The URL to the application's home page. If no homepage is specified this defaults to `https://{name}`.
        :param pulumi.Input[list] identifier_uris: A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
        :param pulumi.Input[str] logout_url: The URL of the logout page.
        :param pulumi.Input[str] name: The display name for the application.
        :param pulumi.Input[bool] oauth2_allow_implicit_flow: Does this Azure AD Application allow OAuth2.0 implicit flow tokens? Defaults to `false`.
        :param pulumi.Input[list] oauth2_permissions: A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by a `oauth2_permission` block as documented below.
        :param pulumi.Input[str] object_id: The Application's Object ID.
        :param pulumi.Input[list] owners: A list of Azure AD Object IDs that will be granted ownership of the application. Defaults to the Object ID of the caller creating the application. If a list is specified the caller Object ID will no longer be included unless explicitly added to the list. 
        :param pulumi.Input[bool] public_client: Is this Azure AD Application a public client? Defaults to `false`.
        :param pulumi.Input[list] reply_urls: A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.
        :param pulumi.Input[list] required_resource_accesses: A collection of `required_resource_access` blocks as documented below.
        :param pulumi.Input[str] type: Type of an application: `webapp/api` or `native`. Defaults to `webapp/api`. For `native` apps type `identifier_uris` property can not not be set.

        The **app_roles** object supports the following:

          * `allowedMemberTypes` (`pulumi.Input[list]`) - Specifies whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications (that are accessing this application in daemon service scenarios) by setting to `Application`, or to both.
          * `description` (`pulumi.Input[str]`) - Permission help text that appears in the admin app assignment and consent experiences.
          * `display_name` (`pulumi.Input[str]`) - Display name for the permission that appears in the admin consent and app assignment experiences.
          * `id` (`pulumi.Input[str]`) - The unique identifier of the `app_role`.
          * `isEnabled` (`pulumi.Input[bool]`) - Determines if the app role is enabled: Defaults to `true`.
          * `value` (`pulumi.Input[str]`) - Specifies the value of the roles claim that the application should expect in the authentication and access tokens.

        The **oauth2_permissions** object supports the following:

          * `adminConsentDescription` (`pulumi.Input[str]`) - The description of the admin consent.
          * `adminConsentDisplayName` (`pulumi.Input[str]`) - The display name of the admin consent.
          * `id` (`pulumi.Input[str]`) - The unique identifier for one of the `OAuth2Permission` or `AppRole` instances that the resource application exposes.
          * `isEnabled` (`pulumi.Input[bool]`) - Determines if the app role is enabled: Defaults to `true`.
          * `type` (`pulumi.Input[str]`) - Type of an application: `webapp/api` or `native`. Defaults to `webapp/api`. For `native` apps type `identifier_uris` property can not not be set.
          * `userConsentDescription` (`pulumi.Input[str]`) - The description of the user consent.
          * `userConsentDisplayName` (`pulumi.Input[str]`) - The display name of the user consent.
          * `value` (`pulumi.Input[str]`) - Specifies the value of the roles claim that the application should expect in the authentication and access tokens.

        The **required_resource_accesses** object supports the following:

          * `resourceAccesses` (`pulumi.Input[list]`) - A collection of `resource_access` blocks as documented below.
            * `id` (`pulumi.Input[str]`) - The unique identifier for one of the `OAuth2Permission` or `AppRole` instances that the resource application exposes.
            * `type` (`pulumi.Input[str]`) - Specifies whether the id property references an `OAuth2Permission` or an `AppRole`. Possible values are `Scope` or `Role`.

          * `resourceAppId` (`pulumi.Input[str]`) - The unique identifier for the resource that the application requires access to. This should be equal to the appId declared on the target resource application.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["app_roles"] = app_roles
        __props__["application_id"] = application_id
        __props__["available_to_other_tenants"] = available_to_other_tenants
        __props__["group_membership_claims"] = group_membership_claims
        __props__["homepage"] = homepage
        __props__["identifier_uris"] = identifier_uris
        __props__["logout_url"] = logout_url
        __props__["name"] = name
        __props__["oauth2_allow_implicit_flow"] = oauth2_allow_implicit_flow
        __props__["oauth2_permissions"] = oauth2_permissions
        __props__["object_id"] = object_id
        __props__["owners"] = owners
        __props__["public_client"] = public_client
        __props__["reply_urls"] = reply_urls
        __props__["required_resource_accesses"] = required_resource_accesses
        __props__["type"] = type
        return Application(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

