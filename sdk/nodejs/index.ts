// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export { ChallengeDynamicIaCArgs, ChallengeDynamicIaCState } from "./challengeDynamicIaC";
export type ChallengeDynamicIaC = import("./challengeDynamicIaC").ChallengeDynamicIaC;
export const ChallengeDynamicIaC: typeof import("./challengeDynamicIaC").ChallengeDynamicIaC = null as any;
utilities.lazyLoad(exports, ["ChallengeDynamicIaC"], () => require("./challengeDynamicIaC"));

export * from "./provider";
import { Provider } from "./provider";


// Export sub-modules:
import * as config from "./config";
import * as types from "./types";

export {
    config,
    types,
};

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "ctfdcm:index/challengeDynamicIaC:ChallengeDynamicIaC":
                return new ChallengeDynamicIaC(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("ctfdcm", "index/challengeDynamicIaC", _module)
pulumi.runtime.registerResourcePackage("ctfdcm", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:ctfdcm") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
