// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class ApplicationPreAuthorizedArgs extends com.pulumi.resources.ResourceArgs {

    public static final ApplicationPreAuthorizedArgs Empty = new ApplicationPreAuthorizedArgs();

    /**
     * The object ID of the application for which permissions are being authorized. Changing this field forces a new resource to be created.
     * 
     */
    @Import(name="applicationObjectId", required=true)
    private Output<String> applicationObjectId;

    /**
     * @return The object ID of the application for which permissions are being authorized. Changing this field forces a new resource to be created.
     * 
     */
    public Output<String> applicationObjectId() {
        return this.applicationObjectId;
    }

    /**
     * The application ID of the pre-authorized application
     * 
     */
    @Import(name="authorizedAppId", required=true)
    private Output<String> authorizedAppId;

    /**
     * @return The application ID of the pre-authorized application
     * 
     */
    public Output<String> authorizedAppId() {
        return this.authorizedAppId;
    }

    /**
     * A set of permission scope IDs required by the authorized application.
     * 
     */
    @Import(name="permissionIds", required=true)
    private Output<List<String>> permissionIds;

    /**
     * @return A set of permission scope IDs required by the authorized application.
     * 
     */
    public Output<List<String>> permissionIds() {
        return this.permissionIds;
    }

    private ApplicationPreAuthorizedArgs() {}

    private ApplicationPreAuthorizedArgs(ApplicationPreAuthorizedArgs $) {
        this.applicationObjectId = $.applicationObjectId;
        this.authorizedAppId = $.authorizedAppId;
        this.permissionIds = $.permissionIds;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ApplicationPreAuthorizedArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ApplicationPreAuthorizedArgs $;

        public Builder() {
            $ = new ApplicationPreAuthorizedArgs();
        }

        public Builder(ApplicationPreAuthorizedArgs defaults) {
            $ = new ApplicationPreAuthorizedArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param applicationObjectId The object ID of the application for which permissions are being authorized. Changing this field forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder applicationObjectId(Output<String> applicationObjectId) {
            $.applicationObjectId = applicationObjectId;
            return this;
        }

        /**
         * @param applicationObjectId The object ID of the application for which permissions are being authorized. Changing this field forces a new resource to be created.
         * 
         * @return builder
         * 
         */
        public Builder applicationObjectId(String applicationObjectId) {
            return applicationObjectId(Output.of(applicationObjectId));
        }

        /**
         * @param authorizedAppId The application ID of the pre-authorized application
         * 
         * @return builder
         * 
         */
        public Builder authorizedAppId(Output<String> authorizedAppId) {
            $.authorizedAppId = authorizedAppId;
            return this;
        }

        /**
         * @param authorizedAppId The application ID of the pre-authorized application
         * 
         * @return builder
         * 
         */
        public Builder authorizedAppId(String authorizedAppId) {
            return authorizedAppId(Output.of(authorizedAppId));
        }

        /**
         * @param permissionIds A set of permission scope IDs required by the authorized application.
         * 
         * @return builder
         * 
         */
        public Builder permissionIds(Output<List<String>> permissionIds) {
            $.permissionIds = permissionIds;
            return this;
        }

        /**
         * @param permissionIds A set of permission scope IDs required by the authorized application.
         * 
         * @return builder
         * 
         */
        public Builder permissionIds(List<String> permissionIds) {
            return permissionIds(Output.of(permissionIds));
        }

        /**
         * @param permissionIds A set of permission scope IDs required by the authorized application.
         * 
         * @return builder
         * 
         */
        public Builder permissionIds(String... permissionIds) {
            return permissionIds(List.of(permissionIds));
        }

        public ApplicationPreAuthorizedArgs build() {
            $.applicationObjectId = Objects.requireNonNull($.applicationObjectId, "expected parameter 'applicationObjectId' to be non-null");
            $.authorizedAppId = Objects.requireNonNull($.authorizedAppId, "expected parameter 'authorizedAppId' to be non-null");
            $.permissionIds = Objects.requireNonNull($.permissionIds, "expected parameter 'permissionIds' to be non-null");
            return $;
        }
    }

}
