// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * Federated Identity Credentials can be imported using the object ID of the associated application and the ID of the federated identity credential, e.g.
 *
 * ```sh
 *  $ pulumi import azuread:index/applicationFederatedIdentityCredential:ApplicationFederatedIdentityCredential test 00000000-0000-0000-0000-000000000000/federatedIdentityCredential/11111111-1111-1111-1111-111111111111
 * ```
 *
 *  -> This ID format is unique to Terraform and is composed of the application's object ID, the string "federatedIdentityCredential" and the credential ID in the format `{ObjectId}/federatedIdentityCredential/{CredentialId}`.
 */
export class ApplicationFederatedIdentityCredential extends pulumi.CustomResource {
    /**
     * Get an existing ApplicationFederatedIdentityCredential resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationFederatedIdentityCredentialState, opts?: pulumi.CustomResourceOptions): ApplicationFederatedIdentityCredential {
        return new ApplicationFederatedIdentityCredential(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/applicationFederatedIdentityCredential:ApplicationFederatedIdentityCredential';

    /**
     * Returns true if the given object is an instance of ApplicationFederatedIdentityCredential.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApplicationFederatedIdentityCredential {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApplicationFederatedIdentityCredential.__pulumiType;
    }

    /**
     * The object ID of the application for which this federated identity credential should be created. Changing this field forces a new resource to be created.
     */
    public readonly applicationObjectId!: pulumi.Output<string>;
    /**
     * List of audiences that can appear in the external token. This specifies what should be accepted in the `aud` claim of incoming tokens.
     */
    public readonly audiences!: pulumi.Output<string[]>;
    /**
     * A UUID used to uniquely identify this federated identity credential.
     */
    public /*out*/ readonly credentialId!: pulumi.Output<string>;
    /**
     * A description for the federated identity credential.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A unique display name for the federated identity credential. Changing this forces a new resource to be created.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The URL of the external identity provider, which must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app.
     */
    public readonly issuer!: pulumi.Output<string>;
    /**
     * The identifier of the external software workload within the external identity provider. The combination of issuer and subject must be unique on the app.
     */
    public readonly subject!: pulumi.Output<string>;

    /**
     * Create a ApplicationFederatedIdentityCredential resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApplicationFederatedIdentityCredentialArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationFederatedIdentityCredentialArgs | ApplicationFederatedIdentityCredentialState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApplicationFederatedIdentityCredentialState | undefined;
            resourceInputs["applicationObjectId"] = state ? state.applicationObjectId : undefined;
            resourceInputs["audiences"] = state ? state.audiences : undefined;
            resourceInputs["credentialId"] = state ? state.credentialId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["issuer"] = state ? state.issuer : undefined;
            resourceInputs["subject"] = state ? state.subject : undefined;
        } else {
            const args = argsOrState as ApplicationFederatedIdentityCredentialArgs | undefined;
            if ((!args || args.applicationObjectId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationObjectId'");
            }
            if ((!args || args.audiences === undefined) && !opts.urn) {
                throw new Error("Missing required property 'audiences'");
            }
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.issuer === undefined) && !opts.urn) {
                throw new Error("Missing required property 'issuer'");
            }
            if ((!args || args.subject === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subject'");
            }
            resourceInputs["applicationObjectId"] = args ? args.applicationObjectId : undefined;
            resourceInputs["audiences"] = args ? args.audiences : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["issuer"] = args ? args.issuer : undefined;
            resourceInputs["subject"] = args ? args.subject : undefined;
            resourceInputs["credentialId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ApplicationFederatedIdentityCredential.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApplicationFederatedIdentityCredential resources.
 */
export interface ApplicationFederatedIdentityCredentialState {
    /**
     * The object ID of the application for which this federated identity credential should be created. Changing this field forces a new resource to be created.
     */
    applicationObjectId?: pulumi.Input<string>;
    /**
     * List of audiences that can appear in the external token. This specifies what should be accepted in the `aud` claim of incoming tokens.
     */
    audiences?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A UUID used to uniquely identify this federated identity credential.
     */
    credentialId?: pulumi.Input<string>;
    /**
     * A description for the federated identity credential.
     */
    description?: pulumi.Input<string>;
    /**
     * A unique display name for the federated identity credential. Changing this forces a new resource to be created.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The URL of the external identity provider, which must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app.
     */
    issuer?: pulumi.Input<string>;
    /**
     * The identifier of the external software workload within the external identity provider. The combination of issuer and subject must be unique on the app.
     */
    subject?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ApplicationFederatedIdentityCredential resource.
 */
export interface ApplicationFederatedIdentityCredentialArgs {
    /**
     * The object ID of the application for which this federated identity credential should be created. Changing this field forces a new resource to be created.
     */
    applicationObjectId: pulumi.Input<string>;
    /**
     * List of audiences that can appear in the external token. This specifies what should be accepted in the `aud` claim of incoming tokens.
     */
    audiences: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A description for the federated identity credential.
     */
    description?: pulumi.Input<string>;
    /**
     * A unique display name for the federated identity credential. Changing this forces a new resource to be created.
     */
    displayName: pulumi.Input<string>;
    /**
     * The URL of the external identity provider, which must match the issuer claim of the external token being exchanged. The combination of the values of issuer and subject must be unique on the app.
     */
    issuer: pulumi.Input<string>;
    /**
     * The identifier of the external software workload within the external identity provider. The combination of issuer and subject must be unique on the app.
     */
    subject: pulumi.Input<string>;
}
