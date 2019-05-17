const id = "clip_go";
const title = "Clip Go";
const type = "normal";
const contexts = ["selection"];

let code;

chrome.runtime.onMessage.addListener(message => {
  const result = formatter(message.code);
  if (result.err) {
    chrome.contextMenus.update(id, {
      title: `${title} [Error] ${result.output}`,
      enabled: false
    });
  } else {
    code = result.output;
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
