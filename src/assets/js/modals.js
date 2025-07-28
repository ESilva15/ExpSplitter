let pendingDeleteUrl = null;
let deleteTarget = null;

function confirmDelete(button) {
  let modal = document.getElementById("confirm-modal");
  pendingDeleteUrl = button.getAttribute("data-delete-url");
  deleteTarget = button.closest("tr");
  modal.showModal();
}

document.addEventListener("DOMContentLoaded", function() {
  document.body.addEventListener("click", function(evt) {
    const modalCancelBtn = evt.target.closest("#modal-cancel");
    if (!modalCancelBtn) {
      return;
    }

    let modal = document.getElementById("confirm-modal");
    modal.close();
    pendingDeleteUrl = null;
    deleteTarget = null;
  });

  document.body.addEventListener("click", function(evt) {
    const modalConfirmBtn = evt.target.closest("#modal-confirm");
    if (!modalConfirmBtn) {
      return;
    }

    let modal = document.getElementById("confirm-modal");
    if (pendingDeleteUrl) {
      htmx.ajax('DELETE', pendingDeleteUrl, {
        target: deleteTarget,
        swap: 'outerHTML'
      });
    }
    modal.close();
  });
});
