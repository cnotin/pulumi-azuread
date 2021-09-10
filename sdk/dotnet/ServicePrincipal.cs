// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD
{
    /// <summary>
    /// Manages a service principal associated with an application within Azure Active Directory.
    /// 
    /// ## API Permissions
    /// 
    /// The following API permissions are required in order to use this resource.
    /// 
    /// When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.All` or `Directory.ReadWrite.All`
    /// 
    /// It is not currently possible to manage service principals whilst having only the `Application.ReadWrite.OwnedBy` role granted.
    /// 
    /// When authenticated with a user principal, this resource requires one of the following directory roles: `Application Administrator` or `Global Administrator`
    /// 
    /// ## Example Usage
    /// 
    /// *Create a service principal for an application*
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AzureAD = Pulumi.AzureAD;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var current = Output.Create(AzureAD.GetClientConfig.InvokeAsync());
    ///         var exampleApplication = new AzureAD.Application("exampleApplication", new AzureAD.ApplicationArgs
    ///         {
    ///             DisplayName = "example",
    ///             Owners = 
    ///             {
    ///                 current.Apply(current =&gt; current.ObjectId),
    ///             },
    ///         });
    ///         var exampleServicePrincipal = new AzureAD.ServicePrincipal("exampleServicePrincipal", new AzureAD.ServicePrincipalArgs
    ///         {
    ///             ApplicationId = exampleApplication.ApplicationId,
    ///             AppRoleAssignmentRequired = false,
    ///             Owners = 
    ///             {
    ///                 current.Apply(current =&gt; current.ObjectId),
    ///             },
    ///             Tags = 
    ///             {
    ///                 "example",
    ///                 "tags",
    ///                 "here",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// *Manage a service principal for a first-party Microsoft application*
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AzureAD = Pulumi.AzureAD;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var wellKnown = Output.Create(AzureAD.GetApplicationPublishedAppIds.InvokeAsync());
    ///         var msgraph = new AzureAD.ServicePrincipal("msgraph", new AzureAD.ServicePrincipalArgs
    ///         {
    ///             ApplicationId = wellKnown.Apply(wellKnown =&gt; wellKnown.Result?.MicrosoftGraph),
    ///             UseExisting = true,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// *Create a service principal for an application created from a gallery template*
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AzureAD = Pulumi.AzureAD;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleApplicationTemplate = Output.Create(AzureAD.GetApplicationTemplate.InvokeAsync(new AzureAD.GetApplicationTemplateArgs
    ///         {
    ///             DisplayName = "Marketo",
    ///         }));
    ///         var exampleApplication = new AzureAD.Application("exampleApplication", new AzureAD.ApplicationArgs
    ///         {
    ///             DisplayName = "example",
    ///             TemplateId = exampleApplicationTemplate.Apply(exampleApplicationTemplate =&gt; exampleApplicationTemplate.TemplateId),
    ///         });
    ///         var exampleServicePrincipal = new AzureAD.ServicePrincipal("exampleServicePrincipal", new AzureAD.ServicePrincipalArgs
    ///         {
    ///             ApplicationId = exampleApplication.ApplicationId,
    ///             UseExisting = true,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Service principals can be imported using their object ID, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azuread:index/servicePrincipal:ServicePrincipal test 00000000-0000-0000-0000-000000000000
    /// ```
    /// </summary>
    [AzureADResourceType("azuread:index/servicePrincipal:ServicePrincipal")]
    public partial class ServicePrincipal : Pulumi.CustomResource
    {
        /// <summary>
        /// Whether or not the service principal account is enabled. Defaults to `true`.
        /// </summary>
        [Output("accountEnabled")]
        public Output<bool?> AccountEnabled { get; private set; } = null!;

        /// <summary>
        /// A set of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.
        /// </summary>
        [Output("alternativeNames")]
        public Output<ImmutableArray<string>> AlternativeNames { get; private set; } = null!;

        /// <summary>
        /// Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application. Defaults to `false`.
        /// </summary>
        [Output("appRoleAssignmentRequired")]
        public Output<bool?> AppRoleAssignmentRequired { get; private set; } = null!;

        /// <summary>
        /// A mapping of app role values to app role IDs, as published by the associated application, intended to be useful when referencing app roles in other resources in your configuration.
        /// </summary>
        [Output("appRoleIds")]
        public Output<ImmutableDictionary<string, string>> AppRoleIds { get; private set; } = null!;

        /// <summary>
        /// A list of app roles published by the associated application, as documented below. For more information [official documentation](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
        /// </summary>
        [Output("appRoles")]
        public Output<ImmutableArray<Outputs.ServicePrincipalAppRole>> AppRoles { get; private set; } = null!;

        /// <summary>
        /// The application ID (client ID) of the application for which to create a service principal.
        /// </summary>
        [Output("applicationId")]
        public Output<string> ApplicationId { get; private set; } = null!;

        /// <summary>
        /// The tenant ID where the associated application is registered.
        /// </summary>
        [Output("applicationTenantId")]
        public Output<string> ApplicationTenantId { get; private set; } = null!;

        /// <summary>
        /// A description of the service principal provided for internal end-users.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Display name for the app role that appears during app role assignment and in consent experiences.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Home page or landing page of the associated application.
        /// </summary>
        [Output("homepageUrl")]
        public Output<string> HomepageUrl { get; private set; } = null!;

        /// <summary>
        /// The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on.
        /// </summary>
        [Output("loginUrl")]
        public Output<string?> LoginUrl { get; private set; } = null!;

        /// <summary>
        /// The URL that will be used by Microsoft's authorization service to logout an user using OpenId Connect front-channel, back-channel or SAML logout protocols, taken from the associated application.
        /// </summary>
        [Output("logoutUrl")]
        public Output<string> LogoutUrl { get; private set; } = null!;

        /// <summary>
        /// A free text field to capture information about the service principal, typically used for operational purposes.
        /// </summary>
        [Output("notes")]
        public Output<string?> Notes { get; private set; } = null!;

        /// <summary>
        /// A set of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
        /// </summary>
        [Output("notificationEmailAddresses")]
        public Output<ImmutableArray<string>> NotificationEmailAddresses { get; private set; } = null!;

        /// <summary>
        /// A mapping of OAuth2.0 permission scope values to scope IDs, as exposed by the associated application, intended to be useful when referencing permission scopes in other resources in your configuration.
        /// </summary>
        [Output("oauth2PermissionScopeIds")]
        public Output<ImmutableDictionary<string, string>> Oauth2PermissionScopeIds { get; private set; } = null!;

        /// <summary>
        /// A list of OAuth 2.0 delegated permission scopes exposed by the associated application, as documented below.
        /// </summary>
        [Output("oauth2PermissionScopes")]
        public Output<ImmutableArray<Outputs.ServicePrincipalOauth2PermissionScope>> Oauth2PermissionScopes { get; private set; } = null!;

        /// <summary>
        /// The object ID of the service principal.
        /// </summary>
        [Output("objectId")]
        public Output<string> ObjectId { get; private set; } = null!;

        /// <summary>
        /// A set of object IDs of principals that will be granted ownership of the service principal. Supported object types are users or service principals. By default, no owners are assigned.
        /// </summary>
        [Output("owners")]
        public Output<ImmutableArray<string>> Owners { get; private set; } = null!;

        /// <summary>
        /// The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps. Supported values are `oidc`, `password`, `saml` or `notSupported`. Omit this property or specify a blank string to unset.
        /// </summary>
        [Output("preferredSingleSignOnMode")]
        public Output<string?> PreferredSingleSignOnMode { get; private set; } = null!;

        /// <summary>
        /// A list of URLs where user tokens are sent for sign-in with the associated application, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent for the associated application.
        /// </summary>
        [Output("redirectUris")]
        public Output<ImmutableArray<string>> RedirectUris { get; private set; } = null!;

        /// <summary>
        /// The URL where the service exposes SAML metadata for federation.
        /// </summary>
        [Output("samlMetadataUrl")]
        public Output<string> SamlMetadataUrl { get; private set; } = null!;

        /// <summary>
        /// A `saml_single_sign_on` block as documented below.
        /// </summary>
        [Output("samlSingleSignOn")]
        public Output<Outputs.ServicePrincipalSamlSingleSignOn?> SamlSingleSignOn { get; private set; } = null!;

        /// <summary>
        /// A list of identifier URI(s), copied over from the associated application.
        /// </summary>
        [Output("servicePrincipalNames")]
        public Output<ImmutableArray<string>> ServicePrincipalNames { get; private set; } = null!;

        /// <summary>
        /// The Microsoft account types that are supported for the associated application. Possible values include `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`.
        /// </summary>
        [Output("signInAudience")]
        public Output<string> SignInAudience { get; private set; } = null!;

        /// <summary>
        /// A set of tags to apply to the service principal.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are `User` or `Admin`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// When true, any existing service principal linked to the same application will be automatically imported. When false, an import error will be raised for any pre-existing service principal.
        /// </summary>
        [Output("useExisting")]
        public Output<bool?> UseExisting { get; private set; } = null!;


        /// <summary>
        /// Create a ServicePrincipal resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServicePrincipal(string name, ServicePrincipalArgs args, CustomResourceOptions? options = null)
            : base("azuread:index/servicePrincipal:ServicePrincipal", name, args ?? new ServicePrincipalArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServicePrincipal(string name, Input<string> id, ServicePrincipalState? state = null, CustomResourceOptions? options = null)
            : base("azuread:index/servicePrincipal:ServicePrincipal", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServicePrincipal resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServicePrincipal Get(string name, Input<string> id, ServicePrincipalState? state = null, CustomResourceOptions? options = null)
        {
            return new ServicePrincipal(name, id, state, options);
        }
    }

    public sealed class ServicePrincipalArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether or not the service principal account is enabled. Defaults to `true`.
        /// </summary>
        [Input("accountEnabled")]
        public Input<bool>? AccountEnabled { get; set; }

        [Input("alternativeNames")]
        private InputList<string>? _alternativeNames;

        /// <summary>
        /// A set of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.
        /// </summary>
        public InputList<string> AlternativeNames
        {
            get => _alternativeNames ?? (_alternativeNames = new InputList<string>());
            set => _alternativeNames = value;
        }

        /// <summary>
        /// Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application. Defaults to `false`.
        /// </summary>
        [Input("appRoleAssignmentRequired")]
        public Input<bool>? AppRoleAssignmentRequired { get; set; }

        /// <summary>
        /// The application ID (client ID) of the application for which to create a service principal.
        /// </summary>
        [Input("applicationId", required: true)]
        public Input<string> ApplicationId { get; set; } = null!;

        /// <summary>
        /// A description of the service principal provided for internal end-users.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on.
        /// </summary>
        [Input("loginUrl")]
        public Input<string>? LoginUrl { get; set; }

        /// <summary>
        /// A free text field to capture information about the service principal, typically used for operational purposes.
        /// </summary>
        [Input("notes")]
        public Input<string>? Notes { get; set; }

        [Input("notificationEmailAddresses")]
        private InputList<string>? _notificationEmailAddresses;

        /// <summary>
        /// A set of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
        /// </summary>
        public InputList<string> NotificationEmailAddresses
        {
            get => _notificationEmailAddresses ?? (_notificationEmailAddresses = new InputList<string>());
            set => _notificationEmailAddresses = value;
        }

        [Input("owners")]
        private InputList<string>? _owners;

        /// <summary>
        /// A set of object IDs of principals that will be granted ownership of the service principal. Supported object types are users or service principals. By default, no owners are assigned.
        /// </summary>
        public InputList<string> Owners
        {
            get => _owners ?? (_owners = new InputList<string>());
            set => _owners = value;
        }

        /// <summary>
        /// The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps. Supported values are `oidc`, `password`, `saml` or `notSupported`. Omit this property or specify a blank string to unset.
        /// </summary>
        [Input("preferredSingleSignOnMode")]
        public Input<string>? PreferredSingleSignOnMode { get; set; }

        /// <summary>
        /// A `saml_single_sign_on` block as documented below.
        /// </summary>
        [Input("samlSingleSignOn")]
        public Input<Inputs.ServicePrincipalSamlSingleSignOnArgs>? SamlSingleSignOn { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A set of tags to apply to the service principal.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// When true, any existing service principal linked to the same application will be automatically imported. When false, an import error will be raised for any pre-existing service principal.
        /// </summary>
        [Input("useExisting")]
        public Input<bool>? UseExisting { get; set; }

        public ServicePrincipalArgs()
        {
        }
    }

    public sealed class ServicePrincipalState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether or not the service principal account is enabled. Defaults to `true`.
        /// </summary>
        [Input("accountEnabled")]
        public Input<bool>? AccountEnabled { get; set; }

        [Input("alternativeNames")]
        private InputList<string>? _alternativeNames;

        /// <summary>
        /// A set of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.
        /// </summary>
        public InputList<string> AlternativeNames
        {
            get => _alternativeNames ?? (_alternativeNames = new InputList<string>());
            set => _alternativeNames = value;
        }

        /// <summary>
        /// Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application. Defaults to `false`.
        /// </summary>
        [Input("appRoleAssignmentRequired")]
        public Input<bool>? AppRoleAssignmentRequired { get; set; }

        [Input("appRoleIds")]
        private InputMap<string>? _appRoleIds;

        /// <summary>
        /// A mapping of app role values to app role IDs, as published by the associated application, intended to be useful when referencing app roles in other resources in your configuration.
        /// </summary>
        public InputMap<string> AppRoleIds
        {
            get => _appRoleIds ?? (_appRoleIds = new InputMap<string>());
            set => _appRoleIds = value;
        }

        [Input("appRoles")]
        private InputList<Inputs.ServicePrincipalAppRoleGetArgs>? _appRoles;

        /// <summary>
        /// A list of app roles published by the associated application, as documented below. For more information [official documentation](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
        /// </summary>
        public InputList<Inputs.ServicePrincipalAppRoleGetArgs> AppRoles
        {
            get => _appRoles ?? (_appRoles = new InputList<Inputs.ServicePrincipalAppRoleGetArgs>());
            set => _appRoles = value;
        }

        /// <summary>
        /// The application ID (client ID) of the application for which to create a service principal.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// The tenant ID where the associated application is registered.
        /// </summary>
        [Input("applicationTenantId")]
        public Input<string>? ApplicationTenantId { get; set; }

        /// <summary>
        /// A description of the service principal provided for internal end-users.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Display name for the app role that appears during app role assignment and in consent experiences.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Home page or landing page of the associated application.
        /// </summary>
        [Input("homepageUrl")]
        public Input<string>? HomepageUrl { get; set; }

        /// <summary>
        /// The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on.
        /// </summary>
        [Input("loginUrl")]
        public Input<string>? LoginUrl { get; set; }

        /// <summary>
        /// The URL that will be used by Microsoft's authorization service to logout an user using OpenId Connect front-channel, back-channel or SAML logout protocols, taken from the associated application.
        /// </summary>
        [Input("logoutUrl")]
        public Input<string>? LogoutUrl { get; set; }

        /// <summary>
        /// A free text field to capture information about the service principal, typically used for operational purposes.
        /// </summary>
        [Input("notes")]
        public Input<string>? Notes { get; set; }

        [Input("notificationEmailAddresses")]
        private InputList<string>? _notificationEmailAddresses;

        /// <summary>
        /// A set of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
        /// </summary>
        public InputList<string> NotificationEmailAddresses
        {
            get => _notificationEmailAddresses ?? (_notificationEmailAddresses = new InputList<string>());
            set => _notificationEmailAddresses = value;
        }

        [Input("oauth2PermissionScopeIds")]
        private InputMap<string>? _oauth2PermissionScopeIds;

        /// <summary>
        /// A mapping of OAuth2.0 permission scope values to scope IDs, as exposed by the associated application, intended to be useful when referencing permission scopes in other resources in your configuration.
        /// </summary>
        public InputMap<string> Oauth2PermissionScopeIds
        {
            get => _oauth2PermissionScopeIds ?? (_oauth2PermissionScopeIds = new InputMap<string>());
            set => _oauth2PermissionScopeIds = value;
        }

        [Input("oauth2PermissionScopes")]
        private InputList<Inputs.ServicePrincipalOauth2PermissionScopeGetArgs>? _oauth2PermissionScopes;

        /// <summary>
        /// A list of OAuth 2.0 delegated permission scopes exposed by the associated application, as documented below.
        /// </summary>
        public InputList<Inputs.ServicePrincipalOauth2PermissionScopeGetArgs> Oauth2PermissionScopes
        {
            get => _oauth2PermissionScopes ?? (_oauth2PermissionScopes = new InputList<Inputs.ServicePrincipalOauth2PermissionScopeGetArgs>());
            set => _oauth2PermissionScopes = value;
        }

        /// <summary>
        /// The object ID of the service principal.
        /// </summary>
        [Input("objectId")]
        public Input<string>? ObjectId { get; set; }

        [Input("owners")]
        private InputList<string>? _owners;

        /// <summary>
        /// A set of object IDs of principals that will be granted ownership of the service principal. Supported object types are users or service principals. By default, no owners are assigned.
        /// </summary>
        public InputList<string> Owners
        {
            get => _owners ?? (_owners = new InputList<string>());
            set => _owners = value;
        }

        /// <summary>
        /// The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps. Supported values are `oidc`, `password`, `saml` or `notSupported`. Omit this property or specify a blank string to unset.
        /// </summary>
        [Input("preferredSingleSignOnMode")]
        public Input<string>? PreferredSingleSignOnMode { get; set; }

        [Input("redirectUris")]
        private InputList<string>? _redirectUris;

        /// <summary>
        /// A list of URLs where user tokens are sent for sign-in with the associated application, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent for the associated application.
        /// </summary>
        public InputList<string> RedirectUris
        {
            get => _redirectUris ?? (_redirectUris = new InputList<string>());
            set => _redirectUris = value;
        }

        /// <summary>
        /// The URL where the service exposes SAML metadata for federation.
        /// </summary>
        [Input("samlMetadataUrl")]
        public Input<string>? SamlMetadataUrl { get; set; }

        /// <summary>
        /// A `saml_single_sign_on` block as documented below.
        /// </summary>
        [Input("samlSingleSignOn")]
        public Input<Inputs.ServicePrincipalSamlSingleSignOnGetArgs>? SamlSingleSignOn { get; set; }

        [Input("servicePrincipalNames")]
        private InputList<string>? _servicePrincipalNames;

        /// <summary>
        /// A list of identifier URI(s), copied over from the associated application.
        /// </summary>
        public InputList<string> ServicePrincipalNames
        {
            get => _servicePrincipalNames ?? (_servicePrincipalNames = new InputList<string>());
            set => _servicePrincipalNames = value;
        }

        /// <summary>
        /// The Microsoft account types that are supported for the associated application. Possible values include `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`.
        /// </summary>
        [Input("signInAudience")]
        public Input<string>? SignInAudience { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A set of tags to apply to the service principal.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are `User` or `Admin`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// When true, any existing service principal linked to the same application will be automatically imported. When false, an import error will be raised for any pre-existing service principal.
        /// </summary>
        [Input("useExisting")]
        public Input<bool>? UseExisting { get; set; }

        public ServicePrincipalState()
        {
        }
    }
}
