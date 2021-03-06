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
    EntBillEdges,
    EntBillEdgesFromJSON,
    EntBillEdgesFromJSONTyped,
    EntBillEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntBill
 */
export interface EntBill {
    /**
     * PayNumber holds the value of the "Pay_number" field.
     * @type {number}
     * @memberof EntBill
     */
    payNumber?: number;
    /**
     * PayTime holds the value of the "Pay_time" field.
     * @type {string}
     * @memberof EntBill
     */
    payTime?: string;
    /**
     * PayTotal holds the value of the "Pay_total" field.
     * @type {number}
     * @memberof EntBill
     */
    payTotal?: number;
    /**
     * 
     * @type {EntBillEdges}
     * @memberof EntBill
     */
    edges?: EntBillEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntBill
     */
    id?: number;
}

export function EntBillFromJSON(json: any): EntBill {
    return EntBillFromJSONTyped(json, false);
}

export function EntBillFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntBill {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'payNumber': !exists(json, 'Pay_number') ? undefined : json['Pay_number'],
        'payTime': !exists(json, 'Pay_time') ? undefined : json['Pay_time'],
        'payTotal': !exists(json, 'Pay_total') ? undefined : json['Pay_total'],
        'edges': !exists(json, 'edges') ? undefined : EntBillEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntBillToJSON(value?: EntBill | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Pay_number': value.payNumber,
        'Pay_time': value.payTime,
        'Pay_total': value.payTotal,
        'edges': EntBillEdgesToJSON(value.edges),
        'id': value.id,
    };
}


