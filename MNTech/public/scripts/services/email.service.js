angular.module('mnTechSite')
       .service('EmailService', EmailService);

function EmailService($http) {

    const service = this;
    const route = '/sendEmail';
    const quoteRoute = '/quoteMailer';
    const memberRoute = '/memberMailer';
    const surveyRoute = '/surveyMailer';

    service.sendQuoteData = function(data) {
      console.log("In the service: ", data);
      service.emailData = {
        name: data.name,
        subject: data.subject,
        sender: data.email,
        email: data.email,
        to: 'mntechcoop@gmail.com',
        phone: data.phone,
        organization: data.organization,
        description: data.description,
      };
      return $http.post(route + quoteRoute, service.emailData)
      .then(function(response){
        return repsonse;
      });
    }

    service.sendMemberData = function functionName(data) {
      console.log("In the service: ", data);
      service.emailData = {
        name: data.name,
        address: data.address,
        subject: data.subject,
        sender: data.email,
        email: data.email,
        to: 'mntechcoop@gmail.com',
        phone: data.phone,
        sample: data.sample,
        info: data.info,
        };
      return $http.post(route + memberRoute, service.emailData)
      .then(function(response){
        return repsonse;
      });
     }

      service.sendSurveyData = function(data) {
        console.log("In the service: ", data);
        service.emailData = {
          to: 'mntechcoop@gmail.com',
          subject:data.subject,
          goals:data.goals,
          websitesWeLike:data.website,
          listOfWhatsImportantFirstToLast:data.visiterList,
          Pages:"We would like moved:",
          Home:data.home,
          aboutBNA:data.aboutBNA,
          NeighborHoodHistory:data.history,
          mission:data.mission,
          meetingMinutes:data.meetingMinutes,
          meetOurBoard:data.meetOurBoard,
          takeATour:data.takeATour,
          volunteer:data.volunteer,
          blockClubs:data.blockClubs,
          bnaCommittees:data.bnaCommittees,
          inTheNeighborhood:data.inTheNeighborhood,
          resourcesAndForms:data.resourcesAndForms,
          commercialImprovement:data.commercialImprovement,
          NewFeatures:"We would like add:",
          LiveChat:data.LiveChat,
          appointmentBooking:data.appointmentBooking,
          eventManagementSystem:data.eventManagementSystem,
          expandedBusinessDirectory:data.expandedBusinessDirectory,
          formsForPrograms:data.formsForPrograms,
          store:data.store,
          surveysPolling:data.surveysPolling,
          otherNeeds:data.other,
        };
        return $http.post(route + surveyRoute, service.emailData)
        .then(function(response){
          return repsonse;
        });
      }

  }//end of service
