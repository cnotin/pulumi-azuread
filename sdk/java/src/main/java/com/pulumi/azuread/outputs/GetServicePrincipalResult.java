// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.azuread.outputs.GetServicePrincipalAppRole;
import com.pulumi.azuread.outputs.GetServicePrincipalFeature;
import com.pulumi.azuread.outputs.GetServicePrincipalFeatureTag;
import com.pulumi.azuread.outputs.GetServicePrincipalOauth2PermissionScope;
import com.pulumi.azuread.outputs.GetServicePrincipalSamlSingleSignOn;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetServicePrincipalResult {
    /**
     * @return Whether or not the service principal account is enabled.
     * 
     */
    private Boolean accountEnabled;
    /**
     * @return A list of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.
     * 
     */
    private List<String> alternativeNames;
    /**
     * @return Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application.
     * 
     */
    private Boolean appRoleAssignmentRequired;
    /**
     * @return A mapping of app role values to app role IDs, as published by the associated application, intended to be useful when referencing app roles in other resources in your configuration.
     * 
     */
    private Map<String,String> appRoleIds;
    /**
     * @return A list of app roles published by the associated application, as documented below. For more information [official documentation](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
     * 
     */
    private List<GetServicePrincipalAppRole> appRoles;
    /**
     * @return The application ID (client ID) of the application associated with this service principal.
     * 
     */
    private String applicationId;
    /**
     * @return The tenant ID where the associated application is registered.
     * 
     */
    private String applicationTenantId;
    /**
     * @return Permission help text that appears in the admin app assignment and consent experiences.
     * 
     */
    private String description;
    /**
     * @return Display name for the permission that appears in the admin consent and app assignment experiences.
     * 
     */
    private String displayName;
    private List<GetServicePrincipalFeatureTag> featureTags;
    /**
     * @return A `features` block as described below.
     * 
     * @deprecated
     * This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider
     * 
     */
    @Deprecated /* This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider */
    private List<GetServicePrincipalFeature> features;
    /**
     * @return Home page or landing page of the associated application.
     * 
     */
    private String homepageUrl;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps.
     * 
     */
    private String loginUrl;
    /**
     * @return The URL that will be used by Microsoft&#39;s authorization service to logout an user using OpenId Connect front-channel, back-channel or SAML logout protocols, taken from the associated application.
     * 
     */
    private String logoutUrl;
    /**
     * @return A free text field to capture information about the service principal, typically used for operational purposes.
     * 
     */
    private String notes;
    /**
     * @return A list of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
     * 
     */
    private List<String> notificationEmailAddresses;
    /**
     * @return A mapping of OAuth2.0 permission scope values to scope IDs, as exposed by the associated application, intended to be useful when referencing permission scopes in other resources in your configuration.
     * 
     */
    private Map<String,String> oauth2PermissionScopeIds;
    /**
     * @return A collection of OAuth 2.0 delegated permissions exposed by the associated application. Each permission is covered by an `oauth2_permission_scopes` block as documented below.
     * 
     */
    private List<GetServicePrincipalOauth2PermissionScope> oauth2PermissionScopes;
    /**
     * @return The object ID of the service principal.
     * 
     */
    private String objectId;
    /**
     * @return The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps.
     * 
     */
    private String preferredSingleSignOnMode;
    /**
     * @return A list of URLs where user tokens are sent for sign-in with the associated application, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent for the associated application.
     * 
     */
    private List<String> redirectUris;
    /**
     * @return The URL where the service exposes SAML metadata for federation.
     * 
     */
    private String samlMetadataUrl;
    /**
     * @return A `saml_single_sign_on` block as documented below.
     * 
     */
    private List<GetServicePrincipalSamlSingleSignOn> samlSingleSignOns;
    /**
     * @return A list of identifier URI(s), copied over from the associated application.
     * 
     */
    private List<String> servicePrincipalNames;
    /**
     * @return The Microsoft account types that are supported for the associated application. Possible values include `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`.
     * 
     */
    private String signInAudience;
    /**
     * @return A list of tags applied to the service principal.
     * 
     */
    private List<String> tags;
    /**
     * @return Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are `User` or `Admin`.
     * 
     */
    private String type;

    private GetServicePrincipalResult() {}
    /**
     * @return Whether or not the service principal account is enabled.
     * 
     */
    public Boolean accountEnabled() {
        return this.accountEnabled;
    }
    /**
     * @return A list of alternative names, used to retrieve service principals by subscription, identify resource group and full resource ids for managed identities.
     * 
     */
    public List<String> alternativeNames() {
        return this.alternativeNames;
    }
    /**
     * @return Whether this service principal requires an app role assignment to a user or group before Azure AD will issue a user or access token to the application.
     * 
     */
    public Boolean appRoleAssignmentRequired() {
        return this.appRoleAssignmentRequired;
    }
    /**
     * @return A mapping of app role values to app role IDs, as published by the associated application, intended to be useful when referencing app roles in other resources in your configuration.
     * 
     */
    public Map<String,String> appRoleIds() {
        return this.appRoleIds;
    }
    /**
     * @return A list of app roles published by the associated application, as documented below. For more information [official documentation](https://docs.microsoft.com/en-us/azure/architecture/multitenant-identity/app-roles).
     * 
     */
    public List<GetServicePrincipalAppRole> appRoles() {
        return this.appRoles;
    }
    /**
     * @return The application ID (client ID) of the application associated with this service principal.
     * 
     */
    public String applicationId() {
        return this.applicationId;
    }
    /**
     * @return The tenant ID where the associated application is registered.
     * 
     */
    public String applicationTenantId() {
        return this.applicationTenantId;
    }
    /**
     * @return Permission help text that appears in the admin app assignment and consent experiences.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return Display name for the permission that appears in the admin consent and app assignment experiences.
     * 
     */
    public String displayName() {
        return this.displayName;
    }
    public List<GetServicePrincipalFeatureTag> featureTags() {
        return this.featureTags;
    }
    /**
     * @return A `features` block as described below.
     * 
     * @deprecated
     * This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider
     * 
     */
    @Deprecated /* This block has been renamed to `feature_tags` and will be removed in version 3.0 of the provider */
    public List<GetServicePrincipalFeature> features() {
        return this.features;
    }
    /**
     * @return Home page or landing page of the associated application.
     * 
     */
    public String homepageUrl() {
        return this.homepageUrl;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The URL where the service provider redirects the user to Azure AD to authenticate. Azure AD uses the URL to launch the application from Microsoft 365 or the Azure AD My Apps.
     * 
     */
    public String loginUrl() {
        return this.loginUrl;
    }
    /**
     * @return The URL that will be used by Microsoft&#39;s authorization service to logout an user using OpenId Connect front-channel, back-channel or SAML logout protocols, taken from the associated application.
     * 
     */
    public String logoutUrl() {
        return this.logoutUrl;
    }
    /**
     * @return A free text field to capture information about the service principal, typically used for operational purposes.
     * 
     */
    public String notes() {
        return this.notes;
    }
    /**
     * @return A list of email addresses where Azure AD sends a notification when the active certificate is near the expiration date. This is only for the certificates used to sign the SAML token issued for Azure AD Gallery applications.
     * 
     */
    public List<String> notificationEmailAddresses() {
        return this.notificationEmailAddresses;
    }
    /**
     * @return A mapping of OAuth2.0 permission scope values to scope IDs, as exposed by the associated application, intended to be useful when referencing permission scopes in other resources in your configuration.
     * 
     */
    public Map<String,String> oauth2PermissionScopeIds() {
        return this.oauth2PermissionScopeIds;
    }
    /**
     * @return A collection of OAuth 2.0 delegated permissions exposed by the associated application. Each permission is covered by an `oauth2_permission_scopes` block as documented below.
     * 
     */
    public List<GetServicePrincipalOauth2PermissionScope> oauth2PermissionScopes() {
        return this.oauth2PermissionScopes;
    }
    /**
     * @return The object ID of the service principal.
     * 
     */
    public String objectId() {
        return this.objectId;
    }
    /**
     * @return The single sign-on mode configured for this application. Azure AD uses the preferred single sign-on mode to launch the application from Microsoft 365 or the Azure AD My Apps.
     * 
     */
    public String preferredSingleSignOnMode() {
        return this.preferredSingleSignOnMode;
    }
    /**
     * @return A list of URLs where user tokens are sent for sign-in with the associated application, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent for the associated application.
     * 
     */
    public List<String> redirectUris() {
        return this.redirectUris;
    }
    /**
     * @return The URL where the service exposes SAML metadata for federation.
     * 
     */
    public String samlMetadataUrl() {
        return this.samlMetadataUrl;
    }
    /**
     * @return A `saml_single_sign_on` block as documented below.
     * 
     */
    public List<GetServicePrincipalSamlSingleSignOn> samlSingleSignOns() {
        return this.samlSingleSignOns;
    }
    /**
     * @return A list of identifier URI(s), copied over from the associated application.
     * 
     */
    public List<String> servicePrincipalNames() {
        return this.servicePrincipalNames;
    }
    /**
     * @return The Microsoft account types that are supported for the associated application. Possible values include `AzureADMyOrg`, `AzureADMultipleOrgs`, `AzureADandPersonalMicrosoftAccount` or `PersonalMicrosoftAccount`.
     * 
     */
    public String signInAudience() {
        return this.signInAudience;
    }
    /**
     * @return A list of tags applied to the service principal.
     * 
     */
    public List<String> tags() {
        return this.tags;
    }
    /**
     * @return Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Possible values are `User` or `Admin`.
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetServicePrincipalResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean accountEnabled;
        private List<String> alternativeNames;
        private Boolean appRoleAssignmentRequired;
        private Map<String,String> appRoleIds;
        private List<GetServicePrincipalAppRole> appRoles;
        private String applicationId;
        private String applicationTenantId;
        private String description;
        private String displayName;
        private List<GetServicePrincipalFeatureTag> featureTags;
        private List<GetServicePrincipalFeature> features;
        private String homepageUrl;
        private String id;
        private String loginUrl;
        private String logoutUrl;
        private String notes;
        private List<String> notificationEmailAddresses;
        private Map<String,String> oauth2PermissionScopeIds;
        private List<GetServicePrincipalOauth2PermissionScope> oauth2PermissionScopes;
        private String objectId;
        private String preferredSingleSignOnMode;
        private List<String> redirectUris;
        private String samlMetadataUrl;
        private List<GetServicePrincipalSamlSingleSignOn> samlSingleSignOns;
        private List<String> servicePrincipalNames;
        private String signInAudience;
        private List<String> tags;
        private String type;
        public Builder() {}
        public Builder(GetServicePrincipalResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accountEnabled = defaults.accountEnabled;
    	      this.alternativeNames = defaults.alternativeNames;
    	      this.appRoleAssignmentRequired = defaults.appRoleAssignmentRequired;
    	      this.appRoleIds = defaults.appRoleIds;
    	      this.appRoles = defaults.appRoles;
    	      this.applicationId = defaults.applicationId;
    	      this.applicationTenantId = defaults.applicationTenantId;
    	      this.description = defaults.description;
    	      this.displayName = defaults.displayName;
    	      this.featureTags = defaults.featureTags;
    	      this.features = defaults.features;
    	      this.homepageUrl = defaults.homepageUrl;
    	      this.id = defaults.id;
    	      this.loginUrl = defaults.loginUrl;
    	      this.logoutUrl = defaults.logoutUrl;
    	      this.notes = defaults.notes;
    	      this.notificationEmailAddresses = defaults.notificationEmailAddresses;
    	      this.oauth2PermissionScopeIds = defaults.oauth2PermissionScopeIds;
    	      this.oauth2PermissionScopes = defaults.oauth2PermissionScopes;
    	      this.objectId = defaults.objectId;
    	      this.preferredSingleSignOnMode = defaults.preferredSingleSignOnMode;
    	      this.redirectUris = defaults.redirectUris;
    	      this.samlMetadataUrl = defaults.samlMetadataUrl;
    	      this.samlSingleSignOns = defaults.samlSingleSignOns;
    	      this.servicePrincipalNames = defaults.servicePrincipalNames;
    	      this.signInAudience = defaults.signInAudience;
    	      this.tags = defaults.tags;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder accountEnabled(Boolean accountEnabled) {
            this.accountEnabled = Objects.requireNonNull(accountEnabled);
            return this;
        }
        @CustomType.Setter
        public Builder alternativeNames(List<String> alternativeNames) {
            this.alternativeNames = Objects.requireNonNull(alternativeNames);
            return this;
        }
        public Builder alternativeNames(String... alternativeNames) {
            return alternativeNames(List.of(alternativeNames));
        }
        @CustomType.Setter
        public Builder appRoleAssignmentRequired(Boolean appRoleAssignmentRequired) {
            this.appRoleAssignmentRequired = Objects.requireNonNull(appRoleAssignmentRequired);
            return this;
        }
        @CustomType.Setter
        public Builder appRoleIds(Map<String,String> appRoleIds) {
            this.appRoleIds = Objects.requireNonNull(appRoleIds);
            return this;
        }
        @CustomType.Setter
        public Builder appRoles(List<GetServicePrincipalAppRole> appRoles) {
            this.appRoles = Objects.requireNonNull(appRoles);
            return this;
        }
        public Builder appRoles(GetServicePrincipalAppRole... appRoles) {
            return appRoles(List.of(appRoles));
        }
        @CustomType.Setter
        public Builder applicationId(String applicationId) {
            this.applicationId = Objects.requireNonNull(applicationId);
            return this;
        }
        @CustomType.Setter
        public Builder applicationTenantId(String applicationTenantId) {
            this.applicationTenantId = Objects.requireNonNull(applicationTenantId);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder displayName(String displayName) {
            this.displayName = Objects.requireNonNull(displayName);
            return this;
        }
        @CustomType.Setter
        public Builder featureTags(List<GetServicePrincipalFeatureTag> featureTags) {
            this.featureTags = Objects.requireNonNull(featureTags);
            return this;
        }
        public Builder featureTags(GetServicePrincipalFeatureTag... featureTags) {
            return featureTags(List.of(featureTags));
        }
        @CustomType.Setter
        public Builder features(List<GetServicePrincipalFeature> features) {
            this.features = Objects.requireNonNull(features);
            return this;
        }
        public Builder features(GetServicePrincipalFeature... features) {
            return features(List.of(features));
        }
        @CustomType.Setter
        public Builder homepageUrl(String homepageUrl) {
            this.homepageUrl = Objects.requireNonNull(homepageUrl);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder loginUrl(String loginUrl) {
            this.loginUrl = Objects.requireNonNull(loginUrl);
            return this;
        }
        @CustomType.Setter
        public Builder logoutUrl(String logoutUrl) {
            this.logoutUrl = Objects.requireNonNull(logoutUrl);
            return this;
        }
        @CustomType.Setter
        public Builder notes(String notes) {
            this.notes = Objects.requireNonNull(notes);
            return this;
        }
        @CustomType.Setter
        public Builder notificationEmailAddresses(List<String> notificationEmailAddresses) {
            this.notificationEmailAddresses = Objects.requireNonNull(notificationEmailAddresses);
            return this;
        }
        public Builder notificationEmailAddresses(String... notificationEmailAddresses) {
            return notificationEmailAddresses(List.of(notificationEmailAddresses));
        }
        @CustomType.Setter
        public Builder oauth2PermissionScopeIds(Map<String,String> oauth2PermissionScopeIds) {
            this.oauth2PermissionScopeIds = Objects.requireNonNull(oauth2PermissionScopeIds);
            return this;
        }
        @CustomType.Setter
        public Builder oauth2PermissionScopes(List<GetServicePrincipalOauth2PermissionScope> oauth2PermissionScopes) {
            this.oauth2PermissionScopes = Objects.requireNonNull(oauth2PermissionScopes);
            return this;
        }
        public Builder oauth2PermissionScopes(GetServicePrincipalOauth2PermissionScope... oauth2PermissionScopes) {
            return oauth2PermissionScopes(List.of(oauth2PermissionScopes));
        }
        @CustomType.Setter
        public Builder objectId(String objectId) {
            this.objectId = Objects.requireNonNull(objectId);
            return this;
        }
        @CustomType.Setter
        public Builder preferredSingleSignOnMode(String preferredSingleSignOnMode) {
            this.preferredSingleSignOnMode = Objects.requireNonNull(preferredSingleSignOnMode);
            return this;
        }
        @CustomType.Setter
        public Builder redirectUris(List<String> redirectUris) {
            this.redirectUris = Objects.requireNonNull(redirectUris);
            return this;
        }
        public Builder redirectUris(String... redirectUris) {
            return redirectUris(List.of(redirectUris));
        }
        @CustomType.Setter
        public Builder samlMetadataUrl(String samlMetadataUrl) {
            this.samlMetadataUrl = Objects.requireNonNull(samlMetadataUrl);
            return this;
        }
        @CustomType.Setter
        public Builder samlSingleSignOns(List<GetServicePrincipalSamlSingleSignOn> samlSingleSignOns) {
            this.samlSingleSignOns = Objects.requireNonNull(samlSingleSignOns);
            return this;
        }
        public Builder samlSingleSignOns(GetServicePrincipalSamlSingleSignOn... samlSingleSignOns) {
            return samlSingleSignOns(List.of(samlSingleSignOns));
        }
        @CustomType.Setter
        public Builder servicePrincipalNames(List<String> servicePrincipalNames) {
            this.servicePrincipalNames = Objects.requireNonNull(servicePrincipalNames);
            return this;
        }
        public Builder servicePrincipalNames(String... servicePrincipalNames) {
            return servicePrincipalNames(List.of(servicePrincipalNames));
        }
        @CustomType.Setter
        public Builder signInAudience(String signInAudience) {
            this.signInAudience = Objects.requireNonNull(signInAudience);
            return this;
        }
        @CustomType.Setter
        public Builder tags(List<String> tags) {
            this.tags = Objects.requireNonNull(tags);
            return this;
        }
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        public GetServicePrincipalResult build() {
            final var o = new GetServicePrincipalResult();
            o.accountEnabled = accountEnabled;
            o.alternativeNames = alternativeNames;
            o.appRoleAssignmentRequired = appRoleAssignmentRequired;
            o.appRoleIds = appRoleIds;
            o.appRoles = appRoles;
            o.applicationId = applicationId;
            o.applicationTenantId = applicationTenantId;
            o.description = description;
            o.displayName = displayName;
            o.featureTags = featureTags;
            o.features = features;
            o.homepageUrl = homepageUrl;
            o.id = id;
            o.loginUrl = loginUrl;
            o.logoutUrl = logoutUrl;
            o.notes = notes;
            o.notificationEmailAddresses = notificationEmailAddresses;
            o.oauth2PermissionScopeIds = oauth2PermissionScopeIds;
            o.oauth2PermissionScopes = oauth2PermissionScopes;
            o.objectId = objectId;
            o.preferredSingleSignOnMode = preferredSingleSignOnMode;
            o.redirectUris = redirectUris;
            o.samlMetadataUrl = samlMetadataUrl;
            o.samlSingleSignOns = samlSingleSignOns;
            o.servicePrincipalNames = servicePrincipalNames;
            o.signInAudience = signInAudience;
            o.tags = tags;
            o.type = type;
            return o;
        }
    }
}
