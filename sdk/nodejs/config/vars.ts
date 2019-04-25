// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

let __config = new pulumi.Config("azuread");

export let clientCertificatePassword: string | undefined = __config.get("clientCertificatePassword");
export let clientCertificatePath: string | undefined = __config.get("clientCertificatePath");
export let clientId: string | undefined = __config.get("clientId");
export let clientSecret: string | undefined = __config.get("clientSecret");
export let environment: string = __config.require("environment");
export let msiEndpoint: string | undefined = __config.get("msiEndpoint");
export let subscriptionId: string | undefined = __config.get("subscriptionId");
export let tenantId: string | undefined = __config.get("tenantId");
export let useMsi: boolean | undefined = __config.getObject<boolean>("useMsi");
