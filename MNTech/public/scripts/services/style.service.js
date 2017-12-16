angular.module('mnTechSite')
       .service('StyleService', StyleService);

function StyleService() {

  const service = this;

  service.isActive = {};

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
    switch (route) {
      case 'home': service.isActive.home = true;
                   service.isActive.about = false;
                   service.isActive.services = false;
                   service.isActive.join = false;
        break;
      case 'about': service.isActive.home = false;
                   service.isActive.about = true;
                   service.isActive.services = false;
                   service.isActive.join = false;
        break;
      case 'services': service.isActive.home = false;
                   service.isActive.about = false;
                   service.isActive.services = true;
                   service.isActive.join = false;
        break;
      case 'join': service.isActive.home = false;
                   service.isActive.about = false;
                   service.isActive.services = false;
                   service.isActive.join = true;
        break;
      default: service.isActive.home = true;
    }
    console.log(service.isActive);
    return service.isActive;
  }

}// End of Service
