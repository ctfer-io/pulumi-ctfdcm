// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getTeams(opts?: pulumi.InvokeOptions): Promise<GetTeamsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ctfd:index/getTeams:getTeams", {
    }, opts);
}

/**
 * A collection of values returned by getTeams.
 */
export interface GetTeamsResult {
    /**
     * Affiliation to a company or agency.
     */
    readonly affiliation: string;
    /**
     * Is true if the team is banned from the CTF.
     */
    readonly banned: boolean;
    /**
     * Member who is captain of the team. Must be part of the members too. Note it could cause a fatal error in case of resource import with an inconsistent CTFd configuration i.e. if a team has no captain yet (should not be possible).
     */
    readonly captain: string;
    /**
     * Country the team represent or is hail from.
     */
    readonly country: string;
    /**
     * Email of the team.
     */
    readonly email: string;
    /**
     * Is true if the team is hidden to the participants.
     */
    readonly hidden: boolean;
    /**
     * Identifier of the user.
     */
    readonly id: string;
    /**
     * List of members (User), defined by their IDs.
     */
    readonly members: string[];
    /**
     * Name of the team.
     */
    readonly name: string;
    /**
     * Password of the team. Notice that during a CTF you may not want to update those to avoid defaulting team accesses.
     */
    readonly password: string;
    /**
     * Website, blog, or anything similar (displayed to other participants).
     */
    readonly website: string;
}
export function getTeamsOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetTeamsResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ctfd:index/getTeams:getTeams", {
    }, opts);
}
