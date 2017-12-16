angular.module('mnTechSite')
       .controller('ServicesController', ServicesController);

function ServicesController(EmailService, StyleService) {
  console.log('ServicesController Loaded');

  const ctrl = this;
 //
 //  ctrl.setActive = function(route) {
 //    StyleService.setActive(route)
 // }
 ctrl.isFocused;

  ctrl.sendQuoteData = function(data) {
    data.subject = "Requesting A Quote"
    console.log(ServicesController.name);
    EmailService.sendQuoteData(data).then(function(response){
      console.log("Response from submit: ", repsonse);
    });
  }

  ctrl.setFocus = function(bool) {
    ctrl.isFocused = bool;
  }
}
