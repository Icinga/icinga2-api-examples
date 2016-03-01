# General

[Clippy.js](https://www.smore.com/clippy-js) is a full
Javascript implementation of Microsoft Agent (AKA
Clippy and friends), ready to be embedded in any website.

`query.php` is a simple wrapper for the Icinga 2 API.
I had to come up with that proxy in the middle to
workaround CORS issues hindering direct javascript calls.

`index.php` contains the jQuery ajax requests to poll
the API status in a loop and then selectively let clippy
tell you about problems and run the animation.

**It certainly is not an example for production usage,
this demo is just for fun.**

If - at some later point - the Icinga 2 API is fully
integrated into Icinga Web 2 as a query resource one
could think of realizing this nifty widget for notifications
and such (https://dev.icinga.org/issues/8084).

# Configuration

Edit `query.php` and add your basic auth credentials and connection
info for the Icinga 2 API.

Put everything somewhere on your webserver's root directory
and call it from your browser.

A short demo is available on [Youtube](https://www.youtube.com/watch?v=e3enywTuAX8).
