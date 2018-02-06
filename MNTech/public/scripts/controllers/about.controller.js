angular.module('mnTechSite')
       .controller('AboutController', AboutController);

function AboutController(StyleService) {

   var ctrl = this;

   ctrl.setActive = function(route) {
     StyleService.setActive(route);
   }

   ctrl.setActive('about');
}//End of controller
