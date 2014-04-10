/// <reference path="../../typed/angularjs/angular.d.ts" />


var catalogApp = angular.module('catalogApp', [
    'ngRoute'
]);


catalogApp.config(['$routeProvider',
    function($routeProvider) {
        $routeProvider.
            when('/collections', {
                templateUrl: '/views/collections.html',
                controller: 'collectionController'
            }).
            when('/collection/add', {
                templateUrl: '/views/add_collection.html',
                controller: 'addCollectionController'
            }).
            when('/collection/:collectionId', {
                templateUrl: '/views/display_collection.html',
                controller: 'displayCollectionController'
            }).
            when('/products', {
                templateUrl: '/views/collection_old.html',
                controller: 'productController'
            }).
            when('/pages', {
                templateUrl: '/views/pages.html',
                controller: 'pagesController'
            }).
            when('/themes/:themeId', {
                templateUrl: '/views/theme_detail.html',
                controller: 'themeDetailController'
            }).
            otherwise({
                redirectTo: '/'
            });
    }]);