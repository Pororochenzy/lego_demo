'use strict';
const MANIFEST = 'flutter-app-manifest';
const TEMP = 'flutter-temp-cache';
const CACHE_NAME = 'flutter-app-cache';
const RESOURCES = {
  "version.json": "65599b2146d3271b98f39fbeb65984b5",
"index.html": "c95f6c3b18d1e7fa54887ee9aefd451e",
"/": "c95f6c3b18d1e7fa54887ee9aefd451e",
"main.dart.js": "0fc6491557a415913045a75edb0cac44",
"favicon.png": "5dcef449791fa27946b3d35ad8803796",
"icons/Icon-192.png": "ac9a721a12bbc803b44f645561ecb1e1",
"icons/Icon-maskable-192.png": "c457ef57daa1d16f64b27b786ec2ea3c",
"icons/Icon-maskable-512.png": "301a7604d45b3e739efc881eb04896ea",
"icons/Icon-512.png": "96e752610906ba2a93c65f8abe1645f1",
"manifest.json": "47a68c1617e6c6ba25f4e156f6d80de9",
"assets/AssetManifest.json": "cf94df5fbfcb15313cd74f2578315438",
"assets/NOTICES": "5e00ce96e96cd4e98cc551810099a97b",
"assets/FontManifest.json": "dc3d03800ccca4601324923c0b1d6d57",
"assets/packages/cupertino_icons/assets/CupertinoIcons.ttf": "6d342eb68f170c97609e9da345464e5e",
"assets/fonts/MaterialIcons-Regular.otf": "4e6447691c9509f7acdbf8a931a85ca1",
"assets/assets/images/undraw_clean_up_ucm0.svg": "e80ae0b787b1a1a0f30586b5d214a875",
"assets/assets/images/%25E4%25BB%25BB%25E5%258A%25A1.svg": "2689c2b1dbf9f842895761c8c8d40774",
"assets/assets/images/userImage.png": "afc26cb27789a846c7047e57cb9c6422",
"assets/assets/images/undraw_interview_re_e5jn.svg": "b3362ab3f55a5d54284fec4f6ccd7e71",
"assets/assets/images/%25E5%2586%2585%25E5%25AD%2598.svg": "2a906cd58d85e87b3024c8d416274b62",
"assets/assets/images/lego.png": "4f1433882808cee5b481f301758cfde3",
"assets/assets/images/splash_01.png": "307c773d181ceeb899559add51c7acb4",
"assets/assets/images/splash_03.png": "b4529fcfd305f5843e17d08bf361f1bf",
"assets/assets/images/splash_02.png": "162d339138dd70b904d8751ac13fa739",
"assets/assets/images/CPU.svg": "db8a68abad90e6b58cd83d176df40311",
"assets/assets/images/splash_05.png": "ca71826f550c9f58d6f163770f0755d5",
"assets/assets/images/splash_04.png": "e73afcfe32ac2e4188baa5b62c7b2074",
"assets/assets/images/%25E7%25AE%25AD%25E5%25A4%25B4_%25E5%2590%2591%25E5%258F%25B3.svg": "c79ecae4aec00cf7fec19fec66841247",
"assets/assets/images/login_02.png": "6f66821af0e7a8029aa3f7849d843e14",
"assets/assets/images/%25E7%259B%2591%25E6%258E%25A7.svg": "40b15c8ff44aedaa60e57d5dc9930b48",
"assets/assets/images/%25E9%25A1%25B9%25E7%259B%25AE%25E4%25BF%25A1%25E6%2581%25AF.svg": "e91038fb124b228e15c8fbaba7f54ff5",
"assets/assets/images/login_03.png": "b336317ddc33f9541a4af6296021aca1",
"assets/assets/images/login_01.png": "c7a492a02755e8d6786a8b77ad743848",
"assets/assets/images/undraw_surveillance_re_8tkl.svg": "5097555541464350594ac59385c0ef71",
"assets/assets/images/undraw_monitor_iqpq.svg": "340e1ec51bcaf95ab156aa61f6d00d6e",
"assets/assets/images/%25E7%25AE%25AD%25E5%25A4%25B4.svg": "b0e432061a3874656a391e4a88e3d93d",
"assets/assets/images/404%25E7%259B%2591%25E6%258E%25A7%25E3%2580%2581%25E6%2591%2584%25E5%2583%258F%25E5%25A4%25B4-%25E7%25BA%25BF%25E6%2580%25A7.svg": "439855d428618534f14e3d2053431bb5",
"canvaskit/canvaskit.js": "43fa9e17039a625450b6aba93baf521e",
"canvaskit/profiling/canvaskit.js": "f3bfccc993a1e0bfdd3440af60d99df4",
"canvaskit/profiling/canvaskit.wasm": "a9610cf39260f60fbe7524a785c66101",
"canvaskit/canvaskit.wasm": "04ed3c745ff1dee16504be01f9623498"
};

// The application shell files that are downloaded before a service worker can
// start.
const CORE = [
  "/",
"main.dart.js",
"index.html",
"assets/NOTICES",
"assets/AssetManifest.json",
"assets/FontManifest.json"];
// During install, the TEMP cache is populated with the application shell files.
self.addEventListener("install", (event) => {
  self.skipWaiting();
  return event.waitUntil(
    caches.open(TEMP).then((cache) => {
      return cache.addAll(
        CORE.map((value) => new Request(value, {'cache': 'reload'})));
    })
  );
});

// During activate, the cache is populated with the temp files downloaded in
// install. If this service worker is upgrading from one with a saved
// MANIFEST, then use this to retain unchanged resource files.
self.addEventListener("activate", function(event) {
  return event.waitUntil(async function() {
    try {
      var contentCache = await caches.open(CACHE_NAME);
      var tempCache = await caches.open(TEMP);
      var manifestCache = await caches.open(MANIFEST);
      var manifest = await manifestCache.match('manifest');
      // When there is no prior manifest, clear the entire cache.
      if (!manifest) {
        await caches.delete(CACHE_NAME);
        contentCache = await caches.open(CACHE_NAME);
        for (var request of await tempCache.keys()) {
          var response = await tempCache.match(request);
          await contentCache.put(request, response);
        }
        await caches.delete(TEMP);
        // Save the manifest to make future upgrades efficient.
        await manifestCache.put('manifest', new Response(JSON.stringify(RESOURCES)));
        return;
      }
      var oldManifest = await manifest.json();
      var origin = self.location.origin;
      for (var request of await contentCache.keys()) {
        var key = request.url.substring(origin.length + 1);
        if (key == "") {
          key = "/";
        }
        // If a resource from the old manifest is not in the new cache, or if
        // the MD5 sum has changed, delete it. Otherwise the resource is left
        // in the cache and can be reused by the new service worker.
        if (!RESOURCES[key] || RESOURCES[key] != oldManifest[key]) {
          await contentCache.delete(request);
        }
      }
      // Populate the cache with the app shell TEMP files, potentially overwriting
      // cache files preserved above.
      for (var request of await tempCache.keys()) {
        var response = await tempCache.match(request);
        await contentCache.put(request, response);
      }
      await caches.delete(TEMP);
      // Save the manifest to make future upgrades efficient.
      await manifestCache.put('manifest', new Response(JSON.stringify(RESOURCES)));
      return;
    } catch (err) {
      // On an unhandled exception the state of the cache cannot be guaranteed.
      console.error('Failed to upgrade service worker: ' + err);
      await caches.delete(CACHE_NAME);
      await caches.delete(TEMP);
      await caches.delete(MANIFEST);
    }
  }());
});

// The fetch handler redirects requests for RESOURCE files to the service
// worker cache.
self.addEventListener("fetch", (event) => {
  if (event.request.method !== 'GET') {
    return;
  }
  var origin = self.location.origin;
  var key = event.request.url.substring(origin.length + 1);
  // Redirect URLs to the index.html
  if (key.indexOf('?v=') != -1) {
    key = key.split('?v=')[0];
  }
  if (event.request.url == origin || event.request.url.startsWith(origin + '/#') || key == '') {
    key = '/';
  }
  // If the URL is not the RESOURCE list then return to signal that the
  // browser should take over.
  if (!RESOURCES[key]) {
    return;
  }
  // If the URL is the index.html, perform an online-first request.
  if (key == '/') {
    return onlineFirst(event);
  }
  event.respondWith(caches.open(CACHE_NAME)
    .then((cache) =>  {
      return cache.match(event.request).then((response) => {
        // Either respond with the cached resource, or perform a fetch and
        // lazily populate the cache.
        return response || fetch(event.request).then((response) => {
          cache.put(event.request, response.clone());
          return response;
        });
      })
    })
  );
});

self.addEventListener('message', (event) => {
  // SkipWaiting can be used to immediately activate a waiting service worker.
  // This will also require a page refresh triggered by the main worker.
  if (event.data === 'skipWaiting') {
    self.skipWaiting();
    return;
  }
  if (event.data === 'downloadOffline') {
    downloadOffline();
    return;
  }
});

// Download offline will check the RESOURCES for all files not in the cache
// and populate them.
async function downloadOffline() {
  var resources = [];
  var contentCache = await caches.open(CACHE_NAME);
  var currentContent = {};
  for (var request of await contentCache.keys()) {
    var key = request.url.substring(origin.length + 1);
    if (key == "") {
      key = "/";
    }
    currentContent[key] = true;
  }
  for (var resourceKey of Object.keys(RESOURCES)) {
    if (!currentContent[resourceKey]) {
      resources.push(resourceKey);
    }
  }
  return contentCache.addAll(resources);
}

// Attempt to download the resource online before falling back to
// the offline cache.
function onlineFirst(event) {
  return event.respondWith(
    fetch(event.request).then((response) => {
      return caches.open(CACHE_NAME).then((cache) => {
        cache.put(event.request, response.clone());
        return response;
      });
    }).catch((error) => {
      return caches.open(CACHE_NAME).then((cache) => {
        return cache.match(event.request).then((response) => {
          if (response != null) {
            return response;
          }
          throw error;
        });
      });
    })
  );
}
