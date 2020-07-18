function tryConnectToReload(address) {
    var conn = new WebSocket(address);
  
    conn.onclose = function() {
      setTimeout(function() {
        tryConnectToReload(address);
      }, 2000);
    };
  
    conn.onmessage = function(evt) {
      location.reload()
    };
}

try {
  if (window["WebSocket"]) {
    // The reload endpoint is hosted on a statically defined port.
    try {
      tryConnectToReload("ws://%v/reload");
    }
    catch (ex) {
      // If an exception is thrown, that means that we couldn't connect to to WebSockets because of mixed content
      // security restrictions, so we try to connect using wss.
      tryConnectToReload("wss://%v/reload");
    }
  } else {
    console.log("Your browser does not support WebSockets, cannot connect to the Reload service.");
  }
} catch (ex) {
  console.error('Exception during connecting to Reload:', ex);
}