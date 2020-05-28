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
}
.dark {
    background:#000000;
    color:#cccccc;
}
</style>
<body class="dark">
    <input onclick="switchThemeClick()" type="checkbox" id="themeSwitcher" style="float:right;">

{{.Content}}

<script>
  let theme = localStorage.getItem("theme");
  if (theme === "light") {
    document.getElementById("themeSwitcher").checked = false;
    document.body.className = theme;
  }

  function switchThemeClick() {
    theme = localStorage.getItem("theme");
    if (theme === "light") {
      document.getElementById("themeSwitcher").checked = true;
      theme = "dark";
      localStorage.setItem("theme", theme);
    } else {
      document.getElementById("themeSwitcher").checked = false;
      theme = "light";
      localStorage.setItem("theme", theme);
    }
    document.body.className = theme;
  }
</script>

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
}
.dark {
    background:#000000;
    color:#cccccc;
}
pre, code{
    padding-left: 3px;
    padding-right: 3px;
    border-radius: 3px;
    background: #cccccc;
}
.dark pre, .dark code{
    padding-left: 3px;
    padding-right: 3px;
    border-radius: 3px;
    background: #333333;
}
pre{
    padding: 5px;
}
pre>code {
    color: #000000;
}
.dark pre>code {
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
.dark th{
    padding: 3px;
    border: 1px solid #234b4f;
    background-color: #143134;
}
td{
    padding: 3px;
    border: 1px solid #234b4f;
}
.container {
  width: 100%;
  padding-right: 1rem;
  padding-left: 1rem;
  margin-right: auto;
  margin-left: auto;
}

@media (min-width: 576px) {
  .container {
    max-width: 540px;
  }
}

@media (min-width: 768px) {
  .container {
    max-width: 720px;
  }
}

@media (min-width: 992px) {
  .container {
    max-width: 900px;
  }
}

@media (min-width: 1200px) {
  .container {
    max-width: 900px;
  }
}

@media (min-width: 1400px) {
  .container {
    max-width: 900px;
  }
}
</style>
<body class="dark">
<input onclick="switchThemeClick()" type="checkbox" id="themeSwitcher" style="float:right;">
<a href="/" title="go to root">root</a>
<div id="content" class="container">
{{.Content}}
</div>

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
	    // console.log("evt", evt);
	    document.getElementById("content").innerHTML = evt.data;
	}
})();


  let theme = localStorage.getItem("theme");
  if (theme === "light") {
    document.getElementById("themeSwitcher").checked = false;
    document.body.className = theme;
  }

  function switchThemeClick() {
    theme = localStorage.getItem("theme");
    if (theme === "light") {
      document.getElementById("themeSwitcher").checked = true;
      theme = "dark";
      localStorage.setItem("theme", theme);
    } else {
      document.getElementById("themeSwitcher").checked = false;
      theme = "light";
      localStorage.setItem("theme", theme);
    }
    document.body.className = theme;
  }

</script>
</body>
</html>
`

const errTemplate = `
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
    <a href="/" title="go to root">root</a>

<h3>File doesn't exist</h3>
<br><br>
<a href="/">Back to dir menu</a>

</body>
</html>
`
