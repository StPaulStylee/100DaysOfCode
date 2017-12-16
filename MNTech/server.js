const express = require('express');
const bodyParser = require('body-parser');
const app = express();
const path = require('path');
require('dotenv').config();
//const router = express.Router();
// const router = require('router'); Router placed on this line

//routes
const nodeMailerRoute = require('./routes/nodeMailerRoute');

app.use(express.static('public'));
app.use(bodyParser.urlencoded({extended: true}));
app.use(bodyParser.json());// You must have to this parse req.body

app.use('/sendEmail', nodeMailerRoute);

app.get('/*', function(req, res) {
  res.sendFile(path.join(__dirname, 'public/views/index.html'));
});

const port = process.env.PORT || 5000;
const server = app.listen(port, function () {
  console.log('MnTech is listening on port', server.address().port);
});
