// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a delegated permission grant for a service principal, on behalf of a single user, or all users.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this resource.
 *
 * When authenticated with a service principal, this resource requires the following application role: `Directory.ReadWrite.All`
 *
 * When authenticated with a user principal, this resource requires one the following directory role: `Global Administrator`
 *
 * ## Example Usage
 *
 * *Delegated permission grant for all users*
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const wellKnown = azuread.getApplicationPublishedAppIds({});
 * const msgraph = new azuread.ServicePrincipal("msgraph", {
 *     applicationId: wellKnown.then(wellKnown => wellKnown.result?.MicrosoftGraph),
 *     useExisting: true,
 * });
 * const exampleApplication = new azuread.Application("exampleApplication", {
 *     displayName: "example",
 *     requiredResourceAccesses: [{
 *         resourceAppId: wellKnown.then(wellKnown => wellKnown.result?.MicrosoftGraph),
 *         resourceAccesses: [
 *             {
 *                 id: msgraph.oauth2PermissionScopeIds.openid,
 *                 type: "Scope",
 *             },
 *             {
 *                 id: msgraph.oauth2PermissionScopeIds["User.Read"],
 *                 type: "Scope",
 *             },
 *         ],
 *     }],
 * });
 * const exampleServicePrincipal = new azuread.ServicePrincipal("exampleServicePrincipal", {applicationId: exampleApplication.applicationId});
 * const exampleServicePrincipalDelegatedPermissionGrant = new azuread.ServicePrincipalDelegatedPermissionGrant("exampleServicePrincipalDelegatedPermissionGrant", {
 *     servicePrincipalObjectId: exampleServicePrincipal.objectId,
 *     resourceServicePrincipalObjectId: msgraph.objectId,
 *     claimValues: [
 *         "openid",
 *         "User.Read.All",
 *     ],
 * });
 * ```
 *
 * *Delegated permission grant for a single user*
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const wellKnown = azuread.getApplicationPublishedAppIds({});
 * const msgraph = new azuread.ServicePrincipal("msgraph", {
 *     applicationId: wellKnown.then(wellKnown => wellKnown.result?.MicrosoftGraph),
 *     useExisting: true,
 * });
 * const exampleApplication = new azuread.Application("exampleApplication", {
 *     displayName: "example",
 *     requiredResourceAccesses: [{
 *         resourceAppId: wellKnown.then(wellKnown => wellKnown.result?.MicrosoftGraph),
 *         resourceAccesses: [
 *             {
 *                 id: msgraph.oauth2PermissionScopeIds.openid,
 *                 type: "Scope",
 *             },
 *             {
 *                 id: msgraph.oauth2PermissionScopeIds["User.Read"],
 *                 type: "Scope",
 *             },
 *         ],
 *     }],
 * });
 * const exampleServicePrincipal = new azuread.ServicePrincipal("exampleServicePrincipal", {applicationId: exampleApplication.applicationId});
 * const exampleUser = new azuread.User("exampleUser", {
 *     displayName: "J. Doe",
 *     userPrincipalName: "jdoe@hashicorp.com",
 *     mailNickname: "jdoe",
 *     password: "SecretP@sswd99!",
 * });
 * const exampleServicePrincipalDelegatedPermissionGrant = new azuread.ServicePrincipalDelegatedPermissionGrant("exampleServicePrincipalDelegatedPermissionGrant", {
 *     servicePrincipalObjectId: exampleServicePrincipal.objectId,
 *     resourceServicePrincipalObjectId: msgraph.objectId,
 *     claimValues: [
 *         "openid",
 *         "User.Read.All",
 *     ],
 *     userObjectId: exampleUser.objectId,
 * });
 * ```
 *
 * ## Import
 *
 * Delegated permission grants can be imported using their ID, e.g.
 *
 * ```sh
 *  $ pulumi import azuread:index/servicePrincipalDelegatedPermissionGrant:ServicePrincipalDelegatedPermissionGrant example aaBBcDDeFG6h5JKLMN2PQrrssTTUUvWWxxxxxyyyzzz
 * ```
 */
export class ServicePrincipalDelegatedPermissionGrant extends pulumi.CustomResource {
    /**
     * Get an existing ServicePrincipalDelegatedPermissionGrant resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServicePrincipalDelegatedPermissionGrantState, opts?: pulumi.CustomResourceOptions): ServicePrincipalDelegatedPermissionGrant {
        return new ServicePrincipalDelegatedPermissionGrant(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/servicePrincipalDelegatedPermissionGrant:ServicePrincipalDelegatedPermissionGrant';

    /**
     * Returns true if the given object is an instance of ServicePrincipalDelegatedPermissionGrant.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServicePrincipalDelegatedPermissionGrant {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServicePrincipalDelegatedPermissionGrant.__pulumiType;
    }

    /**
     * - A set of claim values for delegated permission scopes which should be included in access tokens for the resource.
     */
    public readonly claimValues!: pulumi.Output<string[]>;
    /**
     * The object ID of the service principal representing the resource to be accessed. Changing this forces a new resource to be created.
     */
    public readonly resourceServicePrincipalObjectId!: pulumi.Output<string>;
    /**
     * The object ID of the service principal for which this delegated permission grant should be created. Changing this forces a new resource to be created.
     */
    public readonly servicePrincipalObjectId!: pulumi.Output<string>;
    /**
     * - The object ID of the user on behalf of whom the service principal is authorized to access the resource. When omitted, the delegated permission grant will be consented for all users. Changing this forces a new resource to be created.
     */
    public readonly userObjectId!: pulumi.Output<string | undefined>;

    /**
     * Create a ServicePrincipalDelegatedPermissionGrant resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServicePrincipalDelegatedPermissionGrantArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServicePrincipalDelegatedPermissionGrantArgs | ServicePrincipalDelegatedPermissionGrantState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServicePrincipalDelegatedPermissionGrantState | undefined;
            resourceInputs["claimValues"] = state ? state.claimValues : undefined;
            resourceInputs["resourceServicePrincipalObjectId"] = state ? state.resourceServicePrincipalObjectId : undefined;
            resourceInputs["servicePrincipalObjectId"] = state ? state.servicePrincipalObjectId : undefined;
            resourceInputs["userObjectId"] = state ? state.userObjectId : undefined;
        } else {
            const args = argsOrState as ServicePrincipalDelegatedPermissionGrantArgs | undefined;
            if ((!args || args.claimValues === undefined) && !opts.urn) {
                throw new Error("Missing required property 'claimValues'");
            }
            if ((!args || args.resourceServicePrincipalObjectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceServicePrincipalObjectId'");
            }
            if ((!args || args.servicePrincipalObjectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'servicePrincipalObjectId'");
            }
            resourceInputs["claimValues"] = args ? args.claimValues : undefined;
            resourceInputs["resourceServicePrincipalObjectId"] = args ? args.resourceServicePrincipalObjectId : undefined;
            resourceInputs["servicePrincipalObjectId"] = args ? args.servicePrincipalObjectId : undefined;
            resourceInputs["userObjectId"] = args ? args.userObjectId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServicePrincipalDelegatedPermissionGrant.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServicePrincipalDelegatedPermissionGrant resources.
 */
export interface ServicePrincipalDelegatedPermissionGrantState {
    /**
     * - A set of claim values for delegated permission scopes which should be included in access tokens for the resource.
     */
    claimValues?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The object ID of the service principal representing the resource to be accessed. Changing this forces a new resource to be created.
     */
    resourceServicePrincipalObjectId?: pulumi.Input<string>;
    /**
     * The object ID of the service principal for which this delegated permission grant should be created. Changing this forces a new resource to be created.
     */
    servicePrincipalObjectId?: pulumi.Input<string>;
    /**
     * - The object ID of the user on behalf of whom the service principal is authorized to access the resource. When omitted, the delegated permission grant will be consented for all users. Changing this forces a new resource to be created.
     */
    userObjectId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServicePrincipalDelegatedPermissionGrant resource.
 */
export interface ServicePrincipalDelegatedPermissionGrantArgs {
    /**
     * - A set of claim values for delegated permission scopes which should be included in access tokens for the resource.
     */
    claimValues: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The object ID of the service principal representing the resource to be accessed. Changing this forces a new resource to be created.
     */
    resourceServicePrincipalObjectId: pulumi.Input<string>;
    /**
     * The object ID of the service principal for which this delegated permission grant should be created. Changing this forces a new resource to be created.
     */
    servicePrincipalObjectId: pulumi.Input<string>;
    /**
     * - The object ID of the user on behalf of whom the service principal is authorized to access the resource. When omitted, the delegated permission grant will be consented for all users. Changing this forces a new resource to be created.
     */
    userObjectId?: pulumi.Input<string>;
}
