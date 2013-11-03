
var bitoolaPolls = angular.module('bitoolaPolls', ['ngRoute', 'ngAnimate']);

bitoolaPolls.config(['$routeProvider', function ($routeProvider) {
	$routeProvider.when('/home', {templateUrl: 'home/index', controller: 'HomeCtrl', title: 'Home'})
	              .when('/signup', {templateUrl: 'account/signupform', controller: 'AccountCtrl', title: 'Sign Up'})
				  .when('/login', {templateUrl: 'account/loginform', controller: 'AccountCtrl'})
				  .when('/logout', {templateUrl: 'account/logoutform', controller: 'AccountCtrl'})
				  .otherwise({ redirectTo: '/home'});
}]);

bitoolaPolls.controller('HomeCtrl', ['$scope', '$http', function($scope, $http) {

}]);

bitoolaPolls.controller('AccountCtrl', ['$scope', '$http', function($scope, $http) {
	
}]);

bitoolaPolls.run(['$location', '$rootScope', function($location, $rootScope) {
    $rootScope.$on('$routeChangeSuccess', function (event, current, previous) {
        $rootScope.title = current.$$route.title;
    });
}]);
