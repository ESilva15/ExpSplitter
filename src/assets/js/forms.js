function addShareRow() {
  const templateRow = document.getElementById("default-payment-row");

  if (!templateRow) {
    console.error("Example row not found");
    return;
  }

  // Clone the hidden row
  const newRow = templateRow.cloneNode(true);
  newRow.removeAttribute("id"); // IDs must be unique
  newRow.style.display = "";    // Make it visible

  // Optional: Reset input values if needed
  const inputs = newRow.querySelectorAll("input");
  inputs.forEach(input => {
    if (input.type === "number") {
      input.value = "";
    }
  });

  // Insert the cloned row after the template row
  templateRow.parentNode.appendChild(newRow);
}

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
