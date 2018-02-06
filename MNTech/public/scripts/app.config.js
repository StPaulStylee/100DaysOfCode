angular.module('mnTechSite')
       .config(function($routeProvider, $locationProvider) {
         $routeProvider.when('/home', {
           templateUrl: 'views/home.html',
           controller: 'HomeController as home'
         }).when('/about', {
           templateUrl: 'views/about.html',
           controller: 'AboutController as about'
         }).when('/services', {
           templateUrl: 'views/services.html',
           controller: 'ServicesController as services'
         }).when('/blog', {
           templateUrl: 'views/blog.html',
           controller: 'BlogController as blog'
         }).when('/portfolio', {
           templateUrl: 'views/portfolio.html',
           controller: 'PortfolioController as portfolio'
         }).when('/survey', {
           templateUrl: 'views/survey.html',
           controller: 'SurveyController as survey'
         }).otherwise({
           redirectTo: '/home'
         });
         $locationProvider.html5Mode(true);
       });
