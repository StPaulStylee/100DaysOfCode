angular.module('mnTechSite')
       .controller('HomeController', HomeController);

function HomeController(StyleService) {

  console.log('HomeController Loaded');

  var ctrl = this;

  // ctrl.setActive = function(route) {
  //   StyleService.setActive(route);
  // }


}
