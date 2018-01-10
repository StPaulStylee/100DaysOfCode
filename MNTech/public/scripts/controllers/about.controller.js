angular.module('mnTechSite')
       .controller('AboutController', AboutController);

function AboutController(StyleService) {

   var ctrl = this;

   ctrl.isExpanded = {
     adrian: false,
     hannah: false,
     jeff: false,
     allFalse: function() {
       ctrl.isExpanded.adrian = false;
       ctrl.isExpanded.hannah = false;
       ctrl.isExpanded.jeff = false;
     }
   }

   ctrl.setActive = function(route) {
     StyleService.setActive(route)
  }

  ctrl.expand = function(member) {
    console.log(member);
    ctrl.isExpanded.allFalse();
    switch(member) {
      case 'adrian': ctrl.isExpanded.adrian = true;
        break;
      case 'hannah': ctrl.isExpanded.hannah = true;
        break;
      case 'jeff': ctrl.isExpanded.jeff = true;
        break
      default: ctrl.isExpanded.allFalse();
    }
  }
}//End of controller
