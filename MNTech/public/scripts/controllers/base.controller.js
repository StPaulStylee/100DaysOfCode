angular.module('mnTechSite')
       .controller('BaseController', BaseController);

function BaseController(StyleService) {
  console.log('BaseController Loaded');

  var ctrl = this;

  ctrl.isActive = {};
  ctrl.isOpen = false;

  ctrl.setActive = function(route) {
    ctrl.isOpen = !ctrl.isOpen;
    ctrl.isActive = StyleService.setActive(route);
  }

  ctrl.setOpen = function() {
    ctrl.isOpen = !ctrl.isOpen;
  }

}
