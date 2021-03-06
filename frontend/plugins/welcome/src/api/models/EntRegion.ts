/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Playlist Vidoe
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntRegionEdges,
    EntRegionEdgesFromJSON,
    EntRegionEdgesFromJSONTyped,
    EntRegionEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntRegion
 */
export interface EntRegion {
    /**
     * 
     * @type {EntRegionEdges}
     * @memberof EntRegion
     */
    edges?: EntRegionEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntRegion
     */
    id?: number;
    /**
     * Name holds the value of the "name" field.
     * @type {string}
     * @memberof EntRegion
     */
    name?: string;
}

export function EntRegionFromJSON(json: any): EntRegion {
    return EntRegionFromJSONTyped(json, false);
}

export function EntRegionFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntRegion {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntRegionEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'name': !exists(json, 'name') ? undefined : json['name'],
    };
}

export function EntRegionToJSON(value?: EntRegion | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntRegionEdgesToJSON(value.edges),
        'id': value.id,
        'name': value.name,
    };
}


