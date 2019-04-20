document.addEventListener("selectionchange", () => {
  chrome.runtime.sendMessage({ code: window.getSelection().toString() });
});
