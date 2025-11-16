import browser from "webextension-polyfill";
import { BackgroundEvent, BgToExtensionMessageType } from "../utils/types";
import { Message } from "./types";
import { setTitle } from "./utils";

let socket: WebSocket | null = null;

browser.tabs.onUpdated.addListener(async (tabId, changeInfo, tab) => {
  if (
    changeInfo.status !== "complete" ||
    !tab.url?.startsWith("https://leetcode.com/problems/") ||
    !socket
  )
    return;

  try {
    const [
      {
        result: { title, difficulty },
      },
    ] = await browser.scripting.executeScript({
      target: { tabId },
      func: () => {
        const titleElement = document.querySelector(
          ".flexlayout__tab > div >div>div>div>div>a"
        );
        const difficultyElement = document.querySelector(
          ".flexlayout__tab > div >div>div:nth-child(2)>div"
        );

        return {
          title: titleElement?.textContent,
          difficulty: difficultyElement?.textContent,
        };
      },
    });

    if (title && difficulty) {
      const message: Message = {
        title: setTitle(title, difficulty),
        url: tab.url,
      };
      socket.send(JSON.stringify(message));
    }
  } catch (e) {
    console.warn("Injection failed", e);
  }
});

browser.runtime.onMessage.addListener((msg) => {
  const response = msg as BackgroundEvent;

  switch (response.type) {
    case BgToExtensionMessageType.CONNECT:
      socket = new WebSocket(`ws://localhost:${response.port}`);
      break;
    case BgToExtensionMessageType.IS_CONNECTED:
      browser.runtime.sendMessage({
        type: "isConnectedValue",
        isConnected: socket !== null,
      });
      break;
    case "disconnect":
      socket?.close();
      socket = null;
      break;
  }
});
