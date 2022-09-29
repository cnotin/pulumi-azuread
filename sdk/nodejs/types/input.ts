// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

export interface ApplicationApi {
    /**
     * A set of application IDs (client IDs), used for bundling consent if you have a solution that contains two parts: a client app and a custom web API app.
     */
    knownClientApplications?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Allows an application to use claims mapping without specifying a custom signing key. Defaults to `false`.
     */
    mappedClaimsEnabled?: pulumi.Input<boolean>;
    /**
     * One or more `oauth2PermissionScope` blocks as documented below, to describe delegated permissions exposed by the web API represented by this application.
     */
    oauth2PermissionScopes?: pulumi.Input<pulumi.Input<inputs.ApplicationApiOauth2PermissionScope>[]>;
    /**
     * The access token version expected by this resource. Must be one of `1` or `2`, and must be `2` when `signInAudience` is either `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount` Defaults to `1`.
     */
    requestedAccessTokenVersion?: pulumi.Input<number>;
}

export interface ApplicationApiOauth2PermissionScope {
    /**
     * Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
     */
    adminConsentDescription?: pulumi.Input<string>;
    /**
     * Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
     */
    adminConsentDisplayName?: pulumi.Input<string>;
    /**
     * Determines if the permission scope is enabled. Defaults to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The unique identifier of the delegated permission. Must be a valid UUID.
     */
    id: pulumi.Input<string>;
    /**
     * Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Defaults to `User`. Possible values are `User` or `Admin`.
     */
    type?: pulumi.Input<string>;
    /**
     * Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
     */
    userConsentDescription?: pulumi.Input<string>;
    /**
     * Display name for the delegated permission that appears in the end user consent experience.
     */
    userConsentDisplayName?: pulumi.Input<string>;
    /**
     * The value that is used for the `scp` claim in OAuth 2.0 access tokens.
     */
    value?: pulumi.Input<string>;
}

export interface ApplicationAppRole {
    /**
     * Specifies whether this app role definition can be assigned to users and groups by setting to `User`, or to other applications (that are accessing this application in a standalone scenario) by setting to `Application`, or to both.
     */
    allowedMemberTypes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Description of the app role that appears when the role is being assigned and, if the role functions as an application permissions, during the consent experiences.
     */
    description: pulumi.Input<string>;
    /**
     * Display name for the app role that appears during app role assignment and in consent experiences.
     */
    displayName: pulumi.Input<string>;
    /**
     * Determines if the app role is enabled. Defaults to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The unique identifier of the app role. Must be a valid UUID.
     */
    id: pulumi.Input<string>;
    /**
     * The value that is used for the `roles` claim in ID tokens and OAuth 2.0 access tokens that are authenticating an assigned service or user principal.
     */
    value?: pulumi.Input<string>;
}

export interface ApplicationFeatureTag {
    /**
     * Whether this application represents a custom SAML application for linked service principals. Enabling this will assign the `WindowsAzureActiveDirectoryCustomSingleSignOnApplication` tag. Defaults to `false`.
     */
    customSingleSignOn?: pulumi.Input<boolean>;
    /**
     * Whether this application represents an Enterprise Application for linked service principals. Enabling this will assign the `WindowsAzureActiveDirectoryIntegratedApp` tag. Defaults to `false`.
     */
    enterprise?: pulumi.Input<boolean>;
    /**
     * Whether this application represents a gallery application for linked service principals. Enabling this will assign the `WindowsAzureActiveDirectoryGalleryApplicationNonPrimaryV1` tag. Defaults to `false`.
     */
    gallery?: pulumi.Input<boolean>;
    /**
     * Whether this app is invisible to users in My Apps and Office 365 Launcher. Enabling this will assign the `HideApp` tag. Defaults to `false`.
     */
    hide?: pulumi.Input<boolean>;
}

export interface ApplicationOptionalClaims {
    /**
     * One or more `accessToken` blocks as documented below.
     */
    accessTokens?: pulumi.Input<pulumi.Input<inputs.ApplicationOptionalClaimsAccessToken>[]>;
    /**
     * One or more `idToken` blocks as documented below.
     */
    idTokens?: pulumi.Input<pulumi.Input<inputs.ApplicationOptionalClaimsIdToken>[]>;
    /**
     * One or more `saml2Token` blocks as documented below.
     */
    saml2Tokens?: pulumi.Input<pulumi.Input<inputs.ApplicationOptionalClaimsSaml2Token>[]>;
}

export interface ApplicationOptionalClaimsAccessToken {
    /**
     * List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim.
     */
    additionalProperties?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether the claim specified by the client is necessary to ensure a smooth authorization experience.
     */
    essential?: pulumi.Input<boolean>;
    /**
     * The name of the optional claim.
     */
    name: pulumi.Input<string>;
    /**
     * The source of the claim. If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object.
     */
    source?: pulumi.Input<string>;
}

export interface ApplicationOptionalClaimsIdToken {
    /**
     * List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim.
     */
    additionalProperties?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether the claim specified by the client is necessary to ensure a smooth authorization experience.
     */
    essential?: pulumi.Input<boolean>;
    /**
     * The name of the optional claim.
     */
    name: pulumi.Input<string>;
    /**
     * The source of the claim. If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object.
     */
    source?: pulumi.Input<string>;
}

export interface ApplicationOptionalClaimsSaml2Token {
    /**
     * List of additional properties of the claim. If a property exists in this list, it modifies the behaviour of the optional claim.
     */
    additionalProperties?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether the claim specified by the client is necessary to ensure a smooth authorization experience.
     */
    essential?: pulumi.Input<boolean>;
    /**
     * The name of the optional claim.
     */
    name: pulumi.Input<string>;
    /**
     * The source of the claim. If `source` is absent, the claim is a predefined optional claim. If `source` is `user`, the value of `name` is the extension property from the user object.
     */
    source?: pulumi.Input<string>;
}

export interface ApplicationPublicClient {
    /**
     * A set of URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent. Must be a valid `https` or `ms-appx-web` URL.
     */
    redirectUris?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ApplicationRequiredResourceAccess {
    /**
     * A collection of `resourceAccess` blocks as documented below, describing OAuth2.0 permission scopes and app roles that the application requires from the specified resource.
     */
    resourceAccesses: pulumi.Input<pulumi.Input<inputs.ApplicationRequiredResourceAccessResourceAccess>[]>;
    /**
     * The unique identifier for the resource that the application requires access to. This should be the Application ID of the target application.
     */
    resourceAppId: pulumi.Input<string>;
}

export interface ApplicationRequiredResourceAccessResourceAccess {
    /**
     * The unique identifier for an app role or OAuth2 permission scope published by the resource application.
     */
    id: pulumi.Input<string>;
    /**
     * Specifies whether the `id` property references an app role or an OAuth2 permission scope. Possible values are `Role` or `Scope`.
     */
    type: pulumi.Input<string>;
}

export interface ApplicationSinglePageApplication {
    /**
     * A set of URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent. Must be a valid `https` URL.
     */
    redirectUris?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ApplicationWeb {
    /**
     * Home page or landing page of the application.
     */
    homepageUrl?: pulumi.Input<string>;
    /**
     * An `implicitGrant` block as documented above.
     */
    implicitGrant?: pulumi.Input<inputs.ApplicationWebImplicitGrant>;
    /**
     * The URL that will be used by Microsoft's authorization service to sign out a user using front-channel, back-channel or SAML logout protocols.
     */
    logoutUrl?: pulumi.Input<string>;
    /**
     * A set of URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent. Must be a valid `http` URL or a URN.
     */
    redirectUris?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ApplicationWebImplicitGrant {
    /**
     * Whether this web application can request an access token using OAuth 2.0 implicit flow.
     */
    accessTokenIssuanceEnabled?: pulumi.Input<boolean>;
    /**
     * Whether this web application can request an ID token using OAuth 2.0 implicit flow.
     */
    idTokenIssuanceEnabled?: pulumi.Input<boolean>;
}

export interface ConditionalAccessPolicyConditions {
    /**
     * An `applications` block as documented below, which specifies applications and user actions included in and excluded from the policy.
     */
    applications: pulumi.Input<inputs.ConditionalAccessPolicyConditionsApplications>;
    /**
     * A list of client application types included in the policy. Possible values are: `all`, `browser`, `mobileAppsAndDesktopClients`, `exchangeActiveSync`, `easSupported` and `other`.
     */
    clientAppTypes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A `devices` block as documented below, which describes devices to be included in and excluded from the policy. A `devices` block can be added to an existing policy, but removing the `devices` block forces a new resource to be created.
     */
    devices?: pulumi.Input<inputs.ConditionalAccessPolicyConditionsDevices>;
    /**
     * A `locations` block as documented below, which specifies locations included in and excluded from the policy.
     */
    locations?: pulumi.Input<inputs.ConditionalAccessPolicyConditionsLocations>;
    /**
     * A `platforms` block as documented below, which specifies platforms included in and excluded from the policy.
     */
    platforms?: pulumi.Input<inputs.ConditionalAccessPolicyConditionsPlatforms>;
    /**
     * A list of sign-in risk levels included in the policy. Possible values are: `low`, `medium`, `high`, `hidden`, `none`, `unknownFutureValue`.
     */
    signInRiskLevels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of user risk levels included in the policy. Possible values are: `low`, `medium`, `high`, `hidden`, `none`, `unknownFutureValue`.
     */
    userRiskLevels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A `users` block as documented below, which specifies users, groups, and roles included in and excluded from the policy.
     */
    users: pulumi.Input<inputs.ConditionalAccessPolicyConditionsUsers>;
}

export interface ConditionalAccessPolicyConditionsApplications {
    /**
     * A list of application IDs explicitly excluded from the policy. Can also be set to `Office365`.
     */
    excludedApplications?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of application IDs the policy applies to, unless explicitly excluded (in `excludedApplications`). Can also be set to `All` or `Office365`. Cannot be specified with `includedUserActions`. One of `includedApplications` or `includedUserActions` must be specified.
     */
    includedApplications?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of user actions to include. Supported values are `urn:user:registerdevice` and `urn:user:registersecurityinfo`. Cannot be specified with `includedApplications`. One of `includedApplications` or `includedUserActions` must be specified.
     */
    includedUserActions?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ConditionalAccessPolicyConditionsDevices {
    /**
     * A `filter` block as described below. A `filter` block can be added to an existing policy, but removing the `filter` block forces a new resource to be created.
     */
    filter?: pulumi.Input<inputs.ConditionalAccessPolicyConditionsDevicesFilter>;
}

export interface ConditionalAccessPolicyConditionsDevicesFilter {
    /**
     * Whether to include in, or exclude from, matching devices from the policy. Supported values are `include` or `exclude`.
     */
    mode: pulumi.Input<string>;
    /**
     * Condition filter to match devices. For more information, see [official documentation](https://docs.microsoft.com/en-us/azure/active-directory/conditional-access/concept-condition-filters-for-devices#supported-operators-and-device-properties-for-filters).
     */
    rule: pulumi.Input<string>;
}

export interface ConditionalAccessPolicyConditionsLocations {
    /**
     * A list of location IDs excluded from scope of policy. Can also be set to `AllTrusted`.
     */
    excludedLocations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of location IDs in scope of policy unless explicitly excluded. Can also be set to `All`, or `AllTrusted`.
     */
    includedLocations: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ConditionalAccessPolicyConditionsPlatforms {
    /**
     * A list of platforms explicitly excluded from the policy. Possible values are: `all`, `android`, `iOS`, `linux`, `macOS`, `windows`, `windowsPhone` or `unknownFutureValue`.
     */
    excludedPlatforms?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of platforms the policy applies to, unless explicitly excluded. Possible values are: `all`, `android`, `iOS`, `linux`, `macOS`, `windows`, `windowsPhone` or `unknownFutureValue`.
     */
    includedPlatforms: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ConditionalAccessPolicyConditionsUsers {
    /**
     * A list of group IDs excluded from scope of policy.
     */
    excludedGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of role IDs excluded from scope of policy.
     */
    excludedRoles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of user IDs excluded from scope of policy and/or `GuestsOrExternalUsers`.
     */
    excludedUsers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of group IDs in scope of policy unless explicitly excluded.
     */
    includedGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of role IDs in scope of policy unless explicitly excluded.
     */
    includedRoles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of user IDs in scope of policy unless explicitly excluded, or `None` or `All` or `GuestsOrExternalUsers`.
     */
    includedUsers?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ConditionalAccessPolicyGrantControls {
    /**
     * List of built-in controls required by the policy. Possible values are: `block`, `mfa`, `approvedApplication`, `compliantApplication`, `compliantDevice`, `domainJoinedDevice`, `passwordChange` or `unknownFutureValue`.
     */
    builtInControls: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of custom controls IDs required by the policy.
     */
    customAuthenticationFactors?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Defines the relationship of the grant controls. Possible values are: `AND`, `OR`.
     */
    operator: pulumi.Input<string>;
    /**
     * List of terms of use IDs required by the policy.
     */
    termsOfUses?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ConditionalAccessPolicySessionControls {
    /**
     * Whether or not application enforced restrictions are enabled. Defaults to `false`.
     */
    applicationEnforcedRestrictionsEnabled?: pulumi.Input<boolean>;
    /**
     * Enables cloud app security and specifies the cloud app security policy to use. Possible values are: `blockDownloads`, `mcasConfigured`, `monitorOnly` or `unknownFutureValue`.
     */
    cloudAppSecurityPolicy?: pulumi.Input<string>;
    /**
     * Session control to define whether to persist cookies or not. Possible values are: `always` or `never`.
     */
    persistentBrowserMode?: pulumi.Input<string>;
    /**
     * Number of days or hours to enforce sign-in frequency. Required when `signInFrequencyPeriod` is specified. Due to an API issue, removing this property forces a new resource to be created.
     */
    signInFrequency?: pulumi.Input<number>;
    /**
     * The time period to enforce sign-in frequency. Possible values are: `hours` or `days`. Required when `signInFrequencyPeriod` is specified. Due to an API issue, removing this property forces a new resource to be created.
     */
    signInFrequencyPeriod?: pulumi.Input<string>;
}

export interface CustomDirectoryRolePermission {
    /**
     * A set of tasks that can be performed on a resource. For more information, see the [Permissions Reference](https://docs.microsoft.com/en-us/azure/active-directory/roles/permissions-reference) documentation.
     */
    allowedResourceActions: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GroupDynamicMembership {
    /**
     * Whether rule processing is "On" (true) or "Paused" (false).
     */
    enabled: pulumi.Input<boolean>;
    /**
     * The rule that determines membership of this group. For more information, see official documentation on [membership rules syntax](https://docs.microsoft.com/en-gb/azure/active-directory/enterprise-users/groups-dynamic-membership).
     */
    rule: pulumi.Input<string>;
}

export interface InvitationMessage {
    /**
     * Email addresses of additional recipients the invitation message should be sent to. Only 1 additional recipient is currently supported by Azure.
     */
    additionalRecipients?: pulumi.Input<string>;
    /**
     * Customized message body you want to send if you don't want to send the default message. Cannot be specified with `language`.
     */
    body?: pulumi.Input<string>;
    /**
     * The language you want to send the default message in. The value specified must be in ISO 639 format. Defaults to `en-US`. Cannot be specified with `body`.
     */
    language?: pulumi.Input<string>;
}

export interface NamedLocationCountry {
    /**
     * List of countries and/or regions in two-letter format specified by ISO 3166-2.
     */
    countriesAndRegions: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether IP addresses that don't map to a country or region should be included in the named location. Defaults to `false`.
     */
    includeUnknownCountriesAndRegions?: pulumi.Input<boolean>;
}

export interface NamedLocationIp {
    /**
     * List of IP address ranges in IPv4 CIDR format (e.g. 1.2.3.4/32) or any allowable IPv6 format from IETF RFC596.
     */
    ipRanges: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether the named location is trusted. Defaults to `false`.
     */
    trusted?: pulumi.Input<boolean>;
}

export interface ServicePrincipalAppRole {
    /**
     * Specifies whether this app role definition can be assigned to users and groups, or to other applications (that are accessing this application in a standalone scenario). Possible values are: `User` and `Application`, or both.
     */
    allowedMemberTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A description of the service principal provided for internal end-users.
     */
    description?: pulumi.Input<string>;
    /**
     * Display name for the app role that appears during app role assignment and in consent experiences.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Specifies whether the permission scope is enabled.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The unique identifier of the delegated permission.
     */
    id?: pulumi.Input<string>;
    /**
     * The value that is used for the `scp` claim in OAuth 2.0 access tokens.
     */
    value?: pulumi.Input<string>;
}

export interface ServicePrincipalFeature {
    customSingleSignOnApp?: pulumi.Input<boolean>;
    enterpriseApplication?: pulumi.Input<boolean>;
    galleryApplication?: pulumi.Input<boolean>;
    visibleToUsers?: pulumi.Input<boolean>;
}

export interface ServicePrincipalFeatureTag {
    /**
     * Whether this service principal represents a custom SAML application. Enabling this will assign the `WindowsAzureActiveDirectoryCustomSingleSignOnApplication` tag. Defaults to `false`.
     */
    customSingleSignOn?: pulumi.Input<boolean>;
    /**
     * Whether this service principal represents an Enterprise Application. Enabling this will assign the `WindowsAzureActiveDirectoryIntegratedApp` tag. Defaults to `false`.
     */
    enterprise?: pulumi.Input<boolean>;
    /**
     * Whether this service principal represents a gallery application. Enabling this will assign the `WindowsAzureActiveDirectoryGalleryApplicationNonPrimaryV1` tag. Defaults to `false`.
     */
    gallery?: pulumi.Input<boolean>;
    /**
     * Whether this app is invisible to users in My Apps and Office 365 Launcher. Enabling this will assign the `HideApp` tag. Defaults to `false`.
     */
    hide?: pulumi.Input<boolean>;
}

export interface ServicePrincipalOauth2PermissionScope {
    /**
     * Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
     */
    adminConsentDescription?: pulumi.Input<string>;
    /**
     * Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
     */
    adminConsentDisplayName?: pulumi.Input<string>;
    /**
     * Specifies whether the permission scope is enabled.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The unique identifier of the delegated permission.
     */
    id?: pulumi.Input<string>;
    /**
     * Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are `User` or `Admin`.
     */
    type?: pulumi.Input<string>;
    /**
     * Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
     */
    userConsentDescription?: pulumi.Input<string>;
    /**
     * Display name for the delegated permission that appears in the end user consent experience.
     */
    userConsentDisplayName?: pulumi.Input<string>;
    /**
     * The value that is used for the `scp` claim in OAuth 2.0 access tokens.
     */
    value?: pulumi.Input<string>;
}

export interface ServicePrincipalSamlSingleSignOn {
    /**
     * The relative URI the service provider would redirect to after completion of the single sign-on flow.
     */
    relayState?: pulumi.Input<string>;
}

