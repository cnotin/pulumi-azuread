// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.azuread.outputs;

import com.pulumi.azuread.outputs.GetGroupDynamicMembership;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetGroupResult {
    /**
     * @return Indicates whether this group can be assigned to an Azure Active Directory role.
     * 
     */
    private Boolean assignableToRole;
    /**
     * @return Indicates whether new members added to the group will be auto-subscribed to receive email notifications. Only set for Unified groups.
     * 
     */
    private Boolean autoSubscribeNewMembers;
    /**
     * @return A list of behaviors for a Microsoft 365 group, such as `AllowOnlyMembersToPost`, `HideGroupInOutlook`, `SubscribeNewGroupMembers` and `WelcomeEmailDisabled`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for more details.
     * 
     */
    private List<String> behaviors;
    /**
     * @return The optional description of the group.
     * 
     */
    private String description;
    /**
     * @return The display name for the group.
     * 
     */
    private String displayName;
    /**
     * @return A `dynamic_membership` block as documented below.
     * 
     */
    private List<GetGroupDynamicMembership> dynamicMemberships;
    /**
     * @return Indicates whether people external to the organization can send messages to the group. Only set for Unified groups.
     * 
     */
    private Boolean externalSendersAllowed;
    /**
     * @return Indicates whether the group is displayed in certain parts of the Outlook user interface: in the Address Book, in address lists for selecting message recipients, and in the Browse Groups dialog for searching groups. Only set for Unified groups.
     * 
     */
    private Boolean hideFromAddressLists;
    /**
     * @return Indicates whether the group is displayed in Outlook clients, such as Outlook for Windows and Outlook on the web. Only set for Unified groups.
     * 
     */
    private Boolean hideFromOutlookClients;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The SMTP address for the group.
     * 
     */
    private String mail;
    /**
     * @return Whether the group is mail-enabled.
     * 
     */
    private Boolean mailEnabled;
    /**
     * @return The mail alias for the group, unique in the organisation.
     * 
     */
    private String mailNickname;
    /**
     * @return List of object IDs of the group members.
     * 
     */
    private List<String> members;
    /**
     * @return The object ID of the group.
     * 
     */
    private String objectId;
    /**
     * @return The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.
     * 
     */
    private String onpremisesDomainName;
    /**
     * @return The on-premises NetBIOS name, synchronised from the on-premises directory when Azure AD Connect is used.
     * 
     */
    private String onpremisesNetbiosName;
    /**
     * @return The on-premises SAM account name, synchronised from the on-premises directory when Azure AD Connect is used.
     * 
     */
    private String onpremisesSamAccountName;
    /**
     * @return The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.
     * 
     */
    private String onpremisesSecurityIdentifier;
    /**
     * @return Whether this group is synchronised from an on-premises directory (`true`), no longer synchronised (`false`), or has never been synchronised (`null`).
     * 
     */
    private Boolean onpremisesSyncEnabled;
    /**
     * @return List of object IDs of the group owners.
     * 
     */
    private List<String> owners;
    /**
     * @return The preferred language for a Microsoft 365 group, in ISO 639-1 notation.
     * 
     */
    private String preferredLanguage;
    /**
     * @return A list of provisioning options for a Microsoft 365 group, such as `Team`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for details.
     * 
     */
    private List<String> provisioningOptions;
    /**
     * @return List of email addresses for the group that direct to the same group mailbox.
     * 
     */
    private List<String> proxyAddresses;
    /**
     * @return Whether the group is a security group.
     * 
     */
    private Boolean securityEnabled;
    /**
     * @return The colour theme for a Microsoft 365 group. Possible values are `Blue`, `Green`, `Orange`, `Pink`, `Purple`, `Red` or `Teal`. When no theme is set, the value is `null`.
     * 
     */
    private String theme;
    /**
     * @return A list of group types configured for the group. Supported values are `DynamicMembership`, which denotes a group with dynamic membership, and `Unified`, which specifies a Microsoft 365 group.
     * 
     */
    private List<String> types;
    /**
     * @return The group join policy and group content visibility. Possible values are `Private`, `Public`, or `Hiddenmembership`. Only Microsoft 365 groups can have `Hiddenmembership` visibility.
     * 
     */
    private String visibility;

    private GetGroupResult() {}
    /**
     * @return Indicates whether this group can be assigned to an Azure Active Directory role.
     * 
     */
    public Boolean assignableToRole() {
        return this.assignableToRole;
    }
    /**
     * @return Indicates whether new members added to the group will be auto-subscribed to receive email notifications. Only set for Unified groups.
     * 
     */
    public Boolean autoSubscribeNewMembers() {
        return this.autoSubscribeNewMembers;
    }
    /**
     * @return A list of behaviors for a Microsoft 365 group, such as `AllowOnlyMembersToPost`, `HideGroupInOutlook`, `SubscribeNewGroupMembers` and `WelcomeEmailDisabled`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for more details.
     * 
     */
    public List<String> behaviors() {
        return this.behaviors;
    }
    /**
     * @return The optional description of the group.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The display name for the group.
     * 
     */
    public String displayName() {
        return this.displayName;
    }
    /**
     * @return A `dynamic_membership` block as documented below.
     * 
     */
    public List<GetGroupDynamicMembership> dynamicMemberships() {
        return this.dynamicMemberships;
    }
    /**
     * @return Indicates whether people external to the organization can send messages to the group. Only set for Unified groups.
     * 
     */
    public Boolean externalSendersAllowed() {
        return this.externalSendersAllowed;
    }
    /**
     * @return Indicates whether the group is displayed in certain parts of the Outlook user interface: in the Address Book, in address lists for selecting message recipients, and in the Browse Groups dialog for searching groups. Only set for Unified groups.
     * 
     */
    public Boolean hideFromAddressLists() {
        return this.hideFromAddressLists;
    }
    /**
     * @return Indicates whether the group is displayed in Outlook clients, such as Outlook for Windows and Outlook on the web. Only set for Unified groups.
     * 
     */
    public Boolean hideFromOutlookClients() {
        return this.hideFromOutlookClients;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The SMTP address for the group.
     * 
     */
    public String mail() {
        return this.mail;
    }
    /**
     * @return Whether the group is mail-enabled.
     * 
     */
    public Boolean mailEnabled() {
        return this.mailEnabled;
    }
    /**
     * @return The mail alias for the group, unique in the organisation.
     * 
     */
    public String mailNickname() {
        return this.mailNickname;
    }
    /**
     * @return List of object IDs of the group members.
     * 
     */
    public List<String> members() {
        return this.members;
    }
    /**
     * @return The object ID of the group.
     * 
     */
    public String objectId() {
        return this.objectId;
    }
    /**
     * @return The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.
     * 
     */
    public String onpremisesDomainName() {
        return this.onpremisesDomainName;
    }
    /**
     * @return The on-premises NetBIOS name, synchronised from the on-premises directory when Azure AD Connect is used.
     * 
     */
    public String onpremisesNetbiosName() {
        return this.onpremisesNetbiosName;
    }
    /**
     * @return The on-premises SAM account name, synchronised from the on-premises directory when Azure AD Connect is used.
     * 
     */
    public String onpremisesSamAccountName() {
        return this.onpremisesSamAccountName;
    }
    /**
     * @return The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.
     * 
     */
    public String onpremisesSecurityIdentifier() {
        return this.onpremisesSecurityIdentifier;
    }
    /**
     * @return Whether this group is synchronised from an on-premises directory (`true`), no longer synchronised (`false`), or has never been synchronised (`null`).
     * 
     */
    public Boolean onpremisesSyncEnabled() {
        return this.onpremisesSyncEnabled;
    }
    /**
     * @return List of object IDs of the group owners.
     * 
     */
    public List<String> owners() {
        return this.owners;
    }
    /**
     * @return The preferred language for a Microsoft 365 group, in ISO 639-1 notation.
     * 
     */
    public String preferredLanguage() {
        return this.preferredLanguage;
    }
    /**
     * @return A list of provisioning options for a Microsoft 365 group, such as `Team`. See [official documentation](https://docs.microsoft.com/en-us/graph/group-set-options) for details.
     * 
     */
    public List<String> provisioningOptions() {
        return this.provisioningOptions;
    }
    /**
     * @return List of email addresses for the group that direct to the same group mailbox.
     * 
     */
    public List<String> proxyAddresses() {
        return this.proxyAddresses;
    }
    /**
     * @return Whether the group is a security group.
     * 
     */
    public Boolean securityEnabled() {
        return this.securityEnabled;
    }
    /**
     * @return The colour theme for a Microsoft 365 group. Possible values are `Blue`, `Green`, `Orange`, `Pink`, `Purple`, `Red` or `Teal`. When no theme is set, the value is `null`.
     * 
     */
    public String theme() {
        return this.theme;
    }
    /**
     * @return A list of group types configured for the group. Supported values are `DynamicMembership`, which denotes a group with dynamic membership, and `Unified`, which specifies a Microsoft 365 group.
     * 
     */
    public List<String> types() {
        return this.types;
    }
    /**
     * @return The group join policy and group content visibility. Possible values are `Private`, `Public`, or `Hiddenmembership`. Only Microsoft 365 groups can have `Hiddenmembership` visibility.
     * 
     */
    public String visibility() {
        return this.visibility;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetGroupResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean assignableToRole;
        private Boolean autoSubscribeNewMembers;
        private List<String> behaviors;
        private String description;
        private String displayName;
        private List<GetGroupDynamicMembership> dynamicMemberships;
        private Boolean externalSendersAllowed;
        private Boolean hideFromAddressLists;
        private Boolean hideFromOutlookClients;
        private String id;
        private String mail;
        private Boolean mailEnabled;
        private String mailNickname;
        private List<String> members;
        private String objectId;
        private String onpremisesDomainName;
        private String onpremisesNetbiosName;
        private String onpremisesSamAccountName;
        private String onpremisesSecurityIdentifier;
        private Boolean onpremisesSyncEnabled;
        private List<String> owners;
        private String preferredLanguage;
        private List<String> provisioningOptions;
        private List<String> proxyAddresses;
        private Boolean securityEnabled;
        private String theme;
        private List<String> types;
        private String visibility;
        public Builder() {}
        public Builder(GetGroupResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.assignableToRole = defaults.assignableToRole;
    	      this.autoSubscribeNewMembers = defaults.autoSubscribeNewMembers;
    	      this.behaviors = defaults.behaviors;
    	      this.description = defaults.description;
    	      this.displayName = defaults.displayName;
    	      this.dynamicMemberships = defaults.dynamicMemberships;
    	      this.externalSendersAllowed = defaults.externalSendersAllowed;
    	      this.hideFromAddressLists = defaults.hideFromAddressLists;
    	      this.hideFromOutlookClients = defaults.hideFromOutlookClients;
    	      this.id = defaults.id;
    	      this.mail = defaults.mail;
    	      this.mailEnabled = defaults.mailEnabled;
    	      this.mailNickname = defaults.mailNickname;
    	      this.members = defaults.members;
    	      this.objectId = defaults.objectId;
    	      this.onpremisesDomainName = defaults.onpremisesDomainName;
    	      this.onpremisesNetbiosName = defaults.onpremisesNetbiosName;
    	      this.onpremisesSamAccountName = defaults.onpremisesSamAccountName;
    	      this.onpremisesSecurityIdentifier = defaults.onpremisesSecurityIdentifier;
    	      this.onpremisesSyncEnabled = defaults.onpremisesSyncEnabled;
    	      this.owners = defaults.owners;
    	      this.preferredLanguage = defaults.preferredLanguage;
    	      this.provisioningOptions = defaults.provisioningOptions;
    	      this.proxyAddresses = defaults.proxyAddresses;
    	      this.securityEnabled = defaults.securityEnabled;
    	      this.theme = defaults.theme;
    	      this.types = defaults.types;
    	      this.visibility = defaults.visibility;
        }

        @CustomType.Setter
        public Builder assignableToRole(Boolean assignableToRole) {
            this.assignableToRole = Objects.requireNonNull(assignableToRole);
            return this;
        }
        @CustomType.Setter
        public Builder autoSubscribeNewMembers(Boolean autoSubscribeNewMembers) {
            this.autoSubscribeNewMembers = Objects.requireNonNull(autoSubscribeNewMembers);
            return this;
        }
        @CustomType.Setter
        public Builder behaviors(List<String> behaviors) {
            this.behaviors = Objects.requireNonNull(behaviors);
            return this;
        }
        public Builder behaviors(String... behaviors) {
            return behaviors(List.of(behaviors));
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
        public Builder dynamicMemberships(List<GetGroupDynamicMembership> dynamicMemberships) {
            this.dynamicMemberships = Objects.requireNonNull(dynamicMemberships);
            return this;
        }
        public Builder dynamicMemberships(GetGroupDynamicMembership... dynamicMemberships) {
            return dynamicMemberships(List.of(dynamicMemberships));
        }
        @CustomType.Setter
        public Builder externalSendersAllowed(Boolean externalSendersAllowed) {
            this.externalSendersAllowed = Objects.requireNonNull(externalSendersAllowed);
            return this;
        }
        @CustomType.Setter
        public Builder hideFromAddressLists(Boolean hideFromAddressLists) {
            this.hideFromAddressLists = Objects.requireNonNull(hideFromAddressLists);
            return this;
        }
        @CustomType.Setter
        public Builder hideFromOutlookClients(Boolean hideFromOutlookClients) {
            this.hideFromOutlookClients = Objects.requireNonNull(hideFromOutlookClients);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder mail(String mail) {
            this.mail = Objects.requireNonNull(mail);
            return this;
        }
        @CustomType.Setter
        public Builder mailEnabled(Boolean mailEnabled) {
            this.mailEnabled = Objects.requireNonNull(mailEnabled);
            return this;
        }
        @CustomType.Setter
        public Builder mailNickname(String mailNickname) {
            this.mailNickname = Objects.requireNonNull(mailNickname);
            return this;
        }
        @CustomType.Setter
        public Builder members(List<String> members) {
            this.members = Objects.requireNonNull(members);
            return this;
        }
        public Builder members(String... members) {
            return members(List.of(members));
        }
        @CustomType.Setter
        public Builder objectId(String objectId) {
            this.objectId = Objects.requireNonNull(objectId);
            return this;
        }
        @CustomType.Setter
        public Builder onpremisesDomainName(String onpremisesDomainName) {
            this.onpremisesDomainName = Objects.requireNonNull(onpremisesDomainName);
            return this;
        }
        @CustomType.Setter
        public Builder onpremisesNetbiosName(String onpremisesNetbiosName) {
            this.onpremisesNetbiosName = Objects.requireNonNull(onpremisesNetbiosName);
            return this;
        }
        @CustomType.Setter
        public Builder onpremisesSamAccountName(String onpremisesSamAccountName) {
            this.onpremisesSamAccountName = Objects.requireNonNull(onpremisesSamAccountName);
            return this;
        }
        @CustomType.Setter
        public Builder onpremisesSecurityIdentifier(String onpremisesSecurityIdentifier) {
            this.onpremisesSecurityIdentifier = Objects.requireNonNull(onpremisesSecurityIdentifier);
            return this;
        }
        @CustomType.Setter
        public Builder onpremisesSyncEnabled(Boolean onpremisesSyncEnabled) {
            this.onpremisesSyncEnabled = Objects.requireNonNull(onpremisesSyncEnabled);
            return this;
        }
        @CustomType.Setter
        public Builder owners(List<String> owners) {
            this.owners = Objects.requireNonNull(owners);
            return this;
        }
        public Builder owners(String... owners) {
            return owners(List.of(owners));
        }
        @CustomType.Setter
        public Builder preferredLanguage(String preferredLanguage) {
            this.preferredLanguage = Objects.requireNonNull(preferredLanguage);
            return this;
        }
        @CustomType.Setter
        public Builder provisioningOptions(List<String> provisioningOptions) {
            this.provisioningOptions = Objects.requireNonNull(provisioningOptions);
            return this;
        }
        public Builder provisioningOptions(String... provisioningOptions) {
            return provisioningOptions(List.of(provisioningOptions));
        }
        @CustomType.Setter
        public Builder proxyAddresses(List<String> proxyAddresses) {
            this.proxyAddresses = Objects.requireNonNull(proxyAddresses);
            return this;
        }
        public Builder proxyAddresses(String... proxyAddresses) {
            return proxyAddresses(List.of(proxyAddresses));
        }
        @CustomType.Setter
        public Builder securityEnabled(Boolean securityEnabled) {
            this.securityEnabled = Objects.requireNonNull(securityEnabled);
            return this;
        }
        @CustomType.Setter
        public Builder theme(String theme) {
            this.theme = Objects.requireNonNull(theme);
            return this;
        }
        @CustomType.Setter
        public Builder types(List<String> types) {
            this.types = Objects.requireNonNull(types);
            return this;
        }
        public Builder types(String... types) {
            return types(List.of(types));
        }
        @CustomType.Setter
        public Builder visibility(String visibility) {
            this.visibility = Objects.requireNonNull(visibility);
            return this;
        }
        public GetGroupResult build() {
            final var o = new GetGroupResult();
            o.assignableToRole = assignableToRole;
            o.autoSubscribeNewMembers = autoSubscribeNewMembers;
            o.behaviors = behaviors;
            o.description = description;
            o.displayName = displayName;
            o.dynamicMemberships = dynamicMemberships;
            o.externalSendersAllowed = externalSendersAllowed;
            o.hideFromAddressLists = hideFromAddressLists;
            o.hideFromOutlookClients = hideFromOutlookClients;
            o.id = id;
            o.mail = mail;
            o.mailEnabled = mailEnabled;
            o.mailNickname = mailNickname;
            o.members = members;
            o.objectId = objectId;
            o.onpremisesDomainName = onpremisesDomainName;
            o.onpremisesNetbiosName = onpremisesNetbiosName;
            o.onpremisesSamAccountName = onpremisesSamAccountName;
            o.onpremisesSecurityIdentifier = onpremisesSecurityIdentifier;
            o.onpremisesSyncEnabled = onpremisesSyncEnabled;
            o.owners = owners;
            o.preferredLanguage = preferredLanguage;
            o.provisioningOptions = provisioningOptions;
            o.proxyAddresses = proxyAddresses;
            o.securityEnabled = securityEnabled;
            o.theme = theme;
            o.types = types;
            o.visibility = visibility;
            return o;
        }
    }
}
