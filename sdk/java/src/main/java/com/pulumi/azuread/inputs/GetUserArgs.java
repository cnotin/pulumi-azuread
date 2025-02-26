// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetUserArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetUserArgs Empty = new GetUserArgs();

    /**
     * The email alias of the user.
     * 
     */
    @Import(name="mailNickname")
    private @Nullable Output<String> mailNickname;

    /**
     * @return The email alias of the user.
     * 
     */
    public Optional<Output<String>> mailNickname() {
        return Optional.ofNullable(this.mailNickname);
    }

    /**
     * The object ID of the user.
     * 
     */
    @Import(name="objectId")
    private @Nullable Output<String> objectId;

    /**
     * @return The object ID of the user.
     * 
     */
    public Optional<Output<String>> objectId() {
        return Optional.ofNullable(this.objectId);
    }

    /**
     * The user principal name (UPN) of the user.
     * 
     */
    @Import(name="userPrincipalName")
    private @Nullable Output<String> userPrincipalName;

    /**
     * @return The user principal name (UPN) of the user.
     * 
     */
    public Optional<Output<String>> userPrincipalName() {
        return Optional.ofNullable(this.userPrincipalName);
    }

    private GetUserArgs() {}

    private GetUserArgs(GetUserArgs $) {
        this.mailNickname = $.mailNickname;
        this.objectId = $.objectId;
        this.userPrincipalName = $.userPrincipalName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetUserArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetUserArgs $;

        public Builder() {
            $ = new GetUserArgs();
        }

        public Builder(GetUserArgs defaults) {
            $ = new GetUserArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param mailNickname The email alias of the user.
         * 
         * @return builder
         * 
         */
        public Builder mailNickname(@Nullable Output<String> mailNickname) {
            $.mailNickname = mailNickname;
            return this;
        }

        /**
         * @param mailNickname The email alias of the user.
         * 
         * @return builder
         * 
         */
        public Builder mailNickname(String mailNickname) {
            return mailNickname(Output.of(mailNickname));
        }

        /**
         * @param objectId The object ID of the user.
         * 
         * @return builder
         * 
         */
        public Builder objectId(@Nullable Output<String> objectId) {
            $.objectId = objectId;
            return this;
        }

        /**
         * @param objectId The object ID of the user.
         * 
         * @return builder
         * 
         */
        public Builder objectId(String objectId) {
            return objectId(Output.of(objectId));
        }

        /**
         * @param userPrincipalName The user principal name (UPN) of the user.
         * 
         * @return builder
         * 
         */
        public Builder userPrincipalName(@Nullable Output<String> userPrincipalName) {
            $.userPrincipalName = userPrincipalName;
            return this;
        }

        /**
         * @param userPrincipalName The user principal name (UPN) of the user.
         * 
         * @return builder
         * 
         */
        public Builder userPrincipalName(String userPrincipalName) {
            return userPrincipalName(Output.of(userPrincipalName));
        }

        public GetUserArgs build() {
            return $;
        }
    }

}
