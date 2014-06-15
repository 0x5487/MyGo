enum ManageInvertoryMethod {
    NoTrack = 1,
    Tracking = 2
}

enum WeightUnit {
    KG = 1,
    LBL = 2
}

class Product {

    private _id: number;

    public get Id(): number {
        return this._id;
    }

    public set Id(value: number) {
        this._id = value;
    }

    private _name: string;

    public get Name() :string {
        return this._name;
    }

    public set Name(value: string) {
        this._name = value;
    }

    private _content: string;

    public get Content(): string {
        return this._content;
    }

    public set Content(value: string) {
        this._content = value;
    }

    private _tags: string;

    public get Tags(): string {
        return this._tags;
    }

    public set Tags(value: string) {
        this._tags = value;
    }

    private _sku: string;

    public get Sku(): string {
        return this._sku;
    }

    public set Sku(value: string) {
        this._sku = value;
    }

    private _vendor: string;

    public get Vendor(): string {
        return this._vendor;
    }

    public set Vendor(value: string) {
        this._vendor = value;
    }

    private _price: number;

    public get Price(): number {
        return this._price;
    }

    public set Price(value: number) {
        this._price = value;
    }


    private _regularPrice: number;

    public get RegularPrice(): number {
        return this._regularPrice;
    }

    public set RegularPrice(value: number) {
        this._regularPrice = value;
    }

    private _manageInvertoryMethod: ManageInvertoryMethod;

    public get ManageInventoryMethod(): ManageInvertoryMethod {
        return this._manageInvertoryMethod;
    }

    public set ManageInventoryMethod(value: ManageInvertoryMethod) {
        this._manageInvertoryMethod = value;
    }

    private _inventoryQuantity: number;

    public get InventoryQuantity(): number {
        return this._inventoryQuantity;
    }

    public set InventoryQuantity(value: number) {
        this._inventoryQuantity = value;
    }

    private _lowLevelQuantity: number;

    public get LowLevelQuantity(): number {
        return this._lowLevelQuantity;
    }

    public set LowLevelQuantity(value: number) {
        this._lowLevelQuantity = value;
    }

    private _isShippingAddressRequired: boolean;

    public get IsShippingAddressRequired(): boolean {
        return this._isShippingAddressRequired;
    }

    public set IsShippingAddressRequired(value: boolean) {
        this._isShippingAddressRequired = value;
    }


    private _weight: number;

    public get Weight(): number {
        return this._weight;
    }

    public set Weight(value: number) {
        this._weight = value;
    }


    private _weightUnit: WeightUnit;

    public get WeightUnit(): WeightUnit {
        return this._weightUnit;
    }

    public set WeightUnit(value: WeightUnit) {
        this._weightUnit = value;
    }


    private _isVisible: boolean;

    public get IsVisible(): boolean {
        return this._isVisible;
    }

    public set IsVisible(value: boolean) {
        this._isVisible = value;
    }

    private _isPurchasable: boolean;

    public get IsPurchasable(): boolean {
        return this._isPurchasable;
    }

    public set IsPurchasable(value: boolean) {
        this._isPurchasable = value;
    }

    private _isBackOrder: boolean;

    public get IsBackOrder(): boolean {
        return this._isBackOrder;
    }

    public set IsBackOrder(value: boolean) {
        this._isBackOrder = value;
    }

    private _isPreOrder: boolean;

    public get IsPreOrder(): boolean {
        return this._isPreOrder;
    }

    public set IsPreOrder(value: boolean) {
        this._isPreOrder = value;
    }


    private _resourceId: string;

    public get ResourceId(): string {
        return this._resourceId;
    }

    public set ResourceId(value: string) {
        this._resourceId = value;
    }


    private _pageTitle: string;

    public get PageTitle(): string {
        return this._pageTitle;
    }

    public set PageTitle(value: string) {
        this._pageTitle = value;
    }


    private _metaDescription: string;

    public get MetaDescription(): string {
        return this._metaDescription;
    }

    public set MetaDescription(value: string) {
        this._metaDescription = value;
    }

    //navigation
    private _customFields: CustomField[];

    public get CustomFields(): CustomField[] {
        return this._customFields;
    }

    public set CustomFields(value: CustomField[]) {
        this._customFields = value;
    }

    public constructor() {
        this.ManageInventoryMethod = ManageInvertoryMethod.NoTrack;
        this.WeightUnit = WeightUnit.KG;
    }
}

class CustomField{

    private _id: number;

    public get Id(): number {
        return this._id;
    }

    public set Id(value: number) {
        this._id = value;
    }


    private _name: string;

    public get Name() :string {
        return this._name;
    }

    public set Name(value: string) {
        this._name = value;
    }


    private _value: string;

    public get Value() :string {
        return this._value;
    }

    public set Value(value: string) {
        this._value = value;
    }
}
