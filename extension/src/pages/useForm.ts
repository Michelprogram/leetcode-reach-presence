import { onMounted, ref } from "vue";
import browser from "webextension-polyfill";
import { BgToExtensionMessageType, BackgroundEvent, ConnectionStatus } from "../utils/types";

export const useForm = () => {
  const connectionState = ref<ConnectionStatus>(ConnectionStatus.IDLE);

  const onDisconnect = async () => {
    connectionState.value = ConnectionStatus.DISCONNECTED;
    await browser.runtime.sendMessage({
      type: BgToExtensionMessageType.DISCONNECT,
    });
  };

  const onConnect = async (port: string) => {
    connectionState.value = ConnectionStatus.CONNECTED;
    await browser.runtime.sendMessage({
      type: BgToExtensionMessageType.CONNECT,
      port: port,
    });
  };

  browser.runtime.onMessage.addListener((msg) => {
    const response = msg as BackgroundEvent;
    if (response.type === BgToExtensionMessageType.IS_CONNECTED_VALUE) {
      connectionState.value = msg.isConnected ? "connected" : "idle";
    }
  });

  onMounted(async () => {
    browser.runtime.sendMessage({
      type: BgToExtensionMessageType.IS_CONNECTED,
    });
  });

  return { connectionState , onDisconnect, onConnect };
};
