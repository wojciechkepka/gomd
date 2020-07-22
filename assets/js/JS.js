function themeChange(cb) {
	var rq = new XMLHttpRequest();

	if (cb.checked) {
		rq.open("GET", "/theme/light", true);
	} else {
		rq.open("GET", "/theme/dark", true);
	}
	rq.onreadystatechange = reload;
	rq.send();
}
function codeHlChange(a_theme) {
	var rq = new XMLHttpRequest();
	rq.open("GET", "/theme/" + a_theme.textContent, true);
	rq.onreadystatechange = reload;
	rq.send();
}
function reload() {
	location.reload();
}

