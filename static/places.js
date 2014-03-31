Handlebars.registerHelper('round', function(sig, num) {
  if(num == undefined) {return }
  return num.toFixed(sig);
});

var Place = {};
(function() {
  var source = '<li data-guid="{{guid}}">'+
                 '<span>{{address}}</span>'+
                 '<span>{{round 3 lat}}</span>'+
                 '<span>{{round 3 lng}}</span>'+
               '</li>';

  var template = Handlebars.compile(source);

  this.new_from_event = function(evt) {
    var message = {}
    message['action'] = 'create'
    message['place'] = {}
    message['place']['guid'] = Helpers.guid()
    message['place']['lat'] = evt.latLng.lat()
    message['place']['lng'] = evt.latLng.lng()
    return message
  }

  this.list_element = function(place) {
    return template(place)
  }

  this.destroy = function(place) {

  }

  this.update = function(place) {
    var latlng = new google.maps.LatLng(place.lat, place.lng);
    var marker = markers[place.guid]
    marker.setPosition(latlng);

    $("li[data-guid="+place.guid+"]").html(Place.list_element(place))
  }

  this.update_address = function(place, address) {
      var message = {}
      message['action'] = 'update'
      message['place'] = place
      message['place']['address'] = address
      Config.conn.send(JSON.stringify(message));
  }
  
  // TODO rename this. This method is more about putting the
  // existing place on the map and setting up even handlers,
  // no about creating places.
  this.create = function(place) {
    places[place.guid] = place
    $('ul.nav').append(Place.list_element(place))

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
      Helpers.reverse_geocode(message['place'], evt.latLng)
      Config.conn.send(JSON.stringify(message));
    });
    markers[place.guid] = marker;
    marker.day_id = place.guid;
    marker.polyline = new google.maps.Polyline({
      map: map,
      strokeOpacity: 0.5,
      strokeWeight: 3
    });
    if(place.lat && place.lng) {
      var latLng = new google.maps.LatLng(place.lat, place.lng);
      marker.setPosition(latLng);
      if(!place.address) {
        console.log("no address");
        //reverse_geocode(place, latlng);
      }
      Helpers.reverse_geocode(place, latLng)
    } else {
      console.log("no location");
      //geocode(place);
    }
  }

}).apply(Place);
