angular.module('mnTechSite')
       .config(function($routeProvider, $locationProvider) {
         $routeProvider.when('/landing', {
           templateUrl: 'views/landing.html',
           controller: 'BaseController as base'
         }).when('/home', {
           templateUrl: 'views/home.html',
           controller: 'HomeController as home'
         }).when('/about', {
           templateUrl: 'views/about.html',
           controller: 'AboutController as about'
         }).when('/services', {
           templateUrl: 'views/services.html',
           controller: 'ServicesController as services'
         }).when('/join', {
           templateUrl: 'views/join.html',
           controller: 'JoinController as join'
         }).when('/blog', {
           templateUrl: 'views/blog.html',
           controller: 'BlogController as blog'
         }).when('/bylaws', {
           templateUrl: 'views/bylaws.html',
           controller: 'BylawsController as laws'
         }).when('/portfolio', {
           templateUrl: 'views/portfolio.html',
           controller: 'PortfolioController as portfolio'
         }).when('/survey', {
           templateUrl: 'views/survey.html',
           controller: 'SurveyController as survey'
         }).otherwise({
           redirectTo: '/landing'
         });
         $locationProvider.html5Mode(true);
       });
