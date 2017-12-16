angular.module('mnTechSite')
       .controller('SurveyController', SurveyController);

function SurveyController(EmailService) {
  console.log('SurveyController Loaded');

  let ctrl = this;

  ctrl.visiterList = [
    'Resident - Looking for programs',
    'Resident - Looking for engagement opportunities',
    'Resident - Looking for events/calendar',
    'Resident - Looking for information (history, minutes, local business, etc.)',
    'Business - Looking for programs',
    'Business - Looking for engagement opportunities',
    'Business - Looking for events/calendar',
    'Business - Looking for information (history, local business, etc.)',
    'Funder - Looking for information about your organization'
  ];

//funtion sends form data to email service
  ctrl.sendSurveyData = function(data) {
    data.subject = "Client Survey"
    console.log(ServicesController.name);
    console.log('Survey is:',data);
    EmailService.sendSurveyData(data).then(function(response){
      console.log("Response from submit: ", repsonse);
    });
  }

}//end of conroller function
