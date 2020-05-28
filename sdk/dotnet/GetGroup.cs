// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureAD
{
    public static class GetGroup
    {
        /// <summary>
        /// Gets information about an Azure Active Directory group.
        /// 
        /// &gt; **NOTE:** If you're authenticating using a Service Principal then it must have permissions to `Read directory data` within the `Windows Azure Active Directory` API.
        /// 
        /// ## Example Usage (by Group Display Name)
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using AzureAD = Pulumi.AzureAD;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(AzureAD.GetGroup.InvokeAsync(new AzureAD.GetGroupArgs
        ///         {
        ///             Name = "A-AD-Group",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// </summary>
        public static Task<GetGroupResult> InvokeAsync(GetGroupArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGroupResult>("azuread:index/getGroup:getGroup", args ?? new GetGroupArgs(), options.WithVersion());
    }


    public sealed class GetGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Name of the AD Group we want to lookup.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// Specifies the Object ID of the AD Group within Azure Active Directory.
        /// </summary>
        [Input("objectId")]
        public string? ObjectId { get; set; }

        public GetGroupArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetGroupResult
    {
        /// <summary>
        /// The description of the AD Group.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Object IDs of the Azure AD Group members.
        /// </summary>
        public readonly ImmutableArray<string> Members;
        /// <summary>
        /// The name of the Azure AD Group.
        /// </summary>
        public readonly string Name;
        public readonly string ObjectId;
        /// <summary>
        /// The Object IDs of the Azure AD Group owners.
        /// </summary>
        public readonly ImmutableArray<string> Owners;

        [OutputConstructor]
        private GetGroupResult(
            string description,

            string id,

            ImmutableArray<string> members,

            string name,

            string objectId,

            ImmutableArray<string> owners)
        {
            Description = description;
            Id = id;
            Members = members;
            Name = name;
            ObjectId = objectId;
            Owners = owners;
        }
    }
}
