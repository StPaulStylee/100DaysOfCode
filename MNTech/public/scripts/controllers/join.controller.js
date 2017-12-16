angular.module('mnTechSite')
       .controller('JoinController', JoinController);

function JoinController(EmailService, StyleService) {
  console.log('JoinController Loaded');

  var ctrl = this;

  ctrl.isFocused;

  ctrl.sendMemberData = function(data) {
    data.subject = "Membership Request"
    console.log(ServicesController.name);
    EmailService.sendMemberData(data).then(function(response){
      console.log("Response from submit: ", repsonse);
    });
  }

  ctrl.setFocus = function(bool) {
    ctrl.isFocused = bool;
  }

  // ctrl.isActive();
}
