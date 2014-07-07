/// <reference path="../../../typings/angularjs/angular.d.ts" />
/// <reference path="../../../typings/jquery/jquery.d.ts" />
/// <reference path="../../../typings/jquery.fileupload/jquery.fileupload.d.ts" />
/// <reference path="../../../typings/bootstrap/bootstrap.d.ts" />
/// <reference path="../../../typings/underscore/underscore.d.ts" />
/// <reference path="../models.ts" />
function displayProductsController($scope, $http) {
    $scope.viewClass = "cl-mcont";

    var products = [];

    for (var i = 0; i < 20; i++) {
        var product1 = new Product();
        product1.Id = i;
        product1.Name = "IPhone";
        product1.ManageInventoryMethod = 1 /* NoTrack */;
        product1.Price = 399;
        product1.PriceWithSymbol = "USD 399";
        product1.InventoryQuantity = i;
        product1.UpdateAt = new Date();
        products.push(product1);
    }

    $scope.products = products;

    //page loaded
    setTimeout(function () {
        $('.dataTables_filter input').addClass('form-control').attr('placeholder', 'Search');
        $('.dataTables_length select').addClass('form-control');
        console.log("controller loaded");
    }, 500);
}

function displayProductController($scope, $routeParams) {
    $scope.viewClass = "cl-mcont";
}

function productAddController($scope, $http) {
    var product = new Product();

    var collection1 = new Collection();
    collection1.Id = 1;
    collection1.Title = "Men";

    var collection2 = new Collection();
    collection2.Id = 2;
    collection2.Title = "Women";

    var collection3 = new Collection();
    collection3.Id = 3;
    collection3.Title = "Kids";

    var option1 = new Option();
    option1.Name = "Color";
    option1.Values = "Black, White";

    var option2 = new Option();
    option2.Name = "Size";
    option2.Values = "8GB, 16GB, 32GB";

    var option3 = new Option();
    option3.Name = "LCD Size";
    option3.Values = "4.7, 5.5";

    product.Options = [option1, option2, option3];

    var field1 = new CustomField();
    field1.Id = 1;
    field1.Name = "Jason1";
    field1.Value = "abc1";

    var field2 = new CustomField();
    field2.Id = 2;
    field2.Name = "Jason2";
    field2.Value = "abc2";

    var field3 = new CustomField();
    field3.Id = 3;
    field3.Name = "Jason3";
    field3.Value = "abc3";

    product.CustomFields = [field1, field2, field3];

    //page properties
    $scope.viewClass = "cl-mcont";
    $scope.optionNumber = 1;
    $scope.checkAllVariations = false;
    $scope.product = product;
    $scope.collections = [collection1, collection2, collection3];

    //*page properties
    //page events
    $scope.optionNumberChange = function () {
        product.Variations = null;
    };

    $scope.saveProduct = function () {
        $scope.isSubmitted = true;

        //redirect to error tab
        if ($scope.productDetailsForm.$invalid) {
            $('#detailTab').tab('show');
        }

        console.log($scope.product);
    };

    $scope.generateSKUs = function () {
        var options = [];
        var opt1, opt2, opt3;

        var opt1Name = $.trim(product.Options[0].Name);
        var opt2Name = $.trim(product.Options[1].Name);
        var opt3Name = $.trim(product.Options[2].Name);

        var optionValues1 = $.trim(product.Options[0].Values).split(',');
        _.each(optionValues1, function (element1) {
            opt1 = new Option();
            opt1.Name = opt1Name;
            opt1.Values = $.trim(element1);

            if ($scope.optionNumber == 1) {
                var tmpOption = [opt1];
                options.push(tmpOption);
            }
            ;

            if ($scope.optionNumber >= 2) {
                var optionValues2 = $.trim(product.Options[1].Values).split(',');

                _.each(optionValues2, function (element2) {
                    opt2 = new Option();
                    opt2.Name = opt2Name;
                    opt2.Values = $.trim(element2);
                    if ($scope.optionNumber == 2) {
                        var tmpOption = [opt1, opt2];
                        options.push(tmpOption);
                    }
                    ;

                    if ($scope.optionNumber >= 3) {
                        var optionValues3 = $.trim(product.Options[2].Values).split(',');

                        _.each(optionValues3, function (element3) {
                            opt3 = new Option();
                            opt3.Name = opt3Name;
                            opt3.Values = $.trim(element3);

                            if ($scope.optionNumber == 3) {
                                var tmpOption = [opt1, opt2, opt3];
                                options.push(tmpOption);
                            }
                            ;
                        });
                    }
                });
            }
        });

        if (options.length > 0) {
            if ($scope.product.Variations == null) {
                $scope.product.Variations = [];
            }

            _.each(options, function (element, index) {
                var variation = new Variation();
                variation.Sku = "sku" + index;
                variation.Options = element;
                $scope.product.Variations.push(variation);
            });
        }

        console.log(options.length);
    };

    $scope.selectedAllVariations = function (e) {
        var $chkAll = angular.element($(e.target));
        var isChecked = $chkAll.is(':checked');

        for (var i = 0; i < product.Variations.length; i++) {
            if (isChecked)
                product.Variations[i].IsSelected = true;
            else
                product.Variations[i].IsSelected = false;
        }
    };

    $scope.removeSelectedVariations = function () {
        var newVariations = product.Variations.slice(0);

        for (var i = 0; i < newVariations.length; i++) {
            var newVariation = newVariations[i];

            if (newVariation.IsSelected == true) {
                var newTemp = [];

                for (var j = 0; j < product.Variations.length; j++) {
                    var scopeVariation = product.Variations[j];

                    if (scopeVariation.Sku != newVariation.Sku) {
                        newTemp.push(scopeVariation);
                    }
                }

                product.Variations = newTemp;
            }
        }

        $scope.checkAllVariations = false;
    };

    $scope.createCustomField = function () {
        var newCustomField = new CustomField();
        newCustomField.IsEditingMode = true;
        product.CustomFields.push(newCustomField);
    };

    $scope.saveCustomField = function (index) {
        product.CustomFields[index].IsEditingMode = false;
    };

    $scope.removeCustomField = function (index) {
        product.CustomFields.splice(index, 1);
    };

    $scope.validatePage = function () {
    };

    //*page events
    $scope.fileList = [];

    var uploadButton = $('<button/>').addClass('btn btn-primary').prop('disabled', true).text('Processing...').on('click', function () {
        var $this = $(this), data = $this.data();
        $this.off('click').text('Abort').on('click', function () {
            $this.remove();
            data.abort();
        });
        data.submit().always(function () {
            $this.remove();
        });
    });

    $('#fileupload').on('fileuploadadd', function (e, data) {
        // Add the files to the list
        data.context = $('<div/>').appendTo('#files');
        $.each(data.files, function (index, file) {
            var node = $('<p/>').append($('<span/>').text(file.name));
            if (!index) {
                node.append('<br>');
            }
            node.appendTo(data.context);
        });
    }).on('fileuploadprocessalways', function (e, data) {
        console.log("fileuploadprocessalways fired");
        var index = data.index, file = data.files[index], node = $(data.context.children()[index]);
        console.log(node);

        if (file.preview) {
            node.prepend('<br>').prepend(file.preview);
        }
        if (file.error) {
            node.append('<br>').append($('<span class="text-danger"/>').text(file.error));
        }
        if (index + 1 === data.files.length) {
            data.context.find('button').text('Upload').prop('disabled', !!data.files.error);
        }
    }).on('fileuploaddone', function (e, data) {
        $.each(data.result.files, function (index, file) {
            if (file.url) {
                var link = $('<a>').attr('target', '_blank').prop('href', file.url);
                $(data.context.children()[index]).wrap(link);
            } else if (file.error) {
                var error = $('<span class="text-danger"/>').text(file.error);
                $(data.context.children()[index]).append('<br>').append(error);
            }
        });
    }).on('fileuploadfail', function (e, data) {
        $.each(data.files, function (index, file) {
            var error = $('<span class="text-danger"/>').text('File upload failed.');
            $(data.context.children()[index]).append('<br>').append(error);
        });
    }).prop('disabled', !$.support.fileInput).parent().addClass($.support.fileInput ? undefined : 'disabled');

    //page loaded
    _.each($scope.collections, function (element) {
        $('#selCollections').multiSelect('addOption', { value: element.Id, text: element.Title, selected: true });
    });

    $('#selCollections').multiSelect('select', ['1', '2']);
}
//# sourceMappingURL=productsController.js.map
