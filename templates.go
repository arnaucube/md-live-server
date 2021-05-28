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
    background:#1d2021;
    color:#cccccc;
}
.dark a, a:hover, a:visited {
	color: #458588;
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
    background:#1d2021;
    color:#cccccc;
}
a, a:hover, a:visited {
	color: #458588;
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
    max-width: 800px;
  }
}

@media (min-width: 1200px) {
  .container {
    max-width: 800px;
  }
}

@media (min-width: 1400px) {
  .container {
    max-width: 800px;
  }
}
</style>

<body class="dark">
<input onclick="switchThemeClick()" type="checkbox" id="themeSwitcher" style="float:right;">
<a href="/" title="go to root">root</a>
<div id="content" class="container">
{{.Content}}
</div>


<!-- LaTex renderization -->
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/katex@0.13.11/dist/katex.min.css" integrity="sha384-Um5gpz1odJg5Z4HAmzPtgZKdTBHZdw8S29IecapCSB31ligYPhHQZMIlWLYQGVoc" crossorigin="anonymous">

<!-- The loading of KaTeX is deferred to speed up page rendering -->
<script defer src="https://cdn.jsdelivr.net/npm/katex@0.13.11/dist/katex.min.js" integrity="sha384-YNHdsYkH6gMx9y3mRkmcJ2mFUjTd0qNQQvY9VYZgQd7DcN7env35GzlmFaZ23JGp" crossorigin="anonymous"></script>

<!-- To automatically render math in text elements, include the auto-render extension: -->
<script defer src="https://cdn.jsdelivr.net/npm/katex@0.13.11/dist/contrib/auto-render.min.js" integrity="sha384-vZTG03m+2yp6N6BNi5iM4rW4oIwk5DfcNdFfxkk9ZWpDriOkXX8voJBFrAO7MpVl" crossorigin="anonymous"></script>

<script>
    document.addEventListener("DOMContentLoaded", function() {
        renderMathInElement(document.body, {
          // customised options
          // • auto-render specific keys, e.g.:
          delimiters: [
              {left: '$$', right: '$$', display: true},
              {left: '$', right: '$', display: true},
              {left: '\\(', right: '\\)', display: false},
              {left: '\\[', right: '\\]', display: true}
          ],
          // • rendering keys, e.g.:
          throwOnError : false
        });
    });
</script>

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

		renderMathInElement(document.body, {
		  // customised options
		  // • auto-render specific keys, e.g.:
		  delimiters: [
		      {left: '$$', right: '$$', display: true},
		      {left: '$', right: '$', display: true},
		      {left: '\\(', right: '\\)', display: false},
		      {left: '\\[', right: '\\]', display: true}
		  ],
		  // • rendering keys, e.g.:
		  throwOnError : false
		});

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
