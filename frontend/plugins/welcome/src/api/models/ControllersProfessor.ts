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
/**
 * 
 * @export
 * @interface ControllersProfessor
 */
export interface ControllersProfessor {
    /**
     * 
     * @type {string}
     * @memberof ControllersProfessor
     */
    email?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersProfessor
     */
    faculty?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersProfessor
     */
    name?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersProfessor
     */
    prefix?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersProfessor
     */
    professorship?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersProfessor
     */
    tel?: string;
}

export function ControllersProfessorFromJSON(json: any): ControllersProfessor {
    return ControllersProfessorFromJSONTyped(json, false);
}

export function ControllersProfessorFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersProfessor {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'email': !exists(json, 'email') ? undefined : json['email'],
        'faculty': !exists(json, 'faculty') ? undefined : json['faculty'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'prefix': !exists(json, 'prefix') ? undefined : json['prefix'],
        'professorship': !exists(json, 'professorship') ? undefined : json['professorship'],
        'tel': !exists(json, 'tel') ? undefined : json['tel'],
    };
}

export function ControllersProfessorToJSON(value?: ControllersProfessor | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'email': value.email,
        'faculty': value.faculty,
        'name': value.name,
        'prefix': value.prefix,
        'professorship': value.professorship,
        'tel': value.tel,
    };
}


