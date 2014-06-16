/// <reference path="../../../typings/angularjs/angular.d.ts" />
function collectionController($scope) {
    $scope.viewClass = "cl-mcont";

    $scope.collections = [
        { Id: 1, Name: "Man", Description: "well....", IsSelected: false },
        { Id: 2, Name: "Women", Description: "blah...blah...", IsSelected: false },
        { Id: 3, Name: "Kids", Description: "for your kids", IsSelected: false }
    ];

    $scope.remove = function (e) {
        var newCollections = $scope.collections.slice(0);

        for (var i = 0; i < newCollections.length; i++) {
            var newCollection = newCollections[i];

            if (newCollection.IsSelected == true) {
                var newTemp = [];

                for (var j = 0; j < $scope.collections.length; j++) {
                    var scopeCollection = $scope.collections[j];

                    if (scopeCollection.Id != newCollection.Id) {
                        newTemp.push(scopeCollection);
                    }
                }

                $scope.collections = newTemp;
            }
        }
    };

    $scope.selectedAll = function (e) {
        var $chkAll = angular.element($(e.target));
        var isChecked = $chkAll.is(':checked');

        for (var i = 0; i < $scope.collections.length; i++) {
            if (isChecked)
                $scope.collections[i].IsSelected = true;
            else
                $scope.collections[i].IsSelected = false;
        }
    };
}

function addCollectionController($scope) {
    $scope.collection = {};

    $scope.create = function () {
        console.log($scope.collection);
    };

    console.log("addCollectionController");
}

function displayCollectionController($scope, $routeParams) {
    $scope.Id = $routeParams.collectionId;
}
//# sourceMappingURL=collectionController.js.map
