/// <reference path="../../typings/angularjs/angular.d.ts" />
var catalogApp = angular.module('catalogApp', ['ngRoute']);

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
            templateUrl: '/views/product_add.html',
            controller: 'productAddController'
        }).when('/products/:productId', {
            templateUrl: '/views/product_list.html',
            controller: 'displayProductController'
        }).when('/products', {
            templateUrl: '/views/product_list.html',
            controller: 'displayProductsController'
        }).when('/options/add', {
            templateUrl: '/views/option_add.html',
            controller: 'addOptionController'
        }).when('/options', {
            templateUrl: '/views/option_list.html',
            controller: 'displayOptionsController'
        }).when('/optionset/add', {
            templateUrl: '/views/optionset_add.html',
            controller: 'addOptionSetController'
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
