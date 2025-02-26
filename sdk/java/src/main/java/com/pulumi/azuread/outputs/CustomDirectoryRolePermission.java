// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class CustomDirectoryRolePermission {
    /**
     * @return A set of tasks that can be performed on a resource. For more information, see the [Permissions Reference](https://docs.microsoft.com/en-us/azure/active-directory/roles/permissions-reference) documentation.
     * 
     */
    private List<String> allowedResourceActions;

    private CustomDirectoryRolePermission() {}
    /**
     * @return A set of tasks that can be performed on a resource. For more information, see the [Permissions Reference](https://docs.microsoft.com/en-us/azure/active-directory/roles/permissions-reference) documentation.
     * 
     */
    public List<String> allowedResourceActions() {
        return this.allowedResourceActions;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(CustomDirectoryRolePermission defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> allowedResourceActions;
        public Builder() {}
        public Builder(CustomDirectoryRolePermission defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allowedResourceActions = defaults.allowedResourceActions;
        }

        @CustomType.Setter
        public Builder allowedResourceActions(List<String> allowedResourceActions) {
            this.allowedResourceActions = Objects.requireNonNull(allowedResourceActions);
            return this;
        }
        public Builder allowedResourceActions(String... allowedResourceActions) {
            return allowedResourceActions(List.of(allowedResourceActions));
        }
        public CustomDirectoryRolePermission build() {
            final var o = new CustomDirectoryRolePermission();
            o.allowedResourceActions = allowedResourceActions;
            return o;
        }
    }
}
