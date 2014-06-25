/// <reference path="../../../typings/angularjs/angular.d.ts" />
/// <reference path="../../../typings/jquery/jquery.d.ts" />
/// <reference path="../../../typings/jquery.fileupload/jquery.fileupload.d.ts" />
/// <reference path="../../../typings/bootstrap/bootstrap.d.ts" />
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
