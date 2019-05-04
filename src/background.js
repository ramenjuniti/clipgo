const id = "clip_go";
const title = "Clip Go";
const type = "normal";
const contexts = ["selection"];

let code;
let err;

chrome.runtime.onMessage.addListener(message => {
  const result = codeFormat(message.code);
  if (result.err) {
    err = result.err;
    chrome.contextMenus.update(id, {
      title: `${title} Error[${err}]`,
      enabled: false
    });
  } else {
    code = result.newCode;
    chrome.contextMenus.update(id, {
      title: title,
      enabled: true
    });
  }
});

chrome.contextMenus.create({
  id: id,
  title: title,
  type: type,
  contexts: contexts
});

chrome.contextMenus.onClicked.addListener(() => {
  chrome.tabs.create({ url: "https://play.golang.org/" }, tab => {
    chrome.tabs.executeScript(tab.id, {
      code: `document.getElementById("code").value = \`${code}\`;document.getElementById("run").click();`
    });
  });
});
