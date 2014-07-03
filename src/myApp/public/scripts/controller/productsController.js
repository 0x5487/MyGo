/// <reference path="../../../typings/angularjs/angular.d.ts" />
/// <reference path="../../../typings/jquery/jquery.d.ts" />
/// <reference path="../../../typings/jquery.fileupload/jquery.fileupload.d.ts" />
/// <reference path="../../../typings/bootstrap/bootstrap.d.ts" />
/// <reference path="../../../typings/underscore/underscore.d.ts" />
/// <reference path="../models.ts" />
function displayProductsController($scope) {
    $scope.viewClass = "cl-mcont";
}

function displayProductController($scope, $routeParams) {
    $scope.viewClass = "cl-mcont";
}

function productAddController($scope) {
    $scope.viewClass = "cl-mcont";

    var product = new Product();

    var optionSet1 = new OptionSet();
    optionSet1.Id = 0;
    optionSet1.Name = "None";

    var optionSet2 = new OptionSet();
    optionSet2.Id = 23;
    optionSet2.Name = "Phone";

    var optionSet3 = new OptionSet();
    optionSet3.Id = 24;
    optionSet3.Name = "NoteBook";

    $scope.optionSet = [optionSet1, optionSet2, optionSet3];
    $scope.selectedOptionSet = optionSet1;
    $scope.optionSetChange = function () {
        product.OptionSetId = $scope.selectedOptionSet.Id;
    };

    $scope.optionNumber = 1;

    var option1 = new Option();
    option1.Name = "Jason1";

    var option2 = new Option();
    option2.Name = "Jason2";

    var option3 = new Option();
    option3.Name = "Jason3";

    product.Options = [option1, option2, option3];
    $scope.optionNumberChange = function () {
    };

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

    var variation1 = new Variation();
    variation1.Sku = "SKU-123";

    var variation_option1 = new Option();
    variation_option1.Name = "Color";
    variation_option1.Values = "Black";

    var variation_option2 = new Option();
    variation_option2.Name = "Size";
    variation_option2.Values = "8G";

    variation1.Options = [variation_option1, variation_option2];
    product.Variations = [variation1];

    $scope.product = product;

    //events
    $scope.save = function () {
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

    $scope.createCustomField = function () {
        var newCustomField = new CustomField();
        newCustomField.UIState = 2 /* Editing */;
        product.CustomFields.push(newCustomField);
    };

    $scope.editCustomField = function (index) {
        product.CustomFields[index].UIState = 2 /* Editing */;
    };

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
}
//# sourceMappingURL=productsController.js.map
