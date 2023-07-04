// Service Worker registration
if ("serviceWorker" in navigator) {
    window.addEventListener("load", () => {
        navigator.serviceWorker
            .register("/assets/sw.js")
            .then((registration) => {
                console.log("Service Worker registered:", registration);
            })
            .catch((error) => {
                console.log("Service Worker registration failed:", error);
            });
    });
}

// Service Worker implementation (service-worker.js file)
self.addEventListener("install", (event) => {
    event.waitUntil(
        caches.open("my-pwa-cache").then((cache) => {
            return cache.addAll([
                "/",
                "/auth"
            ]);
        })
    );
});

self.addEventListener("fetch", (event) => {
    event.respondWith(
        caches.match(event.request).then((response) => {
            return response || fetch(event.request);
        })
    );
});
