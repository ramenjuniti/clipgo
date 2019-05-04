document.addEventListener("contextmenu", () => {
  chrome.runtime.sendMessage({ code: window.getSelection().toString() });
});
