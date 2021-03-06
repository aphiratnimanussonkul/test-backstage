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
    EntPostalEdges,
    EntPostalEdgesFromJSON,
    EntPostalEdgesFromJSONTyped,
    EntPostalEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPostal
 */
export interface EntPostal {
    /**
     * 
     * @type {EntPostalEdges}
     * @memberof EntPostal
     */
    edges?: EntPostalEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntPostal
     */
    id?: number;
    /**
     * Postal holds the value of the "postal" field.
     * @type {string}
     * @memberof EntPostal
     */
    postal?: string;
}

export function EntPostalFromJSON(json: any): EntPostal {
    return EntPostalFromJSONTyped(json, false);
}

export function EntPostalFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPostal {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntPostalEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'postal': !exists(json, 'postal') ? undefined : json['postal'],
    };
}

export function EntPostalToJSON(value?: EntPostal | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntPostalEdgesToJSON(value.edges),
        'id': value.id,
        'postal': value.postal,
    };
}


