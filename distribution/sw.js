self.addEventListener('install', (event) => {
    self.skipWaiting();
    console.log('install event');
    console.log(event);
});
self.addEventListener('activate', (event) => {
  event.waitUntil(clients.claim());
    console.log('activate event');
    console.log(event);
});
