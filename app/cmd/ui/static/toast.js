document.addEventListener("htmx:load", function() {
    // Ensure HTMX is loaded and events are bound
    htmx.on("htmx:responseError", function(event) {
        const errorMessage = event.detail.xhr.responseText || "HTTP response error!";

        Toastify({
            text: errorMessage,
            duration: 3000, // Duration in milliseconds
            close: true, // Show close button
            gravity: "bottom", // Toast position
            position: "center", // Horizontal position
        }).showToast();
    });
  });