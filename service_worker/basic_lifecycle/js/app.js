const APP = {
  SW: null,
  
  init() {
    //called after DOMContentLoaded
    if ("serviceWorker" in navigator) {
      // 1. Register a service worker hosted at the root of the
      // site using the default scope.
      navigator.serviceWorker
        .register("/sw.js", {
          scope: "/",
        })
        .then((registration) => {
          APP.SW =
            registration.installing ||
            registration.waiting ||
            registration.active;
          console.log("service worker registered");
        });
      // 2. See if the page is currently has a service worker.
      if (navigator.serviceWorker.controller) {
        console.log("we have a service worker installed");
      }

      // 3. Register a handler to detect when a new or
      // updated service worker is installed & activate.
      navigator.serviceWorker.oncontrollerchange = (ev) => {
        console.log("New service worker activated");
      };

      // 4. remove/unregister service workers
      // navigator.serviceWorker.getRegistrations().then((regs) => {
      //   for (let reg of regs) {
      //     reg.unregister().then((isUnreg) => console.log(isUnreg));
      //   }
      // });
      // 5. Listen for messages from the service worker
    } else {
      console.log("Service workers are not supported.");
    }
  },
};

document.addEventListener("DOMContentLoaded", APP.init);
