var Place = {};
(function() {
  this.new_from_event = function(evt) {
    var message = {}
    message['action'] = 'create'
    message['place'] = {}
    message['place']['guid'] = Helpers.guid()
    message['place']['lat'] = evt.latLng.lat()
    message['place']['lng'] = evt.latLng.lng()
    return message
  }

  this.destroy = function(place) {
  
  }

  this.update = function(place) {
    var latlng = new google.maps.LatLng(place.lat, place.lng);
    var marker = markers[place.guid]
    marker.setPosition(latlng);
  }

  this.create = function(place) {
    var marker = new google.maps.Marker({
      map: map,
      draggable: true,
      title: place.guid,
    });
    new google.maps.event.addListener(marker, 'dragend', function(evt) {
      console.log(marker, evt) 
      var message = {}
      message['action'] = 'update'
      message['place'] = {}
      message['place']['guid'] = marker.title 
      message['place']['lat'] = evt.latLng.lat()
      message['place']['lng'] = evt.latLng.lng()
      Config.conn.send(JSON.stringify(message));
    })    
    markers[place.guid] = marker;
    marker.day_id = place.guid;
    marker.polyline = new google.maps.Polyline({
      map: map,
      strokeOpacity: 0.5,
      strokeWeight: 3
    });
    if(place.lat && place.lng) {
      var latlng = new google.maps.LatLng(place.lat, place.lng);
      marker.setPosition(latlng);
      if(!place.address) {
        console.log("no address")              
        //reverse_geocode(place, latlng);
      }
    } else {
      console.log("no location");
      //geocode(place);
    }
  }

}).apply(Place);
