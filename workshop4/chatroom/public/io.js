var io = function() {
  return new IO();
}

var IO = function() {
  var host = window.location.hostname;
  var protocol = 'ws:';
  if (window.location.port) {
    host += ':' + window.location.port;
  }
  if (window.location.protocol === 'https:') {
    protocol = 'wss:'
  }

  this._sock = new WebSocket(protocol + '//' + host + '/chat');
  this._eventCallbacks = {};
  this.setup();
};

IO.prototype.setup = function() {
  this._sock.onmessage = this.onmessage.bind(this);
};

IO.prototype.on = function(event, callback) {
  this._eventCallbacks[event] = callback;
};

IO.prototype.emit = function(event, data) {
  var msg = {
    'event': event,
    'data': data,
  };
  this._sock.send(JSON.stringify(msg));
}

IO.prototype.onmessage = function(event) {
  console.log('Received:', event.data);

  var msg = JSON.parse(event.data);
  var callback = this._eventCallbacks[msg.event];
  if (callback) {
    callback(msg.data);
  } else {
    console.log('No callbacks registered for', msg.event);
  }
};
