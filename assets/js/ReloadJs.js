function get(ep) {
	var rq = new XMLHttpRequest();
    rq.open("GET", ep, true);
	rq.send();
    return rq;
}

function get_call(ep, func) {
    var rq = get(ep);    
    rq.onreadystatechange = function(){
        if(rq.readyState === XMLHttpRequest.DONE) {
          var status = rq.status;
          console.log(status)
          if (status === 0 || (status >= 200 && status < 400)) {
              console.log(rq.responseText)
              func(rq.responseText);
          }
        }
    };
}

function showDiff() {
    var currLocation = window.location;
    var path = currLocation.href;
    var btn = document.getElementById("diff-btn");
    
    var ep = "/"
    if (btn.innerText === "diff") {
        ep = ep.concat("diff")
        btn.innerText = "content"
    } else {
        ep = ep.concat("content")
        btn.innerText = "diff"
    }

    get_call(path.concat(ep), function(s) {
        var content = document.getElementsByClassName("content")[0];
        content.innerHTML = s;
        btn
    })
}

function changeNavNotify(state) {
	if (state.includes("open")) {
        get("/sidebar/open");
	} else if (state.includes("close")) {
        get("/sidebar/close");
	}
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
  
    conn.onclose = function(){
      setTimeout(function() {
        tryConnectToReload(address);
      }, 2000);
    };
  
    conn.onmessage = function(evt){
        location.reload()
    };
}


function checkSidebar() {
    get_call("/sidebar/check", function(s){
        if (s.includes("open")) {
          openNav()
        } else if (s.includes("close")) {
          closeNav()
        }
    })
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
