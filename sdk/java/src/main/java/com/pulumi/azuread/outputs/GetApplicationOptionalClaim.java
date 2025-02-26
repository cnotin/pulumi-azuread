// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.azuread.outputs.GetApplicationOptionalClaimAccessToken;
import com.pulumi.azuread.outputs.GetApplicationOptionalClaimIdToken;
import com.pulumi.azuread.outputs.GetApplicationOptionalClaimSaml2Token;
import com.pulumi.core.annotations.CustomType;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetApplicationOptionalClaim {
    /**
     * @return One or more `access_token` blocks as documented below.
     * 
     */
    private @Nullable List<GetApplicationOptionalClaimAccessToken> accessTokens;
    /**
     * @return One or more `id_token` blocks as documented below.
     * 
     */
    private @Nullable List<GetApplicationOptionalClaimIdToken> idTokens;
    /**
     * @return One or more `saml2_token` blocks as documented below.
     * 
     */
    private @Nullable List<GetApplicationOptionalClaimSaml2Token> saml2Tokens;

    private GetApplicationOptionalClaim() {}
    /**
     * @return One or more `access_token` blocks as documented below.
     * 
     */
    public List<GetApplicationOptionalClaimAccessToken> accessTokens() {
        return this.accessTokens == null ? List.of() : this.accessTokens;
    }
    /**
     * @return One or more `id_token` blocks as documented below.
     * 
     */
    public List<GetApplicationOptionalClaimIdToken> idTokens() {
        return this.idTokens == null ? List.of() : this.idTokens;
    }
    /**
     * @return One or more `saml2_token` blocks as documented below.
     * 
     */
    public List<GetApplicationOptionalClaimSaml2Token> saml2Tokens() {
        return this.saml2Tokens == null ? List.of() : this.saml2Tokens;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetApplicationOptionalClaim defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<GetApplicationOptionalClaimAccessToken> accessTokens;
        private @Nullable List<GetApplicationOptionalClaimIdToken> idTokens;
        private @Nullable List<GetApplicationOptionalClaimSaml2Token> saml2Tokens;
        public Builder() {}
        public Builder(GetApplicationOptionalClaim defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accessTokens = defaults.accessTokens;
    	      this.idTokens = defaults.idTokens;
    	      this.saml2Tokens = defaults.saml2Tokens;
        }

        @CustomType.Setter
        public Builder accessTokens(@Nullable List<GetApplicationOptionalClaimAccessToken> accessTokens) {
            this.accessTokens = accessTokens;
            return this;
        }
        public Builder accessTokens(GetApplicationOptionalClaimAccessToken... accessTokens) {
            return accessTokens(List.of(accessTokens));
        }
        @CustomType.Setter
        public Builder idTokens(@Nullable List<GetApplicationOptionalClaimIdToken> idTokens) {
            this.idTokens = idTokens;
            return this;
        }
        public Builder idTokens(GetApplicationOptionalClaimIdToken... idTokens) {
            return idTokens(List.of(idTokens));
        }
        @CustomType.Setter
        public Builder saml2Tokens(@Nullable List<GetApplicationOptionalClaimSaml2Token> saml2Tokens) {
            this.saml2Tokens = saml2Tokens;
            return this;
        }
        public Builder saml2Tokens(GetApplicationOptionalClaimSaml2Token... saml2Tokens) {
            return saml2Tokens(List.of(saml2Tokens));
        }
        public GetApplicationOptionalClaim build() {
            final var o = new GetApplicationOptionalClaim();
            o.accessTokens = accessTokens;
            o.idTokens = idTokens;
            o.saml2Tokens = saml2Tokens;
            return o;
        }
    }
}
