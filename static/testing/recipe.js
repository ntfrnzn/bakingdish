

var recipeApp = angular.module('recipeApp', []);

recipeApp.controller('RecipeCtrl', function ($scope, $http){

    id="dummy_id";
    queryAll = {};
    $scope.display_id=null;
    
    $scope.getSelected = function(recipe_item) {
        console.log(recipe_item.id)
	$http.get('/recipe/'+recipe_item.id).success(
	    function(data) { $scope.recipe = data;}
	    //	function(error){ console.log("something went wrong");}
	);
    };

    $scope.searchAll = function (){
	$http.post('/search', queryAll).success(
	    function(data) {
		$scope.search_results = data;
	    }
	);
    };
});
    



