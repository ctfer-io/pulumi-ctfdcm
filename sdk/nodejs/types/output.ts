// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface ChallengeDynamicIaCRequirements {
    /**
     * Behavior if not unlocked, either hidden or anonymized.
     */
    behavior: string;
    /**
     * List of the challenges ID.
     */
    prerequisites?: string[];
}

