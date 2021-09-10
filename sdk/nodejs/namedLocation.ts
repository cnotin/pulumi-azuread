// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Manages a Named Location within Azure Active Directory.
 *
 * ## API Permissions
 *
 * The following API permissions are required in order to use this resource.
 *
 * When authenticated with a service principal, this resource requires the following application roles: `Policy.ReadWrite.ConditionalAccess` and `Policy.Read.All`
 *
 * When authenticated with a user principal, this resource requires one of the following directory roles: `Conditional Access Administrator` or `Global Administrator`
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azuread from "@pulumi/azuread";
 *
 * const example_ip = new azuread.NamedLocation("example-ip", {
 *     displayName: "IP Named Location",
 *     ip: {
 *         ipRanges: [
 *             "1.1.1.1/32",
 *             "2.2.2.2/32",
 *         ],
 *         trusted: true,
 *     },
 * });
 * const example_country = new azuread.NamedLocation("example-country", {
 *     country: {
 *         countriesAndRegions: [
 *             "GB",
 *             "US",
 *         ],
 *         includeUnknownCountriesAndRegions: false,
 *     },
 *     displayName: "Country Named Location",
 * });
 * ```
 *
 * ## Import
 *
 * Named Locations can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import azuread:index/namedLocation:NamedLocation my_location 00000000-0000-0000-0000-000000000000
 * ```
 */
export class NamedLocation extends pulumi.CustomResource {
    /**
     * Get an existing NamedLocation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NamedLocationState, opts?: pulumi.CustomResourceOptions): NamedLocation {
        return new NamedLocation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azuread:index/namedLocation:NamedLocation';

    /**
     * Returns true if the given object is an instance of NamedLocation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NamedLocation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NamedLocation.__pulumiType;
    }

    /**
     * A `country` block as documented below, which configures a country-based named location.
     */
    public readonly country!: pulumi.Output<outputs.NamedLocationCountry | undefined>;
    /**
     * The friendly name for this named location.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * An `ip` block as documented below, which configures an IP-based named location.
     */
    public readonly ip!: pulumi.Output<outputs.NamedLocationIp | undefined>;

    /**
     * Create a NamedLocation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NamedLocationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NamedLocationArgs | NamedLocationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NamedLocationState | undefined;
            inputs["country"] = state ? state.country : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["ip"] = state ? state.ip : undefined;
        } else {
            const args = argsOrState as NamedLocationArgs | undefined;
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            inputs["country"] = args ? args.country : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["ip"] = args ? args.ip : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(NamedLocation.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NamedLocation resources.
 */
export interface NamedLocationState {
    /**
     * A `country` block as documented below, which configures a country-based named location.
     */
    country?: pulumi.Input<inputs.NamedLocationCountry>;
    /**
     * The friendly name for this named location.
     */
    displayName?: pulumi.Input<string>;
    /**
     * An `ip` block as documented below, which configures an IP-based named location.
     */
    ip?: pulumi.Input<inputs.NamedLocationIp>;
}

/**
 * The set of arguments for constructing a NamedLocation resource.
 */
export interface NamedLocationArgs {
    /**
     * A `country` block as documented below, which configures a country-based named location.
     */
    country?: pulumi.Input<inputs.NamedLocationCountry>;
    /**
     * The friendly name for this named location.
     */
    displayName: pulumi.Input<string>;
    /**
     * An `ip` block as documented below, which configures an IP-based named location.
     */
    ip?: pulumi.Input<inputs.NamedLocationIp>;
}
