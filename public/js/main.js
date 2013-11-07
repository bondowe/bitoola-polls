
var bitoolaPolls = angular.module('bitoolaPolls', ['ngRoute', 'ngAnimate']);

bitoolaPolls.config(['$interpolateProvider', '$routeProvider', '$locationProvider', function ($interpolateProvider, $routeProvider, $locationProvider) {
	$interpolateProvider.startSymbol('[[');
	$interpolateProvider.endSymbol(']]');
	$locationProvider.html5Mode(false)
	                 .hashPrefix('!');
	$routeProvider.when('/home', {templateUrl: 'home/index', controller: 'HomeCtrl'})
	              .when('/sign-up', {templateUrl: 'account/signupform', controller: 'AccountCtrl'})
				  .when('/sign-in', {templateUrl: 'account/signinform', controller: 'AccountCtrl'})
				  .when('/sign-out', {templateUrl: 'account/signout', controller: 'AccountCtrl'})
				  .when('/reset-password', {templateUrl: 'account/passwordresetform', controller: 'AccountCtrl'})
				  .otherwise({ redirectTo: '/home'});
}]);

bitoolaPolls.value('seoService', {
	pageTitle: '',
	pageMetaDescription: ''
});

bitoolaPolls.service('authService', function() {
	this.user = {};
	this.authenticating = false;
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

bitoolaPolls.directive('popover', function() {
	return {
		restrict: 'A',
		link: function(scope, elem, attrs, ctrl) {
			var p = $('#' + attrs.popover);
			if(p) {
				elem.popover({trigger: 'focus', html:true, title: p.find('.title').html(), content: p.find('.content').html()});
			}
		}
	};
});

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

bitoolaPolls.directive('captcha', [function() {
	return {
		restrict: 'AE',
		replace: false,
		link: function(scope, elem, attrs, ctrl) {
			$.get('/account/captcha')
			 .done(function(data) {
			 	elem.replaceWith(data);
			 	createCaptcha();
			 });
		}
	};
}]);


bitoolaPolls.controller('MainCtrl', ['$scope', '$http', 'seoService', 'authService', function($scope, $http, seoService, authService) {
	$scope.pageTitle = '';
	$scope.pageMetaDescription = '';
	$scope.user = null;
	$scope.authenticating = false;
	$scope.$watch(function(){ return seoService.pageTitle }, function(newTitle) {
		$scope.pageTitle = newTitle;
	});
	$scope.$watch(function(){ return seoService.pageMetaDescription }, function(newDescription) {
		$scope.pageMetaDescription = newDescription;
	});
	$scope.$watch(function(){ return authService.user }, function(user) {
		$scope.user = user;
	}, true);
	$scope.$watch(function(){ return authService.authenticating }, function(authenticating) {
		$scope.authenticating = authenticating;
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

bitoolaPolls.controller('AccountCtrl', ['$scope', '$http', 'authService', function($scope, $http, authService) {
	authService.authenticating = true;
}]);

bitoolaPolls.run(['$rootScope', 'authService', function($rootScope, authService) {
    $rootScope.$on('$routeChangeSuccess', function (event, current, previous) {
        authService.authenticating = (current.controller == 'AccountCtrl');
    });
}]);

