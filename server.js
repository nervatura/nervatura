#!/bin/env node

var AppContainer = function () {

  var self = this;

  self.terminator = function (sig) {
    if (typeof sig === "string") {
      console.log('%s: Received %s - terminating sample app ...',
        Date(Date.now()), sig);
      process.exit(1);}
    console.log('%s: Node server stopped.', Date(Date.now()));};

  self.setupTerminationHandlers = function () {
    //  Process on exit and signals.
    process.on('exit', function () {
      self.terminator();});
    // Removed 'SIGPIPE' from the list - bugz 852598.
    ['SIGHUP', 'SIGINT', 'SIGQUIT', 'SIGILL', 'SIGTRAP', 'SIGABRT',
      'SIGBUS', 'SIGFPE', 'SIGUSR1', 'SIGSEGV', 'SIGUSR2', 'SIGTERM'
    ].forEach(function (element, index, array) {
        process.on(element, function () {
          self.terminator(element);});});};

  self.initialize = function () {
    self.setupTerminationHandlers();};

  self.setupServer = function () {

    require('./app')(function(err,app){
      if(!err){
        var port = normalizePort(process.env.PORT || '8080');
        app.set('port', port);

        var server = require('http').createServer(app);
        if(app.get('env') === 'development'){
          self.debug = require('debug')('nervatura:server');}
          
        if(app.get('ip')){
          server.listen(port, app.get('ip'), function () {
          console.log('%s: Node server started on %s:%d ...',
            Date(Date.now()), app.get('ip'), port);});}
        else{
          server.listen(port);}
        server.on('error', onError);
        server.on('listening', onListening);

        function normalizePort(val) {
          var port = parseInt(val, 10);
          if (isNaN(port)) {
            // named pipe
            return val;}
          if (port >= 0) {
            // port number
            return port;}
          return false;}

        function onError(error) {
          if (error.syscall !== 'listen') {
            throw error;}

          var bind = typeof port === 'string'
            ? 'Pipe ' + port
            : 'Port ' + port;

          switch (error.code) {
            case 'EACCES':
              console.error(bind + ' requires elevated privileges');
              process.exit(1);
              break;
            case 'EADDRINUSE':
              console.error(bind + ' is already in use');
              process.exit(1);
              break;
            default:
              throw error;}}

        function onListening() {
          var addr = server.address();
          console.log('Server on port : ' + addr.port);
          if(self.debug){
            var bind = typeof addr === 'string'
              ? 'pipe ' + addr
              : 'port ' + addr.port;
            self.debug('Listening on ' + bind);}}}
      else {
        console.error(err);
        process.exit(1);}});};};

var zapp = new AppContainer();
zapp.initialize();
zapp.setupServer();
