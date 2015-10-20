

var recipeApp = angular.module('recipeApp', []);
recipeApp.controller('RecipeCtrl', function ($scope, $http){
    $http.get('/recipe/dummy_id').success(function(data) {
          $scope.recipe = data;
        });
});



