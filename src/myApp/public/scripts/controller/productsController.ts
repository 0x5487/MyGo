/// <reference path="../../../typings/angularjs/angular.d.ts" />

function products($scope){

    $scope.viewClass = "cl-mcont";

}

function productAdd($scope){

    $scope.product = {};

    $scope.create = function(){
        console.log($scope.product);
    }

    console.log("productAdd");



}