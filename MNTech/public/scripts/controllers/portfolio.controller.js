angular.module('mnTechSite')
       .controller('PortfolioController', PortfolioController);

function PortfolioController() {
  console.log('PortfolioController Loaded');

  var ctrl = this;

  ctrl.setActive = function(route) {
    StyleService.setActive(route)
  }
}
