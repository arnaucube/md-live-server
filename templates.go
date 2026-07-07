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
	display: flex;
	align-items: flex-start;
	min-width: 200px;
	margin: 0;
	padding: 0;
}

.page-layout {
	box-sizing: border-box;
	flex: 1 1 auto;
	max-width: 980px;
	margin-left: 40px;
	margin-right: auto;
	padding: 45px;
}

.toc {
	position: sticky;
	top: 0;
	flex: 0 0 280px;
	box-sizing: border-box;
	width: 280px;
	height: 100vh;
	max-height: 100vh;
	padding: 45px 14px;
	overflow-y: auto;
	border-right: 1px solid var(--color-border-muted);
}

.toc-title {
	margin: 0 0 12px;
	color: var(--color-fg-muted);
	font-size: 12px;
	font-weight: 600;
	text-transform: uppercase;
}

.toc-list {
	margin: 0;
	padding: 0;
	list-style: none;
}

.toc-list li {
	margin: 2px 0;
}

.toc-list a {
	display: block;
	overflow: hidden;
	padding: 4px 0;
	color: var(--color-fg-muted);
	font-size: 13px;
	line-height: 1.35;
	text-overflow: ellipsis;
	white-space: nowrap;
}

.toc-list a:hover {
	color: var(--color-accent-fg);
	text-decoration: none;
}

.toc-level-2 { padding-left: 2px; }
.toc-level-3 { padding-left: 14px; }
.toc-level-4 { padding-left: 26px; }
.toc-level-5 { padding-left: 38px; }
.toc-level-6 { padding-left: 40px; }

.toc-empty {
	display: none;
}

@media (max-width: 899px) {
	.toc {
		flex-basis: 240px;
		width: 240px;
		padding: 24px 14px;
	}

	.page-layout {
		margin-left: 20px;
	}
}

@media (max-width: 767px) {
	.page-layout {
		margin-left: 0;
		padding: 24px;
	}

	.toc {
		flex-basis: 200px;
		width: 200px;
		padding: 24px 12px;
	}
}

</style>

<body class="dark-theme">
<nav id="toc" class="toc markdown-body" aria-label="Table of contents">
	<div class="toc-title">Contents</div>
	<ul id="tocList" class="toc-list"></ul>
</nav>
<main class="page-layout">
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
  mermaid.initialize({ startOnLoad: false });

  window.renderMermaid = async function() {
    document.querySelectorAll('pre > code.language-mermaid').forEach(function(codeBlock) {
      const pre = codeBlock.parentElement;
      pre.className = 'mermaid';
      pre.textContent = codeBlock.textContent;
    });

    await mermaid.run({ querySelector: '.mermaid' });
  };
</script>

<script>
    document.addEventListener("DOMContentLoaded", function() {
        buildTableOfContents();

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

        if (window.renderMermaid) {
          window.renderMermaid().catch(console.error);
        }
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
		buildTableOfContents();

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

		if (window.renderMermaid) {
		  window.renderMermaid().catch(console.error);
		}
	}
})();

  function slugifyHeading(text) {
    return text.trim().toLowerCase()
      .replace(/[^\w\s-]/g, "")
      .replace(/\s+/g, "-")
      .replace(/-+/g, "-") || "heading";
  }

  function buildTableOfContents() {
    const content = document.getElementById("content");
    const toc = document.getElementById("toc");
    const tocList = document.getElementById("tocList");
    if (!content || !toc || !tocList) {
      return;
    }

    const headings = content.querySelectorAll("h1, h2, h3, h4, h5, h6");
    const usedIds = new Set();
    tocList.innerHTML = "";

    headings.forEach(function(heading) {
      const level = parseInt(heading.tagName.substring(1), 10);
      let id = heading.id || slugifyHeading(heading.textContent);
      const baseId = id;
      let suffix = 2;
      while (usedIds.has(id)) {
        id = baseId + "-" + suffix;
        suffix++;
      }
      heading.id = id;
      usedIds.add(id);

      const item = document.createElement("li");
      item.className = "toc-level-" + level;

      const link = document.createElement("a");
      link.href = "#" + id;
      link.textContent = heading.textContent;
      link.title = heading.textContent;

      item.appendChild(link);
      tocList.appendChild(item);
    });

    toc.classList.toggle("toc-empty", headings.length === 0);
  }


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
</main>
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
