self.addEventListener('install', (event) => {
    console.log('install event');
    console.log(event);
    self.skipWaiting();
});
self.addEventListener('activate', (event) => {
    console.log('activate event');
    console.log(event);
    event.waitUntil(clients.claim());
});
