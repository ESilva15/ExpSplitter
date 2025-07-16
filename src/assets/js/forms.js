// forms main function
document.addEventListener("DOMContentLoaded", function() {
  document.body.addEventListener("formState", function(evt) {
    var messageDiv = document.getElementById("form-state")
    messageDiv.innerHTML = evt.detail.value

    setTimeout(() => {
      messageDiv.innerHTML = ""
    }, 3000);
  });

  // Add the spinner
  document.addEventListener("htmx:beforeRequest", function(e) {
    const target = e.target.closest(".subbutton");
    if (!target) {
      return;
    }

    target.classList.add("loading");
    spinner(target)
  });

  document.addEventListener("htmx:afterRequest", function(e) {
    const target = e.target.closest(".subbutton");
    if (!target) {
      return;
    }
    target.classList.remove("loading");
  });
})
