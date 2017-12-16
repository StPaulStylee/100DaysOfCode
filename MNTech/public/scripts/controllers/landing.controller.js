angular.module('mnTechSite')
       .controller('LandingController', LandingController);

function LandingController(StyleService) {
  console.log('LandingController Loaded');

  const ctrl = this;


  ctrl.setInitial = function(route) {
      StyleService.setInitial(route);
  }

} // End of controller
