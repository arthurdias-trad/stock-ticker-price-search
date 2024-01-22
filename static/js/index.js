document.addEventListener("DOMContentLoaded", function () {
  let body = document.querySelector("body");
  body.addEventListener("htmx:afterSwap", function () {
    let symbolInput = document.getElementById("symbol");
    let button = document.getElementById("submit-button");
    symbolInput.value = "";

    button.addEventListener("click", function (event) {
      if (symbolInput.value === "") {
        return;
      }
    });

    symbolInput.focus();
    symbolInput.addEventListener("keyup", function (event) {
      event.target.value = event.target.value.toUpperCase();
      if (event.key === "Enter") {
        if (event.target.value === "") {
          return;
        }
        event.preventDefault();
        button.click();
        return;
      }
    });
  });
});
