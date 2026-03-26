package serverSwagger

import (
	"bytes"
	"html/template"

	"github.com/gin-gonic/gin"
)

func redocHandler(cfg DocsConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		html := buildRedocHTML(cfg)
		c.Data(200, "text/html; charset=utf-8", html)
	}
}

func buildRedocHTML(cfg DocsConfig) []byte {
	if cfg.Product == "" {
		cfg.Product = "API"
	}

	if cfg.Title == "" {
		cfg.Title = "API Docs"
	}

	if cfg.SpecURL == "" {
		cfg.SpecURL = "/swagger/doc.json"
	}

	tmpl := template.Must(template.New("redoc").Parse(redocTemplate))

	var buf bytes.Buffer
	_ = tmpl.Execute(&buf, cfg)

	return buf.Bytes()
}

const redocTemplate = `
<!DOCTYPE html>
<html>
  <head>
    <title>{{.Title}}</title>
    <meta charset="utf-8"/>
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600&display=swap" rel="stylesheet">

    <style>
      :root {
        --primary: #54E0F7;
        --text: #0f172a;
        --text-muted: #64748b;
        --border: #e2e8f0;
        --bg: #ffffff;
      }

      html {
        scroll-behavior: smooth;
      }

      body {
        margin: 0;
        font-family: 'Inter', sans-serif;
        background: var(--bg);
      }

      /* Header */
      .header {
        height: 56px;
        display: flex;
        align-items: center;
        padding: 0 20px;
        border-bottom: 1px solid var(--border);
        background: rgba(255,255,255,0.8);
        backdrop-filter: blur(8px);
        position: sticky;
        top: 0;
        z-index: 1000;
        box-shadow: 0 1px 2px rgba(0,0,0,0.04);
      }

      .header-content {
        width: 100%;
        display: flex;
        justify-content: space-between;
        align-items: center;
      }

      .brand {
        display: flex;
        align-items: center;
        gap: 8px;
      }

      .brand-name {
        font-weight: 600;
      }

      .product {
        color: var(--text-muted);
        font-size: 13px;
      }

      .divider {
        color: var(--text-muted);
      }

      .header-links a {
        margin-left: 16px;
        font-size: 13px;
        color: var(--text-muted);
        text-decoration: none;
        transition: color 0.2s ease;
      }

      .header-links a:hover {
        color: var(--primary);
      }

      .redoc-wrap {
        padding-top: 4px;
        background-color: #f9f4ee;
      }
    </style>

    <link rel="icon" href="/docs/assets/favicon.ico" />
  </head>

  <body>
    <div class="header">
      <div class="header-content">
        <div class="brand">
          <img src="/docs/assets/logo.png" height="22" />
          <span class="brand-name">Onlyizi</span>
          <span class="divider">/</span>
          <span class="product">{{.Product}}</span>
        </div>

        <div class="header-links">
          <a href="/swagger/index.html" target="_blank">Swagger</a>
          <a href="/v1" target="_blank">API</a>
        </div>
      </div>
    </div>

    <div id="redoc-container" class="redoc-wrap"></div>

    <script src="https://cdn.redoc.ly/redoc/latest/bundles/redoc.standalone.js"></script>

    <script>
      Redoc.init("{{.SpecURL}}", {
        expandResponses: "200,201,204",
        theme: {
          colors: {
            primary: {
              main: "#54E0F7"
            },
            text: {
              primary: "#0f172a",
              secondary: "#64748b"
            },
            border: {
              dark: "#e2e8f0"
            },
            background: {
              default: "#ffffff"
            }
          },
          typography: {
            fontFamily: "Inter, sans-serif",
            fontSize: "14px",
            headings: {
              fontWeight: "600"
            }
          },
          sidebar: {
            backgroundColor: "#fafafa",
            textColor: "#475569",
            activeTextColor: "#54E0F7"
          },
          codeBlock: {
            backgroundColor: "#0f172a"
          },
          menu: {
            activeBgColor: "rgba(84,224,247,0.1)"
          }
        }
      }, document.getElementById("redoc-container"));
    </script>
  </body>
</html>
`
