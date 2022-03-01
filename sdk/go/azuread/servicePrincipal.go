// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuread

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a service principal associated with an application within Azure Active Directory.
//
// ## API Permissions
//
// The following API permissions are required in order to use this resource.
//
// When authenticated with a service principal, this resource requires one of the following application roles: `Application.ReadWrite.All` or `Directory.ReadWrite.All`
//
// It is not currently possible to manage service principals whilst having only the `Application.ReadWrite.OwnedBy` role granted.
//
// When authenticated with a user principal, this resource requires one of the following directory roles: `Application Administrator` or `Global Administrator`
//
// ## Example Usage
//
// *Create a service principal for an application*
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
// 		current, err := azuread.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleApplication, err := azuread.NewApplication(ctx, "exampleApplication", &azuread.ApplicationArgs{
// 			DisplayName: pulumi.String("example"),
// 			Owners: pulumi.StringArray{
// 				pulumi.String(current.ObjectId),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = azuread.NewServicePrincipal(ctx, "exampleServicePrincipal", &azuread.ServicePrincipalArgs{
// 			ApplicationId:             exampleApplication.ApplicationId,
// 			AppRoleAssignmentRequired: pulumi.Bool(false),
// 			Owners: pulumi.StringArray{
// 				pulumi.String(current.ObjectId),
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
// *Create a service principal for an enterprise application*
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
// 		current, err := azuread.GetClientConfig(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleApplication, err := azuread.NewApplication(ctx, "exampleApplication", &azuread.ApplicationArgs{
// 			DisplayName: pulumi.String("example"),
// 			Owners: pulumi.StringArray{
// 				pulumi.String(current.ObjectId),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = azuread.NewServicePrincipal(ctx, "exampleServicePrincipal", &azuread.ServicePrincipalArgs{
// 			ApplicationId:             exampleApplication.ApplicationId,
// 			AppRoleAssignmentRequired: pulumi.Bool(false),
// 			Owners: pulumi.StringArray{
// 				pulumi.String(current.ObjectId),
// 			},
// 			FeatureTags: ServicePrincipalFeatureTagArray{
// 				&ServicePrincipalFeatureTagArgs{
// 					Enterprise: pulumi.Bool(true),
// 					Gallery:    pulumi.Bool(true),
// 				},
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
// *Manage a service principal for a first-party Microsoft application*
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
// 		wellKnown, err := azuread.GetApplicationPublishedAppIds(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = azuread.NewServicePrincipal(ctx, "msgraph", &azuread.ServicePrincipalArgs{
// 			ApplicationId: pulumi.String(wellKnown.Result.MicrosoftGraph),
// 			UseExisting:   pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// *Create a service principal for an application created from a gallery template*
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
// 		exampleApplicationTemplate, err := azuread.GetApplicationTemplate(ctx, &GetApplicationTemplateArgs{
// 			DisplayName: pulumi.StringRef("Marketo"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleApplication, err := azuread.NewApplication(ctx, "exampleApplication", &azuread.ApplicationArgs{
// 			DisplayName: pulumi.String("example"),
// 			TemplateId:  pulumi.String(exampleApplicationTemplate.TemplateId),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = azuread.NewServicePrincipal(ctx, "exampleServicePrincipal", &azuread.ServicePrincipalArgs{
// 			ApplicationId: exampleApplication.ApplicationId,
// 			UseExisting:   pulumi.Bool(true),
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
// Service principals can be imported using their object ID, e.g.
//
// ```sh
//  $ pulumi import azuread:index/servicePrincipal:ServicePrincipal test 00000000-0000-0000-0000-000000000000
// ```
type ServicePrincipal struct {
	pulumi.CustomResourceState

	// Whether or not the service principal account is enabled. Defaults to `true`.
	AccountEnabled pulumi.BoolPtrOutput `pulumi:"accountEnabled"`
	// A set of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.
	AlternativeNames pulumi.StringArrayOutput `pulumi:"alternativeNames"`
	// Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application. Defaults to `false`.
	AppRoleAssignmentRequired pulumi.BoolPtrOutput `pulumi:"appRoleAssignmentRequired"`
	// A mapping of app role values to app role IDs, as published by the associated application, intended to be useful when referencing app roles in other resources in your configuration.
	AppRoleIds pulumi.StringMapOutput `pulumi:"appRoleIds"`
	// A list of app roles published by the associated application, as documented below. For more information [official documentation](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
	AppRoles ServicePrincipalAppRoleArrayOutput `pulumi:"appRoles"`
	// The application ID (client ID) of the application for which to create a service principal.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// The tenant ID where the associated application is registered.
	ApplicationTenantId pulumi.StringOutput `pulumi:"applicationTenantId"`
	// A description of the service principal provided for internal end-users.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Display name for the app role that appears during app role assignment and in consent experiences.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// A `featureTags` block as described below. Cannot be used together with the `tags` property.
	FeatureTags ServicePrincipalFeatureTagArrayOutput `pulumi:"featureTags"`
	// Block of features to configure for this service principal using tags
	//
	// Deprecated: This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider
	Features ServicePrincipalFeatureArrayOutput `pulumi:"features"`
	// Home page or landing page of the associated application.
	HomepageUrl pulumi.StringOutput `pulumi:"homepageUrl"`
	// The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on.
	LoginUrl pulumi.StringPtrOutput `pulumi:"loginUrl"`
	// The URL that will be used by Microsoft's authorization service to log out an user using OpenId Connect front-channel, back-channel or SAML logout protocols, taken from the associated application.
	LogoutUrl pulumi.StringOutput `pulumi:"logoutUrl"`
	// A free text field to capture information about the service principal, typically used for operational purposes.
	Notes pulumi.StringPtrOutput `pulumi:"notes"`
	// A set of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
	NotificationEmailAddresses pulumi.StringArrayOutput `pulumi:"notificationEmailAddresses"`
	// A mapping of OAuth2.0 permission scope values to scope IDs, as exposed by the associated application, intended to be useful when referencing permission scopes in other resources in your configuration.
	Oauth2PermissionScopeIds pulumi.StringMapOutput `pulumi:"oauth2PermissionScopeIds"`
	// A list of OAuth 2.0 delegated permission scopes exposed by the associated application, as documented below.
	Oauth2PermissionScopes ServicePrincipalOauth2PermissionScopeArrayOutput `pulumi:"oauth2PermissionScopes"`
	// The object ID of the service principal.
	ObjectId pulumi.StringOutput `pulumi:"objectId"`
	// A set of object IDs of principals that will be granted ownership of the service principal. Supported object types are users or service principals. By default, no owners are assigned.
	Owners pulumi.StringArrayOutput `pulumi:"owners"`
	// The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps. Supported values are `oidc`, `password`, `saml` or `notSupported`. Omit this property or specify a blank string to unset.
	PreferredSingleSignOnMode pulumi.StringPtrOutput `pulumi:"preferredSingleSignOnMode"`
	// A list of URLs where user tokens are sent for sign-in with the associated application, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent for the associated application.
	RedirectUris pulumi.StringArrayOutput `pulumi:"redirectUris"`
	// The URL where the service exposes SAML metadata for federation.
	SamlMetadataUrl pulumi.StringOutput `pulumi:"samlMetadataUrl"`
	// A `samlSingleSignOn` block as documented below.
	SamlSingleSignOn ServicePrincipalSamlSingleSignOnPtrOutput `pulumi:"samlSingleSignOn"`
	// A list of identifier URI(s), copied over from the associated application.
	ServicePrincipalNames pulumi.StringArrayOutput `pulumi:"servicePrincipalNames"`
	// The Microsoft account types that are supported for the associated application. Possible values include `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`.
	SignInAudience pulumi.StringOutput `pulumi:"signInAudience"`
	// A set of tags to apply to the service principal. Cannot be used together with the `featureTags` block.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are `User` or `Admin`.
	Type pulumi.StringOutput `pulumi:"type"`
	// When true, any existing service principal linked to the same application will be automatically imported. When false, an import error will be raised for any pre-existing service principal.
	UseExisting pulumi.BoolPtrOutput `pulumi:"useExisting"`
}

// NewServicePrincipal registers a new resource with the given unique name, arguments, and options.
func NewServicePrincipal(ctx *pulumi.Context,
	name string, args *ServicePrincipalArgs, opts ...pulumi.ResourceOption) (*ServicePrincipal, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	var resource ServicePrincipal
	err := ctx.RegisterResource("azuread:index/servicePrincipal:ServicePrincipal", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServicePrincipal gets an existing ServicePrincipal resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServicePrincipal(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServicePrincipalState, opts ...pulumi.ResourceOption) (*ServicePrincipal, error) {
	var resource ServicePrincipal
	err := ctx.ReadResource("azuread:index/servicePrincipal:ServicePrincipal", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServicePrincipal resources.
type servicePrincipalState struct {
	// Whether or not the service principal account is enabled. Defaults to `true`.
	AccountEnabled *bool `pulumi:"accountEnabled"`
	// A set of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.
	AlternativeNames []string `pulumi:"alternativeNames"`
	// Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application. Defaults to `false`.
	AppRoleAssignmentRequired *bool `pulumi:"appRoleAssignmentRequired"`
	// A mapping of app role values to app role IDs, as published by the associated application, intended to be useful when referencing app roles in other resources in your configuration.
	AppRoleIds map[string]string `pulumi:"appRoleIds"`
	// A list of app roles published by the associated application, as documented below. For more information [official documentation](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
	AppRoles []ServicePrincipalAppRole `pulumi:"appRoles"`
	// The application ID (client ID) of the application for which to create a service principal.
	ApplicationId *string `pulumi:"applicationId"`
	// The tenant ID where the associated application is registered.
	ApplicationTenantId *string `pulumi:"applicationTenantId"`
	// A description of the service principal provided for internal end-users.
	Description *string `pulumi:"description"`
	// Display name for the app role that appears during app role assignment and in consent experiences.
	DisplayName *string `pulumi:"displayName"`
	// A `featureTags` block as described below. Cannot be used together with the `tags` property.
	FeatureTags []ServicePrincipalFeatureTag `pulumi:"featureTags"`
	// Block of features to configure for this service principal using tags
	//
	// Deprecated: This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider
	Features []ServicePrincipalFeature `pulumi:"features"`
	// Home page or landing page of the associated application.
	HomepageUrl *string `pulumi:"homepageUrl"`
	// The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on.
	LoginUrl *string `pulumi:"loginUrl"`
	// The URL that will be used by Microsoft's authorization service to log out an user using OpenId Connect front-channel, back-channel or SAML logout protocols, taken from the associated application.
	LogoutUrl *string `pulumi:"logoutUrl"`
	// A free text field to capture information about the service principal, typically used for operational purposes.
	Notes *string `pulumi:"notes"`
	// A set of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
	NotificationEmailAddresses []string `pulumi:"notificationEmailAddresses"`
	// A mapping of OAuth2.0 permission scope values to scope IDs, as exposed by the associated application, intended to be useful when referencing permission scopes in other resources in your configuration.
	Oauth2PermissionScopeIds map[string]string `pulumi:"oauth2PermissionScopeIds"`
	// A list of OAuth 2.0 delegated permission scopes exposed by the associated application, as documented below.
	Oauth2PermissionScopes []ServicePrincipalOauth2PermissionScope `pulumi:"oauth2PermissionScopes"`
	// The object ID of the service principal.
	ObjectId *string `pulumi:"objectId"`
	// A set of object IDs of principals that will be granted ownership of the service principal. Supported object types are users or service principals. By default, no owners are assigned.
	Owners []string `pulumi:"owners"`
	// The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps. Supported values are `oidc`, `password`, `saml` or `notSupported`. Omit this property or specify a blank string to unset.
	PreferredSingleSignOnMode *string `pulumi:"preferredSingleSignOnMode"`
	// A list of URLs where user tokens are sent for sign-in with the associated application, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent for the associated application.
	RedirectUris []string `pulumi:"redirectUris"`
	// The URL where the service exposes SAML metadata for federation.
	SamlMetadataUrl *string `pulumi:"samlMetadataUrl"`
	// A `samlSingleSignOn` block as documented below.
	SamlSingleSignOn *ServicePrincipalSamlSingleSignOn `pulumi:"samlSingleSignOn"`
	// A list of identifier URI(s), copied over from the associated application.
	ServicePrincipalNames []string `pulumi:"servicePrincipalNames"`
	// The Microsoft account types that are supported for the associated application. Possible values include `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`.
	SignInAudience *string `pulumi:"signInAudience"`
	// A set of tags to apply to the service principal. Cannot be used together with the `featureTags` block.
	Tags []string `pulumi:"tags"`
	// Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are `User` or `Admin`.
	Type *string `pulumi:"type"`
	// When true, any existing service principal linked to the same application will be automatically imported. When false, an import error will be raised for any pre-existing service principal.
	UseExisting *bool `pulumi:"useExisting"`
}

type ServicePrincipalState struct {
	// Whether or not the service principal account is enabled. Defaults to `true`.
	AccountEnabled pulumi.BoolPtrInput
	// A set of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.
	AlternativeNames pulumi.StringArrayInput
	// Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application. Defaults to `false`.
	AppRoleAssignmentRequired pulumi.BoolPtrInput
	// A mapping of app role values to app role IDs, as published by the associated application, intended to be useful when referencing app roles in other resources in your configuration.
	AppRoleIds pulumi.StringMapInput
	// A list of app roles published by the associated application, as documented below. For more information [official documentation](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
	AppRoles ServicePrincipalAppRoleArrayInput
	// The application ID (client ID) of the application for which to create a service principal.
	ApplicationId pulumi.StringPtrInput
	// The tenant ID where the associated application is registered.
	ApplicationTenantId pulumi.StringPtrInput
	// A description of the service principal provided for internal end-users.
	Description pulumi.StringPtrInput
	// Display name for the app role that appears during app role assignment and in consent experiences.
	DisplayName pulumi.StringPtrInput
	// A `featureTags` block as described below. Cannot be used together with the `tags` property.
	FeatureTags ServicePrincipalFeatureTagArrayInput
	// Block of features to configure for this service principal using tags
	//
	// Deprecated: This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider
	Features ServicePrincipalFeatureArrayInput
	// Home page or landing page of the associated application.
	HomepageUrl pulumi.StringPtrInput
	// The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on.
	LoginUrl pulumi.StringPtrInput
	// The URL that will be used by Microsoft's authorization service to log out an user using OpenId Connect front-channel, back-channel or SAML logout protocols, taken from the associated application.
	LogoutUrl pulumi.StringPtrInput
	// A free text field to capture information about the service principal, typically used for operational purposes.
	Notes pulumi.StringPtrInput
	// A set of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
	NotificationEmailAddresses pulumi.StringArrayInput
	// A mapping of OAuth2.0 permission scope values to scope IDs, as exposed by the associated application, intended to be useful when referencing permission scopes in other resources in your configuration.
	Oauth2PermissionScopeIds pulumi.StringMapInput
	// A list of OAuth 2.0 delegated permission scopes exposed by the associated application, as documented below.
	Oauth2PermissionScopes ServicePrincipalOauth2PermissionScopeArrayInput
	// The object ID of the service principal.
	ObjectId pulumi.StringPtrInput
	// A set of object IDs of principals that will be granted ownership of the service principal. Supported object types are users or service principals. By default, no owners are assigned.
	Owners pulumi.StringArrayInput
	// The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps. Supported values are `oidc`, `password`, `saml` or `notSupported`. Omit this property or specify a blank string to unset.
	PreferredSingleSignOnMode pulumi.StringPtrInput
	// A list of URLs where user tokens are sent for sign-in with the associated application, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent for the associated application.
	RedirectUris pulumi.StringArrayInput
	// The URL where the service exposes SAML metadata for federation.
	SamlMetadataUrl pulumi.StringPtrInput
	// A `samlSingleSignOn` block as documented below.
	SamlSingleSignOn ServicePrincipalSamlSingleSignOnPtrInput
	// A list of identifier URI(s), copied over from the associated application.
	ServicePrincipalNames pulumi.StringArrayInput
	// The Microsoft account types that are supported for the associated application. Possible values include `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`.
	SignInAudience pulumi.StringPtrInput
	// A set of tags to apply to the service principal. Cannot be used together with the `featureTags` block.
	Tags pulumi.StringArrayInput
	// Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are `User` or `Admin`.
	Type pulumi.StringPtrInput
	// When true, any existing service principal linked to the same application will be automatically imported. When false, an import error will be raised for any pre-existing service principal.
	UseExisting pulumi.BoolPtrInput
}

func (ServicePrincipalState) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePrincipalState)(nil)).Elem()
}

type servicePrincipalArgs struct {
	// Whether or not the service principal account is enabled. Defaults to `true`.
	AccountEnabled *bool `pulumi:"accountEnabled"`
	// A set of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.
	AlternativeNames []string `pulumi:"alternativeNames"`
	// Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application. Defaults to `false`.
	AppRoleAssignmentRequired *bool `pulumi:"appRoleAssignmentRequired"`
	// The application ID (client ID) of the application for which to create a service principal.
	ApplicationId string `pulumi:"applicationId"`
	// A description of the service principal provided for internal end-users.
	Description *string `pulumi:"description"`
	// A `featureTags` block as described below. Cannot be used together with the `tags` property.
	FeatureTags []ServicePrincipalFeatureTag `pulumi:"featureTags"`
	// Block of features to configure for this service principal using tags
	//
	// Deprecated: This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider
	Features []ServicePrincipalFeature `pulumi:"features"`
	// The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on.
	LoginUrl *string `pulumi:"loginUrl"`
	// A free text field to capture information about the service principal, typically used for operational purposes.
	Notes *string `pulumi:"notes"`
	// A set of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
	NotificationEmailAddresses []string `pulumi:"notificationEmailAddresses"`
	// A set of object IDs of principals that will be granted ownership of the service principal. Supported object types are users or service principals. By default, no owners are assigned.
	Owners []string `pulumi:"owners"`
	// The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps. Supported values are `oidc`, `password`, `saml` or `notSupported`. Omit this property or specify a blank string to unset.
	PreferredSingleSignOnMode *string `pulumi:"preferredSingleSignOnMode"`
	// A `samlSingleSignOn` block as documented below.
	SamlSingleSignOn *ServicePrincipalSamlSingleSignOn `pulumi:"samlSingleSignOn"`
	// A set of tags to apply to the service principal. Cannot be used together with the `featureTags` block.
	Tags []string `pulumi:"tags"`
	// When true, any existing service principal linked to the same application will be automatically imported. When false, an import error will be raised for any pre-existing service principal.
	UseExisting *bool `pulumi:"useExisting"`
}

// The set of arguments for constructing a ServicePrincipal resource.
type ServicePrincipalArgs struct {
	// Whether or not the service principal account is enabled. Defaults to `true`.
	AccountEnabled pulumi.BoolPtrInput
	// A set of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.
	AlternativeNames pulumi.StringArrayInput
	// Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application. Defaults to `false`.
	AppRoleAssignmentRequired pulumi.BoolPtrInput
	// The application ID (client ID) of the application for which to create a service principal.
	ApplicationId pulumi.StringInput
	// A description of the service principal provided for internal end-users.
	Description pulumi.StringPtrInput
	// A `featureTags` block as described below. Cannot be used together with the `tags` property.
	FeatureTags ServicePrincipalFeatureTagArrayInput
	// Block of features to configure for this service principal using tags
	//
	// Deprecated: This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider
	Features ServicePrincipalFeatureArrayInput
	// The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps. When blank, Azure AD performs IdP-initiated sign-on for applications configured with SAML-based single sign-on.
	LoginUrl pulumi.StringPtrInput
	// A free text field to capture information about the service principal, typically used for operational purposes.
	Notes pulumi.StringPtrInput
	// A set of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
	NotificationEmailAddresses pulumi.StringArrayInput
	// A set of object IDs of principals that will be granted ownership of the service principal. Supported object types are users or service principals. By default, no owners are assigned.
	Owners pulumi.StringArrayInput
	// The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps. Supported values are `oidc`, `password`, `saml` or `notSupported`. Omit this property or specify a blank string to unset.
	PreferredSingleSignOnMode pulumi.StringPtrInput
	// A `samlSingleSignOn` block as documented below.
	SamlSingleSignOn ServicePrincipalSamlSingleSignOnPtrInput
	// A set of tags to apply to the service principal. Cannot be used together with the `featureTags` block.
	Tags pulumi.StringArrayInput
	// When true, any existing service principal linked to the same application will be automatically imported. When false, an import error will be raised for any pre-existing service principal.
	UseExisting pulumi.BoolPtrInput
}

func (ServicePrincipalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*servicePrincipalArgs)(nil)).Elem()
}

type ServicePrincipalInput interface {
	pulumi.Input

	ToServicePrincipalOutput() ServicePrincipalOutput
	ToServicePrincipalOutputWithContext(ctx context.Context) ServicePrincipalOutput
}

func (*ServicePrincipal) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipal)(nil)).Elem()
}

func (i *ServicePrincipal) ToServicePrincipalOutput() ServicePrincipalOutput {
	return i.ToServicePrincipalOutputWithContext(context.Background())
}

func (i *ServicePrincipal) ToServicePrincipalOutputWithContext(ctx context.Context) ServicePrincipalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalOutput)
}

// ServicePrincipalArrayInput is an input type that accepts ServicePrincipalArray and ServicePrincipalArrayOutput values.
// You can construct a concrete instance of `ServicePrincipalArrayInput` via:
//
//          ServicePrincipalArray{ ServicePrincipalArgs{...} }
type ServicePrincipalArrayInput interface {
	pulumi.Input

	ToServicePrincipalArrayOutput() ServicePrincipalArrayOutput
	ToServicePrincipalArrayOutputWithContext(context.Context) ServicePrincipalArrayOutput
}

type ServicePrincipalArray []ServicePrincipalInput

func (ServicePrincipalArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServicePrincipal)(nil)).Elem()
}

func (i ServicePrincipalArray) ToServicePrincipalArrayOutput() ServicePrincipalArrayOutput {
	return i.ToServicePrincipalArrayOutputWithContext(context.Background())
}

func (i ServicePrincipalArray) ToServicePrincipalArrayOutputWithContext(ctx context.Context) ServicePrincipalArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalArrayOutput)
}

// ServicePrincipalMapInput is an input type that accepts ServicePrincipalMap and ServicePrincipalMapOutput values.
// You can construct a concrete instance of `ServicePrincipalMapInput` via:
//
//          ServicePrincipalMap{ "key": ServicePrincipalArgs{...} }
type ServicePrincipalMapInput interface {
	pulumi.Input

	ToServicePrincipalMapOutput() ServicePrincipalMapOutput
	ToServicePrincipalMapOutputWithContext(context.Context) ServicePrincipalMapOutput
}

type ServicePrincipalMap map[string]ServicePrincipalInput

func (ServicePrincipalMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServicePrincipal)(nil)).Elem()
}

func (i ServicePrincipalMap) ToServicePrincipalMapOutput() ServicePrincipalMapOutput {
	return i.ToServicePrincipalMapOutputWithContext(context.Background())
}

func (i ServicePrincipalMap) ToServicePrincipalMapOutputWithContext(ctx context.Context) ServicePrincipalMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicePrincipalMapOutput)
}

type ServicePrincipalOutput struct{ *pulumi.OutputState }

func (ServicePrincipalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePrincipal)(nil)).Elem()
}

func (o ServicePrincipalOutput) ToServicePrincipalOutput() ServicePrincipalOutput {
	return o
}

func (o ServicePrincipalOutput) ToServicePrincipalOutputWithContext(ctx context.Context) ServicePrincipalOutput {
	return o
}

type ServicePrincipalArrayOutput struct{ *pulumi.OutputState }

func (ServicePrincipalArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServicePrincipal)(nil)).Elem()
}

func (o ServicePrincipalArrayOutput) ToServicePrincipalArrayOutput() ServicePrincipalArrayOutput {
	return o
}

func (o ServicePrincipalArrayOutput) ToServicePrincipalArrayOutputWithContext(ctx context.Context) ServicePrincipalArrayOutput {
	return o
}

func (o ServicePrincipalArrayOutput) Index(i pulumi.IntInput) ServicePrincipalOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServicePrincipal {
		return vs[0].([]*ServicePrincipal)[vs[1].(int)]
	}).(ServicePrincipalOutput)
}

type ServicePrincipalMapOutput struct{ *pulumi.OutputState }

func (ServicePrincipalMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServicePrincipal)(nil)).Elem()
}

func (o ServicePrincipalMapOutput) ToServicePrincipalMapOutput() ServicePrincipalMapOutput {
	return o
}

func (o ServicePrincipalMapOutput) ToServicePrincipalMapOutputWithContext(ctx context.Context) ServicePrincipalMapOutput {
	return o
}

func (o ServicePrincipalMapOutput) MapIndex(k pulumi.StringInput) ServicePrincipalOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServicePrincipal {
		return vs[0].(map[string]*ServicePrincipal)[vs[1].(string)]
	}).(ServicePrincipalOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePrincipalInput)(nil)).Elem(), &ServicePrincipal{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePrincipalArrayInput)(nil)).Elem(), ServicePrincipalArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServicePrincipalMapInput)(nil)).Elem(), ServicePrincipalMap{})
	pulumi.RegisterOutputType(ServicePrincipalOutput{})
	pulumi.RegisterOutputType(ServicePrincipalArrayOutput{})
	pulumi.RegisterOutputType(ServicePrincipalMapOutput{})
}
