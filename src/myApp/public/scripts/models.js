var ManageInventoryMethod;
(function (ManageInventoryMethod) {
    ManageInventoryMethod[ManageInventoryMethod["NoTrack"] = 1] = "NoTrack";
    ManageInventoryMethod[ManageInventoryMethod["Tracking"] = 2] = "Tracking";
})(ManageInventoryMethod || (ManageInventoryMethod = {}));

var WeightUnit;
(function (WeightUnit) {
    WeightUnit[WeightUnit["KG"] = 1] = "KG";
    WeightUnit[WeightUnit["LBL"] = 2] = "LBL";
})(WeightUnit || (WeightUnit = {}));

var UIState;
(function (UIState) {
    UIState[UIState["Normal"] = 1] = "Normal";
    UIState[UIState["Editing"] = 2] = "Editing";
})(UIState || (UIState = {}));

var Product = (function () {
    function Product() {
        this.ManageInventoryMethod = 1 /* NoTrack */;
        this.WeightUnit = 1 /* KG */;
    }
    Object.defineProperty(Product.prototype, "Id", {
        get: function () {
            return this._id;
        },
        set: function (value) {
            this._id = value;
        },
        enumerable: true,
        configurable: true
    });


    Object.defineProperty(Product.prototype, "Name", {
        get: function () {
            return this._name;
        },
        set: function (value) {
            this._name = value;
        },
        enumerable: true,
        configurable: true
    });


    Object.defineProperty(Product.prototype, "Content", {
        get: function () {
            return this._content;
        },
        set: function (value) {
            this._content = value;
        },
        enumerable: true,
        configurable: true
    });


    Object.defineProperty(Product.prototype, "Tags", {
        get: function () {
            return this._tags;
        },
        set: function (value) {
            this._tags = value;
        },
        enumerable: true,
        configurable: true
    });


    Object.defineProperty(Product.prototype, "Sku", {
        get: function () {
            return this._sku;
        },
        set: function (value) {
            this._sku = value;
        },
        enumerable: true,
        configurable: true
    });


    Object.defineProperty(Product.prototype, "Vendor", {
        get: function () {
            return this._vendor;
        },
        set: function (value) {
            this._vendor = value;
        },
        enumerable: true,
        configurable: true
    });


    Object.defineProperty(Product.prototype, "Price", {
        get: function () {
            return this._price;
        },
        set: function (value) {
            this._price = value;
        },
        enumerable: true,
        configurable: true
    });


    Object.defineProperty(Product.prototype, "RegularPrice", {
        get: function () {
            return this._regularPrice;
        },
        set: function (value) {
            this._regularPrice = value;
        },
        enumerable: true,
        configurable: true
    });


    Object.defineProperty(Product.prototype, "ManageInventoryMethod", {
        get: function () {
            return this._manageInvertoryMethod;
        },
        set: function (value) {
            this._manageInvertoryMethod = value;
        },
        enumerable: true,
        configurable: true
    });


    Object.defineProperty(Product.prototype, "InventoryQuantity", {
        get: function () {
            return this._inventoryQuantity;
        },
        set: function (value) {
            this._inventoryQuantity = value;
        },
        enumerable: true,
        configurable: true
    });


    Object.defineProperty(Product.prototype, "LowLevelQuantity", {
        get: function () {
            return this._lowLevelQuantity;
        },
        set: function (value) {
            this._lowLevelQuantity = value;
        },
        enumerable: true,
        configurable: true
    });


    Object.defineProperty(Product.prototype, "IsShippingAddressRequired", {
        get: function () {
            return this._isShippingAddressRequired;
        },
        set: function (value) {
            this._isShippingAddressRequired = value;
        },
        enumerable: true,
        configurable: true
    });


    Object.defineProperty(Product.prototype, "Weight", {
        get: function () {
            return this._weight;
        },
        set: function (value) {
            this._weight = value;
        },
        enumerable: true,
        configurable: true
    });


    Object.defineProperty(Product.prototype, "WeightUnit", {
        get: function () {
            return this._weightUnit;
        },
        set: function (value) {
            this._weightUnit = value;
        },
        enumerable: true,
        configurable: true
    });


    Object.defineProperty(Product.prototype, "IsVisible", {
        get: function () {
            return this._isVisible;
        },
        set: function (value) {
            this._isVisible = value;
        },
        enumerable: true,
        configurable: true
    });


    Object.defineProperty(Product.prototype, "IsPurchasable", {
        get: function () {
            return this._isPurchasable;
        },
        set: function (value) {
            this._isPurchasable = value;
        },
        enumerable: true,
        configurable: true
    });


    Object.defineProperty(Product.prototype, "IsBackOrder", {
        get: function () {
            return this._isBackOrder;
        },
        set: function (value) {
            this._isBackOrder = value;
        },
        enumerable: true,
        configurable: true
    });


    Object.defineProperty(Product.prototype, "IsPreOrder", {
        get: function () {
            return this._isPreOrder;
        },
        set: function (value) {
            this._isPreOrder = value;
        },
        enumerable: true,
        configurable: true
    });


    Object.defineProperty(Product.prototype, "OptionSetId", {
        get: function () {
            return this._optionSetId;
        },
        set: function (value) {
            this._optionSetId = value;
        },
        enumerable: true,
        configurable: true
    });


    Object.defineProperty(Product.prototype, "ResourceId", {
        get: function () {
            return this._resourceId;
        },
        set: function (value) {
            this._resourceId = value;
        },
        enumerable: true,
        configurable: true
    });


    Object.defineProperty(Product.prototype, "PageTitle", {
        get: function () {
            return this._pageTitle;
        },
        set: function (value) {
            this._pageTitle = value;
        },
        enumerable: true,
        configurable: true
    });


    Object.defineProperty(Product.prototype, "MetaDescription", {
        get: function () {
            return this._metaDescription;
        },
        set: function (value) {
            this._metaDescription = value;
        },
        enumerable: true,
        configurable: true
    });


    Object.defineProperty(Product.prototype, "CustomFields", {
        get: function () {
            return this._customFields;
        },
        set: function (value) {
            this._customFields = value;
        },
        enumerable: true,
        configurable: true
    });


    Object.defineProperty(Product.prototype, "Variations", {
        get: function () {
            return this._variations;
        },
        set: function (value) {
            this._variations = value;
        },
        enumerable: true,
        configurable: true
    });

    return Product;
})();

var CustomField = (function () {
    function CustomField() {
        this._uistate = 1 /* Normal */;
    }
    Object.defineProperty(CustomField.prototype, "Id", {
        get: function () {
            return this._id;
        },
        set: function (value) {
            this._id = value;
        },
        enumerable: true,
        configurable: true
    });


    Object.defineProperty(CustomField.prototype, "Name", {
        get: function () {
            return this._name;
        },
        set: function (value) {
            this._name = value;
        },
        enumerable: true,
        configurable: true
    });


    Object.defineProperty(CustomField.prototype, "Value", {
        get: function () {
            return this._value;
        },
        set: function (value) {
            this._value = value;
        },
        enumerable: true,
        configurable: true
    });


    Object.defineProperty(CustomField.prototype, "UIState", {
        get: function () {
            return this._uistate;
        },
        set: function (value) {
            this._uistate = value;
        },
        enumerable: true,
        configurable: true
    });

    return CustomField;
})();

var Variation = (function () {
    function Variation() {
    }
    Object.defineProperty(Variation.prototype, "Id", {
        get: function () {
            return this._id;
        },
        set: function (value) {
            this._id = value;
        },
        enumerable: true,
        configurable: true
    });


    Object.defineProperty(Variation.prototype, "Sku", {
        get: function () {
            return this._sku;
        },
        set: function (value) {
            this._sku = value;
        },
        enumerable: true,
        configurable: true
    });


    Object.defineProperty(Variation.prototype, "Price", {
        get: function () {
            return this._price;
        },
        set: function (value) {
            this._price = value;
        },
        enumerable: true,
        configurable: true
    });


    Object.defineProperty(Variation.prototype, "ManageInventoryMethod", {
        get: function () {
            return this._manageInvertoryMethod;
        },
        set: function (value) {
            this._manageInvertoryMethod = value;
        },
        enumerable: true,
        configurable: true
    });


    Object.defineProperty(Variation.prototype, "InventoryQuantity", {
        get: function () {
            return this._inventoryQuantity;
        },
        set: function (value) {
            this._inventoryQuantity = value;
        },
        enumerable: true,
        configurable: true
    });


    Object.defineProperty(Variation.prototype, "LowLevelQuantity", {
        get: function () {
            return this._lowLevelQuantity;
        },
        set: function (value) {
            this._lowLevelQuantity = value;
        },
        enumerable: true,
        configurable: true
    });

    return Variation;
})();

var OptionSet = (function () {
    function OptionSet() {
    }
    Object.defineProperty(OptionSet.prototype, "Id", {
        get: function () {
            return this._id;
        },
        set: function (value) {
            this._id = value;
        },
        enumerable: true,
        configurable: true
    });


    Object.defineProperty(OptionSet.prototype, "Name", {
        get: function () {
            return this._name;
        },
        set: function (value) {
            this._name = value;
        },
        enumerable: true,
        configurable: true
    });

    return OptionSet;
})();
//# sourceMappingURL=models.js.map
