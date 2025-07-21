function addRow(tableNameID, templateRowID) {
  const table = document.getElementById(tableNameID);
  const templateRow = document.getElementById(templateRowID);
  const clone = templateRow.cloneNode(true);

  // Remove the ID so we don't end up duplicating it and unhide it
  clone.removeAttribute("id");
  clone.style.display = ""; 

  // Enable all the disabled input fields
  // They are disabled so that they don't get grabbed by the backend
  clone.querySelectorAll('input, select').forEach(el => el.disabled = false);

  table.appendChild(clone);
}

function addShareRow() {
  addRow("shares-table", "template-share-row");
}

function addPaymentRow() {
  addRow("payments-table", "template-payment-row");
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
