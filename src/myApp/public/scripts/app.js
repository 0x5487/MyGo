/// <reference path="../../typed/angularjs/angular.d.ts" />
var catalogApp = angular.module('catalogApp', [
    'ngRoute'
]);

catalogApp.config([
    '$routeProvider',
    function ($routeProvider) {
        $routeProvider.when('/collection/add', {
            templateUrl: '/views/add_collection.html',
            controller: 'addCollection'
        }).when('/collection/:collectionId', {
            templateUrl: '/views/display_collection.html',
            controller: 'displayCollectionController'
        }).when('/collections', {
            templateUrl: '/views/collection_list.html',
            controller: 'collectionController'
        }).when('/products/add', {
            templateUrl: '/views/add_product.html',
            controller: 'productAdd'
        }).when('/products', {
            templateUrl: '/views/product_list.html',
            controller: 'products'
        }).when('/pages', {
            templateUrl: '/views/pages.html',
            controller: 'pagesController'
        }).when('/themes/:themeId', {
            templateUrl: '/views/theme_detail.html',
            controller: 'themeDetailController'
        }).otherwise({
            redirectTo: '/'
        });
    }]);
//# sourceMappingURL=app.js.map
