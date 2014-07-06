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


    Object.defineProperty(Product.prototype, "Options", {
        get: function () {
            return this._options;
        },
        set: function (value) {
            this._options = value;
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
        this._isEditingMode = false;
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


    Object.defineProperty(CustomField.prototype, "IsEditingMode", {
        get: function () {
            return this._isEditingMode;
        },
        set: function (value) {
            this._isEditingMode = value;
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


    Object.defineProperty(Variation.prototype, "Options", {
        get: function () {
            return this._options;
        },
        set: function (value) {
            this._options = value;
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


    Object.defineProperty(Variation.prototype, "IsSelected", {
        get: function () {
            return this._isSelected;
        },
        set: function (value) {
            this._isSelected = value;
        },
        enumerable: true,
        configurable: true
    });

    return Variation;
})();

var Option = (function () {
    function Option() {
    }
    Object.defineProperty(Option.prototype, "Name", {
        get: function () {
            return this._name;
        },
        set: function (value) {
            this._name = value;
        },
        enumerable: true,
        configurable: true
    });


    Object.defineProperty(Option.prototype, "Values", {
        get: function () {
            return this._values;
        },
        set: function (value) {
            this._values = value;
        },
        enumerable: true,
        configurable: true
    });

    return Option;
})();

var Collection = (function () {
    function Collection() {
    }
    Object.defineProperty(Collection.prototype, "Id", {
        get: function () {
            return this._id;
        },
        set: function (value) {
            this._id = value;
        },
        enumerable: true,
        configurable: true
    });


    Object.defineProperty(Collection.prototype, "Title", {
        get: function () {
            return this._title;
        },
        set: function (value) {
            this._title = value;
        },
        enumerable: true,
        configurable: true
    });

    return Collection;
})();
//# sourceMappingURL=models.js.map
