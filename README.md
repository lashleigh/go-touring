go-touring
==========

A rewrite of my fantastic bicycle touring app idea in go!

Also an excuse to figure out what web development in pure go looks like. So far, quite pleasant.

Thanks to Gary Burd for the awesome blog post, this would have taken way longer to get going without the examples from http://gary.burd.info/go-websocket-chat


TODO
-----
* Should probably incorporate an actual data binding library, rolling
everything myself is going to be a pain
* Support for destroy operation
* Attach trip_ids to days so there can be multipl trips
* Navigation and polylines (all of this is written in meteor-touring, but it's pretty damn dense)
 - see https://github.com/lashleigh/meteor-touring/blob/master/client/assets/js/calcRoute.js
