angular.module('mnTechSite')
       .controller('BaseController', BaseController);

function BaseController(StyleService) {
  console.log('BaseController Loaded');

  var ctrl = this;

  //Non-mobile
  ctrl.isActive = {};
  //For mobile
  ctrl.isOpen = false;

  ctrl.setActive = function(route) {
    ctrl.isOpen = !ctrl.isOpen;
    ctrl.isActive = StyleService.setActive(route);
  }

  ctrl.setOpen = function() {
    ctrl.isOpen = !ctrl.isOpen;
  }

  ctrl.getActive = function() {
    ctrl.isActive = StyleService.getActive()
    }

  ctrl.getActive();

}
