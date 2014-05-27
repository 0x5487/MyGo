/// <reference path="../../../typings/angularjs/angular.d.ts" />

function homeController($scope, $window) {

    $scope.login = function () {

        if ($scope.username == "jason" && $scope.password == "123123") {

            $window.location = "/AngularCatalog/views/main.html";
        }

    }
}