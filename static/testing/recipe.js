function Recipe($scope, $http) {
    $http.get('/recipe').
        success(function(data) {
            $scope.recipe = data;
        });
}

