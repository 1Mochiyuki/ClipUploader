// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {types} from '../models';

export function ChooseFile():Promise<string>;

export function GetHost():Promise<string>;

export function GetTimeout():Promise<number>;

export function HostList():Promise<{[key: string]: types.Uploader}>;

export function RemovePathFromFile(arg1:string):Promise<string>;

export function SaveHost(arg1:string):Promise<void>;

export function SaveTimeoutDuration(arg1:number):Promise<void>;
