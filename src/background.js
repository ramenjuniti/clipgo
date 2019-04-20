let code;

chrome.runtime.onMessage.addListener(message => {
  code = message.code;
});

chrome.contextMenus.create({
  type: "normal",
  id: "clip_go",
  title: "Clip Go",
  contexts: ["selection"]
});

chrome.contextMenus.onClicked.addListener(() => {
  chrome.tabs.create({ url: "https://play.golang.org/" }, tab => {
    chrome.tabs.executeScript(tab.id, {
      code: `document.getElementById("code").value = \`${code}\`;document.getElementById("run").click();`
    });
  });
});
