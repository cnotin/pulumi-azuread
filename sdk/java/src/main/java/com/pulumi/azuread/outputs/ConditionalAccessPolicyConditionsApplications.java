// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class ConditionalAccessPolicyConditionsApplications {
    /**
     * @return A list of application IDs explicitly excluded from the policy.
     * 
     */
    private @Nullable List<String> excludedApplications;
    /**
     * @return A list of application IDs the policy applies to, unless explicitly excluded (in `excluded_applications`). Can also be set to `All`. Cannot be specified with `included_user_actions`. One of `included_applications` or `included_user_actions` must be specified.
     * 
     */
    private @Nullable List<String> includedApplications;
    /**
     * @return A list of user actions to include. Supported values are `urn:user:registerdevice` and `urn:user:registersecurityinfo`. Cannot be specified with `included_applications`. One of `included_applications` or `included_user_actions` must be specified.
     * 
     */
    private @Nullable List<String> includedUserActions;

    private ConditionalAccessPolicyConditionsApplications() {}
    /**
     * @return A list of application IDs explicitly excluded from the policy.
     * 
     */
    public List<String> excludedApplications() {
        return this.excludedApplications == null ? List.of() : this.excludedApplications;
    }
    /**
     * @return A list of application IDs the policy applies to, unless explicitly excluded (in `excluded_applications`). Can also be set to `All`. Cannot be specified with `included_user_actions`. One of `included_applications` or `included_user_actions` must be specified.
     * 
     */
    public List<String> includedApplications() {
        return this.includedApplications == null ? List.of() : this.includedApplications;
    }
    /**
     * @return A list of user actions to include. Supported values are `urn:user:registerdevice` and `urn:user:registersecurityinfo`. Cannot be specified with `included_applications`. One of `included_applications` or `included_user_actions` must be specified.
     * 
     */
    public List<String> includedUserActions() {
        return this.includedUserActions == null ? List.of() : this.includedUserActions;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ConditionalAccessPolicyConditionsApplications defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> excludedApplications;
        private @Nullable List<String> includedApplications;
        private @Nullable List<String> includedUserActions;
        public Builder() {}
        public Builder(ConditionalAccessPolicyConditionsApplications defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.excludedApplications = defaults.excludedApplications;
    	      this.includedApplications = defaults.includedApplications;
    	      this.includedUserActions = defaults.includedUserActions;
        }

        @CustomType.Setter
        public Builder excludedApplications(@Nullable List<String> excludedApplications) {
            this.excludedApplications = excludedApplications;
            return this;
        }
        public Builder excludedApplications(String... excludedApplications) {
            return excludedApplications(List.of(excludedApplications));
        }
        @CustomType.Setter
        public Builder includedApplications(@Nullable List<String> includedApplications) {
            this.includedApplications = includedApplications;
            return this;
        }
        public Builder includedApplications(String... includedApplications) {
            return includedApplications(List.of(includedApplications));
        }
        @CustomType.Setter
        public Builder includedUserActions(@Nullable List<String> includedUserActions) {
            this.includedUserActions = includedUserActions;
            return this;
        }
        public Builder includedUserActions(String... includedUserActions) {
            return includedUserActions(List.of(includedUserActions));
        }
        public ConditionalAccessPolicyConditionsApplications build() {
            final var o = new ConditionalAccessPolicyConditionsApplications();
            o.excludedApplications = excludedApplications;
            o.includedApplications = includedApplications;
            o.includedUserActions = includedUserActions;
            return o;
        }
    }
}
