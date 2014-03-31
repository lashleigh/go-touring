var Config = {};
(function() {
  this.rendererOptions = {
    draggable: true,
    suppressInfoWindows: true,
    preserveViewport: false,
    markerOptions: {draggable: false}
  };

  this.mapOptions = {
    center: new google.maps.LatLng(46.09277427, -116.2267156),
    zoom: 7,
    mapTypeId: google.maps.MapTypeId.ROADMAP
  }

  // The connection must be globally accessible. otherwise bound events
  // wil try to fire on a closed connection in the case of disconnects
  this.conn;
  
}).apply(Config);
