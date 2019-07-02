(function(window, document, $) {
    "use strict";

    const ajaxCalls = {
        ip: { container: "div#sw-ip-container", uri: "ip", onLoad: true },
        services: {
            container: "div#sw-services-container",
            uri: "services",
            onLoad: true
        },
        scripts: {
            container: "div#sw-scripts-container",
            uri: "scripts",
            onLoad: true
        },
        memory: {
            container: "div#sw-memory-container",
            uri: "memory-usage",
            onLoad: true
        }
    };
    $(function() {
        console.log("running");

        loadInitialAjax();
    });

    function loadInitialAjax() {
        const endKeys = Object.keys(ajaxCalls);

        for (let key of endKeys) {
            const endpoint = ajaxCalls[key];
            if (endpoint.onLoad) {
                ajaxCall(endpoint);
            }
        }
    }

    function ajaxCall(ajax) {
        const url = "/ajax/" + ajax.uri;

        showLoading(ajax.container);
        fetch(url)
            .then(resp => {
                return resp.text();
            })
            .then(html => {
                $(ajax.container).html(html);
                setupListenersOnAjaxButtons();
            })
            .catch(err => {
                console.error(`ajaxCall() :: Error retrieving ${url}. ${err}`);
            });
    }

    function ajaxRunScript(script) {
        const encoded = encodeURIComponent(script);
        const url = `/ajax/run-script?script=${encoded}`;
        fetch(url)
            .then(resp => {
                return resp.text();
            })
            .then(html => {
                $("#sw-tty-container").prepend(html);
            })
            .catch(err => {
                console.log(`ajaxRunScript() :: Error running script ${url}`);
            });
    }

    function setupListenersOnAjaxButtons() {
        $(".sw-ajax-button").off();
        $(".sw-ajax-button").on("click", evt => {
            const $target = $(evt.currentTarget);
            const action = $target.data("action");
            switch (action) {
                case "run-script":
                    const script = $target.data("script");
                    ajaxRunScript(script);
                    break;
                default:
                    break;
            }
        });
    }
    function showLoading(container) {
        $(container).html(
            "<div class='sw-loading'><img src='/assets/graphics/loading-64x64.gif'/><br>Loading...</div>"
        );
    }
})(window, document, cash);
