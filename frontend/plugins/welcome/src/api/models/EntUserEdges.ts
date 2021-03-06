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
    EntBill,
    EntBillFromJSON,
    EntBillFromJSONTyped,
    EntBillToJSON,
    EntConfirmation,
    EntConfirmationFromJSON,
    EntConfirmationFromJSONTyped,
    EntConfirmationToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntUserEdges
 */
export interface EntUserEdges {
    /**
     * Bill holds the value of the bill edge.
     * @type {Array<EntBill>}
     * @memberof EntUserEdges
     */
    bill?: Array<EntBill>;
    /**
     * Confirmation holds the value of the confirmation edge.
     * @type {Array<EntConfirmation>}
     * @memberof EntUserEdges
     */
    confirmation?: Array<EntConfirmation>;
}

export function EntUserEdgesFromJSON(json: any): EntUserEdges {
    return EntUserEdgesFromJSONTyped(json, false);
}

export function EntUserEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntUserEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'bill': !exists(json, 'bill') ? undefined : ((json['bill'] as Array<any>).map(EntBillFromJSON)),
        'confirmation': !exists(json, 'confirmation') ? undefined : ((json['confirmation'] as Array<any>).map(EntConfirmationFromJSON)),
    };
}

export function EntUserEdgesToJSON(value?: EntUserEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'bill': value.bill === undefined ? undefined : ((value.bill as Array<any>).map(EntBillToJSON)),
        'confirmation': value.confirmation === undefined ? undefined : ((value.confirmation as Array<any>).map(EntConfirmationToJSON)),
    };
}


