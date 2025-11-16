export const ConnectionStatus = {
  IDLE: "idle",
  CONNECTED: "connected",
  DISCONNECTED: "disconnected",
} as const;
export type ConnectionStatus = (typeof ConnectionStatus)[keyof typeof ConnectionStatus];

export const BgToExtensionMessageType = {
  CONNECT: "connect",
  DISCONNECT: "disconnect",
  IS_CONNECTED: "isConnected",
  IS_CONNECTED_VALUE: "isConnectedValue",
} as const;
export type BgToExtensionMessageType =
  (typeof BgToExtensionMessageType)[keyof typeof BgToExtensionMessageType];

export type BackgroundEvent =
  | {
      type: Exclude<BgToExtensionMessageType, typeof BgToExtensionMessageType.CONNECT>;
    }
  | {
      type: typeof BgToExtensionMessageType.CONNECT;
      port: string;
    };
