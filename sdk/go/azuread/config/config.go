// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// Base64 encoded PKCS#12 certificate bundle to use when authenticating as a Service Principal using a Client Certificate
func GetClientCertificate(ctx *pulumi.Context) string {
	return config.Get(ctx, "azuread:clientCertificate")
}

// The password to decrypt the Client Certificate. For use when authenticating as a Service Principal using a Client
// Certificate
func GetClientCertificatePassword(ctx *pulumi.Context) string {
	return config.Get(ctx, "azuread:clientCertificatePassword")
}

// The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service
// Principal using a Client Certificate
func GetClientCertificatePath(ctx *pulumi.Context) string {
	return config.Get(ctx, "azuread:clientCertificatePath")
}

// The Client ID which should be used for service principal authentication
func GetClientId(ctx *pulumi.Context) string {
	return config.Get(ctx, "azuread:clientId")
}

// The application password to use when authenticating as a Service Principal using a Client Secret
func GetClientSecret(ctx *pulumi.Context) string {
	return config.Get(ctx, "azuread:clientSecret")
}

// Disable the Terraform Partner ID, which is used if a custom `partner_id` isn't specified
func GetDisableTerraformPartnerId(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "azuread:disableTerraformPartnerId")
}

// The cloud environment which should be used. Possible values are: `global` (also `public`), `usgovernmentl4` (also
// `usgovernment`), `usgovernmentl5` (also `dod`), and `china`. Defaults to `global`
func GetEnvironment(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "azuread:environment")
	if err == nil {
		return v
	}
	return getEnvOrDefault("public", nil, "ARM_ENVIRONMENT").(string)
}

// The path to a custom endpoint for Managed Identity - in most circumstances this should be detected automatically
func GetMsiEndpoint(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "azuread:msiEndpoint")
	if err == nil {
		return v
	}
	return getEnvOrDefault("", nil, "ARM_MSI_ENDPOINT").(string)
}

// The bearer token for the request to the OIDC provider. For use when authenticating as a Service Principal using OpenID
// Connect.
func GetOidcRequestToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "azuread:oidcRequestToken")
}

// The URL for the OIDC provider from which to request an ID token. For use when authenticating as a Service Principal
// using OpenID Connect.
func GetOidcRequestUrl(ctx *pulumi.Context) string {
	return config.Get(ctx, "azuread:oidcRequestUrl")
}

// The ID token for use when authenticating as a Service Principal using OpenID Connect.
func GetOidcToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "azuread:oidcToken")
}

// The path to a file containing an ID token for use when authenticating as a Service Principal using OpenID Connect.
func GetOidcTokenFilePath(ctx *pulumi.Context) string {
	return config.Get(ctx, "azuread:oidcTokenFilePath")
}

// A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution
func GetPartnerId(ctx *pulumi.Context) string {
	return config.Get(ctx, "azuread:partnerId")
}

// The Tenant ID which should be used. Works with all authentication methods except Managed Identity
func GetTenantId(ctx *pulumi.Context) string {
	return config.Get(ctx, "azuread:tenantId")
}

// Allow Azure CLI to be used for Authentication
func GetUseCli(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "azuread:useCli")
}

// Allow Managed Identity to be used for Authentication
func GetUseMsi(ctx *pulumi.Context) bool {
	v, err := config.TryBool(ctx, "azuread:useMsi")
	if err == nil {
		return v
	}
	return getEnvOrDefault(false, parseEnvBool, "ARM_USE_MSI").(bool)
}

// Allow OpenID Connect to be used for authentication
func GetUseOidc(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "azuread:useOidc")
}
