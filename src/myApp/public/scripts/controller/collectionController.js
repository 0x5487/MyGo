/// <reference path="../../../typings/angularjs/angular.d.ts" />
/// <reference path="../../../typings/jquery/jquery.d.ts" />
/// <reference path="../../../typings/jquery.fileupload/jquery.fileupload.d.ts" />
/// <reference path="../../../typings/bootstrap/bootstrap.d.ts" />
/// <reference path="../../../typings/underscore/underscore.d.ts" />
/// <reference path="../models.ts" />
function displayCollectionsController($scope, $http) {
    var collections = [];
    for (var i = 0; i < 20; i++) {
        var collection = new Collection();
        collection.Id = i;
        collection.Name = "Men";
        collection.ProductCount = 50;

        if (i > 10) {
            collection.Name = "Kids";
        }
        collections.push(collection);
    }

    //page properties
    $scope.viewClass = "cl-mcont";
    $scope.collections = collections;

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

    //page loaded
    setTimeout(function () {
        $('.dataTables_filter input').addClass('form-control').attr('placeholder', 'Search');
        $('.dataTables_length select').addClass('form-control');
        console.log("controller loaded");
    }, 300);
}

function collectionAddController($scope) {
    $scope.viewClass = "cl-mcont";
}

function displayCollectionController($scope, $routeParams) {
    $scope.viewClass = "cl-mcont";
    $scope.Id = $routeParams.collectionId;
}
//# sourceMappingURL=collectionController.js.map
