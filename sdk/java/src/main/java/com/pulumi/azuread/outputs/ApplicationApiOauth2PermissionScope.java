// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ApplicationApiOauth2PermissionScope {
    /**
     * @return Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
     * 
     */
    private @Nullable String adminConsentDescription;
    /**
     * @return Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
     * 
     */
    private @Nullable String adminConsentDisplayName;
    /**
     * @return Determines if the permission scope is enabled. Defaults to `true`.
     * 
     */
    private @Nullable Boolean enabled;
    /**
     * @return The unique identifier of the delegated permission. Must be a valid UUID.
     * 
     */
    private String id;
    /**
     * @return Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Defaults to `User`. Possible values are `User` or `Admin`.
     * 
     */
    private @Nullable String type;
    /**
     * @return Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
     * 
     */
    private @Nullable String userConsentDescription;
    /**
     * @return Display name for the delegated permission that appears in the end user consent experience.
     * 
     */
    private @Nullable String userConsentDisplayName;
    /**
     * @return The value that is used for the `scp` claim in OAuth 2.0 access tokens.
     * 
     */
    private @Nullable String value;

    private ApplicationApiOauth2PermissionScope() {}
    /**
     * @return Delegated permission description that appears in all tenant-wide admin consent experiences, intended to be read by an administrator granting the permission on behalf of all users.
     * 
     */
    public Optional<String> adminConsentDescription() {
        return Optional.ofNullable(this.adminConsentDescription);
    }
    /**
     * @return Display name for the delegated permission, intended to be read by an administrator granting the permission on behalf of all users.
     * 
     */
    public Optional<String> adminConsentDisplayName() {
        return Optional.ofNullable(this.adminConsentDisplayName);
    }
    /**
     * @return Determines if the permission scope is enabled. Defaults to `true`.
     * 
     */
    public Optional<Boolean> enabled() {
        return Optional.ofNullable(this.enabled);
    }
    /**
     * @return The unique identifier of the delegated permission. Must be a valid UUID.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator should be required for consent to the permissions. Defaults to `User`. Possible values are `User` or `Admin`.
     * 
     */
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }
    /**
     * @return Delegated permission description that appears in the end user consent experience, intended to be read by a user consenting on their own behalf.
     * 
     */
    public Optional<String> userConsentDescription() {
        return Optional.ofNullable(this.userConsentDescription);
    }
    /**
     * @return Display name for the delegated permission that appears in the end user consent experience.
     * 
     */
    public Optional<String> userConsentDisplayName() {
        return Optional.ofNullable(this.userConsentDisplayName);
    }
    /**
     * @return The value that is used for the `scp` claim in OAuth 2.0 access tokens.
     * 
     */
    public Optional<String> value() {
        return Optional.ofNullable(this.value);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ApplicationApiOauth2PermissionScope defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String adminConsentDescription;
        private @Nullable String adminConsentDisplayName;
        private @Nullable Boolean enabled;
        private String id;
        private @Nullable String type;
        private @Nullable String userConsentDescription;
        private @Nullable String userConsentDisplayName;
        private @Nullable String value;
        public Builder() {}
        public Builder(ApplicationApiOauth2PermissionScope defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.adminConsentDescription = defaults.adminConsentDescription;
    	      this.adminConsentDisplayName = defaults.adminConsentDisplayName;
    	      this.enabled = defaults.enabled;
    	      this.id = defaults.id;
    	      this.type = defaults.type;
    	      this.userConsentDescription = defaults.userConsentDescription;
    	      this.userConsentDisplayName = defaults.userConsentDisplayName;
    	      this.value = defaults.value;
        }

        @CustomType.Setter
        public Builder adminConsentDescription(@Nullable String adminConsentDescription) {
            this.adminConsentDescription = adminConsentDescription;
            return this;
        }
        @CustomType.Setter
        public Builder adminConsentDisplayName(@Nullable String adminConsentDisplayName) {
            this.adminConsentDisplayName = adminConsentDisplayName;
            return this;
        }
        @CustomType.Setter
        public Builder enabled(@Nullable Boolean enabled) {
            this.enabled = enabled;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder type(@Nullable String type) {
            this.type = type;
            return this;
        }
        @CustomType.Setter
        public Builder userConsentDescription(@Nullable String userConsentDescription) {
            this.userConsentDescription = userConsentDescription;
            return this;
        }
        @CustomType.Setter
        public Builder userConsentDisplayName(@Nullable String userConsentDisplayName) {
            this.userConsentDisplayName = userConsentDisplayName;
            return this;
        }
        @CustomType.Setter
        public Builder value(@Nullable String value) {
            this.value = value;
            return this;
        }
        public ApplicationApiOauth2PermissionScope build() {
            final var o = new ApplicationApiOauth2PermissionScope();
            o.adminConsentDescription = adminConsentDescription;
            o.adminConsentDisplayName = adminConsentDisplayName;
            o.enabled = enabled;
            o.id = id;
            o.type = type;
            o.userConsentDescription = userConsentDescription;
            o.userConsentDisplayName = userConsentDisplayName;
            o.value = value;
            return o;
        }
    }
}
