

var recipeApp = angular.module('recipeApp', []);

recipeApp.controller('RecipeCtrl', function ($scope, $http){
    $http.get('/recipe/dummy_id').success(
	function(data) { $scope.recipe = data;}
//	function(error){ console.log("something went wrong");}
    );
});

recipeApp.controller('SearchCtrl', function ($scope, $http){
    $http.post('/search', {}).success(
	function(data) {}
//	function(error){}
    );
    
});



