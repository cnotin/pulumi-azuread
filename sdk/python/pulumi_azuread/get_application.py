# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class GetApplicationResult:
    """
    A collection of values returned by getApplication.
    """
    def __init__(__self__, application_id=None, available_to_other_tenants=None, group_membership_claims=None, homepage=None, identifier_uris=None, name=None, oauth2_allow_implicit_flow=None, oauth2_permissions=None, object_id=None, reply_urls=None, required_resource_accesses=None, type=None, id=None):
        if application_id and not isinstance(application_id, str):
            raise TypeError("Expected argument 'application_id' to be a str")
        __self__.application_id = application_id
        """
        the Application ID of the Azure Active Directory Application.
        """
        if available_to_other_tenants and not isinstance(available_to_other_tenants, bool):
            raise TypeError("Expected argument 'available_to_other_tenants' to be a bool")
        __self__.available_to_other_tenants = available_to_other_tenants
        """
        Is this Azure AD Application available to other tenants?
        """
        if group_membership_claims and not isinstance(group_membership_claims, str):
            raise TypeError("Expected argument 'group_membership_claims' to be a str")
        __self__.group_membership_claims = group_membership_claims
        """
        The `groups` claim issued in a user or OAuth 2.0 access token that the app expects.
        """
        if homepage and not isinstance(homepage, str):
            raise TypeError("Expected argument 'homepage' to be a str")
        __self__.homepage = homepage
        if identifier_uris and not isinstance(identifier_uris, list):
            raise TypeError("Expected argument 'identifier_uris' to be a list")
        __self__.identifier_uris = identifier_uris
        """
        A list of user-defined URI(s) that uniquely identify a Web application within it's Azure AD tenant, or within a verified custom domain if the application is multi-tenant.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if oauth2_allow_implicit_flow and not isinstance(oauth2_allow_implicit_flow, bool):
            raise TypeError("Expected argument 'oauth2_allow_implicit_flow' to be a bool")
        __self__.oauth2_allow_implicit_flow = oauth2_allow_implicit_flow
        """
        Does this Azure AD Application allow OAuth2.0 implicit flow tokens?
        """
        if oauth2_permissions and not isinstance(oauth2_permissions, list):
            raise TypeError("Expected argument 'oauth2_permissions' to be a list")
        __self__.oauth2_permissions = oauth2_permissions
        """
        A collection of OAuth 2.0 permission scopes that the web API (resource) app exposes to client apps. Each permission is covered by a `oauth2_permission` block as documented below.
        """
        if object_id and not isinstance(object_id, str):
            raise TypeError("Expected argument 'object_id' to be a str")
        __self__.object_id = object_id
        """
        the Object ID of the Azure Active Directory Application.
        """
        if reply_urls and not isinstance(reply_urls, list):
            raise TypeError("Expected argument 'reply_urls' to be a list")
        __self__.reply_urls = reply_urls
        """
        A list of URLs that user tokens are sent to for sign in, or the redirect URIs that OAuth 2.0 authorization codes and access tokens are sent to.
        """
        if required_resource_accesses and not isinstance(required_resource_accesses, list):
            raise TypeError("Expected argument 'required_resource_accesses' to be a list")
        __self__.required_resource_accesses = required_resource_accesses
        """
        A collection of `required_resource_access` blocks as documented below.
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        The type of the permission
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_application(name=None,oauth2_permissions=None,object_id=None,opts=None):
    """
    Use this data source to access information about an existing Application within Azure Active Directory.
    
    > **NOTE:** If you're authenticating using a Service Principal then it must have permissions to both `Read and write all applications` and `Sign in and read user profile` within the `Windows Azure Active Directory` API.
    """
    __args__ = dict()

    __args__['name'] = name
    __args__['oauth2Permissions'] = oauth2_permissions
    __args__['objectId'] = object_id
    __ret__ = await pulumi.runtime.invoke('azuread:index/getApplication:getApplication', __args__, opts=opts)

    return GetApplicationResult(
        application_id=__ret__.get('applicationId'),
        available_to_other_tenants=__ret__.get('availableToOtherTenants'),
        group_membership_claims=__ret__.get('groupMembershipClaims'),
        homepage=__ret__.get('homepage'),
        identifier_uris=__ret__.get('identifierUris'),
        name=__ret__.get('name'),
        oauth2_allow_implicit_flow=__ret__.get('oauth2AllowImplicitFlow'),
        oauth2_permissions=__ret__.get('oauth2Permissions'),
        object_id=__ret__.get('objectId'),
        reply_urls=__ret__.get('replyUrls'),
        required_resource_accesses=__ret__.get('requiredResourceAccesses'),
        type=__ret__.get('type'),
        id=__ret__.get('id'))
