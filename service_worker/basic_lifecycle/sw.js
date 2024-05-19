//our service worker
// console.log('sw running');
//new

//console.log({ self });
self.addEventListener("install", (ev) => {
  //service worker is installed.
  console.log("installed");
});
self.addEventListener("activate", (ev) => {
  //service worker is activated
  console.log("activated");
});

self.addEventListener("fetch", (ev) => {
  //service worker intercepted a fetch call
  console.log("intercepted a http request", ev.request);
});

self.addEventListener("message", (ev) => {
  //message from webpage
});
