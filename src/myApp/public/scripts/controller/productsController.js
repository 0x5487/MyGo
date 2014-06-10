function products($scope) {
    $scope.viewClass = "cl-mcont";
}

function productAdd($scope) {
    $scope.viewClass = "cl-mcont";

    $scope.invertoryMethodChange = function () {
        console.log("aa");
    };

    $scope.product = {};

    $scope.create = function () {
        console.log($scope.product);
    };
}
//# sourceMappingURL=productsController.js.map
