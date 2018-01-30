angular.module('mnTechSite')
       .service('StyleService', StyleService);

function StyleService() {

  const service = this;

  service.isActive = {};
  service.isExpanded = false;
  service.isLanding = false;

  service.isLandingActive = function(bool) {
    service.isLanding = bool;
  }

  service.setInitial = function(route) {
    // switch (route) {
    //   case 'home': service.isActive.home = true;
    //     break;
    //   case 'join': service.isActive.join = true;
    //     break;
    //   default: service.isActive.home = true;
    service.setActive(route);
  }

  service.setActive = function(route) {
    service.isActiveFalse();
    switch (route) {
      case 'home': service.isActive.home = true;
        break;
      case 'about': service.isActive.about = true;
        break;
      case 'services': service.isActive.services = true;
        break;
      case 'join': service.isActive.join = true;
        break;
      case 'portfolio': service.isActive.portfolio = true;
        break;
      default: service.isActive.home = true;
    }
    console.log(service.isActive);
    return service.isActive;
  }

  service.isActiveFalse = function() {
    service.isActive.home = false;
    service.isActive.about = false;
    service.isActive.services = false;
    service.isActive.join = false;
    service.isActive.portfolio = false;
  }
  //
  // service.expand = function() {
  //   service.isExpanded = !service.isExpanded;
  //   return service.isExpanded;
  // }

}// End of Service
