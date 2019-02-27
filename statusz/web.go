package statusz

import (
	"net/http"
	"strings"
)

const (
	statuszPath     = "/statusz"
	jsVarStatusAddr = `{{GETSTATUS_ADDR}}`
)

const webpage = `<!DOCTYPE html>
<html lang="en-US">
<head>
  <meta charset="utf-8">
  <title>Statusz</title>
</head>
<body>

<h1 id="message"></h1>
<h2 id="timestamp"></h2>
<div id="metrics"></div>

<script>
  console.error = console.error || console.log;
  (function (document) {
    "use strict";

    function message(s, id) {
      id = id || "message";
      document.getElementById(id).textContent = s;
    }

    function update() {
      const request = new XMLHttpRequest();
      request.onreadystatechange = function () {
        if (request.readyState === XMLHttpRequest.DONE) {
          if (request.status === 200) {
            const body = JSON.parse(request.responseText);
            handleSuccess(body);
          } else {
            message("unexpected status: " + request.status)
          }
        }
      };
      const now = new Date().getTime();
      request.open("POST", "http://{{GETSTATUS_ADDR}}/GetStatus?ts=" + now, true);
      request.send('{}');
    }

    function handleSuccess(body) {
      const status = body.status;
      message("");
      message(new Date(status.timestamp.nanoseconds / 1000000).toString(), "timestamp");
      const sortedGroups = sortedByName(status.groups);
      const el = document.getElementById("metrics");
      el.innerHTML = "";
      handleMetrics(el, "System", sortedGroups.map["system"]);
      for (const index in sortedGroups.names) {
        const name = sortedGroups.names[index];
        if (name !== "system") {
          handleMetrics(el, name, sortedGroups.map[name]);
        }
      }
    }

    function handleMetrics(el, name, system) {
      el.innerHTML += "<h1>" + name + "</h1><dl>";
      handleSortedByName(system.metrics, function (name, value) {
        el.innerHTML += "<dt>" + name + "</dt><dd>" + metricValueString(value) + "</dd>";
      });
      el.innerHTML += "</dl>";
    }

    function metricValueString(metric) {
      const value = getOneof(metric, ["i64", "u64", "f64", "string", "bool"]);
      if (value !== undefined) {
        return value;
      }
      if (metric.duration !== undefined) {
        const d = metric.duration;
        if (d.nanoseconds !== undefined) {
          return normalizeNanos(d.nanoseconds);
        } else if (d.milliseconds !== undefined) {
          return normalizeNanos(d.milliseconds * 1000000);
        } else if (d.seconds !== undefined) {
          return normalizeNanos(d.seconds * 1000000000);
        } else {
          console.error("unknown duration type");
        }
      }
    }

    function normalizeNanos(n) {
      if (n < 1000) {
        return n + "ns";
      } else if (n < 1000000) {
        return n / 1000 + "μs";
      } else if (n < 1000000000) {
        return n / 1000000 + "ms";
      } else if (n < 1000000000000) {
        return n / 1000000000 + "s";
      }
      return n;
    }

    function getOneof(obj, props) {
      for (const index in props) {
        const type = props[index];
        if (obj[type] !== undefined) {
          return obj[type];
        }
      }
    }

    function sortedByName(obj) {
      const result = {names: [], map: {}};
      for (const index in obj) {
        const group = obj[index];
        const name = group.name;
        result.names.push(name);
        result.map[name] = group;
      }
      return result;
    }

    function handleSortedByName(obj, f) {
      const sorted = sortedByName(obj);
      for (const index in sorted.names) {
        const name = sorted.names[index];
        f(name, sorted.map[name]);
      }
    }

    window.updateStatusz = function () {
      update();
      setInterval(update, 1000);
    }
  })(document);
  window.updateStatusz();
</script>
</body>
</html>`

func RegisterStatuszWebpage(statuszAddr string, mux *http.ServeMux) {
	handler := NewStatusPageHandler(statuszAddr)
	mux.Handle(statuszPath, handler)
}

func NewStatusPageHandler(statuszAddr string) http.Handler {
	webpageBytes := []byte(strings.Replace(webpage, jsVarStatusAddr, statuszAddr, -1))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "text/html")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write(webpageBytes); err != nil {
			silentlyDropError(err)
		}
	})
}

// silentlyDropError deliberately ignores an error and prevent linters from complaining.
func silentlyDropError(err error) {
}
