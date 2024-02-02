package main

import (
	"strings"
)

// html templates are here instead of an .html file to avoid depending on external files
// in this way, everything is inside the binary

func fillCSSintoTemplate(template string) string {
	if strings.Contains(template, "[CSS]") {
		template = strings.Replace(template, "[CSS]", cssTemplate, -1)
	}
	return template
}

var dirTemplate = `
<!DOCTYPE html>
<html>
<title>{{.Title}}</title>
[CSS]
<style>
body {
	box-sizing: border-box;
	min-width: 200px;
	max-width: 980px;
	margin: 0 auto;
	padding: 45px;
}

@media (prefers-color-scheme: dark) {
	body {
		background-color: #0d1117;
	}
}
</style>
<body class="dark-theme">
<article class="markdown-body">
    <input onclick="switchThemeClick()" type="checkbox" id="themeSwitcher" style="float:right;">

{{.Content}}

<script>
  let theme = localStorage.getItem("theme");
  if (theme === "light") {
    document.getElementById("themeSwitcher").checked = false;
    document.body.className = theme;
  } else {
    document.getElementById("themeSwitcher").checked = true;
  }

  function switchThemeClick() {
    theme = localStorage.getItem("theme");
    if (theme === "light") {
      document.getElementById("themeSwitcher").checked = true;
      theme = "dark-theme";
      localStorage.setItem("theme", theme);
    } else {
      document.getElementById("themeSwitcher").checked = false;
      theme = "light";
      localStorage.setItem("theme", theme);
    }
    document.body.className = theme;
  }
</script>

</article>
</body>
</html>
`

var htmlTemplate = `
<!DOCTYPE html>
<html>
<title>{{.Title}}</title>
[CSS]
<style>
body {
	box-sizing: border-box;
	min-width: 200px;
	max-width: 980px;
	margin: 0 auto;
	padding: 45px;
}

@media (prefers-color-scheme: dark) {
	body {
		background-color: #0d1117;
	}
}
</style>

<body class="dark-theme">
<article class="markdown-body">
<input onclick="switchThemeClick()" type="checkbox" id="themeSwitcher" style="float:right;">
<a href="/" title="go to root">root</a>
<div id="content" class="container">
{{.Content}}
</div>


<!-- TODO: make all the imports locally, to allow working offline -->

<!-- LaTex renderization -->
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/katex@0.13.11/dist/katex.min.css" integrity="sha384-Um5gpz1odJg5Z4HAmzPtgZKdTBHZdw8S29IecapCSB31ligYPhHQZMIlWLYQGVoc" crossorigin="anonymous">

<!-- The loading of KaTeX is deferred to speed up page rendering -->
<script defer src="https://cdn.jsdelivr.net/npm/katex@0.13.11/dist/katex.min.js" integrity="sha384-YNHdsYkH6gMx9y3mRkmcJ2mFUjTd0qNQQvY9VYZgQd7DcN7env35GzlmFaZ23JGp" crossorigin="anonymous"></script>

<!-- To automatically render math in text elements, include the auto-render extension: -->
<script defer src="https://cdn.jsdelivr.net/npm/katex@0.13.11/dist/contrib/auto-render.min.js" integrity="sha384-vZTG03m+2yp6N6BNi5iM4rW4oIwk5DfcNdFfxkk9ZWpDriOkXX8voJBFrAO7MpVl" crossorigin="anonymous"></script>

<!-- Mermaidjs -->
<script type="module">
  import mermaid from 'https://cdn.jsdelivr.net/npm/mermaid@10/dist/mermaid.esm.min.mjs';
</script>

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
		      {left: '$', right: '$', display: false},
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
  } else {
    document.getElementById("themeSwitcher").checked = true;
  }

  function switchThemeClick() {
    theme = localStorage.getItem("theme");
    if (theme === "light") {
      document.getElementById("themeSwitcher").checked = true;
      theme = "dark-theme";
      localStorage.setItem("theme", theme);
    } else {
      document.getElementById("themeSwitcher").checked = false;
      theme = "light";
      localStorage.setItem("theme", theme);
    }
    document.body.className = theme;
  }

</script>
</article>
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
