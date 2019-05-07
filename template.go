package main

// html templates are here instead of an .html file to avoid depending on external files
// in this way, everything is inside the binary

const dirTemplate = `
<!DOCTYPE html>
<html>
<title>{{.Title}}</title>
<style>
body {
    font-family: Arial, Helvetica, sans-serif;
    background:#000000;
    color:#cccccc;
}
</style>
<body>

{{.Content}}

</body>
</html>
`

const htmlTemplate = `
<!DOCTYPE html>
<html>
<title>{{.Title}}</title>
<style>
body {
    font-family: Arial, Helvetica, sans-serif;
    background:#000000;
    color:#cccccc;
}
pre, code{
    padding-left: 3px;
    padding-right: 3px;
    background: #333333;
    border-radius: 3px;
}
pre>code {
    color: #c0fffa;
}
h1:after{
    content:' ';
    display:block;
    border:1px solid #cccccc;
}
h2:after{
    content:' ';
    display:block;
    border:0.7px solid #cccccc;
}
</style>
<body>

{{.Content}}

<script>
(function() {
	var conn = new WebSocket("ws://127.0.0.1:8080/ws/{{.Title}}");
	conn.onclose = function(evt) {
	    console.log('Connection closed');
	    alert('Connection closed');
	}
	conn.onmessage = function(evt) {
	    console.log('file updated');
	    // location.reload();
	    console.log("evt", evt);
	    document.getElementsByTagName("BODY")[0].innerHTML = evt.data;
	}
})();
</script>
</body>
</html>
`
