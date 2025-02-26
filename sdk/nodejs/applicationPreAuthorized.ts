// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const authorized = new azuread.Application("authorized", {displayName: "example-authorized-app"});
 * const authorizer = new azuread.Application("authorizer", {
 *     displayName: "example-authorizing-app",
 *     api: {
 *         oauth2PermissionScopes: [
 *             {
 *                 adminConsentDescription: "Administer the application",
 *                 adminConsentDisplayName: "Administer",
 *                 enabled: true,
 *                 id: "ced9c4c3-c273-4f0f-ac71-a20377b90f9c",
 *                 type: "Admin",
 *                 value: "administer",
 *             },
 *             {
 *                 adminConsentDescription: "Access the application",
 *                 adminConsentDisplayName: "Access",
 *                 enabled: true,
 *                 id: "2d5e07ca-664d-4d9b-ad61-ec07fd215213",
 *                 type: "User",
 *                 userConsentDescription: "Access the application",
 *                 userConsentDisplayName: "Access",
 *                 value: "user_impersonation",
 *             },
 *         ],
 *     },
 * });
 * const example = new azuread.ApplicationPreAuthorized("example", {
 *     applicationObjectId: authorizer.objectId,
 *     authorizedAppId: authorized.applicationId,
 *     permissionIds: [
 *         "ced9c4c3-c273-4f0f-ac71-a20377b90f9c",
 *         "2d5e07ca-664d-4d9b-ad61-ec07fd215213",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Pre-authorized applications can be imported using the object ID of the authorizing application and the application ID of the application being authorized, e.g.
 *
 * ```sh
 *  $ pulumi import azuread:index/applicationPreAuthorized:ApplicationPreAuthorized example 00000000-0000-0000-0000-000000000000/preAuthorizedApplication/11111111-1111-1111-1111-111111111111
 * ```
 *
 *  -> This ID format is unique to Terraform and is composed of the authorizing application's object ID, the string "preAuthorizedApplication" and the authorized application's application ID (client ID) in the format `{ObjectId}/preAuthorizedApplication/{ApplicationId}`.
 */
export class ApplicationPreAuthorized extends pulumi.CustomResource {
    /**
     * Get an existing ApplicationPreAuthorized resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationPreAuthorizedState, opts?: pulumi.CustomResourceOptions): ApplicationPreAuthorized {
        return new ApplicationPreAuthorized(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/applicationPreAuthorized:ApplicationPreAuthorized';

    /**
     * Returns true if the given object is an instance of ApplicationPreAuthorized.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApplicationPreAuthorized {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApplicationPreAuthorized.__pulumiType;
    }

    /**
     * The object ID of the application for which permissions are being authorized. Changing this field forces a new resource to be created.
     */
    public readonly applicationObjectId!: pulumi.Output<string>;
    /**
     * The application ID of the pre-authorized application
     */
    public readonly authorizedAppId!: pulumi.Output<string>;
    /**
     * A set of permission scope IDs required by the authorized application.
     */
    public readonly permissionIds!: pulumi.Output<string[]>;

    /**
     * Create a ApplicationPreAuthorized resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationPreAuthorizedArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationPreAuthorizedArgs | ApplicationPreAuthorizedState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApplicationPreAuthorizedState | undefined;
            resourceInputs["applicationObjectId"] = state ? state.applicationObjectId : undefined;
            resourceInputs["authorizedAppId"] = state ? state.authorizedAppId : undefined;
            resourceInputs["permissionIds"] = state ? state.permissionIds : undefined;
        } else {
            const args = argsOrState as ApplicationPreAuthorizedArgs | undefined;
            if ((!args || args.applicationObjectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationObjectId'");
            }
            if ((!args || args.authorizedAppId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authorizedAppId'");
            }
            if ((!args || args.permissionIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'permissionIds'");
            }
            resourceInputs["applicationObjectId"] = args ? args.applicationObjectId : undefined;
            resourceInputs["authorizedAppId"] = args ? args.authorizedAppId : undefined;
            resourceInputs["permissionIds"] = args ? args.permissionIds : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ApplicationPreAuthorized.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApplicationPreAuthorized resources.
 */
export interface ApplicationPreAuthorizedState {
    /**
     * The object ID of the application for which permissions are being authorized. Changing this field forces a new resource to be created.
     */
    applicationObjectId?: pulumi.Input<string>;
    /**
     * The application ID of the pre-authorized application
     */
    authorizedAppId?: pulumi.Input<string>;
    /**
     * A set of permission scope IDs required by the authorized application.
     */
    permissionIds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a ApplicationPreAuthorized resource.
 */
export interface ApplicationPreAuthorizedArgs {
    /**
     * The object ID of the application for which permissions are being authorized. Changing this field forces a new resource to be created.
     */
    applicationObjectId: pulumi.Input<string>;
    /**
     * The application ID of the pre-authorized application
     */
    authorizedAppId: pulumi.Input<string>;
    /**
     * A set of permission scope IDs required by the authorized application.
     */
    permissionIds: pulumi.Input<pulumi.Input<string>[]>;
}
