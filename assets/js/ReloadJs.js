function changeNavNotify(state) {
	var rq = new XMLHttpRequest();

	if (state.includes("open")) {
		rq.open("GET", "/sidebar/open", true);
	} else if (state.includes("close")) {
		rq.open("GET", "/sidebar/close", true);
	}
	rq.send();
}

function openNav() {
  document.getElementsByClassName("sidebar")[0].style.width = "250px";
	document.getElementsByClassName("main")[0].style.marginLeft = "250px";
	changeNavNotify("open");
}
  
function closeNav() {
	document.getElementsByClassName("sidebar")[0].style.width = "0";
	document.getElementsByClassName("main")[0].style.marginLeft= "0";
	changeNavNotify("close");
}

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

function checkSidebar() {
  var rq = new XMLHttpRequest();
	rq.open("GET", "/sidebar/check", true);
  rq.send();
  rq.onreadystatechange = function () {
    if(rq.readyState === XMLHttpRequest.DONE) {
      var status = rq.status;
      if (status === 0 || (status >= 200 && status < 400)) {
        if (rq.responseText.includes("open")) {
          openNav()
        } else if (rq.responseText.includes("close")) {
          closeNav()
        }
      }
    }
  };
}

try {
  if (window["WebSocket"]) {
    // The reload endpoint is hosted on a statically defined port.
    try {
      tryConnectToReload("ws://{{.BindAddr}}/reload");
    }
    catch (ex) {
      // If an exception is thrown, that means that we couldn't connect to to WebSockets because of mixed content
      // security restrictions, so we try to connect using wss.
      tryConnectToReload("wss://{{.BindAddr}}/reload");

    }
  } else {
    console.log("Your browser does not support WebSockets, cannot connect to the Reload service.");
  }
} catch (ex) {
  console.error('Exception during connecting to Reload:', ex);
}

document.addEventListener("DOMContentLoaded", checkSidebar());
