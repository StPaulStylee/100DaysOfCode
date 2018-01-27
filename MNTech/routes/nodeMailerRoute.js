const router = require('express').Router();
const nodemailer = require('nodemailer');

let transporter = nodemailer.createTransport({
  host: 'mntech.co', // Hope this is right!
  service: 'gmail',
  auth: {
      user: process.env.EMAIL_USERNAME,
      pass: process.env.EMAIL_PASSWORD
  },
});

// Trying to get this to return the response value.. I think I am
// running into some sort of hoisting issue.
let sendMail = function(mailOptions) {
  transporter.sendMail(mailOptions, (error, info) => {
    if (error) {
        return console.log("From sendMail: ", error, "Object: ", mailOptions);
    } else {
      console.log('Message %s sent: %s', info.messageId, info.response);
      return info.response;
    }
  });
}

router.post('/quoteMailer', function(req, res) {
  // setup email data with unicode symbols
  let mailOptions = {
      from: req.body.from, // sender address
      to: req.body.to, // list of receivers
      subject: req.body.subject, // Subject line
      name: req.body.name,
      sender: req.body.sender,
      phone: req.body.phone,
      text: req.body.description, // plain text body
      html: '<h4>Name:</h4><p> ' + req.body.name + '</p>' +
            '<h4>Phone:</h4><p> ' + req.body.phone + '</p>' +
            '<h4>Organization:</h4><p> ' + req.body.organization + '</p>' +
            '<h4>Email:</h4><p> ' + req.body.email + '</p>' +
            '<h4>Project Description:</h4><p> ' + req.body.description + '</p>'
  };
  //send mail with defined transport object
  sendMail(mailOptions);
});

router.post('/memberMailer', function(req, res) {

  // setup email data with unicode symbols
  let mailOptions = {
      from: req.body.from, // sender address
      to: req.body.to, // list of receivers
      subject: req.body.subject, // Subject line
      name: req.body.name,
      sender: req.body.sender,
      phone: req.body.phone,
      text: req.body.description, // plain text body
      html: '<h4>Name:</h4><p> ' + req.body.name + '</p>' +
            '<h4>Address:</h4><p> ' + req.body.address + '</p>' +
            '<h4>Phone:</h4><p> ' + req.body.phone + '</p>' +
            '<h4>Email:</h4><p> ' + req.body.email + '</p>' +
            '<h4>Work Sample:</h4><p> ' + req.body.sample + '</p>' +
            '<h4>Info:</h4><p> ' + req.body.info + '</p>'
  };
  // send mail with defined transport object
  sendMail(mailOptions);
});

router.post('/surveyMailer', function(req, res) {
  console.log('What is the req.body:',req.body);
  // setup email data with unicode symbols
  let mailOptions = {
      to: req.body.to, // list of receivers
      subject: req.body.subject, // Subject line
      html:
      '<h4>The Main Goals:</h4><p> ' + req.body.goals+ '</p>' +
      '<h4>Websites We Like:</h4><p> ' + req.body.websitesWeLike+ '</p>' +
      '<h4>List Of Whats Important First To Last:</h4><p> ' + req.body.listOfWhatsImportantFirstToLast+ '</p>' +
      '<h4>The pages </h4><p> ' + req.body.Pages + '</p>' +
      '<h4>HOME:</h4><p> ' + req.body.Home + '</p>' +
      '<h4>About BNA:</h4><p> ' + req.body.aboutBNA + '</p>' +
      '<h4>Bancroft NeighborHood History:</h4><p> ' + req.body.NeighborHoodHistory + '</p>' +
      '<h4>Mission:</h4><p> ' + req.body.mission + '</p>' +
      '<h4>Meeting Minutes:</h4><p> ' + req.body.meetingMinutes + '</p>' +
      '<h4>Meet Our Board:</h4><p> ' + req.body.meetOurBoard + '</p>' +
      '<h4>Take a Tour:</h4><p> ' + req.body.takeATour + '</p>' +
      '<h4>Volunteer:</h4><p> ' + req.body.volunteer + '</p>' +
      '<h4>Block Clubs:</h4><p> ' + req.body.blockClubs + '</p>' +
      '<h4>BNA Committees:</h4><p> ' + req.body.bnaCommittees + '</p>' +
      '<h4>In the Neighborhood:</h4><p> ' + req.body.inTheNeighborhood + '</p>' +
      '<h4>Resources and Forms:</h4><p> ' + req.body.resourcesAndForms + '</p>' +
      '<h4>Comercial Improvements:</h4><p> ' + req.body.commercialImprovement + '</p>' +
      '<h4>The new features </h4><p> ' + req.body.NewFeatures + '</p>' +
      '<h4>Live Chat:</h4><p> ' + req.body.LiveChat + '</p>' +
      '<h4>Appointment Booking:</h4><p> ' + req.body.appointmentBooking + '</p>' +
      '<h4>Event Mangement System:</h4><p> ' + req.body.eventManagementSystem + '</p>' +
      '<h4>Expanded Business Directory:</h4><p> ' + req.body.expandedBusinessDirectory + '</p>' +
      '<h4>Forms for Programs:</h4><p> ' + req.body.formsForPrograms + '</p>' +
      '<h4>An E-store:</h4><p> ' + req.body.store + '</p>' +
      '<h4>Survey and Polling:</h4><p> ' + req.body.surveysPolling + '</p>' +
      '<h4>Other Needs and Wants:</h4><p> ' + req.body.otherNeeds + '</p>'
  };
  //send mail with defined transport object
  sendMail(mailOptions);
});



module.exports = router;
