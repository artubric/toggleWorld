export interface FeatureToggle {
    id:string;
    isActive:boolean;
    displayName?:string;
    technicalName:string;
    expiresOn?:Date;
    description?:string;
    isInverted:boolean;
    customerIds: string[];
}

export interface FTdataEdit {
    data: FeatureToggle,
    submitFn(ft:FeatureToggle):void
}
