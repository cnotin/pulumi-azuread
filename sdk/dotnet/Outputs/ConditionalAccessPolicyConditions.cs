// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD.Outputs
{

    [OutputType]
    public sealed class ConditionalAccessPolicyConditions
    {
        /// <summary>
        /// An `applications` block as documented below, which specifies applications and user actions included in and excluded from the policy.
        /// </summary>
        public readonly Outputs.ConditionalAccessPolicyConditionsApplications Applications;
        /// <summary>
        /// A list of client application types included in the policy. Possible values are: `all`, `browser`, `mobileAppsAndDesktopClients`, `exchangeActiveSync`, `easSupported` and `other`.
        /// </summary>
        public readonly ImmutableArray<string> ClientAppTypes;
        /// <summary>
        /// A `devices` block as documented below, which describes devices to be included in and excluded from the policy. A `devices` block can be added to an existing policy, but removing the `devices` block forces a new resource to be created.
        /// </summary>
        public readonly Outputs.ConditionalAccessPolicyConditionsDevices? Devices;
        /// <summary>
        /// A `locations` block as documented below, which specifies locations included in and excluded from the policy.
        /// </summary>
        public readonly Outputs.ConditionalAccessPolicyConditionsLocations Locations;
        /// <summary>
        /// A `platforms` block as documented below, which specifies platforms included in and excluded from the policy.
        /// </summary>
        public readonly Outputs.ConditionalAccessPolicyConditionsPlatforms Platforms;
        /// <summary>
        /// A list of sign-in risk levels included in the policy. Possible values are: `low`, `medium`, `high`, `hidden`, `none`, `unknownFutureValue`.
        /// </summary>
        public readonly ImmutableArray<string> SignInRiskLevels;
        /// <summary>
        /// A list of user risk levels included in the policy. Possible values are: `low`, `medium`, `high`, `hidden`, `none`, `unknownFutureValue`.
        /// </summary>
        public readonly ImmutableArray<string> UserRiskLevels;
        /// <summary>
        /// A `users` block as documented below, which specifies users, groups, and roles included in and excluded from the policy.
        /// </summary>
        public readonly Outputs.ConditionalAccessPolicyConditionsUsers Users;

        [OutputConstructor]
        private ConditionalAccessPolicyConditions(
            Outputs.ConditionalAccessPolicyConditionsApplications applications,

            ImmutableArray<string> clientAppTypes,

            Outputs.ConditionalAccessPolicyConditionsDevices? devices,

            Outputs.ConditionalAccessPolicyConditionsLocations locations,

            Outputs.ConditionalAccessPolicyConditionsPlatforms platforms,

            ImmutableArray<string> signInRiskLevels,

            ImmutableArray<string> userRiskLevels,

            Outputs.ConditionalAccessPolicyConditionsUsers users)
        {
            Applications = applications;
            ClientAppTypes = clientAppTypes;
            Devices = devices;
            Locations = locations;
            Platforms = platforms;
            SignInRiskLevels = signInRiskLevels;
            UserRiskLevels = userRiskLevels;
            Users = users;
        }
    }
}
