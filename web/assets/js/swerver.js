(function(window, document, $) {
    "use strict";

    const ajaxCalls = {
        ip: { container: "div#sw-ip-container", uri: "ip", onLoad: true }
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
            })
            .catch(err => {
                console.error(`ajaxCall() :: Error retrieving ${url}. ${err}`);
            });
    }

    function showLoading(container) {
        $(container).html(
            "<div class='sw-loading'><img src='/assets/graphics/loading-64x64.gif'/><br>Loading...</div>"
        );
    }
})(window, document, cash);
