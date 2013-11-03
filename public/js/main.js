
var bitoolaPolls = angular.module('bitoolaPolls', ['ngRoute', 'ngAnimate']);

bitoolaPolls.config(['$interpolateProvider', '$routeProvider', '$locationProvider', function ($interpolateProvider, $routeProvider, $locationProvider) {
	$interpolateProvider.startSymbol('[[');
	$interpolateProvider.endSymbol(']]');
	$locationProvider.html5Mode(false)
	                 .hashPrefix('!');
	$routeProvider.when('/home', {templateUrl: 'home/index', controller: 'HomeCtrl'})
	              .when('/signup', {templateUrl: 'account/signupform', controller: 'AccountCtrl'})
				  .when('/signin', {templateUrl: 'account/signinform', controller: 'AccountCtrl'})
				  .when('/signout', {templateUrl: 'account/signout', controller: 'AccountCtrl'})
				  .otherwise({ redirectTo: '/home'});
}]);

bitoolaPolls.value('seoService', {
	pageTitle: '',
	pageMetaDescription: ''
});

bitoolaPolls.service('authService', function() {
	this.user = {};
	this.isUserAuthenticated = function() {
		return this.user && this.user.email;
	}
    
});

bitoolaPolls.directive('pageTitle', ['seoService', function(seoService) {
	return {
		restrict: 'AE',
		link: function(scope, elem, attrs, ctrl) {
			var pageTitle = elem.text();
			elem.remove();
			seoService.pageTitle = pageTitle;
		}
	};
}]);

bitoolaPolls.directive('pageMetaDescription', ['seoService', function(seoService) {
	return {
		restrict: 'AE',
		link: function(scope, elem, attrs, ctrl) {
			var pageMetaDescription = elem.text();
			elem.remove();
			seoService.pageMetaDescription = pageMetaDescription;
		}
	};
}]);

bitoolaPolls.controller('MainCtrl', ['$scope', '$http', 'seoService', 'authService', function($scope, $http, seoService, authService) {
	$scope.$watch(function(){ return seoService.pageTitle }, function(newTitle) {
		$scope.pageTitle = newTitle;
	});
	$scope.$watch(function(){ return seoService.pageMetaDescription }, function(newDescription) {
		$scope.pageMetaDescription = newDescription;
	});
	$scope.$watch(function(){ return authService.user }, function(user) {
		$scope.user = user;
	}, true);

	$scope.isUserAuthenticated = function() {
		return this.user && this.user.email;
	}

	$scope.setUser = function(user) {
		authService.user = user;
	}
}]);

bitoolaPolls.controller('HomeCtrl', ['$scope', '$http', 'seoService', 'authService', function($scope, $http, seoService, authService) {

}]);

bitoolaPolls.controller('AccountCtrl', ['$scope', '$http', function($scope, $http) {
	
}]);

bitoolaPolls.run(['$location', '$rootScope', function($location, $rootScope) {
    $rootScope.$on('$routeChangeSuccess', function (event, current, previous) {
        // $rootScope.title = current.title;
    });
}]);
