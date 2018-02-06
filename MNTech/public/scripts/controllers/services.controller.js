angular.module('mnTechSite')
       .controller('ServicesController', ServicesController);

function ServicesController(EmailService, StyleService) {

  const ctrl = this;

  ctrl.sendQuoteData = function(data) {
    data.subject = "Requesting A Quote"
    ctrl.quoteForm.$setPristine();
    ctrl.quoteForm.$setUntouched();
    EmailService.sendQuoteData(data).then(function(response){
      console.log("Response from submit: ", response);
    });
  }
  ctrl.setActive = function(route) {
    StyleService.setActive(route);
  }

  ctrl.setActive('services');
}
