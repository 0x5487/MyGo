function products($scope) {
    $scope.viewClass = "cl-mcont";
}

function productAdd($scope) {
    $scope.viewClass = "cl-mcont";
    $scope.selectedInvertoryMethod = 0;

    $scope.invertoryMethodChange = function () {
        var $selInvertoryMethod = $("#selInvertoryMethod");
        var $invertoryMethodPanel = $("#invertoryMethodPanel");

        if ($selInvertoryMethod.val() == "1") {
            $invertoryMethodPanel.show();
        } else {
            $invertoryMethodPanel.hide();
        }

        console.log("aa");
    };

    $scope.product = {};

    $scope.create = function () {
        console.log($scope.product);
    };
}
//# sourceMappingURL=productsController.js.map
